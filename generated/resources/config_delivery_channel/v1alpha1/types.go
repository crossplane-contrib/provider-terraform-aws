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

// ConfigDeliveryChannel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigDeliveryChannel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigDeliveryChannelSpec   `json:"spec"`
	Status ConfigDeliveryChannelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigDeliveryChannel contains a list of ConfigDeliveryChannelList
type ConfigDeliveryChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigDeliveryChannel `json:"items"`
}

// A ConfigDeliveryChannelSpec defines the desired state of a ConfigDeliveryChannel
type ConfigDeliveryChannelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigDeliveryChannelParameters `json:",inline"`
}

// A ConfigDeliveryChannelParameters defines the desired state of a ConfigDeliveryChannel
type ConfigDeliveryChannelParameters struct {
	SnsTopicArn                string                     `json:"sns_topic_arn"`
	Id                         string                     `json:"id"`
	Name                       string                     `json:"name"`
	S3BucketName               string                     `json:"s3_bucket_name"`
	S3KeyPrefix                string                     `json:"s3_key_prefix"`
	SnapshotDeliveryProperties SnapshotDeliveryProperties `json:"snapshot_delivery_properties"`
}

type SnapshotDeliveryProperties struct {
	DeliveryFrequency string `json:"delivery_frequency"`
}

// A ConfigDeliveryChannelStatus defines the observed state of a ConfigDeliveryChannel
type ConfigDeliveryChannelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigDeliveryChannelObservation `json:",inline"`
}

// A ConfigDeliveryChannelObservation records the observed state of a ConfigDeliveryChannel
type ConfigDeliveryChannelObservation struct{}