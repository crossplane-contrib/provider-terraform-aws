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

// KinesisFirehoseDeliveryStream is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KinesisFirehoseDeliveryStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KinesisFirehoseDeliveryStreamSpec   `json:"spec"`
	Status KinesisFirehoseDeliveryStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisFirehoseDeliveryStream contains a list of KinesisFirehoseDeliveryStreamList
type KinesisFirehoseDeliveryStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisFirehoseDeliveryStream `json:"items"`
}

// A KinesisFirehoseDeliveryStreamSpec defines the desired state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KinesisFirehoseDeliveryStreamParameters `json:",inline"`
}

// A KinesisFirehoseDeliveryStreamParameters defines the desired state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamParameters struct {
	Name        string `json:"name"`
	Destination string `json:"destination"`
}

// A KinesisFirehoseDeliveryStreamStatus defines the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisFirehoseDeliveryStreamObservation `json:",inline"`
}

// A KinesisFirehoseDeliveryStreamObservation records the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamObservation struct {
	VersionId     string `json:"version_id"`
	Arn           string `json:"arn"`
	DestinationId string `json:"destination_id"`
	Id            string `json:"id"`
}