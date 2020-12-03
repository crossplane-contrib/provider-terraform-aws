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

// KinesisStream is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KinesisStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KinesisStreamSpec   `json:"spec"`
	Status KinesisStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisStream contains a list of KinesisStreamList
type KinesisStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisStream `json:"items"`
}

// A KinesisStreamSpec defines the desired state of a KinesisStream
type KinesisStreamSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KinesisStreamParameters `json:",inline"`
}

// A KinesisStreamParameters defines the desired state of a KinesisStream
type KinesisStreamParameters struct {
	Arn                     string            `json:"arn"`
	EncryptionType          string            `json:"encryption_type"`
	EnforceConsumerDeletion bool              `json:"enforce_consumer_deletion"`
	Id                      string            `json:"id"`
	KmsKeyId                string            `json:"kms_key_id"`
	Name                    string            `json:"name"`
	RetentionPeriod         int               `json:"retention_period"`
	ShardCount              int               `json:"shard_count"`
	ShardLevelMetrics       []string          `json:"shard_level_metrics"`
	Tags                    map[string]string `json:"tags"`
	Timeouts                []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A KinesisStreamStatus defines the observed state of a KinesisStream
type KinesisStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisStreamObservation `json:",inline"`
}

// A KinesisStreamObservation records the observed state of a KinesisStream
type KinesisStreamObservation struct{}