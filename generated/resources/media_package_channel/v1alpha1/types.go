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

// MediaPackageChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MediaPackageChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MediaPackageChannelSpec   `json:"spec"`
	Status MediaPackageChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MediaPackageChannel contains a list of MediaPackageChannelList
type MediaPackageChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MediaPackageChannel `json:"items"`
}

// A MediaPackageChannelSpec defines the desired state of a MediaPackageChannel
type MediaPackageChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MediaPackageChannelParameters `json:",inline"`
}

// A MediaPackageChannelParameters defines the desired state of a MediaPackageChannel
type MediaPackageChannelParameters struct {
	Description string            `json:"description"`
	Id          string            `json:"id"`
	Tags        map[string]string `json:"tags"`
	ChannelId   string            `json:"channel_id"`
}

// A MediaPackageChannelStatus defines the observed state of a MediaPackageChannel
type MediaPackageChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MediaPackageChannelObservation `json:",inline"`
}

// A MediaPackageChannelObservation records the observed state of a MediaPackageChannel
type MediaPackageChannelObservation struct {
	HlsIngest []HlsIngest `json:"hls_ingest"`
	Arn       string      `json:"arn"`
}

type HlsIngest struct {
	IngestEndpoints []IngestEndpoints `json:"ingest_endpoints"`
}

type IngestEndpoints struct {
	Password string `json:"password"`
	Url      string `json:"url"`
	Username string `json:"username"`
}