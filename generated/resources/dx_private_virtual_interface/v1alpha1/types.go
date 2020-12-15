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

// DxPrivateVirtualInterface is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxPrivateVirtualInterfaceSpec   `json:"spec"`
	Status DxPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxPrivateVirtualInterface contains a list of DxPrivateVirtualInterfaceList
type DxPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxPrivateVirtualInterface `json:"items"`
}

// A DxPrivateVirtualInterfaceSpec defines the desired state of a DxPrivateVirtualInterface
type DxPrivateVirtualInterfaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxPrivateVirtualInterfaceParameters `json:"forProvider"`
}

// A DxPrivateVirtualInterfaceParameters defines the desired state of a DxPrivateVirtualInterface
type DxPrivateVirtualInterfaceParameters struct {
	Tags            map[string]string `json:"tags"`
	VpnGatewayId    string            `json:"vpn_gateway_id"`
	AddressFamily   string            `json:"address_family"`
	AmazonAddress   string            `json:"amazon_address"`
	ConnectionId    string            `json:"connection_id"`
	DxGatewayId     string            `json:"dx_gateway_id"`
	Mtu             int64             `json:"mtu"`
	Vlan            int64             `json:"vlan"`
	BgpAsn          int64             `json:"bgp_asn"`
	BgpAuthKey      string            `json:"bgp_auth_key"`
	CustomerAddress string            `json:"customer_address"`
	Name            string            `json:"name"`
	Timeouts        Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A DxPrivateVirtualInterfaceStatus defines the observed state of a DxPrivateVirtualInterface
type DxPrivateVirtualInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxPrivateVirtualInterfaceObservation `json:"atProvider"`
}

// A DxPrivateVirtualInterfaceObservation records the observed state of a DxPrivateVirtualInterface
type DxPrivateVirtualInterfaceObservation struct {
	AmazonSideAsn     string `json:"amazon_side_asn"`
	Arn               string `json:"arn"`
	AwsDevice         string `json:"aws_device"`
	JumboFrameCapable bool   `json:"jumbo_frame_capable"`
}