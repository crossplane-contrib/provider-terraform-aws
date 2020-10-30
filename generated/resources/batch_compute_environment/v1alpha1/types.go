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

// BatchComputeEnvironment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BatchComputeEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchComputeEnvironmentSpec   `json:"spec"`
	Status BatchComputeEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BatchComputeEnvironment contains a list of BatchComputeEnvironmentList
type BatchComputeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchComputeEnvironment `json:"items"`
}

// A BatchComputeEnvironmentSpec defines the desired state of a BatchComputeEnvironment
type BatchComputeEnvironmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BatchComputeEnvironmentParameters `json:",inline"`
}

// A BatchComputeEnvironmentParameters defines the desired state of a BatchComputeEnvironment
type BatchComputeEnvironmentParameters struct {
	ComputeEnvironmentNamePrefix string `json:"compute_environment_name_prefix"`
	ServiceRole                  string `json:"service_role"`
	State                        string `json:"state"`
	Type                         string `json:"type"`
}

// A BatchComputeEnvironmentStatus defines the observed state of a BatchComputeEnvironment
type BatchComputeEnvironmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BatchComputeEnvironmentObservation `json:",inline"`
}

// A BatchComputeEnvironmentObservation records the observed state of a BatchComputeEnvironment
type BatchComputeEnvironmentObservation struct {
	Status                 string `json:"status"`
	StatusReason           string `json:"status_reason"`
	ComputeEnvironmentName string `json:"compute_environment_name"`
	EcsClusterArn          string `json:"ecs_cluster_arn"`
	Id                     string `json:"id"`
	Arn                    string `json:"arn"`
}