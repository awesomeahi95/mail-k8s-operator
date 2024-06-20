package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// EmailSpec defines the desired state of Email
type EmailSpec struct {
	Body            string `json:"body,omitempty"`
	RecipientEmail  string `json:"recipientEmail,omitempty"`
	SenderConfigRef string `json:"senderConfigRef,omitempty"`
	Subject         string `json:"subject,omitempty"`
}

// EmailStatus defines the observed state of Email
type EmailStatus struct {
	DeliveryStatus string `json:"deliveryStatus,omitempty"`
	Error          string `json:"error,omitempty"`
	MessageID      string `json:"messageId,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Email is the Schema for the emails API
type Email struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmailSpec   `json:"spec,omitempty"`
	Status EmailStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EmailList contains a list of Email
type EmailList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Email `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Email{}, &EmailList{})
}
