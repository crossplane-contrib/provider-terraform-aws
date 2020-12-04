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
	"github.com/zclconf/go-cty/cty"
)

func EncodeCloudwatchMetricAlarm(r CloudwatchMetricAlarm) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchMetricAlarm_ComparisonOperator(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Threshold(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_ThresholdMetricId(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Period(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_TreatMissingData(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_ActionsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_DatapointsToAlarm(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Dimensions(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Statistic(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_AlarmActions(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_EvaluateLowSampleCountPercentiles(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_InsufficientDataActions(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_OkActions(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Unit(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_AlarmDescription(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_AlarmName(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_EvaluationPeriods(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_ExtendedStatistic(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_Namespace(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery(r.Spec.ForProvider.MetricQuery, ctyVal)
	EncodeCloudwatchMetricAlarm_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchMetricAlarm_ComparisonOperator(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeCloudwatchMetricAlarm_Threshold(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["threshold"] = cty.NumberIntVal(p.Threshold)
}

func EncodeCloudwatchMetricAlarm_ThresholdMetricId(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["threshold_metric_id"] = cty.StringVal(p.ThresholdMetricId)
}

func EncodeCloudwatchMetricAlarm_Period(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["period"] = cty.NumberIntVal(p.Period)
}

func EncodeCloudwatchMetricAlarm_TreatMissingData(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["treat_missing_data"] = cty.StringVal(p.TreatMissingData)
}

func EncodeCloudwatchMetricAlarm_ActionsEnabled(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["actions_enabled"] = cty.BoolVal(p.ActionsEnabled)
}

func EncodeCloudwatchMetricAlarm_DatapointsToAlarm(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["datapoints_to_alarm"] = cty.NumberIntVal(p.DatapointsToAlarm)
}

func EncodeCloudwatchMetricAlarm_Dimensions(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Dimensions {
		mVals[key] = cty.StringVal(value)
	}
	vals["dimensions"] = cty.MapVal(mVals)
}

func EncodeCloudwatchMetricAlarm_Id(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchMetricAlarm_MetricName(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeCloudwatchMetricAlarm_Statistic(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["statistic"] = cty.StringVal(p.Statistic)
}

func EncodeCloudwatchMetricAlarm_Tags(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudwatchMetricAlarm_AlarmActions(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AlarmActions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["alarm_actions"] = cty.SetVal(colVals)
}

func EncodeCloudwatchMetricAlarm_EvaluateLowSampleCountPercentiles(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["evaluate_low_sample_count_percentiles"] = cty.StringVal(p.EvaluateLowSampleCountPercentiles)
}

func EncodeCloudwatchMetricAlarm_InsufficientDataActions(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.InsufficientDataActions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["insufficient_data_actions"] = cty.SetVal(colVals)
}

func EncodeCloudwatchMetricAlarm_OkActions(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.OkActions {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ok_actions"] = cty.SetVal(colVals)
}

func EncodeCloudwatchMetricAlarm_Unit(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["unit"] = cty.StringVal(p.Unit)
}

func EncodeCloudwatchMetricAlarm_AlarmDescription(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["alarm_description"] = cty.StringVal(p.AlarmDescription)
}

func EncodeCloudwatchMetricAlarm_AlarmName(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["alarm_name"] = cty.StringVal(p.AlarmName)
}

func EncodeCloudwatchMetricAlarm_EvaluationPeriods(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["evaluation_periods"] = cty.NumberIntVal(p.EvaluationPeriods)
}

func EncodeCloudwatchMetricAlarm_ExtendedStatistic(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["extended_statistic"] = cty.StringVal(p.ExtendedStatistic)
}

func EncodeCloudwatchMetricAlarm_Namespace(p CloudwatchMetricAlarmParameters, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeCloudwatchMetricAlarm_MetricQuery(p MetricQuery, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchMetricAlarm_MetricQuery_ReturnData(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Expression(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Id(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Label(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Metric(p.Metric, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["metric_query"] = cty.SetVal(valsForCollection)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_ReturnData(p MetricQuery, vals map[string]cty.Value) {
	vals["return_data"] = cty.BoolVal(p.ReturnData)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Expression(p MetricQuery, vals map[string]cty.Value) {
	vals["expression"] = cty.StringVal(p.Expression)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Id(p MetricQuery, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Label(p MetricQuery, vals map[string]cty.Value) {
	vals["label"] = cty.StringVal(p.Label)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Metric(p Metric, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Period(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Stat(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Unit(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Dimensions(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Metric_MetricName(p, ctyVal)
	EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Namespace(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["metric"] = cty.ListVal(valsForCollection)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Period(p Metric, vals map[string]cty.Value) {
	vals["period"] = cty.NumberIntVal(p.Period)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Stat(p Metric, vals map[string]cty.Value) {
	vals["stat"] = cty.StringVal(p.Stat)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Unit(p Metric, vals map[string]cty.Value) {
	vals["unit"] = cty.StringVal(p.Unit)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Dimensions(p Metric, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Dimensions {
		mVals[key] = cty.StringVal(value)
	}
	vals["dimensions"] = cty.MapVal(mVals)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Metric_MetricName(p Metric, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeCloudwatchMetricAlarm_MetricQuery_Metric_Namespace(p Metric, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeCloudwatchMetricAlarm_Arn(p CloudwatchMetricAlarmObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}