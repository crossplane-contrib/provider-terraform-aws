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

// AppautoscalingTarget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppautoscalingTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppautoscalingTargetSpec   `json:"spec"`
	Status AppautoscalingTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppautoscalingTarget contains a list of AppautoscalingTargetList
type AppautoscalingTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppautoscalingTarget `json:"items"`
}

// A AppautoscalingTargetSpec defines the desired state of a AppautoscalingTarget
type AppautoscalingTargetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppautoscalingTargetParameters `json:"forProvider"`
}

// A AppautoscalingTargetParameters defines the desired state of a AppautoscalingTarget
type AppautoscalingTargetParameters struct {
	ServiceNamespace  string `json:"service_namespace"`
	Id                string `json:"id"`
	MaxCapacity       int64  `json:"max_capacity"`
	MinCapacity       int64  `json:"min_capacity"`
	ResourceId        string `json:"resource_id"`
	RoleArn           string `json:"role_arn"`
	ScalableDimension string `json:"scalable_dimension"`
}

// A AppautoscalingTargetStatus defines the observed state of a AppautoscalingTarget
type AppautoscalingTargetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppautoscalingTargetObservation `json:"atProvider"`
}

// A AppautoscalingTargetObservation records the observed state of a AppautoscalingTarget
type AppautoscalingTargetObservation struct{}