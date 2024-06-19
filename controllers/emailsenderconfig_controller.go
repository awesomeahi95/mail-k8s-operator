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

// EmailSenderConfigReconciler reconciles an EmailSenderConfig object
type EmailSenderConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=email.mailerlitetask.com,resources=emailsenderconfigs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=email.mailerlitetask.com,resources=emailsenderconfigs/status,verbs=get;update;patch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *EmailSenderConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	// Fetch the EmailSenderConfig instance
	var emailSenderConfig emailv1.EmailSenderConfig
	if err := r.Get(ctx, req.NamespacedName, &emailSenderConfig); err != nil {
		if errors.IsNotFound(err) {
			// EmailSenderConfig resource not found. Ignoring since object must be deleted.
			log.Info("EmailSenderConfig resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get EmailSenderConfig")
		return ctrl.Result{}, err
	}

	log.Info("EmailSenderConfig resource found", "emailsenderconfig", emailSenderConfig)

	emailSenderConfig.Status.Error = ""

	// Update the status of the EmailSenderConfig resource
	if err := r.Status().Update(ctx, &emailSenderConfig); err != nil {
		log.Error(err, "Failed to update EmailSenderConfig status")
		return ctrl.Result{}, err
	}

	log.Info("EmailSenderConfig status updated successfully")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *EmailSenderConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&emailv1.EmailSenderConfig{}).
		Complete(r)
}
