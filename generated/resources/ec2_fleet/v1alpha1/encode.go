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

func EncodeEc2Fleet(r Ec2Fleet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Fleet_ReplaceUnhealthyInstances(r.Spec.ForProvider, ctyVal)
	EncodeEc2Fleet_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2Fleet_TerminateInstances(r.Spec.ForProvider, ctyVal)
	EncodeEc2Fleet_TerminateInstancesWithExpiration(r.Spec.ForProvider, ctyVal)
	EncodeEc2Fleet_Type(r.Spec.ForProvider, ctyVal)
	EncodeEc2Fleet_ExcessCapacityTerminationPolicy(r.Spec.ForProvider, ctyVal)
	EncodeEc2Fleet_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2Fleet_TargetCapacitySpecification(r.Spec.ForProvider.TargetCapacitySpecification, ctyVal)
	EncodeEc2Fleet_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeEc2Fleet_LaunchTemplateConfig(r.Spec.ForProvider.LaunchTemplateConfig, ctyVal)
	EncodeEc2Fleet_OnDemandOptions(r.Spec.ForProvider.OnDemandOptions, ctyVal)
	EncodeEc2Fleet_SpotOptions(r.Spec.ForProvider.SpotOptions, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEc2Fleet_ReplaceUnhealthyInstances(p Ec2FleetParameters, vals map[string]cty.Value) {
	vals["replace_unhealthy_instances"] = cty.BoolVal(p.ReplaceUnhealthyInstances)
}

func EncodeEc2Fleet_Tags(p Ec2FleetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2Fleet_TerminateInstances(p Ec2FleetParameters, vals map[string]cty.Value) {
	vals["terminate_instances"] = cty.BoolVal(p.TerminateInstances)
}

func EncodeEc2Fleet_TerminateInstancesWithExpiration(p Ec2FleetParameters, vals map[string]cty.Value) {
	vals["terminate_instances_with_expiration"] = cty.BoolVal(p.TerminateInstancesWithExpiration)
}

func EncodeEc2Fleet_Type(p Ec2FleetParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEc2Fleet_ExcessCapacityTerminationPolicy(p Ec2FleetParameters, vals map[string]cty.Value) {
	vals["excess_capacity_termination_policy"] = cty.StringVal(p.ExcessCapacityTerminationPolicy)
}

func EncodeEc2Fleet_Id(p Ec2FleetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2Fleet_TargetCapacitySpecification(p TargetCapacitySpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Fleet_TargetCapacitySpecification_DefaultTargetCapacityType(p, ctyVal)
	EncodeEc2Fleet_TargetCapacitySpecification_OnDemandTargetCapacity(p, ctyVal)
	EncodeEc2Fleet_TargetCapacitySpecification_SpotTargetCapacity(p, ctyVal)
	EncodeEc2Fleet_TargetCapacitySpecification_TotalTargetCapacity(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["target_capacity_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEc2Fleet_TargetCapacitySpecification_DefaultTargetCapacityType(p TargetCapacitySpecification, vals map[string]cty.Value) {
	vals["default_target_capacity_type"] = cty.StringVal(p.DefaultTargetCapacityType)
}

func EncodeEc2Fleet_TargetCapacitySpecification_OnDemandTargetCapacity(p TargetCapacitySpecification, vals map[string]cty.Value) {
	vals["on_demand_target_capacity"] = cty.NumberIntVal(p.OnDemandTargetCapacity)
}

func EncodeEc2Fleet_TargetCapacitySpecification_SpotTargetCapacity(p TargetCapacitySpecification, vals map[string]cty.Value) {
	vals["spot_target_capacity"] = cty.NumberIntVal(p.SpotTargetCapacity)
}

func EncodeEc2Fleet_TargetCapacitySpecification_TotalTargetCapacity(p TargetCapacitySpecification, vals map[string]cty.Value) {
	vals["total_target_capacity"] = cty.NumberIntVal(p.TotalTargetCapacity)
}

func EncodeEc2Fleet_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Fleet_Timeouts_Update(p, ctyVal)
	EncodeEc2Fleet_Timeouts_Create(p, ctyVal)
	EncodeEc2Fleet_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeEc2Fleet_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeEc2Fleet_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeEc2Fleet_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeEc2Fleet_LaunchTemplateConfig(p LaunchTemplateConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification(p.LaunchTemplateSpecification, ctyVal)
	EncodeEc2Fleet_LaunchTemplateConfig_Override(p.Override, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["launch_template_config"] = cty.ListVal(valsForCollection)
}

func EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification(p LaunchTemplateSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification_LaunchTemplateId(p, ctyVal)
	EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification_LaunchTemplateName(p, ctyVal)
	EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification_Version(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["launch_template_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification_LaunchTemplateId(p LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["launch_template_id"] = cty.StringVal(p.LaunchTemplateId)
}

func EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification_LaunchTemplateName(p LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["launch_template_name"] = cty.StringVal(p.LaunchTemplateName)
}

func EncodeEc2Fleet_LaunchTemplateConfig_LaunchTemplateSpecification_Version(p LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeEc2Fleet_LaunchTemplateConfig_Override(p []Override, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeEc2Fleet_LaunchTemplateConfig_Override_InstanceType(v, ctyVal)
		EncodeEc2Fleet_LaunchTemplateConfig_Override_MaxPrice(v, ctyVal)
		EncodeEc2Fleet_LaunchTemplateConfig_Override_Priority(v, ctyVal)
		EncodeEc2Fleet_LaunchTemplateConfig_Override_SubnetId(v, ctyVal)
		EncodeEc2Fleet_LaunchTemplateConfig_Override_WeightedCapacity(v, ctyVal)
		EncodeEc2Fleet_LaunchTemplateConfig_Override_AvailabilityZone(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["override"] = cty.ListVal(valsForCollection)
}

func EncodeEc2Fleet_LaunchTemplateConfig_Override_InstanceType(p Override, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEc2Fleet_LaunchTemplateConfig_Override_MaxPrice(p Override, vals map[string]cty.Value) {
	vals["max_price"] = cty.StringVal(p.MaxPrice)
}

func EncodeEc2Fleet_LaunchTemplateConfig_Override_Priority(p Override, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeEc2Fleet_LaunchTemplateConfig_Override_SubnetId(p Override, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeEc2Fleet_LaunchTemplateConfig_Override_WeightedCapacity(p Override, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.NumberIntVal(p.WeightedCapacity)
}

func EncodeEc2Fleet_LaunchTemplateConfig_Override_AvailabilityZone(p Override, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeEc2Fleet_OnDemandOptions(p OnDemandOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Fleet_OnDemandOptions_AllocationStrategy(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["on_demand_options"] = cty.ListVal(valsForCollection)
}

func EncodeEc2Fleet_OnDemandOptions_AllocationStrategy(p OnDemandOptions, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEc2Fleet_SpotOptions(p SpotOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeEc2Fleet_SpotOptions_AllocationStrategy(p, ctyVal)
	EncodeEc2Fleet_SpotOptions_InstanceInterruptionBehavior(p, ctyVal)
	EncodeEc2Fleet_SpotOptions_InstancePoolsToUseCount(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["spot_options"] = cty.ListVal(valsForCollection)
}

func EncodeEc2Fleet_SpotOptions_AllocationStrategy(p SpotOptions, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEc2Fleet_SpotOptions_InstanceInterruptionBehavior(p SpotOptions, vals map[string]cty.Value) {
	vals["instance_interruption_behavior"] = cty.StringVal(p.InstanceInterruptionBehavior)
}

func EncodeEc2Fleet_SpotOptions_InstancePoolsToUseCount(p SpotOptions, vals map[string]cty.Value) {
	vals["instance_pools_to_use_count"] = cty.NumberIntVal(p.InstancePoolsToUseCount)
}