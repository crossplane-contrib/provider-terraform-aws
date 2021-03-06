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

// RamResourceAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RamResourceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RamResourceAssociationSpec   `json:"spec"`
	Status RamResourceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RamResourceAssociation contains a list of RamResourceAssociationList
type RamResourceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RamResourceAssociation `json:"items"`
}

// A RamResourceAssociationSpec defines the desired state of a RamResourceAssociation
type RamResourceAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RamResourceAssociationParameters `json:"forProvider"`
}

// A RamResourceAssociationParameters defines the desired state of a RamResourceAssociation
type RamResourceAssociationParameters struct {
	ResourceArn      string `json:"resource_arn"`
	ResourceShareArn string `json:"resource_share_arn"`
}

// A RamResourceAssociationStatus defines the observed state of a RamResourceAssociation
type RamResourceAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RamResourceAssociationObservation `json:"atProvider"`
}

// A RamResourceAssociationObservation records the observed state of a RamResourceAssociation
type RamResourceAssociationObservation struct{}