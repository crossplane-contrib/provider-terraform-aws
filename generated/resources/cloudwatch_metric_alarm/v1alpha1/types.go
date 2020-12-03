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

// CloudwatchMetricAlarm is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchMetricAlarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchMetricAlarmSpec   `json:"spec"`
	Status CloudwatchMetricAlarmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchMetricAlarm contains a list of CloudwatchMetricAlarmList
type CloudwatchMetricAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchMetricAlarm `json:"items"`
}

// A CloudwatchMetricAlarmSpec defines the desired state of a CloudwatchMetricAlarm
type CloudwatchMetricAlarmSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchMetricAlarmParameters `json:",inline"`
}

// A CloudwatchMetricAlarmParameters defines the desired state of a CloudwatchMetricAlarm
type CloudwatchMetricAlarmParameters struct {
	Threshold                         int               `json:"threshold"`
	Unit                              string            `json:"unit"`
	EvaluationPeriods                 int               `json:"evaluation_periods"`
	ExtendedStatistic                 string            `json:"extended_statistic"`
	MetricName                        string            `json:"metric_name"`
	Period                            int               `json:"period"`
	OkActions                         []string          `json:"ok_actions"`
	ThresholdMetricId                 string            `json:"threshold_metric_id"`
	TreatMissingData                  string            `json:"treat_missing_data"`
	AlarmName                         string            `json:"alarm_name"`
	DatapointsToAlarm                 int               `json:"datapoints_to_alarm"`
	InsufficientDataActions           []string          `json:"insufficient_data_actions"`
	Namespace                         string            `json:"namespace"`
	EvaluateLowSampleCountPercentiles string            `json:"evaluate_low_sample_count_percentiles"`
	Tags                              map[string]string `json:"tags"`
	ActionsEnabled                    bool              `json:"actions_enabled"`
	AlarmDescription                  string            `json:"alarm_description"`
	ComparisonOperator                string            `json:"comparison_operator"`
	AlarmActions                      []string          `json:"alarm_actions"`
	Dimensions                        map[string]string `json:"dimensions"`
	Id                                string            `json:"id"`
	Statistic                         string            `json:"statistic"`
	MetricQuery                       []MetricQuery     `json:"metric_query"`
}

type MetricQuery struct {
	Expression string `json:"expression"`
	Id         string `json:"id"`
	Label      string `json:"label"`
	ReturnData bool   `json:"return_data"`
	Metric     Metric `json:"metric"`
}

type Metric struct {
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Period     int               `json:"period"`
	Stat       string            `json:"stat"`
	Unit       string            `json:"unit"`
	Dimensions map[string]string `json:"dimensions"`
}

// A CloudwatchMetricAlarmStatus defines the observed state of a CloudwatchMetricAlarm
type CloudwatchMetricAlarmStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchMetricAlarmObservation `json:",inline"`
}

// A CloudwatchMetricAlarmObservation records the observed state of a CloudwatchMetricAlarm
type CloudwatchMetricAlarmObservation struct {
	Arn string `json:"arn"`
}