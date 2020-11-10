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

// LambdaEventSourceMapping is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LambdaEventSourceMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LambdaEventSourceMappingSpec   `json:"spec"`
	Status LambdaEventSourceMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaEventSourceMapping contains a list of LambdaEventSourceMappingList
type LambdaEventSourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaEventSourceMapping `json:"items"`
}

// A LambdaEventSourceMappingSpec defines the desired state of a LambdaEventSourceMapping
type LambdaEventSourceMappingSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LambdaEventSourceMappingParameters `json:",inline"`
}

// A LambdaEventSourceMappingParameters defines the desired state of a LambdaEventSourceMapping
type LambdaEventSourceMappingParameters struct {
	EventSourceArn                 string            `json:"event_source_arn"`
	Enabled                        bool              `json:"enabled"`
	BatchSize                      int               `json:"batch_size"`
	StartingPosition               string            `json:"starting_position"`
	StartingPositionTimestamp      string            `json:"starting_position_timestamp"`
	BisectBatchOnFunctionError     bool              `json:"bisect_batch_on_function_error"`
	FunctionName                   string            `json:"function_name"`
	MaximumBatchingWindowInSeconds int               `json:"maximum_batching_window_in_seconds"`
	DestinationConfig              DestinationConfig `json:"destination_config"`
}

type DestinationConfig struct {
	OnFailure OnFailure `json:"on_failure"`
}

type OnFailure struct {
	DestinationArn string `json:"destination_arn"`
}

// A LambdaEventSourceMappingStatus defines the observed state of a LambdaEventSourceMapping
type LambdaEventSourceMappingStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaEventSourceMappingObservation `json:",inline"`
}

// A LambdaEventSourceMappingObservation records the observed state of a LambdaEventSourceMapping
type LambdaEventSourceMappingObservation struct {
	ParallelizationFactor     int    `json:"parallelization_factor"`
	LastProcessingResult      string `json:"last_processing_result"`
	StateTransitionReason     string `json:"state_transition_reason"`
	Uuid                      string `json:"uuid"`
	FunctionArn               string `json:"function_arn"`
	Id                        string `json:"id"`
	LastModified              string `json:"last_modified"`
	State                     string `json:"state"`
	MaximumRetryAttempts      int    `json:"maximum_retry_attempts"`
	MaximumRecordAgeInSeconds int    `json:"maximum_record_age_in_seconds"`
}