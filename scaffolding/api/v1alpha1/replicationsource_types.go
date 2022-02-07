/*
Copyright 2022 The VolSync authors.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReplicationSourceSpec defines the desired state of ReplicationSource
type ReplicationSourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ReplicationSource. Edit replicationsource_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ReplicationSourceStatus defines the observed state of ReplicationSource
type ReplicationSourceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ReplicationSource is the Schema for the replicationsources API
type ReplicationSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReplicationSourceSpec   `json:"spec,omitempty"`
	Status ReplicationSourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReplicationSourceList contains a list of ReplicationSource
type ReplicationSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReplicationSource{}, &ReplicationSourceList{})
}
