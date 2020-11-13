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

// PinpointAdmChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointAdmChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointAdmChannelSpec   `json:"spec"`
	Status PinpointAdmChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointAdmChannel contains a list of PinpointAdmChannelList
type PinpointAdmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointAdmChannel `json:"items"`
}

// A PinpointAdmChannelSpec defines the desired state of a PinpointAdmChannel
type PinpointAdmChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointAdmChannelParameters `json:",inline"`
}

// A PinpointAdmChannelParameters defines the desired state of a PinpointAdmChannel
type PinpointAdmChannelParameters struct {
	ApplicationId string `json:"application_id"`
	ClientId      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	Enabled       bool   `json:"enabled"`
	Id            string `json:"id"`
}

// A PinpointAdmChannelStatus defines the observed state of a PinpointAdmChannel
type PinpointAdmChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointAdmChannelObservation `json:",inline"`
}

// A PinpointAdmChannelObservation records the observed state of a PinpointAdmChannel
type PinpointAdmChannelObservation struct{}