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

func EncodeAutoscalingPolicy(r AutoscalingPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingPolicy_AdjustmentType(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_AutoscalingGroupName(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_Cooldown(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_EstimatedInstanceWarmup(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_MetricAggregationType(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_MinAdjustmentMagnitude(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_PolicyType(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_ScalingAdjustment(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingPolicy_StepAdjustment(r.Spec.ForProvider.StepAdjustment, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration(r.Spec.ForProvider.TargetTrackingConfiguration, ctyVal)
	EncodeAutoscalingPolicy_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAutoscalingPolicy_AdjustmentType(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["adjustment_type"] = cty.StringVal(p.AdjustmentType)
}

func EncodeAutoscalingPolicy_AutoscalingGroupName(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["autoscaling_group_name"] = cty.StringVal(p.AutoscalingGroupName)
}

func EncodeAutoscalingPolicy_Cooldown(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["cooldown"] = cty.NumberIntVal(p.Cooldown)
}

func EncodeAutoscalingPolicy_EstimatedInstanceWarmup(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["estimated_instance_warmup"] = cty.NumberIntVal(p.EstimatedInstanceWarmup)
}

func EncodeAutoscalingPolicy_Id(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAutoscalingPolicy_MetricAggregationType(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["metric_aggregation_type"] = cty.StringVal(p.MetricAggregationType)
}

func EncodeAutoscalingPolicy_Name(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAutoscalingPolicy_MinAdjustmentMagnitude(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["min_adjustment_magnitude"] = cty.NumberIntVal(p.MinAdjustmentMagnitude)
}

func EncodeAutoscalingPolicy_PolicyType(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["policy_type"] = cty.StringVal(p.PolicyType)
}

func EncodeAutoscalingPolicy_ScalingAdjustment(p AutoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["scaling_adjustment"] = cty.NumberIntVal(p.ScalingAdjustment)
}

func EncodeAutoscalingPolicy_StepAdjustment(p StepAdjustment, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingPolicy_StepAdjustment_MetricIntervalLowerBound(p, ctyVal)
	EncodeAutoscalingPolicy_StepAdjustment_MetricIntervalUpperBound(p, ctyVal)
	EncodeAutoscalingPolicy_StepAdjustment_ScalingAdjustment(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["step_adjustment"] = cty.SetVal(valsForCollection)
}

func EncodeAutoscalingPolicy_StepAdjustment_MetricIntervalLowerBound(p StepAdjustment, vals map[string]cty.Value) {
	vals["metric_interval_lower_bound"] = cty.StringVal(p.MetricIntervalLowerBound)
}

func EncodeAutoscalingPolicy_StepAdjustment_MetricIntervalUpperBound(p StepAdjustment, vals map[string]cty.Value) {
	vals["metric_interval_upper_bound"] = cty.StringVal(p.MetricIntervalUpperBound)
}

func EncodeAutoscalingPolicy_StepAdjustment_ScalingAdjustment(p StepAdjustment, vals map[string]cty.Value) {
	vals["scaling_adjustment"] = cty.NumberIntVal(p.ScalingAdjustment)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration(p TargetTrackingConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_DisableScaleIn(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_TargetValue(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_PredefinedMetricSpecification(p.PredefinedMetricSpecification, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification(p.CustomizedMetricSpecification, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["target_tracking_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_DisableScaleIn(p TargetTrackingConfiguration, vals map[string]cty.Value) {
	vals["disable_scale_in"] = cty.BoolVal(p.DisableScaleIn)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_TargetValue(p TargetTrackingConfiguration, vals map[string]cty.Value) {
	vals["target_value"] = cty.NumberIntVal(p.TargetValue)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_PredefinedMetricSpecification(p PredefinedMetricSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_PredefinedMetricSpecification_PredefinedMetricType(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_PredefinedMetricSpecification_ResourceLabel(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["predefined_metric_specification"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_PredefinedMetricSpecification_PredefinedMetricType(p PredefinedMetricSpecification, vals map[string]cty.Value) {
	vals["predefined_metric_type"] = cty.StringVal(p.PredefinedMetricType)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_PredefinedMetricSpecification_ResourceLabel(p PredefinedMetricSpecification, vals map[string]cty.Value) {
	vals["resource_label"] = cty.StringVal(p.ResourceLabel)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricName(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_Namespace(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_Statistic(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_Unit(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricDimension(p.MetricDimension, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["customized_metric_specification"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricName(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_Namespace(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_Statistic(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["statistic"] = cty.StringVal(p.Statistic)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_Unit(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["unit"] = cty.StringVal(p.Unit)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricDimension(p MetricDimension, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricDimension_Name(p, ctyVal)
	EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricDimension_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["metric_dimension"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricDimension_Name(p MetricDimension, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAutoscalingPolicy_TargetTrackingConfiguration_CustomizedMetricSpecification_MetricDimension_Value(p MetricDimension, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeAutoscalingPolicy_Arn(p AutoscalingPolicyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}