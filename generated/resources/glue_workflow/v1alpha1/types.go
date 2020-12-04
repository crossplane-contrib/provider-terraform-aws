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

// GlueWorkflow is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueWorkflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueWorkflowSpec   `json:"spec"`
	Status GlueWorkflowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueWorkflow contains a list of GlueWorkflowList
type GlueWorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueWorkflow `json:"items"`
}

// A GlueWorkflowSpec defines the desired state of a GlueWorkflow
type GlueWorkflowSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueWorkflowParameters `json:",inline"`
}

// A GlueWorkflowParameters defines the desired state of a GlueWorkflow
type GlueWorkflowParameters struct {
	MaxConcurrentRuns    int64             `json:"max_concurrent_runs"`
	Name                 string            `json:"name"`
	Tags                 map[string]string `json:"tags"`
	DefaultRunProperties map[string]string `json:"default_run_properties"`
	Description          string            `json:"description"`
	Id                   string            `json:"id"`
}

// A GlueWorkflowStatus defines the observed state of a GlueWorkflow
type GlueWorkflowStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueWorkflowObservation `json:",inline"`
}

// A GlueWorkflowObservation records the observed state of a GlueWorkflow
type GlueWorkflowObservation struct {
	Arn string `json:"arn"`
}