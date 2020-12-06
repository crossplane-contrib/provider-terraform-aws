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

// DxBgpPeer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxBgpPeer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxBgpPeerSpec   `json:"spec"`
	Status DxBgpPeerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxBgpPeer contains a list of DxBgpPeerList
type DxBgpPeerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxBgpPeer `json:"items"`
}

// A DxBgpPeerSpec defines the desired state of a DxBgpPeer
type DxBgpPeerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxBgpPeerParameters `json:",inline"`
}

// A DxBgpPeerParameters defines the desired state of a DxBgpPeer
type DxBgpPeerParameters struct {
	Id                 string   `json:"id"`
	AmazonAddress      string   `json:"amazon_address"`
	BgpAsn             int64    `json:"bgp_asn"`
	BgpAuthKey         string   `json:"bgp_auth_key"`
	CustomerAddress    string   `json:"customer_address"`
	AddressFamily      string   `json:"address_family"`
	VirtualInterfaceId string   `json:"virtual_interface_id"`
	Timeouts           Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A DxBgpPeerStatus defines the observed state of a DxBgpPeer
type DxBgpPeerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxBgpPeerObservation `json:",inline"`
}

// A DxBgpPeerObservation records the observed state of a DxBgpPeer
type DxBgpPeerObservation struct {
	AwsDevice string `json:"aws_device"`
	BgpPeerId string `json:"bgp_peer_id"`
	BgpStatus string `json:"bgp_status"`
}