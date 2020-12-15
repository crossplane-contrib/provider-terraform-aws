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

// DxHostedTransitVirtualInterface is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxHostedTransitVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxHostedTransitVirtualInterfaceSpec   `json:"spec"`
	Status DxHostedTransitVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedTransitVirtualInterface contains a list of DxHostedTransitVirtualInterfaceList
type DxHostedTransitVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedTransitVirtualInterface `json:"items"`
}

// A DxHostedTransitVirtualInterfaceSpec defines the desired state of a DxHostedTransitVirtualInterface
type DxHostedTransitVirtualInterfaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxHostedTransitVirtualInterfaceParameters `json:"forProvider"`
}

// A DxHostedTransitVirtualInterfaceParameters defines the desired state of a DxHostedTransitVirtualInterface
type DxHostedTransitVirtualInterfaceParameters struct {
	Mtu             int64    `json:"mtu"`
	Name            string   `json:"name"`
	AddressFamily   string   `json:"address_family"`
	AmazonAddress   string   `json:"amazon_address"`
	BgpAuthKey      string   `json:"bgp_auth_key"`
	CustomerAddress string   `json:"customer_address"`
	ConnectionId    string   `json:"connection_id"`
	OwnerAccountId  string   `json:"owner_account_id"`
	BgpAsn          int64    `json:"bgp_asn"`
	Vlan            int64    `json:"vlan"`
	Timeouts        Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A DxHostedTransitVirtualInterfaceStatus defines the observed state of a DxHostedTransitVirtualInterface
type DxHostedTransitVirtualInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxHostedTransitVirtualInterfaceObservation `json:"atProvider"`
}

// A DxHostedTransitVirtualInterfaceObservation records the observed state of a DxHostedTransitVirtualInterface
type DxHostedTransitVirtualInterfaceObservation struct {
	Arn               string `json:"arn"`
	JumboFrameCapable bool   `json:"jumbo_frame_capable"`
	AmazonSideAsn     string `json:"amazon_side_asn"`
	AwsDevice         string `json:"aws_device"`
}