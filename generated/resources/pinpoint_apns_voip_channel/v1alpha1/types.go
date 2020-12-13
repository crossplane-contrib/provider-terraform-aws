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

// PinpointApnsVoipChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointApnsVoipChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointApnsVoipChannelSpec   `json:"spec"`
	Status PinpointApnsVoipChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsVoipChannel contains a list of PinpointApnsVoipChannelList
type PinpointApnsVoipChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApnsVoipChannel `json:"items"`
}

// A PinpointApnsVoipChannelSpec defines the desired state of a PinpointApnsVoipChannel
type PinpointApnsVoipChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointApnsVoipChannelParameters `json:"forProvider"`
}

// A PinpointApnsVoipChannelParameters defines the desired state of a PinpointApnsVoipChannel
type PinpointApnsVoipChannelParameters struct {
	Id                          string `json:"id"`
	TeamId                      string `json:"team_id"`
	TokenKey                    string `json:"token_key"`
	TokenKeyId                  string `json:"token_key_id"`
	ApplicationId               string `json:"application_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	Enabled                     bool   `json:"enabled"`
	PrivateKey                  string `json:"private_key"`
	BundleId                    string `json:"bundle_id"`
}

// A PinpointApnsVoipChannelStatus defines the observed state of a PinpointApnsVoipChannel
type PinpointApnsVoipChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointApnsVoipChannelObservation `json:"atProvider"`
}

// A PinpointApnsVoipChannelObservation records the observed state of a PinpointApnsVoipChannel
type PinpointApnsVoipChannelObservation struct{}