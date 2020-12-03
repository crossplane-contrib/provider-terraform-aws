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

package v1alpha1func EncodeAutoscalingGroup(r AutoscalingGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAutoscalingGroup_MaxInstanceLifetime(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_MaxSize(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_TargetGroupArns(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_VpcZoneIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_WaitForCapacityTimeout(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_DesiredCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_MinSize(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_TerminationPolicies(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_EnabledMetrics(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_AvailabilityZones(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_ProtectFromScaleIn(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_SuspendedProcesses(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_WaitForElbCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_LoadBalancers(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_HealthCheckGracePeriod(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_HealthCheckType(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_MinElbCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_DefaultCooldown(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_MetricsGranularity(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_LaunchConfiguration(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_ForceDelete(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_PlacementGroup(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_ServiceLinkedRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingGroup_InitialLifecycleHook(r.Spec.ForProvider.InitialLifecycleHook, ctyVal)
	EncodeAutoscalingGroup_LaunchTemplate(r.Spec.ForProvider.LaunchTemplate, ctyVal)
	EncodeAutoscalingGroup_MixedInstancesPolicy(r.Spec.ForProvider.MixedInstancesPolicy, ctyVal)
	EncodeAutoscalingGroup_Tag(r.Spec.ForProvider.Tag, ctyVal)
	EncodeAutoscalingGroup_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeAutoscalingGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAutoscalingGroup_MaxInstanceLifetime(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["max_instance_lifetime"] = cty.IntVal(p.MaxInstanceLifetime)
}

func EncodeAutoscalingGroup_MaxSize(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["max_size"] = cty.IntVal(p.MaxSize)
}

func EncodeAutoscalingGroup_Name(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAutoscalingGroup_TargetGroupArns(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TargetGroupArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["target_group_arns"] = cty.SetVal(colVals)
}

func EncodeAutoscalingGroup_VpcZoneIdentifier(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcZoneIdentifier {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_zone_identifier"] = cty.SetVal(colVals)
}

func EncodeAutoscalingGroup_WaitForCapacityTimeout(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["wait_for_capacity_timeout"] = cty.StringVal(p.WaitForCapacityTimeout)
}

func EncodeAutoscalingGroup_DesiredCapacity(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["desired_capacity"] = cty.IntVal(p.DesiredCapacity)
}

func EncodeAutoscalingGroup_MinSize(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["min_size"] = cty.IntVal(p.MinSize)
}

func EncodeAutoscalingGroup_TerminationPolicies(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TerminationPolicies {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["termination_policies"] = cty.ListVal(colVals)
}

func EncodeAutoscalingGroup_EnabledMetrics(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EnabledMetrics {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["enabled_metrics"] = cty.SetVal(colVals)
}

func EncodeAutoscalingGroup_NamePrefix(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeAutoscalingGroup_AvailabilityZones(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeAutoscalingGroup_ProtectFromScaleIn(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["protect_from_scale_in"] = cty.BoolVal(p.ProtectFromScaleIn)
}

func EncodeAutoscalingGroup_SuspendedProcesses(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SuspendedProcesses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["suspended_processes"] = cty.SetVal(colVals)
}

func EncodeAutoscalingGroup_WaitForElbCapacity(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["wait_for_elb_capacity"] = cty.IntVal(p.WaitForElbCapacity)
}

func EncodeAutoscalingGroup_LoadBalancers(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LoadBalancers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["load_balancers"] = cty.SetVal(colVals)
}

func EncodeAutoscalingGroup_HealthCheckGracePeriod(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["health_check_grace_period"] = cty.IntVal(p.HealthCheckGracePeriod)
}

func EncodeAutoscalingGroup_HealthCheckType(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["health_check_type"] = cty.StringVal(p.HealthCheckType)
}

func EncodeAutoscalingGroup_Id(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAutoscalingGroup_MinElbCapacity(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["min_elb_capacity"] = cty.IntVal(p.MinElbCapacity)
}

func EncodeAutoscalingGroup_DefaultCooldown(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["default_cooldown"] = cty.IntVal(p.DefaultCooldown)
}

func EncodeAutoscalingGroup_MetricsGranularity(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["metrics_granularity"] = cty.StringVal(p.MetricsGranularity)
}

func EncodeAutoscalingGroup_LaunchConfiguration(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["launch_configuration"] = cty.StringVal(p.LaunchConfiguration)
}

func EncodeAutoscalingGroup_ForceDelete(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["force_delete"] = cty.BoolVal(p.ForceDelete)
}

func EncodeAutoscalingGroup_PlacementGroup(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["placement_group"] = cty.StringVal(p.PlacementGroup)
}

func EncodeAutoscalingGroup_ServiceLinkedRoleArn(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	vals["service_linked_role_arn"] = cty.StringVal(p.ServiceLinkedRoleArn)
}

func EncodeAutoscalingGroup_Tags(p *AutoscalingGroupParameters, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Tags {
		ctyVal = make(map[string]cty.Value)

		valsForCollection = append(valsForCollection, cty.MapVal(ctyVal))
	}
	vals["tags"] = cty.SetVal(valsForCollection)
}

func EncodeAutoscalingGroup_InitialLifecycleHook(p *InitialLifecycleHook, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InitialLifecycleHook {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_InitialLifecycleHook_HeartbeatTimeout(v, ctyVal)
		EncodeAutoscalingGroup_InitialLifecycleHook_LifecycleTransition(v, ctyVal)
		EncodeAutoscalingGroup_InitialLifecycleHook_Name(v, ctyVal)
		EncodeAutoscalingGroup_InitialLifecycleHook_NotificationMetadata(v, ctyVal)
		EncodeAutoscalingGroup_InitialLifecycleHook_NotificationTargetArn(v, ctyVal)
		EncodeAutoscalingGroup_InitialLifecycleHook_RoleArn(v, ctyVal)
		EncodeAutoscalingGroup_InitialLifecycleHook_DefaultResult(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["initial_lifecycle_hook"] = cty.SetVal(valsForCollection)
}

func EncodeAutoscalingGroup_InitialLifecycleHook_HeartbeatTimeout(p *InitialLifecycleHook, vals map[string]cty.Value) {
	vals["heartbeat_timeout"] = cty.IntVal(p.HeartbeatTimeout)
}

func EncodeAutoscalingGroup_InitialLifecycleHook_LifecycleTransition(p *InitialLifecycleHook, vals map[string]cty.Value) {
	vals["lifecycle_transition"] = cty.StringVal(p.LifecycleTransition)
}

func EncodeAutoscalingGroup_InitialLifecycleHook_Name(p *InitialLifecycleHook, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAutoscalingGroup_InitialLifecycleHook_NotificationMetadata(p *InitialLifecycleHook, vals map[string]cty.Value) {
	vals["notification_metadata"] = cty.StringVal(p.NotificationMetadata)
}

func EncodeAutoscalingGroup_InitialLifecycleHook_NotificationTargetArn(p *InitialLifecycleHook, vals map[string]cty.Value) {
	vals["notification_target_arn"] = cty.StringVal(p.NotificationTargetArn)
}

func EncodeAutoscalingGroup_InitialLifecycleHook_RoleArn(p *InitialLifecycleHook, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeAutoscalingGroup_InitialLifecycleHook_DefaultResult(p *InitialLifecycleHook, vals map[string]cty.Value) {
	vals["default_result"] = cty.StringVal(p.DefaultResult)
}

func EncodeAutoscalingGroup_LaunchTemplate(p *LaunchTemplate, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchTemplate {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_LaunchTemplate_Id(v, ctyVal)
		EncodeAutoscalingGroup_LaunchTemplate_Name(v, ctyVal)
		EncodeAutoscalingGroup_LaunchTemplate_Version(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_template"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingGroup_LaunchTemplate_Id(p *LaunchTemplate, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAutoscalingGroup_LaunchTemplate_Name(p *LaunchTemplate, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAutoscalingGroup_LaunchTemplate_Version(p *LaunchTemplate, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy(p *MixedInstancesPolicy, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MixedInstancesPolicy {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution(v.InstancesDistribution, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate(v.LaunchTemplate, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["mixed_instances_policy"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution(p *InstancesDistribution, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InstancesDistribution {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_OnDemandAllocationStrategy(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_OnDemandBaseCapacity(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_OnDemandPercentageAboveBaseCapacity(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_SpotAllocationStrategy(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_SpotInstancePools(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_SpotMaxPrice(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["instances_distribution"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_OnDemandAllocationStrategy(p *InstancesDistribution, vals map[string]cty.Value) {
	vals["on_demand_allocation_strategy"] = cty.StringVal(p.OnDemandAllocationStrategy)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_OnDemandBaseCapacity(p *InstancesDistribution, vals map[string]cty.Value) {
	vals["on_demand_base_capacity"] = cty.IntVal(p.OnDemandBaseCapacity)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_OnDemandPercentageAboveBaseCapacity(p *InstancesDistribution, vals map[string]cty.Value) {
	vals["on_demand_percentage_above_base_capacity"] = cty.IntVal(p.OnDemandPercentageAboveBaseCapacity)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_SpotAllocationStrategy(p *InstancesDistribution, vals map[string]cty.Value) {
	vals["spot_allocation_strategy"] = cty.StringVal(p.SpotAllocationStrategy)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_SpotInstancePools(p *InstancesDistribution, vals map[string]cty.Value) {
	vals["spot_instance_pools"] = cty.IntVal(p.SpotInstancePools)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_InstancesDistribution_SpotMaxPrice(p *InstancesDistribution, vals map[string]cty.Value) {
	vals["spot_max_price"] = cty.StringVal(p.SpotMaxPrice)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate(p *LaunchTemplate, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchTemplate {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification(v.LaunchTemplateSpecification, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_Override(v.Override, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_template"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchTemplateSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification_LaunchTemplateId(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification_LaunchTemplateName(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification_Version(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_template_specification"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification_LaunchTemplateId(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["launch_template_id"] = cty.StringVal(p.LaunchTemplateId)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification_LaunchTemplateName(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["launch_template_name"] = cty.StringVal(p.LaunchTemplateName)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_LaunchTemplateSpecification_Version(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_Override(p *Override, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Override {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_Override_InstanceType(v, ctyVal)
		EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_Override_WeightedCapacity(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["override"] = cty.ListVal(valsForCollection)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_Override_InstanceType(p *Override, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeAutoscalingGroup_MixedInstancesPolicy_LaunchTemplate_Override_WeightedCapacity(p *Override, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.StringVal(p.WeightedCapacity)
}

func EncodeAutoscalingGroup_Tag(p *Tag, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Tag {
		ctyVal = make(map[string]cty.Value)
		EncodeAutoscalingGroup_Tag_Key(v, ctyVal)
		EncodeAutoscalingGroup_Tag_PropagateAtLaunch(v, ctyVal)
		EncodeAutoscalingGroup_Tag_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["tag"] = cty.SetVal(valsForCollection)
}

func EncodeAutoscalingGroup_Tag_Key(p *Tag, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeAutoscalingGroup_Tag_PropagateAtLaunch(p *Tag, vals map[string]cty.Value) {
	vals["propagate_at_launch"] = cty.BoolVal(p.PropagateAtLaunch)
}

func EncodeAutoscalingGroup_Tag_Value(p *Tag, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeAutoscalingGroup_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeAutoscalingGroup_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeAutoscalingGroup_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeAutoscalingGroup_Arn(p *AutoscalingGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}