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

// PinpointBaiduChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PinpointBaiduChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PinpointBaiduChannelSpec   `json:"spec"`
	Status PinpointBaiduChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PinpointBaiduChannel contains a list of PinpointBaiduChannelList
type PinpointBaiduChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PinpointBaiduChannel `json:"items"`
}

// A PinpointBaiduChannelSpec defines the desired state of a PinpointBaiduChannel
type PinpointBaiduChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PinpointBaiduChannelParameters `json:",inline"`
}

// A PinpointBaiduChannelParameters defines the desired state of a PinpointBaiduChannel
type PinpointBaiduChannelParameters struct {
	ApiKey        string `json:"api_key"`
	ApplicationId string `json:"application_id"`
	Enabled       bool   `json:"enabled"`
	SecretKey     string `json:"secret_key"`
}

// A PinpointBaiduChannelStatus defines the observed state of a PinpointBaiduChannel
type PinpointBaiduChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PinpointBaiduChannelObservation `json:",inline"`
}

// A PinpointBaiduChannelObservation records the observed state of a PinpointBaiduChannel
type PinpointBaiduChannelObservation struct {
	Id string `json:"id"`
}