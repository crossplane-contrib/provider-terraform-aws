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

// PinpointApnsChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointApnsChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointApnsChannelSpec   `json:"spec"`
	Status PinpointApnsChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointApnsChannel contains a list of PinpointApnsChannelList
type PinpointApnsChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointApnsChannel `json:"items"`
}

// A PinpointApnsChannelSpec defines the desired state of a PinpointApnsChannel
type PinpointApnsChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointApnsChannelParameters `json:",inline"`
}

// A PinpointApnsChannelParameters defines the desired state of a PinpointApnsChannel
type PinpointApnsChannelParameters struct {
	ApplicationId               string `json:"application_id"`
	Enabled                     bool   `json:"enabled"`
	TokenKey                    string `json:"token_key"`
	BundleId                    string `json:"bundle_id"`
	Certificate                 string `json:"certificate"`
	DefaultAuthenticationMethod string `json:"default_authentication_method"`
	PrivateKey                  string `json:"private_key"`
	TeamId                      string `json:"team_id"`
	TokenKeyId                  string `json:"token_key_id"`
}

// A PinpointApnsChannelStatus defines the observed state of a PinpointApnsChannel
type PinpointApnsChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointApnsChannelObservation `json:",inline"`
}

// A PinpointApnsChannelObservation records the observed state of a PinpointApnsChannel
type PinpointApnsChannelObservation struct {
	Id string `json:"id"`
}