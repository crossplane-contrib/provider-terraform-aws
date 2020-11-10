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
	ForProvider                  DxTransitVirtualInterfaceParameters `json:",inline"`
}

// A DxTransitVirtualInterfaceParameters defines the desired state of a DxTransitVirtualInterface
type DxTransitVirtualInterfaceParameters struct {
	AddressFamily string            `json:"address_family"`
	Mtu           int               `json:"mtu"`
	Tags          map[string]string `json:"tags"`
	BgpAsn        int               `json:"bgp_asn"`
	ConnectionId  string            `json:"connection_id"`
	DxGatewayId   string            `json:"dx_gateway_id"`
	Name          string            `json:"name"`
	Vlan          int               `json:"vlan"`
	Timeouts      []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Update string `json:"update"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A DxTransitVirtualInterfaceStatus defines the observed state of a DxTransitVirtualInterface
type DxTransitVirtualInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxTransitVirtualInterfaceObservation `json:",inline"`
}

// A DxTransitVirtualInterfaceObservation records the observed state of a DxTransitVirtualInterface
type DxTransitVirtualInterfaceObservation struct {
	JumboFrameCapable bool   `json:"jumbo_frame_capable"`
	BgpAuthKey        string `json:"bgp_auth_key"`
	Id                string `json:"id"`
	AwsDevice         string `json:"aws_device"`
	CustomerAddress   string `json:"customer_address"`
	AmazonAddress     string `json:"amazon_address"`
	AmazonSideAsn     string `json:"amazon_side_asn"`
	Arn               string `json:"arn"`
}