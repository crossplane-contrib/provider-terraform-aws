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

// PinpointApnsSandboxChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointApnsSandboxChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointApnsSandboxChannelSpec   `json:"spec"`
	Status PinpointApnsSandboxChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsSandboxChannel contains a list of PinpointApnsSandboxChannelList
type PinpointApnsSandboxChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApnsSandboxChannel `json:"items"`
}

// A PinpointApnsSandboxChannelSpec defines the desired state of a PinpointApnsSandboxChannel
type PinpointApnsSandboxChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointApnsSandboxChannelParameters `json:"forProvider"`
}

// A PinpointApnsSandboxChannelParameters defines the desired state of a PinpointApnsSandboxChannel
type PinpointApnsSandboxChannelParameters struct {
	PrivateKey                  string `json:"private_key"`
	TokenKey                    string `json:"token_key"`
	TokenKeyId                  string `json:"token_key_id"`
	ApplicationId               string `json:"application_id"`
	BundleId                    string `json:"bundle_id"`
	Enabled                     bool   `json:"enabled"`
	TeamId                      string `json:"team_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
}

// A PinpointApnsSandboxChannelStatus defines the observed state of a PinpointApnsSandboxChannel
type PinpointApnsSandboxChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointApnsSandboxChannelObservation `json:"atProvider"`
}

// A PinpointApnsSandboxChannelObservation records the observed state of a PinpointApnsSandboxChannel
type PinpointApnsSandboxChannelObservation struct{}