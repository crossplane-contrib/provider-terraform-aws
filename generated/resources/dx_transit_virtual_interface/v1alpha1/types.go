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

// DxTransitVirtualInterface is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxTransitVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxTransitVirtualInterfaceSpec   `json:"spec"`
	Status DxTransitVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxTransitVirtualInterface contains a list of DxTransitVirtualInterfaceList
type DxTransitVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxTransitVirtualInterface `json:"items"`
}

// A DxTransitVirtualInterfaceSpec defines the desired state of a DxTransitVirtualInterface
type DxTransitVirtualInterfaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxTransitVirtualInterfaceParameters `json:"forProvider"`
}

// A DxTransitVirtualInterfaceParameters defines the desired state of a DxTransitVirtualInterface
type DxTransitVirtualInterfaceParameters struct {
	Mtu             int64             `json:"mtu"`
	AddressFamily   string            `json:"address_family"`
	AmazonAddress   string            `json:"amazon_address"`
	BgpAsn          int64             `json:"bgp_asn"`
	Tags            map[string]string `json:"tags"`
	CustomerAddress string            `json:"customer_address"`
	DxGatewayId     string            `json:"dx_gateway_id"`
	Name            string            `json:"name"`
	BgpAuthKey      string            `json:"bgp_auth_key"`
	ConnectionId    string            `json:"connection_id"`
	Vlan            int64             `json:"vlan"`
	Timeouts        Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DxTransitVirtualInterfaceStatus defines the observed state of a DxTransitVirtualInterface
type DxTransitVirtualInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxTransitVirtualInterfaceObservation `json:"atProvider"`
}

// A DxTransitVirtualInterfaceObservation records the observed state of a DxTransitVirtualInterface
type DxTransitVirtualInterfaceObservation struct {
	Arn               string `json:"arn"`
	AmazonSideAsn     string `json:"amazon_side_asn"`
	AwsDevice         string `json:"aws_device"`
	JumboFrameCapable bool   `json:"jumbo_frame_capable"`
}