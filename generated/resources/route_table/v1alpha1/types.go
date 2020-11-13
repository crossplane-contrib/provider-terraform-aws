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

// RouteTable is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteTableSpec   `json:"spec"`
	Status RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTable contains a list of RouteTableList
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// A RouteTableSpec defines the desired state of a RouteTable
type RouteTableSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RouteTableParameters `json:",inline"`
}

// A RouteTableParameters defines the desired state of a RouteTable
type RouteTableParameters struct {
	Tags            map[string]string `json:"tags"`
	VpcId           string            `json:"vpc_id"`
	Id              string            `json:"id"`
	PropagatingVgws []string          `json:"propagating_vgws"`
	Route           []Route           `json:"route"`
}

type Route struct {
	TransitGatewayId       string `json:"transit_gateway_id"`
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
	CidrBlock              string `json:"cidr_block"`
	GatewayId              string `json:"gateway_id"`
	LocalGatewayId         string `json:"local_gateway_id"`
	NetworkInterfaceId     string `json:"network_interface_id"`
	NatGatewayId           string `json:"nat_gateway_id"`
	EgressOnlyGatewayId    string `json:"egress_only_gateway_id"`
	InstanceId             string `json:"instance_id"`
	Ipv6CidrBlock          string `json:"ipv6_cidr_block"`
}

// A RouteTableStatus defines the observed state of a RouteTable
type RouteTableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RouteTableObservation `json:",inline"`
}

// A RouteTableObservation records the observed state of a RouteTable
type RouteTableObservation struct {
	OwnerId string `json:"owner_id"`
}