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

// CloudwatchLogMetricFilter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchLogMetricFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchLogMetricFilterSpec   `json:"spec"`
	Status CloudwatchLogMetricFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogMetricFilter contains a list of CloudwatchLogMetricFilterList
type CloudwatchLogMetricFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogMetricFilter `json:"items"`
}

// A CloudwatchLogMetricFilterSpec defines the desired state of a CloudwatchLogMetricFilter
type CloudwatchLogMetricFilterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchLogMetricFilterParameters `json:",inline"`
}

// A CloudwatchLogMetricFilterParameters defines the desired state of a CloudwatchLogMetricFilter
type CloudwatchLogMetricFilterParameters struct {
	LogGroupName         string               `json:"log_group_name"`
	Name                 string               `json:"name"`
	Pattern              string               `json:"pattern"`
	Id                   string               `json:"id"`
	MetricTransformation MetricTransformation `json:"metric_transformation"`
}

type MetricTransformation struct {
	DefaultValue string `json:"default_value"`
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Value        string `json:"value"`
}

// A CloudwatchLogMetricFilterStatus defines the observed state of a CloudwatchLogMetricFilter
type CloudwatchLogMetricFilterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchLogMetricFilterObservation `json:",inline"`
}

// A CloudwatchLogMetricFilterObservation records the observed state of a CloudwatchLogMetricFilter
type CloudwatchLogMetricFilterObservation struct{}