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

// GameliftGameSessionQueue is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GameliftGameSessionQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GameliftGameSessionQueueSpec   `json:"spec"`
	Status GameliftGameSessionQueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftGameSessionQueue contains a list of GameliftGameSessionQueueList
type GameliftGameSessionQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameliftGameSessionQueue `json:"items"`
}

// A GameliftGameSessionQueueSpec defines the desired state of a GameliftGameSessionQueue
type GameliftGameSessionQueueSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GameliftGameSessionQueueParameters `json:"forProvider"`
}

// A GameliftGameSessionQueueParameters defines the desired state of a GameliftGameSessionQueue
type GameliftGameSessionQueueParameters struct {
	Destinations        []string            `json:"destinations"`
	Id                  string              `json:"id"`
	Name                string              `json:"name"`
	Tags                map[string]string   `json:"tags"`
	TimeoutInSeconds    int64               `json:"timeout_in_seconds"`
	PlayerLatencyPolicy PlayerLatencyPolicy `json:"player_latency_policy"`
}

type PlayerLatencyPolicy struct {
	MaximumIndividualPlayerLatencyMilliseconds int64 `json:"maximum_individual_player_latency_milliseconds"`
	PolicyDurationSeconds                      int64 `json:"policy_duration_seconds"`
}

// A GameliftGameSessionQueueStatus defines the observed state of a GameliftGameSessionQueue
type GameliftGameSessionQueueStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GameliftGameSessionQueueObservation `json:"atProvider"`
}

// A GameliftGameSessionQueueObservation records the observed state of a GameliftGameSessionQueue
type GameliftGameSessionQueueObservation struct {
	Arn string `json:"arn"`
}