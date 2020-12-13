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
	ForProvider                  CloudwatchMetricAlarmParameters `json:"forProvider"`
}

// A CloudwatchMetricAlarmParameters defines the desired state of a CloudwatchMetricAlarm
type CloudwatchMetricAlarmParameters struct {
	ActionsEnabled                    bool              `json:"actions_enabled"`
	DatapointsToAlarm                 int64             `json:"datapoints_to_alarm"`
	ThresholdMetricId                 string            `json:"threshold_metric_id"`
	Statistic                         string            `json:"statistic"`
	Threshold                         int64             `json:"threshold"`
	AlarmDescription                  string            `json:"alarm_description"`
	Dimensions                        map[string]string `json:"dimensions"`
	Id                                string            `json:"id"`
	OkActions                         []string          `json:"ok_actions"`
	Period                            int64             `json:"period"`
	Tags                              map[string]string `json:"tags"`
	Unit                              string            `json:"unit"`
	AlarmActions                      []string          `json:"alarm_actions"`
	AlarmName                         string            `json:"alarm_name"`
	InsufficientDataActions           []string          `json:"insufficient_data_actions"`
	MetricName                        string            `json:"metric_name"`
	Namespace                         string            `json:"namespace"`
	TreatMissingData                  string            `json:"treat_missing_data"`
	ComparisonOperator                string            `json:"comparison_operator"`
	EvaluateLowSampleCountPercentiles string            `json:"evaluate_low_sample_count_percentiles"`
	EvaluationPeriods                 int64             `json:"evaluation_periods"`
	ExtendedStatistic                 string            `json:"extended_statistic"`
	MetricQuery                       MetricQuery       `json:"metric_query"`
}

type MetricQuery struct {
	ReturnData bool   `json:"return_data"`
	Expression string `json:"expression"`
	Id         string `json:"id"`
	Label      string `json:"label"`
	Metric     Metric `json:"metric"`
}

type Metric struct {
	Dimensions map[string]string `json:"dimensions"`
	MetricName string            `json:"metric_name"`
	Namespace  string            `json:"namespace"`
	Period     int64             `json:"period"`
	Stat       string            `json:"stat"`
	Unit       string            `json:"unit"`
}

// A CloudwatchMetricAlarmStatus defines the observed state of a CloudwatchMetricAlarm
type CloudwatchMetricAlarmStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchMetricAlarmObservation `json:"atProvider"`
}

// A CloudwatchMetricAlarmObservation records the observed state of a CloudwatchMetricAlarm
type CloudwatchMetricAlarmObservation struct {
	Arn string `json:"arn"`
}