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

// DxGatewayAssociationProposal is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxGatewayAssociationProposalSpec   `json:"spec"`
	Status DxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxGatewayAssociationProposal contains a list of DxGatewayAssociationProposalList
type DxGatewayAssociationProposalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxGatewayAssociationProposal `json:"items"`
}

// A DxGatewayAssociationProposalSpec defines the desired state of a DxGatewayAssociationProposal
type DxGatewayAssociationProposalSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxGatewayAssociationProposalParameters `json:",inline"`
}

// A DxGatewayAssociationProposalParameters defines the desired state of a DxGatewayAssociationProposal
type DxGatewayAssociationProposalParameters struct {
	DxGatewayId             string   `json:"dx_gateway_id"`
	DxGatewayOwnerAccountId string   `json:"dx_gateway_owner_account_id"`
	Id                      string   `json:"id"`
	AllowedPrefixes         []string `json:"allowed_prefixes"`
	AssociatedGatewayId     string   `json:"associated_gateway_id"`
}

// A DxGatewayAssociationProposalStatus defines the observed state of a DxGatewayAssociationProposal
type DxGatewayAssociationProposalStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxGatewayAssociationProposalObservation `json:",inline"`
}

// A DxGatewayAssociationProposalObservation records the observed state of a DxGatewayAssociationProposal
type DxGatewayAssociationProposalObservation struct {
	AssociatedGatewayOwnerAccountId string `json:"associated_gateway_owner_account_id"`
	AssociatedGatewayType           string `json:"associated_gateway_type"`
}