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

// PinpointSmsChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointSmsChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointSmsChannelSpec   `json:"spec"`
	Status PinpointSmsChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointSmsChannel contains a list of PinpointSmsChannelList
type PinpointSmsChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointSmsChannel `json:"items"`
}

// A PinpointSmsChannelSpec defines the desired state of a PinpointSmsChannel
type PinpointSmsChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointSmsChannelParameters `json:"forProvider"`
}

// A PinpointSmsChannelParameters defines the desired state of a PinpointSmsChannel
type PinpointSmsChannelParameters struct {
	SenderId      string `json:"sender_id"`
	ShortCode     string `json:"short_code"`
	ApplicationId string `json:"application_id"`
	Enabled       bool   `json:"enabled"`
	Id            string `json:"id"`
}

// A PinpointSmsChannelStatus defines the observed state of a PinpointSmsChannel
type PinpointSmsChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointSmsChannelObservation `json:"atProvider"`
}

// A PinpointSmsChannelObservation records the observed state of a PinpointSmsChannel
type PinpointSmsChannelObservation struct {
	PromotionalMessagesPerSecond   int64 `json:"promotional_messages_per_second"`
	TransactionalMessagesPerSecond int64 `json:"transactional_messages_per_second"`
}