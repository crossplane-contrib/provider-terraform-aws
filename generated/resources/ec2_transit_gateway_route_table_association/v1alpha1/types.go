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

// Ec2TransitGatewayRouteTableAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGatewayRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewayRouteTableAssociationSpec   `json:"spec"`
	Status Ec2TransitGatewayRouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTableAssociation contains a list of Ec2TransitGatewayRouteTableAssociationList
type Ec2TransitGatewayRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayRouteTableAssociation `json:"items"`
}

// A Ec2TransitGatewayRouteTableAssociationSpec defines the desired state of a Ec2TransitGatewayRouteTableAssociation
type Ec2TransitGatewayRouteTableAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayRouteTableAssociationParameters `json:"forProvider"`
}

// A Ec2TransitGatewayRouteTableAssociationParameters defines the desired state of a Ec2TransitGatewayRouteTableAssociation
type Ec2TransitGatewayRouteTableAssociationParameters struct {
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
}

// A Ec2TransitGatewayRouteTableAssociationStatus defines the observed state of a Ec2TransitGatewayRouteTableAssociation
type Ec2TransitGatewayRouteTableAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayRouteTableAssociationObservation `json:"atProvider"`
}

// A Ec2TransitGatewayRouteTableAssociationObservation records the observed state of a Ec2TransitGatewayRouteTableAssociation
type Ec2TransitGatewayRouteTableAssociationObservation struct {
	ResourceId   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}