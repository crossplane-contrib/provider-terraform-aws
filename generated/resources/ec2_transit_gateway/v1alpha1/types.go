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

// Ec2TransitGateway is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewaySpec   `json:"spec"`
	Status Ec2TransitGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGateway contains a list of Ec2TransitGatewayList
type Ec2TransitGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGateway `json:"items"`
}

// A Ec2TransitGatewaySpec defines the desired state of a Ec2TransitGateway
type Ec2TransitGatewaySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayParameters `json:",inline"`
}

// A Ec2TransitGatewayParameters defines the desired state of a Ec2TransitGateway
type Ec2TransitGatewayParameters struct {
	AutoAcceptSharedAttachments  string            `json:"auto_accept_shared_attachments"`
	DefaultRouteTableAssociation string            `json:"default_route_table_association"`
	DefaultRouteTablePropagation string            `json:"default_route_table_propagation"`
	DnsSupport                   string            `json:"dns_support"`
	AmazonSideAsn                int               `json:"amazon_side_asn"`
	Description                  string            `json:"description"`
	Id                           string            `json:"id"`
	Tags                         map[string]string `json:"tags"`
	VpnEcmpSupport               string            `json:"vpn_ecmp_support"`
}

// A Ec2TransitGatewayStatus defines the observed state of a Ec2TransitGateway
type Ec2TransitGatewayStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayObservation `json:",inline"`
}

// A Ec2TransitGatewayObservation records the observed state of a Ec2TransitGateway
type Ec2TransitGatewayObservation struct {
	PropagationDefaultRouteTableId string `json:"propagation_default_route_table_id"`
	Arn                            string `json:"arn"`
	AssociationDefaultRouteTableId string `json:"association_default_route_table_id"`
	OwnerId                        string `json:"owner_id"`
}