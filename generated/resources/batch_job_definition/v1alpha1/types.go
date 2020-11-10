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

// BatchJobDefinition is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BatchJobDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchJobDefinitionSpec   `json:"spec"`
	Status BatchJobDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BatchJobDefinition contains a list of BatchJobDefinitionList
type BatchJobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchJobDefinition `json:"items"`
}

// A BatchJobDefinitionSpec defines the desired state of a BatchJobDefinition
type BatchJobDefinitionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BatchJobDefinitionParameters `json:",inline"`
}

// A BatchJobDefinitionParameters defines the desired state of a BatchJobDefinition
type BatchJobDefinitionParameters struct {
	Type                string            `json:"type"`
	ContainerProperties string            `json:"container_properties"`
	Name                string            `json:"name"`
	Parameters          map[string]string `json:"parameters"`
	RetryStrategy       RetryStrategy     `json:"retry_strategy"`
	Timeout             Timeout           `json:"timeout"`
}

type RetryStrategy struct {
	Attempts int `json:"attempts"`
}

type Timeout struct {
	AttemptDurationSeconds int `json:"attempt_duration_seconds"`
}

// A BatchJobDefinitionStatus defines the observed state of a BatchJobDefinition
type BatchJobDefinitionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BatchJobDefinitionObservation `json:",inline"`
}

// A BatchJobDefinitionObservation records the observed state of a BatchJobDefinition
type BatchJobDefinitionObservation struct {
	Arn      string `json:"arn"`
	Id       string `json:"id"`
	Revision int    `json:"revision"`
}