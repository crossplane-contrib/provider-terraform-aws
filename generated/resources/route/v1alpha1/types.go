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

// Route is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteSpec   `json:"spec"`
	Status RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route contains a list of RouteList
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// A RouteSpec defines the desired state of a Route
type RouteSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RouteParameters `json:",inline"`
}

// A RouteParameters defines the desired state of a Route
type RouteParameters struct {
	TransitGatewayId         string `json:"transit_gateway_id"`
	VpcPeeringConnectionId   string `json:"vpc_peering_connection_id"`
	RouteTableId             string `json:"route_table_id"`
	DestinationIpv6CidrBlock string `json:"destination_ipv6_cidr_block"`
	DestinationCidrBlock     string `json:"destination_cidr_block"`
}

// A RouteStatus defines the observed state of a Route
type RouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RouteObservation `json:",inline"`
}

// A RouteObservation records the observed state of a Route
type RouteObservation struct {
	GatewayId               string `json:"gateway_id"`
	InstanceId              string `json:"instance_id"`
	EgressOnlyGatewayId     string `json:"egress_only_gateway_id"`
	InstanceOwnerId         string `json:"instance_owner_id"`
	NatGatewayId            string `json:"nat_gateway_id"`
	LocalGatewayId          string `json:"local_gateway_id"`
	NetworkInterfaceId      string `json:"network_interface_id"`
	State                   string `json:"state"`
	DestinationPrefixListId string `json:"destination_prefix_list_id"`
	Id                      string `json:"id"`
	Origin                  string `json:"origin"`
}