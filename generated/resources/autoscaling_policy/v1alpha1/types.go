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

// AutoscalingPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutoscalingPolicySpec   `json:"spec"`
	Status AutoscalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingPolicy contains a list of AutoscalingPolicyList
type AutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingPolicy `json:"items"`
}

// A AutoscalingPolicySpec defines the desired state of a AutoscalingPolicy
type AutoscalingPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AutoscalingPolicyParameters `json:",inline"`
}

// A AutoscalingPolicyParameters defines the desired state of a AutoscalingPolicy
type AutoscalingPolicyParameters struct {
	PolicyType                  string                      `json:"policy_type"`
	AdjustmentType              string                      `json:"adjustment_type"`
	AutoscalingGroupName        string                      `json:"autoscaling_group_name"`
	Cooldown                    int64                       `json:"cooldown"`
	Id                          string                      `json:"id"`
	MetricAggregationType       string                      `json:"metric_aggregation_type"`
	Name                        string                      `json:"name"`
	ScalingAdjustment           int64                       `json:"scaling_adjustment"`
	EstimatedInstanceWarmup     int64                       `json:"estimated_instance_warmup"`
	MinAdjustmentMagnitude      int64                       `json:"min_adjustment_magnitude"`
	TargetTrackingConfiguration TargetTrackingConfiguration `json:"target_tracking_configuration"`
	StepAdjustment              StepAdjustment              `json:"step_adjustment"`
}

type TargetTrackingConfiguration struct {
	DisableScaleIn                bool                          `json:"disable_scale_in"`
	TargetValue                   int64                         `json:"target_value"`
	CustomizedMetricSpecification CustomizedMetricSpecification `json:"customized_metric_specification"`
	PredefinedMetricSpecification PredefinedMetricSpecification `json:"predefined_metric_specification"`
}

type CustomizedMetricSpecification struct {
	Statistic       string          `json:"statistic"`
	Unit            string          `json:"unit"`
	MetricName      string          `json:"metric_name"`
	Namespace       string          `json:"namespace"`
	MetricDimension MetricDimension `json:"metric_dimension"`
}

type MetricDimension struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PredefinedMetricSpecification struct {
	PredefinedMetricType string `json:"predefined_metric_type"`
	ResourceLabel        string `json:"resource_label"`
}

type StepAdjustment struct {
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int64  `json:"scaling_adjustment"`
}

// A AutoscalingPolicyStatus defines the observed state of a AutoscalingPolicy
type AutoscalingPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AutoscalingPolicyObservation `json:",inline"`
}

// A AutoscalingPolicyObservation records the observed state of a AutoscalingPolicy
type AutoscalingPolicyObservation struct {
	Arn string `json:"arn"`
}