package controllers

import (
        "context"

        "k8s.io/apimachinery/pkg/api/errors"
        "k8s.io/apimachinery/pkg/runtime"
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

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *EmailReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
        log := log.FromContext(ctx)

        // Fetch the Email instance
        var email emailv1.Email
        if err := r.Get(ctx, req.NamespacedName, &email); err != nil {
                if errors.IsNotFound(err) {
                        // Email resource not found. Ignoring since object must be deleted.
                        log.Info("Email resource not found. Ignoring since object must be deleted.")
                        return ctrl.Result{}, nil
                }
                // Error reading the object - requeue the request.
                log.Error(err, "Failed to get Email")
                return ctrl.Result{}, err
        }

        log.Info("Email resource found", "email", email)

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

        // Simulate sending email
        log.Info("Simulating email sending", "recipient", email.Spec.RecipientEmail)
        email.Status.DeliveryStatus = "Sent"
        email.Status.MessageID = "some-id"
        email.Status.Error = ""

        // Update the status of the Email resource
        if err := r.Status().Update(ctx, &email); err != nil {
                log.Error(err, "Failed to update Email status")
                return ctrl.Result{}, err
        }

        log.Info("Email status updated successfully")
        return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EmailReconciler) SetupWithManager(mgr ctrl.Manager) error {
        return ctrl.NewControllerManagedBy(mgr).
                For(&emailv1.Email{}).
                Complete(r)
}
