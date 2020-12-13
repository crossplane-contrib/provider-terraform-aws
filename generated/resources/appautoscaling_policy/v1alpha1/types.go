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

// AppautoscalingPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppautoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppautoscalingPolicySpec   `json:"spec"`
	Status AppautoscalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppautoscalingPolicy contains a list of AppautoscalingPolicyList
type AppautoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppautoscalingPolicy `json:"items"`
}

// A AppautoscalingPolicySpec defines the desired state of a AppautoscalingPolicy
type AppautoscalingPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppautoscalingPolicyParameters `json:"forProvider"`
}

// A AppautoscalingPolicyParameters defines the desired state of a AppautoscalingPolicy
type AppautoscalingPolicyParameters struct {
	Id                                       string                                   `json:"id"`
	Name                                     string                                   `json:"name"`
	PolicyType                               string                                   `json:"policy_type"`
	ResourceId                               string                                   `json:"resource_id"`
	ScalableDimension                        string                                   `json:"scalable_dimension"`
	ServiceNamespace                         string                                   `json:"service_namespace"`
	StepScalingPolicyConfiguration           StepScalingPolicyConfiguration           `json:"step_scaling_policy_configuration"`
	TargetTrackingScalingPolicyConfiguration TargetTrackingScalingPolicyConfiguration `json:"target_tracking_scaling_policy_configuration"`
}

type StepScalingPolicyConfiguration struct {
	AdjustmentType         string         `json:"adjustment_type"`
	Cooldown               int64          `json:"cooldown"`
	MetricAggregationType  string         `json:"metric_aggregation_type"`
	MinAdjustmentMagnitude int64          `json:"min_adjustment_magnitude"`
	StepAdjustment         StepAdjustment `json:"step_adjustment"`
}

type StepAdjustment struct {
	MetricIntervalUpperBound string `json:"metric_interval_upper_bound"`
	ScalingAdjustment        int64  `json:"scaling_adjustment"`
	MetricIntervalLowerBound string `json:"metric_interval_lower_bound"`
}

type TargetTrackingScalingPolicyConfiguration struct {
	DisableScaleIn                bool                          `json:"disable_scale_in"`
	ScaleInCooldown               int64                         `json:"scale_in_cooldown"`
	ScaleOutCooldown              int64                         `json:"scale_out_cooldown"`
	TargetValue                   int64                         `json:"target_value"`
	CustomizedMetricSpecification CustomizedMetricSpecification `json:"customized_metric_specification"`
	PredefinedMetricSpecification PredefinedMetricSpecification `json:"predefined_metric_specification"`
}

type CustomizedMetricSpecification struct {
	MetricName string     `json:"metric_name"`
	Namespace  string     `json:"namespace"`
	Statistic  string     `json:"statistic"`
	Unit       string     `json:"unit"`
	Dimensions Dimensions `json:"dimensions"`
}

type Dimensions struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type PredefinedMetricSpecification struct {
	ResourceLabel        string `json:"resource_label"`
	PredefinedMetricType string `json:"predefined_metric_type"`
}

// A AppautoscalingPolicyStatus defines the observed state of a AppautoscalingPolicy
type AppautoscalingPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppautoscalingPolicyObservation `json:"atProvider"`
}

// A AppautoscalingPolicyObservation records the observed state of a AppautoscalingPolicy
type AppautoscalingPolicyObservation struct {
	Arn string `json:"arn"`
}