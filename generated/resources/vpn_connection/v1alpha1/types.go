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

// VpnConnection is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpnConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpnConnectionSpec   `json:"spec"`
	Status VpnConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpnConnection contains a list of VpnConnectionList
type VpnConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnConnection `json:"items"`
}

// A VpnConnectionSpec defines the desired state of a VpnConnection
type VpnConnectionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpnConnectionParameters `json:",inline"`
}

// A VpnConnectionParameters defines the desired state of a VpnConnection
type VpnConnectionParameters struct {
	Tunnel2PresharedKey string            `json:"tunnel2_preshared_key"`
	Tags                map[string]string `json:"tags"`
	Tunnel1PresharedKey string            `json:"tunnel1_preshared_key"`
	Id                  string            `json:"id"`
	StaticRoutesOnly    bool              `json:"static_routes_only"`
	Tunnel1InsideCidr   string            `json:"tunnel1_inside_cidr"`
	Tunnel2InsideCidr   string            `json:"tunnel2_inside_cidr"`
	Type                string            `json:"type"`
	CustomerGatewayId   string            `json:"customer_gateway_id"`
	TransitGatewayId    string            `json:"transit_gateway_id"`
	VpnGatewayId        string            `json:"vpn_gateway_id"`
}

// A VpnConnectionStatus defines the observed state of a VpnConnection
type VpnConnectionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpnConnectionObservation `json:",inline"`
}

// A VpnConnectionObservation records the observed state of a VpnConnection
type VpnConnectionObservation struct {
	Tunnel2VgwInsideAddress      string         `json:"tunnel2_vgw_inside_address"`
	CustomerGatewayConfiguration string         `json:"customer_gateway_configuration"`
	Tunnel1Address               string         `json:"tunnel1_address"`
	Tunnel1BgpAsn                string         `json:"tunnel1_bgp_asn"`
	Tunnel1BgpHoldtime           int64          `json:"tunnel1_bgp_holdtime"`
	Arn                          string         `json:"arn"`
	Tunnel1VgwInsideAddress      string         `json:"tunnel1_vgw_inside_address"`
	Routes                       []Routes       `json:"routes"`
	TransitGatewayAttachmentId   string         `json:"transit_gateway_attachment_id"`
	Tunnel2Address               string         `json:"tunnel2_address"`
	Tunnel1CgwInsideAddress      string         `json:"tunnel1_cgw_inside_address"`
	Tunnel2BgpAsn                string         `json:"tunnel2_bgp_asn"`
	Tunnel2BgpHoldtime           int64          `json:"tunnel2_bgp_holdtime"`
	Tunnel2CgwInsideAddress      string         `json:"tunnel2_cgw_inside_address"`
	VgwTelemetry                 []VgwTelemetry `json:"vgw_telemetry"`
}

type Routes struct {
	Source               string `json:"source"`
	State                string `json:"state"`
	DestinationCidrBlock string `json:"destination_cidr_block"`
}

type VgwTelemetry struct {
	AcceptedRouteCount int64  `json:"accepted_route_count"`
	LastStatusChange   string `json:"last_status_change"`
	OutsideIpAddress   string `json:"outside_ip_address"`
	Status             string `json:"status"`
	StatusMessage      string `json:"status_message"`
}