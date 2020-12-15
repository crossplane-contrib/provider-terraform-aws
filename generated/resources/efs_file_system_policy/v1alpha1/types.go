/*
	Copyright 2019 The Crossplane Authors.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// EfsFileSystemPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EfsFileSystemPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EfsFileSystemPolicySpec   `json:"spec"`
	Status EfsFileSystemPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EfsFileSystemPolicy contains a list of EfsFileSystemPolicyList
type EfsFileSystemPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EfsFileSystemPolicy `json:"items"`
}

// A EfsFileSystemPolicySpec defines the desired state of a EfsFileSystemPolicy
type EfsFileSystemPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EfsFileSystemPolicyParameters `json:"forProvider"`
}

// A EfsFileSystemPolicyParameters defines the desired state of a EfsFileSystemPolicy
type EfsFileSystemPolicyParameters struct {
	FileSystemId string `json:"file_system_id"`
	Policy       string `json:"policy"`
}

// A EfsFileSystemPolicyStatus defines the observed state of a EfsFileSystemPolicy
type EfsFileSystemPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EfsFileSystemPolicyObservation `json:"atProvider"`
}

// A EfsFileSystemPolicyObservation records the observed state of a EfsFileSystemPolicy
type EfsFileSystemPolicyObservation struct{}