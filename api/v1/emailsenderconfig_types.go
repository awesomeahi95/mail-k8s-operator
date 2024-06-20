package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EmailSenderConfigSpec defines the desired state of EmailSenderConfig
type EmailSenderConfigSpec struct {
	ApiTokenSecretRef string `json:"apiTokenSecretRef,omitempty"`
	SenderEmail       string `json:"senderEmail,omitempty"`
}

// EmailSenderConfigStatus defines the observed state of EmailSenderConfig
type EmailSenderConfigStatus struct {
	Error string `json:"error,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EmailSenderConfig is the Schema for the emailsenderconfigs API
type EmailSenderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmailSenderConfigSpec   `json:"spec,omitempty"`
	Status EmailSenderConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EmailSenderConfigList contains a list of EmailSenderConfig
type EmailSenderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailSenderConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EmailSenderConfig{}, &EmailSenderConfigList{})
}
