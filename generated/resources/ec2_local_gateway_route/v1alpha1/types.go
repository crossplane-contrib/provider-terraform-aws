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

// Ec2LocalGatewayRoute is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2LocalGatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2LocalGatewayRouteSpec   `json:"spec"`
	Status Ec2LocalGatewayRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2LocalGatewayRoute contains a list of Ec2LocalGatewayRouteList
type Ec2LocalGatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2LocalGatewayRoute `json:"items"`
}

// A Ec2LocalGatewayRouteSpec defines the desired state of a Ec2LocalGatewayRoute
type Ec2LocalGatewayRouteSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2LocalGatewayRouteParameters `json:",inline"`
}

// A Ec2LocalGatewayRouteParameters defines the desired state of a Ec2LocalGatewayRoute
type Ec2LocalGatewayRouteParameters struct {
	DestinationCidrBlock                string `json:"destination_cidr_block"`
	Id                                  string `json:"id"`
	LocalGatewayRouteTableId            string `json:"local_gateway_route_table_id"`
	LocalGatewayVirtualInterfaceGroupId string `json:"local_gateway_virtual_interface_group_id"`
}

// A Ec2LocalGatewayRouteStatus defines the observed state of a Ec2LocalGatewayRoute
type Ec2LocalGatewayRouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2LocalGatewayRouteObservation `json:",inline"`
}

// A Ec2LocalGatewayRouteObservation records the observed state of a Ec2LocalGatewayRoute
type Ec2LocalGatewayRouteObservation struct{}