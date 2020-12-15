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

// DxHostedPrivateVirtualInterface is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxHostedPrivateVirtualInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxHostedPrivateVirtualInterfaceSpec   `json:"spec"`
	Status DxHostedPrivateVirtualInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPrivateVirtualInterface contains a list of DxHostedPrivateVirtualInterfaceList
type DxHostedPrivateVirtualInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedPrivateVirtualInterface `json:"items"`
}

// A DxHostedPrivateVirtualInterfaceSpec defines the desired state of a DxHostedPrivateVirtualInterface
type DxHostedPrivateVirtualInterfaceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxHostedPrivateVirtualInterfaceParameters `json:"forProvider"`
}

// A DxHostedPrivateVirtualInterfaceParameters defines the desired state of a DxHostedPrivateVirtualInterface
type DxHostedPrivateVirtualInterfaceParameters struct {
	ConnectionId    string   `json:"connection_id"`
	CustomerAddress string   `json:"customer_address"`
	OwnerAccountId  string   `json:"owner_account_id"`
	AddressFamily   string   `json:"address_family"`
	BgpAuthKey      string   `json:"bgp_auth_key"`
	Mtu             int64    `json:"mtu"`
	Name            string   `json:"name"`
	BgpAsn          int64    `json:"bgp_asn"`
	Vlan            int64    `json:"vlan"`
	AmazonAddress   string   `json:"amazon_address"`
	Timeouts        Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Update string `json:"update"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A DxHostedPrivateVirtualInterfaceStatus defines the observed state of a DxHostedPrivateVirtualInterface
type DxHostedPrivateVirtualInterfaceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxHostedPrivateVirtualInterfaceObservation `json:"atProvider"`
}

// A DxHostedPrivateVirtualInterfaceObservation records the observed state of a DxHostedPrivateVirtualInterface
type DxHostedPrivateVirtualInterfaceObservation struct {
	Arn               string `json:"arn"`
	JumboFrameCapable bool   `json:"jumbo_frame_capable"`
	AmazonSideAsn     string `json:"amazon_side_asn"`
	AwsDevice         string `json:"aws_device"`
}