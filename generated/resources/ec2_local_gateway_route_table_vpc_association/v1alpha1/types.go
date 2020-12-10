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

// Ec2LocalGatewayRouteTableVpcAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2LocalGatewayRouteTableVpcAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2LocalGatewayRouteTableVpcAssociationSpec   `json:"spec"`
	Status Ec2LocalGatewayRouteTableVpcAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2LocalGatewayRouteTableVpcAssociation contains a list of Ec2LocalGatewayRouteTableVpcAssociationList
type Ec2LocalGatewayRouteTableVpcAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2LocalGatewayRouteTableVpcAssociation `json:"items"`
}

// A Ec2LocalGatewayRouteTableVpcAssociationSpec defines the desired state of a Ec2LocalGatewayRouteTableVpcAssociation
type Ec2LocalGatewayRouteTableVpcAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2LocalGatewayRouteTableVpcAssociationParameters `json:",inline"`
}

// A Ec2LocalGatewayRouteTableVpcAssociationParameters defines the desired state of a Ec2LocalGatewayRouteTableVpcAssociation
type Ec2LocalGatewayRouteTableVpcAssociationParameters struct {
	VpcId                    string            `json:"vpc_id"`
	Id                       string            `json:"id"`
	LocalGatewayRouteTableId string            `json:"local_gateway_route_table_id"`
	Tags                     map[string]string `json:"tags"`
}

// A Ec2LocalGatewayRouteTableVpcAssociationStatus defines the observed state of a Ec2LocalGatewayRouteTableVpcAssociation
type Ec2LocalGatewayRouteTableVpcAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2LocalGatewayRouteTableVpcAssociationObservation `json:",inline"`
}

// A Ec2LocalGatewayRouteTableVpcAssociationObservation records the observed state of a Ec2LocalGatewayRouteTableVpcAssociation
type Ec2LocalGatewayRouteTableVpcAssociationObservation struct {
	LocalGatewayId string `json:"local_gateway_id"`
}