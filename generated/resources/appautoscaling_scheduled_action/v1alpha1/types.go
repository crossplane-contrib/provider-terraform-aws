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

// AppautoscalingScheduledAction is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppautoscalingScheduledAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppautoscalingScheduledActionSpec   `json:"spec"`
	Status AppautoscalingScheduledActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppautoscalingScheduledAction contains a list of AppautoscalingScheduledActionList
type AppautoscalingScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppautoscalingScheduledAction `json:"items"`
}

// A AppautoscalingScheduledActionSpec defines the desired state of a AppautoscalingScheduledAction
type AppautoscalingScheduledActionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppautoscalingScheduledActionParameters `json:",inline"`
}

// A AppautoscalingScheduledActionParameters defines the desired state of a AppautoscalingScheduledAction
type AppautoscalingScheduledActionParameters struct {
	EndTime              string               `json:"end_time"`
	Name                 string               `json:"name"`
	ResourceId           string               `json:"resource_id"`
	ScalableDimension    string               `json:"scalable_dimension"`
	Schedule             string               `json:"schedule"`
	ServiceNamespace     string               `json:"service_namespace"`
	StartTime            string               `json:"start_time"`
	ScalableTargetAction ScalableTargetAction `json:"scalable_target_action"`
}

type ScalableTargetAction struct {
	MaxCapacity int `json:"max_capacity"`
	MinCapacity int `json:"min_capacity"`
}

// A AppautoscalingScheduledActionStatus defines the observed state of a AppautoscalingScheduledAction
type AppautoscalingScheduledActionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppautoscalingScheduledActionObservation `json:",inline"`
}

// A AppautoscalingScheduledActionObservation records the observed state of a AppautoscalingScheduledAction
type AppautoscalingScheduledActionObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}