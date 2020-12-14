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

// PinpointApnsVoipSandboxChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointApnsVoipSandboxChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointApnsVoipSandboxChannelSpec   `json:"spec"`
	Status PinpointApnsVoipSandboxChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsVoipSandboxChannel contains a list of PinpointApnsVoipSandboxChannelList
type PinpointApnsVoipSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApnsVoipSandboxChannel `json:"items"`
}

// A PinpointApnsVoipSandboxChannelSpec defines the desired state of a PinpointApnsVoipSandboxChannel
type PinpointApnsVoipSandboxChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointApnsVoipSandboxChannelParameters `json:"forProvider"`
}

// A PinpointApnsVoipSandboxChannelParameters defines the desired state of a PinpointApnsVoipSandboxChannel
type PinpointApnsVoipSandboxChannelParameters struct {
	BundleId                    string `json:"bundle_id"`
	PrivateKey                  string `json:"private_key"`
	TokenKeyId                  string `json:"token_key_id"`
	Id                          string `json:"id"`
	TeamId                      string `json:"team_id"`
	TokenKey                    string `json:"token_key"`
	ApplicationId               string `json:"application_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	Enabled                     bool   `json:"enabled"`
}

// A PinpointApnsVoipSandboxChannelStatus defines the observed state of a PinpointApnsVoipSandboxChannel
type PinpointApnsVoipSandboxChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointApnsVoipSandboxChannelObservation `json:"atProvider"`
}

// A PinpointApnsVoipSandboxChannelObservation records the observed state of a PinpointApnsVoipSandboxChannel
type PinpointApnsVoipSandboxChannelObservation struct{}