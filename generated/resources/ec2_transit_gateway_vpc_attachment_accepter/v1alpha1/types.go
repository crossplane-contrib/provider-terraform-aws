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

// Ec2TransitGatewayVpcAttachmentAccepter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGatewayVpcAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewayVpcAttachmentAccepterSpec   `json:"spec"`
	Status Ec2TransitGatewayVpcAttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayVpcAttachmentAccepter contains a list of Ec2TransitGatewayVpcAttachmentAccepterList
type Ec2TransitGatewayVpcAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayVpcAttachmentAccepter `json:"items"`
}

// A Ec2TransitGatewayVpcAttachmentAccepterSpec defines the desired state of a Ec2TransitGatewayVpcAttachmentAccepter
type Ec2TransitGatewayVpcAttachmentAccepterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayVpcAttachmentAccepterParameters `json:",inline"`
}

// A Ec2TransitGatewayVpcAttachmentAccepterParameters defines the desired state of a Ec2TransitGatewayVpcAttachmentAccepter
type Ec2TransitGatewayVpcAttachmentAccepterParameters struct {
	TransitGatewayDefaultRouteTableAssociation bool   `json:"transit_gateway_default_route_table_association"`
	TransitGatewayDefaultRouteTablePropagation bool   `json:"transit_gateway_default_route_table_propagation"`
	TransitGatewayAttachmentId                 string `json:"transit_gateway_attachment_id"`
}

// A Ec2TransitGatewayVpcAttachmentAccepterStatus defines the observed state of a Ec2TransitGatewayVpcAttachmentAccepter
type Ec2TransitGatewayVpcAttachmentAccepterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayVpcAttachmentAccepterObservation `json:",inline"`
}

// A Ec2TransitGatewayVpcAttachmentAccepterObservation records the observed state of a Ec2TransitGatewayVpcAttachmentAccepter
type Ec2TransitGatewayVpcAttachmentAccepterObservation struct {
	VpcId            string   `json:"vpc_id"`
	VpcOwnerId       string   `json:"vpc_owner_id"`
	Id               string   `json:"id"`
	Ipv6Support      string   `json:"ipv6_support"`
	TransitGatewayId string   `json:"transit_gateway_id"`
	DnsSupport       string   `json:"dns_support"`
	SubnetIds        []string `json:"subnet_ids"`
}