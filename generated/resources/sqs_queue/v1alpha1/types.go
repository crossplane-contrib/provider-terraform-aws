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

// SqsQueue is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SqsQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqsQueueSpec   `json:"spec"`
	Status SqsQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SqsQueue contains a list of SqsQueueList
type SqsQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqsQueue `json:"items"`
}

// A SqsQueueSpec defines the desired state of a SqsQueue
type SqsQueueSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SqsQueueParameters `json:",inline"`
}

// A SqsQueueParameters defines the desired state of a SqsQueue
type SqsQueueParameters struct {
	FifoQueue                    bool              `json:"fifo_queue"`
	RedrivePolicy                string            `json:"redrive_policy"`
	MaxMessageSize               int64             `json:"max_message_size"`
	Name                         string            `json:"name"`
	Tags                         map[string]string `json:"tags"`
	Policy                       string            `json:"policy"`
	ReceiveWaitTimeSeconds       int64             `json:"receive_wait_time_seconds"`
	VisibilityTimeoutSeconds     int64             `json:"visibility_timeout_seconds"`
	ContentBasedDeduplication    bool              `json:"content_based_deduplication"`
	Id                           string            `json:"id"`
	KmsMasterKeyId               string            `json:"kms_master_key_id"`
	NamePrefix                   string            `json:"name_prefix"`
	DelaySeconds                 int64             `json:"delay_seconds"`
	KmsDataKeyReusePeriodSeconds int64             `json:"kms_data_key_reuse_period_seconds"`
	MessageRetentionSeconds      int64             `json:"message_retention_seconds"`
}

// A SqsQueueStatus defines the observed state of a SqsQueue
type SqsQueueStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SqsQueueObservation `json:",inline"`
}

// A SqsQueueObservation records the observed state of a SqsQueue
type SqsQueueObservation struct {
	Arn string `json:"arn"`
}