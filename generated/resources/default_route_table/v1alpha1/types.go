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

// DefaultRouteTable is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DefaultRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultRouteTableSpec   `json:"spec"`
	Status DefaultRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultRouteTable contains a list of DefaultRouteTableList
type DefaultRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultRouteTable `json:"items"`
}

// A DefaultRouteTableSpec defines the desired state of a DefaultRouteTable
type DefaultRouteTableSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DefaultRouteTableParameters `json:",inline"`
}

// A DefaultRouteTableParameters defines the desired state of a DefaultRouteTable
type DefaultRouteTableParameters struct {
	DefaultRouteTableId string            `json:"default_route_table_id"`
	Id                  string            `json:"id"`
	PropagatingVgws     []string          `json:"propagating_vgws"`
	Route               []Route           `json:"route"`
	Tags                map[string]string `json:"tags"`
}

type Route struct {
	CidrBlock              string `json:"cidr_block"`
	GatewayId              string `json:"gateway_id"`
	NetworkInterfaceId     string `json:"network_interface_id"`
	VpcPeeringConnectionId string `json:"vpc_peering_connection_id"`
	EgressOnlyGatewayId    string `json:"egress_only_gateway_id"`
	InstanceId             string `json:"instance_id"`
	Ipv6CidrBlock          string `json:"ipv6_cidr_block"`
	NatGatewayId           string `json:"nat_gateway_id"`
	TransitGatewayId       string `json:"transit_gateway_id"`
}

// A DefaultRouteTableStatus defines the observed state of a DefaultRouteTable
type DefaultRouteTableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DefaultRouteTableObservation `json:",inline"`
}

// A DefaultRouteTableObservation records the observed state of a DefaultRouteTable
type DefaultRouteTableObservation struct {
	OwnerId string `json:"owner_id"`
	VpcId   string `json:"vpc_id"`
}