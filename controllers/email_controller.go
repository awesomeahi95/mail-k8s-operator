package controllers

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/mailersend/mailersend-go"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	emailv1 "github.com/awesomeahi95/mailerlite/api/v1"
)

// EmailReconciler reconciles an Email object
type EmailReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=email.mailerlitetask.com,resources=emails,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=email.mailerlitetask.com,resources=emails/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=email.mailerlitetask.com,resources=emails/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=secrets,verbs=get;list;watch

func (r *EmailReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the Email instance
	var email emailv1.Email
	if err := r.Get(ctx, req.NamespacedName, &email); err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("Email resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get Email")
		return ctrl.Result{}, err
	}

	log.Info("Email resource found", "email", email)

	// Check if the email has already been sent
	if email.Status.DeliveryStatus == "Sent" {
		log.Info("Email already sent, skipping", "email", email.Name)
		return ctrl.Result{}, nil
	}

	// Fetch the EmailSenderConfig instance
	var emailSenderConfig emailv1.EmailSenderConfig
	if err := r.Get(ctx, client.ObjectKey{Name: email.Spec.SenderConfigRef, Namespace: req.Namespace}, &emailSenderConfig); err != nil {
		log.Error(err, "Failed to get EmailSenderConfig", "EmailSenderConfig", email.Spec.SenderConfigRef)
		email.Status.DeliveryStatus = "Failed"
		email.Status.Error = "EmailSenderConfig not found"
		if updateErr := r.Status().Update(ctx, &email); updateErr != nil {
			log.Error(updateErr, "Failed to update Email status")
		}
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("EmailSenderConfig found", "EmailSenderConfig", emailSenderConfig)

	// Fetch the API token and from-email from the secret
	apiToken, fromEmail, err := getSecretValues(req.Namespace, emailSenderConfig.Spec.ApiTokenSecretRef)
	if err != nil {
		log.Error(err, "Failed to get API token or from-email from secret")
		email.Status.DeliveryStatus = "Failed"
		email.Status.Error = "Failed to get API token or from-email from secret"
		if updateErr := r.Status().Update(ctx, &email); updateErr != nil {
			log.Error(updateErr, "Failed to update Email status")
		}
		return ctrl.Result{}, err
	}

	// Send the email
	log.Info("Sending email", "recipient", email.Spec.RecipientEmail)
	deliveryStatus, messageID, err := sendEmail(email, apiToken, fromEmail)
	if err != nil {
		email.Status.DeliveryStatus = "Failed"
		email.Status.Error = err.Error()
	} else {
		email.Status.DeliveryStatus = deliveryStatus
		email.Status.MessageID = messageID
		email.Status.Error = ""
	}

	// Update the status of the Email resource
	if err := r.Status().Update(ctx, &email); err != nil {
		log.Error(err, "Failed to update Email status")
		return ctrl.Result{}, err
	}

	log.Info("Email status updated successfully", "status", email.Status)
	return ctrl.Result{}, nil
}

func sendEmail(email emailv1.Email, apiToken, fromEmail string) (string, string, error) {
	return sendEmailUsingMailerSend(email, apiToken, fromEmail)
}

func sendEmailUsingMailerSend(email emailv1.Email, apiToken, fromEmail string) (string, string, error) {
	log := log.Log

	log.Info("Creating MailerSend client")
	ms := mailersend.NewMailersend(apiToken)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  "MailerSend",
		Email: fromEmail,
	}

	recipients := []mailersend.Recipient{
		{
			Name:  "Recipient",
			Email: email.Spec.RecipientEmail,
		},
	}

	subject := email.Spec.Subject
	text := email.Spec.Body
	html := email.Spec.Body

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)

	log.Info("Sending email with MailerSend", "from", from, "recipients", recipients, "subject", subject, "text", text, "html", html)

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		log.Error(err, "Failed to send email with MailerSend")
		return "Failed", "", err
	}

	messageID := res.Header.Get("X-Message-Id")
	log.Info("Email sent successfully", "messageID", messageID)
	return "Sent", messageID, nil
}

func getSecretValues(namespace, secretName string) (string, string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", "", err
	}

	k8sClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", "", err
	}

	secret, err := k8sClient.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		return "", "", err
	}

	apiToken := string(secret.Data["api-token"])
	fromEmail := string(secret.Data["from-email"])

	decodedApiToken, err := base64.StdEncoding.DecodeString(apiToken)
	if err != nil {
		return "", "", err
	}

	decodedFromEmail, err := base64.StdEncoding.DecodeString(fromEmail)
	if err != nil {
		return "", "", err
	}

	return string(decodedApiToken), string(decodedFromEmail), nil
}

func (r *EmailReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&emailv1.Email{}).
		Complete(r)
}
