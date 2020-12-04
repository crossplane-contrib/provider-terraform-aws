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

// Ec2TransitGatewayVpcAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGatewayVpcAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewayVpcAttachmentSpec   `json:"spec"`
	Status Ec2TransitGatewayVpcAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayVpcAttachment contains a list of Ec2TransitGatewayVpcAttachmentList
type Ec2TransitGatewayVpcAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayVpcAttachment `json:"items"`
}

// A Ec2TransitGatewayVpcAttachmentSpec defines the desired state of a Ec2TransitGatewayVpcAttachment
type Ec2TransitGatewayVpcAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayVpcAttachmentParameters `json:",inline"`
}

// A Ec2TransitGatewayVpcAttachmentParameters defines the desired state of a Ec2TransitGatewayVpcAttachment
type Ec2TransitGatewayVpcAttachmentParameters struct {
	TransitGatewayDefaultRouteTableAssociation bool              `json:"transit_gateway_default_route_table_association"`
	TransitGatewayDefaultRouteTablePropagation bool              `json:"transit_gateway_default_route_table_propagation"`
	Id                                         string            `json:"id"`
	Tags                                       map[string]string `json:"tags"`
	SubnetIds                                  []string          `json:"subnet_ids"`
	TransitGatewayId                           string            `json:"transit_gateway_id"`
	VpcId                                      string            `json:"vpc_id"`
	DnsSupport                                 string            `json:"dns_support"`
	Ipv6Support                                string            `json:"ipv6_support"`
}

// A Ec2TransitGatewayVpcAttachmentStatus defines the observed state of a Ec2TransitGatewayVpcAttachment
type Ec2TransitGatewayVpcAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayVpcAttachmentObservation `json:",inline"`
}

// A Ec2TransitGatewayVpcAttachmentObservation records the observed state of a Ec2TransitGatewayVpcAttachment
type Ec2TransitGatewayVpcAttachmentObservation struct {
	VpcOwnerId string `json:"vpc_owner_id"`
}