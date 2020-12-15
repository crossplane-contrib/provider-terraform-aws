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

// Ec2ClientVpnNetworkAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2ClientVpnNetworkAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2ClientVpnNetworkAssociationSpec   `json:"spec"`
	Status Ec2ClientVpnNetworkAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnNetworkAssociation contains a list of Ec2ClientVpnNetworkAssociationList
type Ec2ClientVpnNetworkAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ClientVpnNetworkAssociation `json:"items"`
}

// A Ec2ClientVpnNetworkAssociationSpec defines the desired state of a Ec2ClientVpnNetworkAssociation
type Ec2ClientVpnNetworkAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2ClientVpnNetworkAssociationParameters `json:"forProvider"`
}

// A Ec2ClientVpnNetworkAssociationParameters defines the desired state of a Ec2ClientVpnNetworkAssociation
type Ec2ClientVpnNetworkAssociationParameters struct {
	SecurityGroups      []string `json:"security_groups"`
	SubnetId            string   `json:"subnet_id"`
	ClientVpnEndpointId string   `json:"client_vpn_endpoint_id"`
}

// A Ec2ClientVpnNetworkAssociationStatus defines the observed state of a Ec2ClientVpnNetworkAssociation
type Ec2ClientVpnNetworkAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2ClientVpnNetworkAssociationObservation `json:"atProvider"`
}

// A Ec2ClientVpnNetworkAssociationObservation records the observed state of a Ec2ClientVpnNetworkAssociation
type Ec2ClientVpnNetworkAssociationObservation struct {
	Status        string `json:"status"`
	VpcId         string `json:"vpc_id"`
	AssociationId string `json:"association_id"`
}