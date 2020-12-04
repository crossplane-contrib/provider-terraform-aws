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

func EncodeAppautoscalingPolicy(r AppautoscalingPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingPolicy_PolicyType(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingPolicy_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingPolicy_ScalableDimension(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingPolicy_ServiceNamespace(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration(r.Spec.ForProvider.TargetTrackingScalingPolicyConfiguration, ctyVal)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration(r.Spec.ForProvider.StepScalingPolicyConfiguration, ctyVal)
	EncodeAppautoscalingPolicy_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppautoscalingPolicy_Id(p AppautoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppautoscalingPolicy_Name(p AppautoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppautoscalingPolicy_PolicyType(p AppautoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["policy_type"] = cty.StringVal(p.PolicyType)
}

func EncodeAppautoscalingPolicy_ResourceId(p AppautoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeAppautoscalingPolicy_ScalableDimension(p AppautoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["scalable_dimension"] = cty.StringVal(p.ScalableDimension)
}

func EncodeAppautoscalingPolicy_ServiceNamespace(p AppautoscalingPolicyParameters, vals map[string]cty.Value) {
	vals["service_namespace"] = cty.StringVal(p.ServiceNamespace)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration(p TargetTrackingScalingPolicyConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_ScaleInCooldown(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_ScaleOutCooldown(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_TargetValue(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_DisableScaleIn(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification(p.CustomizedMetricSpecification, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_PredefinedMetricSpecification(p.PredefinedMetricSpecification, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["target_tracking_scaling_policy_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_ScaleInCooldown(p TargetTrackingScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["scale_in_cooldown"] = cty.NumberIntVal(p.ScaleInCooldown)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_ScaleOutCooldown(p TargetTrackingScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["scale_out_cooldown"] = cty.NumberIntVal(p.ScaleOutCooldown)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_TargetValue(p TargetTrackingScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["target_value"] = cty.NumberIntVal(p.TargetValue)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_DisableScaleIn(p TargetTrackingScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["disable_scale_in"] = cty.BoolVal(p.DisableScaleIn)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Unit(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_MetricName(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Namespace(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Statistic(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Dimensions(p.Dimensions, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["customized_metric_specification"] = cty.ListVal(valsForCollection)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Unit(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["unit"] = cty.StringVal(p.Unit)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_MetricName(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Namespace(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["namespace"] = cty.StringVal(p.Namespace)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Statistic(p CustomizedMetricSpecification, vals map[string]cty.Value) {
	vals["statistic"] = cty.StringVal(p.Statistic)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Dimensions(p Dimensions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Dimensions_Name(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Dimensions_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dimensions"] = cty.SetVal(valsForCollection)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Dimensions_Name(p Dimensions, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_CustomizedMetricSpecification_Dimensions_Value(p Dimensions, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_PredefinedMetricSpecification(p PredefinedMetricSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_PredefinedMetricSpecification_PredefinedMetricType(p, ctyVal)
	EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_PredefinedMetricSpecification_ResourceLabel(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["predefined_metric_specification"] = cty.ListVal(valsForCollection)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_PredefinedMetricSpecification_PredefinedMetricType(p PredefinedMetricSpecification, vals map[string]cty.Value) {
	vals["predefined_metric_type"] = cty.StringVal(p.PredefinedMetricType)
}

func EncodeAppautoscalingPolicy_TargetTrackingScalingPolicyConfiguration_PredefinedMetricSpecification_ResourceLabel(p PredefinedMetricSpecification, vals map[string]cty.Value) {
	vals["resource_label"] = cty.StringVal(p.ResourceLabel)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration(p StepScalingPolicyConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_AdjustmentType(p, ctyVal)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_Cooldown(p, ctyVal)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_MetricAggregationType(p, ctyVal)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_MinAdjustmentMagnitude(p, ctyVal)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment(p.StepAdjustment, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["step_scaling_policy_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_AdjustmentType(p StepScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["adjustment_type"] = cty.StringVal(p.AdjustmentType)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_Cooldown(p StepScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["cooldown"] = cty.NumberIntVal(p.Cooldown)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_MetricAggregationType(p StepScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["metric_aggregation_type"] = cty.StringVal(p.MetricAggregationType)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_MinAdjustmentMagnitude(p StepScalingPolicyConfiguration, vals map[string]cty.Value) {
	vals["min_adjustment_magnitude"] = cty.NumberIntVal(p.MinAdjustmentMagnitude)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment(p StepAdjustment, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment_MetricIntervalUpperBound(p, ctyVal)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment_ScalingAdjustment(p, ctyVal)
	EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment_MetricIntervalLowerBound(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["step_adjustment"] = cty.SetVal(valsForCollection)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment_MetricIntervalUpperBound(p StepAdjustment, vals map[string]cty.Value) {
	vals["metric_interval_upper_bound"] = cty.StringVal(p.MetricIntervalUpperBound)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment_ScalingAdjustment(p StepAdjustment, vals map[string]cty.Value) {
	vals["scaling_adjustment"] = cty.NumberIntVal(p.ScalingAdjustment)
}

func EncodeAppautoscalingPolicy_StepScalingPolicyConfiguration_StepAdjustment_MetricIntervalLowerBound(p StepAdjustment, vals map[string]cty.Value) {
	vals["metric_interval_lower_bound"] = cty.StringVal(p.MetricIntervalLowerBound)
}

func EncodeAppautoscalingPolicy_Arn(p AppautoscalingPolicyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}