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

// DxHostedPublicVirtualInterfaceAccepter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DxHostedPublicVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DxHostedPublicVirtualInterfaceAccepterSpec   `json:"spec"`
	Status DxHostedPublicVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxHostedPublicVirtualInterfaceAccepter contains a list of DxHostedPublicVirtualInterfaceAccepterList
type DxHostedPublicVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxHostedPublicVirtualInterfaceAccepter `json:"items"`
}

// A DxHostedPublicVirtualInterfaceAccepterSpec defines the desired state of a DxHostedPublicVirtualInterfaceAccepter
type DxHostedPublicVirtualInterfaceAccepterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DxHostedPublicVirtualInterfaceAccepterParameters `json:",inline"`
}

// A DxHostedPublicVirtualInterfaceAccepterParameters defines the desired state of a DxHostedPublicVirtualInterfaceAccepter
type DxHostedPublicVirtualInterfaceAccepterParameters struct {
	Id                 string            `json:"id"`
	Tags               map[string]string `json:"tags"`
	VirtualInterfaceId string            `json:"virtual_interface_id"`
	Timeouts           Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Create string `json:"create"`
}

// A DxHostedPublicVirtualInterfaceAccepterStatus defines the observed state of a DxHostedPublicVirtualInterfaceAccepter
type DxHostedPublicVirtualInterfaceAccepterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DxHostedPublicVirtualInterfaceAccepterObservation `json:",inline"`
}

// A DxHostedPublicVirtualInterfaceAccepterObservation records the observed state of a DxHostedPublicVirtualInterfaceAccepter
type DxHostedPublicVirtualInterfaceAccepterObservation struct {
	Arn string `json:"arn"`
}