/*
Copyright 2022 The VolSync authors.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReplicationDestinationSpec defines the desired state of ReplicationDestination
type ReplicationDestinationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ReplicationDestination. Edit replicationdestination_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ReplicationDestinationStatus defines the observed state of ReplicationDestination
type ReplicationDestinationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ReplicationDestination is the Schema for the replicationdestinations API
type ReplicationDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReplicationDestinationSpec   `json:"spec,omitempty"`
	Status ReplicationDestinationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ReplicationDestinationList contains a list of ReplicationDestination
type ReplicationDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ReplicationDestination `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ReplicationDestination{}, &ReplicationDestinationList{})
}
