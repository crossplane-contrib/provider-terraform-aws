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

// Ec2TransitGatewayRoute is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TransitGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TransitGatewayRouteSpec   `json:"spec"`
	Status Ec2TransitGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRoute contains a list of Ec2TransitGatewayRouteList
type Ec2TransitGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayRoute `json:"items"`
}

// A Ec2TransitGatewayRouteSpec defines the desired state of a Ec2TransitGatewayRoute
type Ec2TransitGatewayRouteSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TransitGatewayRouteParameters `json:",inline"`
}

// A Ec2TransitGatewayRouteParameters defines the desired state of a Ec2TransitGatewayRoute
type Ec2TransitGatewayRouteParameters struct {
	Id                         string `json:"id"`
	TransitGatewayAttachmentId string `json:"transit_gateway_attachment_id"`
	TransitGatewayRouteTableId string `json:"transit_gateway_route_table_id"`
	Blackhole                  bool   `json:"blackhole"`
	DestinationCidrBlock       string `json:"destination_cidr_block"`
}

// A Ec2TransitGatewayRouteStatus defines the observed state of a Ec2TransitGatewayRoute
type Ec2TransitGatewayRouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TransitGatewayRouteObservation `json:",inline"`
}

// A Ec2TransitGatewayRouteObservation records the observed state of a Ec2TransitGatewayRoute
type Ec2TransitGatewayRouteObservation struct{}