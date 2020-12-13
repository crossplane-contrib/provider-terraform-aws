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

// SsmPatchGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmPatchGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmPatchGroupSpec   `json:"spec"`
	Status SsmPatchGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmPatchGroup contains a list of SsmPatchGroupList
type SsmPatchGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmPatchGroup `json:"items"`
}

// A SsmPatchGroupSpec defines the desired state of a SsmPatchGroup
type SsmPatchGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmPatchGroupParameters `json:"forProvider"`
}

// A SsmPatchGroupParameters defines the desired state of a SsmPatchGroup
type SsmPatchGroupParameters struct {
	Id         string `json:"id"`
	PatchGroup string `json:"patch_group"`
	BaselineId string `json:"baseline_id"`
}

// A SsmPatchGroupStatus defines the observed state of a SsmPatchGroup
type SsmPatchGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmPatchGroupObservation `json:"atProvider"`
}

// A SsmPatchGroupObservation records the observed state of a SsmPatchGroup
type SsmPatchGroupObservation struct{}