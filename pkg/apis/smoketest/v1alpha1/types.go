package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SmokeTestList contains a list of SmokeTest items.
type SmokeTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SmokeTest `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SmokeTest struct has the spec and status of the SmokeTest object.
type SmokeTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SmokeTestSpec   `json:"spec"`
	Status            SmokeTestStatus `json:"status,omitempty"`
}

// SmokeTestSpec will contain the details of the SmokeTest spec.
type SmokeTestSpec struct {
	// Fill me
}

// SmokeTestStatus has the output status of each instance of the smoke test.
type SmokeTestStatus struct {
	TestOutput string
}
