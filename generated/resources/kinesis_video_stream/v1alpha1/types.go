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

// KinesisVideoStream is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KinesisVideoStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KinesisVideoStreamSpec   `json:"spec"`
	Status KinesisVideoStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisVideoStream contains a list of KinesisVideoStreamList
type KinesisVideoStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisVideoStream `json:"items"`
}

// A KinesisVideoStreamSpec defines the desired state of a KinesisVideoStream
type KinesisVideoStreamSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KinesisVideoStreamParameters `json:",inline"`
}

// A KinesisVideoStreamParameters defines the desired state of a KinesisVideoStream
type KinesisVideoStreamParameters struct {
	DataRetentionInHours int               `json:"data_retention_in_hours"`
	MediaType            string            `json:"media_type"`
	Name                 string            `json:"name"`
	Tags                 map[string]string `json:"tags"`
	DeviceName           string            `json:"device_name"`
	Id                   string            `json:"id"`
	KmsKeyId             string            `json:"kms_key_id"`
	Timeouts             []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A KinesisVideoStreamStatus defines the observed state of a KinesisVideoStream
type KinesisVideoStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisVideoStreamObservation `json:",inline"`
}

// A KinesisVideoStreamObservation records the observed state of a KinesisVideoStream
type KinesisVideoStreamObservation struct {
	Arn          string `json:"arn"`
	CreationTime string `json:"creation_time"`
	Version      string `json:"version"`
}