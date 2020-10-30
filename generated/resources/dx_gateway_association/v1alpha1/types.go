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

// DxGatewayAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxGatewayAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxGatewayAssociationSpec   `json:"spec"`
	Status DxGatewayAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxGatewayAssociation contains a list of DxGatewayAssociationList
type DxGatewayAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxGatewayAssociation `json:"items"`
}

// A DxGatewayAssociationSpec defines the desired state of a DxGatewayAssociation
type DxGatewayAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxGatewayAssociationParameters `json:",inline"`
}

// A DxGatewayAssociationParameters defines the desired state of a DxGatewayAssociation
type DxGatewayAssociationParameters struct {
	ProposalId  string `json:"proposal_id"`
	DxGatewayId string `json:"dx_gateway_id"`
}

// A DxGatewayAssociationStatus defines the observed state of a DxGatewayAssociation
type DxGatewayAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxGatewayAssociationObservation `json:",inline"`
}

// A DxGatewayAssociationObservation records the observed state of a DxGatewayAssociation
type DxGatewayAssociationObservation struct {
	AllowedPrefixes                 []string `json:"allowed_prefixes"`
	AssociatedGatewayType           string   `json:"associated_gateway_type"`
	DxGatewayAssociationId          string   `json:"dx_gateway_association_id"`
	DxGatewayOwnerAccountId         string   `json:"dx_gateway_owner_account_id"`
	Id                              string   `json:"id"`
	AssociatedGatewayId             string   `json:"associated_gateway_id"`
	AssociatedGatewayOwnerAccountId string   `json:"associated_gateway_owner_account_id"`
}