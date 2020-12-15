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

// AutoscalingSchedule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AutoscalingSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutoscalingScheduleSpec   `json:"spec"`
	Status AutoscalingScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingSchedule contains a list of AutoscalingScheduleList
type AutoscalingScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingSchedule `json:"items"`
}

// A AutoscalingScheduleSpec defines the desired state of a AutoscalingSchedule
type AutoscalingScheduleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AutoscalingScheduleParameters `json:"forProvider"`
}

// A AutoscalingScheduleParameters defines the desired state of a AutoscalingSchedule
type AutoscalingScheduleParameters struct {
	Id                   string `json:"id"`
	Recurrence           string `json:"recurrence"`
	StartTime            string `json:"start_time"`
	AutoscalingGroupName string `json:"autoscaling_group_name"`
	EndTime              string `json:"end_time"`
	ScheduledActionName  string `json:"scheduled_action_name"`
	DesiredCapacity      int64  `json:"desired_capacity"`
	MaxSize              int64  `json:"max_size"`
	MinSize              int64  `json:"min_size"`
}

// A AutoscalingScheduleStatus defines the observed state of a AutoscalingSchedule
type AutoscalingScheduleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AutoscalingScheduleObservation `json:"atProvider"`
}

// A AutoscalingScheduleObservation records the observed state of a AutoscalingSchedule
type AutoscalingScheduleObservation struct {
	Arn string `json:"arn"`
}