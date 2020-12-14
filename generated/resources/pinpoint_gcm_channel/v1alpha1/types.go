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

// PinpointGcmChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointGcmChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointGcmChannelSpec   `json:"spec"`
	Status PinpointGcmChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointGcmChannel contains a list of PinpointGcmChannelList
type PinpointGcmChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointGcmChannel `json:"items"`
}

// A PinpointGcmChannelSpec defines the desired state of a PinpointGcmChannel
type PinpointGcmChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointGcmChannelParameters `json:"forProvider"`
}

// A PinpointGcmChannelParameters defines the desired state of a PinpointGcmChannel
type PinpointGcmChannelParameters struct {
	ApiKey        string `json:"api_key"`
	ApplicationId string `json:"application_id"`
	Enabled       bool   `json:"enabled"`
	Id            string `json:"id"`
}

// A PinpointGcmChannelStatus defines the observed state of a PinpointGcmChannel
type PinpointGcmChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointGcmChannelObservation `json:"atProvider"`
}

// A PinpointGcmChannelObservation records the observed state of a PinpointGcmChannel
type PinpointGcmChannelObservation struct{}