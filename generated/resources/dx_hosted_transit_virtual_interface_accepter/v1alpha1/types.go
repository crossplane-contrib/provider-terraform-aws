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

// DxHostedTransitVirtualInterfaceAccepter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxHostedTransitVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxHostedTransitVirtualInterfaceAccepterSpec   `json:"spec"`
	Status DxHostedTransitVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedTransitVirtualInterfaceAccepter contains a list of DxHostedTransitVirtualInterfaceAccepterList
type DxHostedTransitVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedTransitVirtualInterfaceAccepter `json:"items"`
}

// A DxHostedTransitVirtualInterfaceAccepterSpec defines the desired state of a DxHostedTransitVirtualInterfaceAccepter
type DxHostedTransitVirtualInterfaceAccepterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxHostedTransitVirtualInterfaceAccepterParameters `json:"forProvider"`
}

// A DxHostedTransitVirtualInterfaceAccepterParameters defines the desired state of a DxHostedTransitVirtualInterfaceAccepter
type DxHostedTransitVirtualInterfaceAccepterParameters struct {
	DxGatewayId        string            `json:"dx_gateway_id"`
	Tags               map[string]string `json:"tags"`
	VirtualInterfaceId string            `json:"virtual_interface_id"`
	Timeouts           Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Create string `json:"create"`
}

// A DxHostedTransitVirtualInterfaceAccepterStatus defines the observed state of a DxHostedTransitVirtualInterfaceAccepter
type DxHostedTransitVirtualInterfaceAccepterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxHostedTransitVirtualInterfaceAccepterObservation `json:"atProvider"`
}

// A DxHostedTransitVirtualInterfaceAccepterObservation records the observed state of a DxHostedTransitVirtualInterfaceAccepter
type DxHostedTransitVirtualInterfaceAccepterObservation struct {
	Arn string `json:"arn"`
}