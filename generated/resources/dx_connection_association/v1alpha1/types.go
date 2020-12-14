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

// DxConnectionAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxConnectionAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxConnectionAssociationSpec   `json:"spec"`
	Status DxConnectionAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxConnectionAssociation contains a list of DxConnectionAssociationList
type DxConnectionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxConnectionAssociation `json:"items"`
}

// A DxConnectionAssociationSpec defines the desired state of a DxConnectionAssociation
type DxConnectionAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxConnectionAssociationParameters `json:"forProvider"`
}

// A DxConnectionAssociationParameters defines the desired state of a DxConnectionAssociation
type DxConnectionAssociationParameters struct {
	LagId        string `json:"lag_id"`
	ConnectionId string `json:"connection_id"`
	Id           string `json:"id"`
}

// A DxConnectionAssociationStatus defines the observed state of a DxConnectionAssociation
type DxConnectionAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxConnectionAssociationObservation `json:"atProvider"`
}

// A DxConnectionAssociationObservation records the observed state of a DxConnectionAssociation
type DxConnectionAssociationObservation struct{}