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

package v1alpha1func EncodeEmrCluster(r EmrCluster) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEmrCluster_KeepJobFlowAliveWhenNoSteps(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_LogUri(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_SecurityConfiguration(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_Applications(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_AutoscalingRole(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_ConfigurationsJson(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_EbsRootVolumeSize(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_ServiceRole(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_StepConcurrencyLevel(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_VisibleToAllUsers(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_ReleaseLabel(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_Configurations(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_CustomAmiId(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_Id(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_Name(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_ScaleDownBehavior(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_Step(r.Spec.ForProvider.Step, ctyVal)
	EncodeEmrCluster_AdditionalInfo(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_TerminationProtection(r.Spec.ForProvider, ctyVal)
	EncodeEmrCluster_BootstrapAction(r.Spec.ForProvider.BootstrapAction, ctyVal)
	EncodeEmrCluster_CoreInstanceFleet(r.Spec.ForProvider.CoreInstanceFleet, ctyVal)
	EncodeEmrCluster_CoreInstanceGroup(r.Spec.ForProvider.CoreInstanceGroup, ctyVal)
	EncodeEmrCluster_Ec2Attributes(r.Spec.ForProvider.Ec2Attributes, ctyVal)
	EncodeEmrCluster_KerberosAttributes(r.Spec.ForProvider.KerberosAttributes, ctyVal)
	EncodeEmrCluster_MasterInstanceFleet(r.Spec.ForProvider.MasterInstanceFleet, ctyVal)
	EncodeEmrCluster_MasterInstanceGroup(r.Spec.ForProvider.MasterInstanceGroup, ctyVal)
	EncodeEmrCluster_ClusterState(r.Status.AtProvider, ctyVal)
	EncodeEmrCluster_MasterPublicDns(r.Status.AtProvider, ctyVal)
	EncodeEmrCluster_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEmrCluster_KeepJobFlowAliveWhenNoSteps(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["keep_job_flow_alive_when_no_steps"] = cty.BoolVal(p.KeepJobFlowAliveWhenNoSteps)
}

func EncodeEmrCluster_LogUri(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["log_uri"] = cty.StringVal(p.LogUri)
}

func EncodeEmrCluster_SecurityConfiguration(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["security_configuration"] = cty.StringVal(p.SecurityConfiguration)
}

func EncodeEmrCluster_Applications(p *EmrClusterParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Applications {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["applications"] = cty.SetVal(colVals)
}

func EncodeEmrCluster_AutoscalingRole(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["autoscaling_role"] = cty.StringVal(p.AutoscalingRole)
}

func EncodeEmrCluster_ConfigurationsJson(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["configurations_json"] = cty.StringVal(p.ConfigurationsJson)
}

func EncodeEmrCluster_EbsRootVolumeSize(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["ebs_root_volume_size"] = cty.IntVal(p.EbsRootVolumeSize)
}

func EncodeEmrCluster_ServiceRole(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["service_role"] = cty.StringVal(p.ServiceRole)
}

func EncodeEmrCluster_StepConcurrencyLevel(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["step_concurrency_level"] = cty.IntVal(p.StepConcurrencyLevel)
}

func EncodeEmrCluster_VisibleToAllUsers(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["visible_to_all_users"] = cty.BoolVal(p.VisibleToAllUsers)
}

func EncodeEmrCluster_ReleaseLabel(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["release_label"] = cty.StringVal(p.ReleaseLabel)
}

func EncodeEmrCluster_Configurations(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["configurations"] = cty.StringVal(p.Configurations)
}

func EncodeEmrCluster_CustomAmiId(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["custom_ami_id"] = cty.StringVal(p.CustomAmiId)
}

func EncodeEmrCluster_Id(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrCluster_Name(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrCluster_ScaleDownBehavior(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["scale_down_behavior"] = cty.StringVal(p.ScaleDownBehavior)
}

func EncodeEmrCluster_Step(p *Step, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Step {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_Step_Name(v, ctyVal)
		EncodeEmrCluster_Step_ActionOnFailure(v, ctyVal)
		EncodeEmrCluster_Step_HadoopJarStep(v.HadoopJarStep, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["step"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_Step_Name(p *Step, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrCluster_Step_ActionOnFailure(p *Step, vals map[string]cty.Value) {
	vals["action_on_failure"] = cty.StringVal(p.ActionOnFailure)
}

func EncodeEmrCluster_Step_HadoopJarStep(p *HadoopJarStep, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.HadoopJarStep {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_Step_HadoopJarStep_Args(v, ctyVal)
		EncodeEmrCluster_Step_HadoopJarStep_Jar(v, ctyVal)
		EncodeEmrCluster_Step_HadoopJarStep_MainClass(v, ctyVal)
		EncodeEmrCluster_Step_HadoopJarStep_Properties(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["hadoop_jar_step"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_Step_HadoopJarStep_Args(p *HadoopJarStep, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Args {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["args"] = cty.ListVal(colVals)
}

func EncodeEmrCluster_Step_HadoopJarStep_Jar(p *HadoopJarStep, vals map[string]cty.Value) {
	vals["jar"] = cty.StringVal(p.Jar)
}

func EncodeEmrCluster_Step_HadoopJarStep_MainClass(p *HadoopJarStep, vals map[string]cty.Value) {
	vals["main_class"] = cty.StringVal(p.MainClass)
}

func EncodeEmrCluster_Step_HadoopJarStep_Properties(p *HadoopJarStep, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Properties {
		mVals[key] = cty.StringVal(value)
	}
	vals["properties"] = cty.MapVal(mVals)
}

func EncodeEmrCluster_AdditionalInfo(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["additional_info"] = cty.StringVal(p.AdditionalInfo)
}

func EncodeEmrCluster_Tags(p *EmrClusterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEmrCluster_TerminationProtection(p *EmrClusterParameters, vals map[string]cty.Value) {
	vals["termination_protection"] = cty.BoolVal(p.TerminationProtection)
}

func EncodeEmrCluster_BootstrapAction(p *BootstrapAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.BootstrapAction {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_BootstrapAction_Args(v, ctyVal)
		EncodeEmrCluster_BootstrapAction_Name(v, ctyVal)
		EncodeEmrCluster_BootstrapAction_Path(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["bootstrap_action"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_BootstrapAction_Args(p *BootstrapAction, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Args {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["args"] = cty.ListVal(colVals)
}

func EncodeEmrCluster_BootstrapAction_Name(p *BootstrapAction, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrCluster_BootstrapAction_Path(p *BootstrapAction, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeEmrCluster_CoreInstanceFleet(p *CoreInstanceFleet, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CoreInstanceFleet {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceFleet_Id(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_Name(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_ProvisionedOnDemandCapacity(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_ProvisionedSpotCapacity(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_TargetOnDemandCapacity(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_TargetSpotCapacity(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs(v.InstanceTypeConfigs, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications(v.LaunchSpecifications, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["core_instance_fleet"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceFleet_Id(p *CoreInstanceFleet, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrCluster_CoreInstanceFleet_Name(p *CoreInstanceFleet, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrCluster_CoreInstanceFleet_ProvisionedOnDemandCapacity(p *CoreInstanceFleet, vals map[string]cty.Value) {
	vals["provisioned_on_demand_capacity"] = cty.IntVal(p.ProvisionedOnDemandCapacity)
}

func EncodeEmrCluster_CoreInstanceFleet_ProvisionedSpotCapacity(p *CoreInstanceFleet, vals map[string]cty.Value) {
	vals["provisioned_spot_capacity"] = cty.IntVal(p.ProvisionedSpotCapacity)
}

func EncodeEmrCluster_CoreInstanceFleet_TargetOnDemandCapacity(p *CoreInstanceFleet, vals map[string]cty.Value) {
	vals["target_on_demand_capacity"] = cty.IntVal(p.TargetOnDemandCapacity)
}

func EncodeEmrCluster_CoreInstanceFleet_TargetSpotCapacity(p *CoreInstanceFleet, vals map[string]cty.Value) {
	vals["target_spot_capacity"] = cty.IntVal(p.TargetSpotCapacity)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InstanceTypeConfigs {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_BidPrice(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_InstanceType(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_WeightedCapacity(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_Configurations(v.Configurations, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig(v.EbsConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["instance_type_configs"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_BidPrice(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price_as_percentage_of_on_demand_price"] = cty.IntVal(p.BidPriceAsPercentageOfOnDemandPrice)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_InstanceType(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_WeightedCapacity(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.IntVal(p.WeightedCapacity)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_Configurations(p *Configurations, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Configurations {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_Configurations_Classification(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_Configurations_Properties(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["configurations"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_Configurations_Classification(p *Configurations, vals map[string]cty.Value) {
	vals["classification"] = cty.StringVal(p.Classification)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_Configurations_Properties(p *Configurations, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Properties {
		mVals[key] = cty.StringVal(value)
	}
	vals["properties"] = cty.MapVal(mVals)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig(p *EbsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(p *EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(p *EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(p *EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrCluster_CoreInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(p *EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.IntVal(p.VolumesPerInstance)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications(p *LaunchSpecifications, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchSpecifications {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_OnDemandSpecification(v.OnDemandSpecification, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification(v.SpotSpecification, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_specifications"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_OnDemandSpecification(p *OnDemandSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OnDemandSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["on_demand_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(p *OnDemandSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification(p *SpotSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SpotSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(v, ctyVal)
		EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["spot_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(p *SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_action"] = cty.StringVal(p.TimeoutAction)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(p *SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_duration_minutes"] = cty.IntVal(p.TimeoutDurationMinutes)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(p *SpotSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrCluster_CoreInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(p *SpotSpecification, vals map[string]cty.Value) {
	vals["block_duration_minutes"] = cty.IntVal(p.BlockDurationMinutes)
}

func EncodeEmrCluster_CoreInstanceGroup(p *CoreInstanceGroup, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CoreInstanceGroup {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceGroup_InstanceType(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_Name(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_AutoscalingPolicy(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_BidPrice(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_Id(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_InstanceCount(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_EbsConfig(v.EbsConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["core_instance_group"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceGroup_InstanceType(p *CoreInstanceGroup, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrCluster_CoreInstanceGroup_Name(p *CoreInstanceGroup, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrCluster_CoreInstanceGroup_AutoscalingPolicy(p *CoreInstanceGroup, vals map[string]cty.Value) {
	vals["autoscaling_policy"] = cty.StringVal(p.AutoscalingPolicy)
}

func EncodeEmrCluster_CoreInstanceGroup_BidPrice(p *CoreInstanceGroup, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrCluster_CoreInstanceGroup_Id(p *CoreInstanceGroup, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrCluster_CoreInstanceGroup_InstanceCount(p *CoreInstanceGroup, vals map[string]cty.Value) {
	vals["instance_count"] = cty.IntVal(p.InstanceCount)
}

func EncodeEmrCluster_CoreInstanceGroup_EbsConfig(p *EbsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_CoreInstanceGroup_EbsConfig_VolumesPerInstance(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_EbsConfig_Iops(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_EbsConfig_Size(v, ctyVal)
		EncodeEmrCluster_CoreInstanceGroup_EbsConfig_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_CoreInstanceGroup_EbsConfig_VolumesPerInstance(p *EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.IntVal(p.VolumesPerInstance)
}

func EncodeEmrCluster_CoreInstanceGroup_EbsConfig_Iops(p *EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeEmrCluster_CoreInstanceGroup_EbsConfig_Size(p *EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeEmrCluster_CoreInstanceGroup_EbsConfig_Type(p *EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrCluster_Ec2Attributes(p *Ec2Attributes, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Ec2Attributes {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_Ec2Attributes_EmrManagedMasterSecurityGroup(v, ctyVal)
		EncodeEmrCluster_Ec2Attributes_EmrManagedSlaveSecurityGroup(v, ctyVal)
		EncodeEmrCluster_Ec2Attributes_InstanceProfile(v, ctyVal)
		EncodeEmrCluster_Ec2Attributes_KeyName(v, ctyVal)
		EncodeEmrCluster_Ec2Attributes_ServiceAccessSecurityGroup(v, ctyVal)
		EncodeEmrCluster_Ec2Attributes_SubnetId(v, ctyVal)
		EncodeEmrCluster_Ec2Attributes_AdditionalMasterSecurityGroups(v, ctyVal)
		EncodeEmrCluster_Ec2Attributes_AdditionalSlaveSecurityGroups(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ec2_attributes"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_Ec2Attributes_EmrManagedMasterSecurityGroup(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["emr_managed_master_security_group"] = cty.StringVal(p.EmrManagedMasterSecurityGroup)
}

func EncodeEmrCluster_Ec2Attributes_EmrManagedSlaveSecurityGroup(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["emr_managed_slave_security_group"] = cty.StringVal(p.EmrManagedSlaveSecurityGroup)
}

func EncodeEmrCluster_Ec2Attributes_InstanceProfile(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["instance_profile"] = cty.StringVal(p.InstanceProfile)
}

func EncodeEmrCluster_Ec2Attributes_KeyName(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["key_name"] = cty.StringVal(p.KeyName)
}

func EncodeEmrCluster_Ec2Attributes_ServiceAccessSecurityGroup(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["service_access_security_group"] = cty.StringVal(p.ServiceAccessSecurityGroup)
}

func EncodeEmrCluster_Ec2Attributes_SubnetId(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeEmrCluster_Ec2Attributes_AdditionalMasterSecurityGroups(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["additional_master_security_groups"] = cty.StringVal(p.AdditionalMasterSecurityGroups)
}

func EncodeEmrCluster_Ec2Attributes_AdditionalSlaveSecurityGroups(p *Ec2Attributes, vals map[string]cty.Value) {
	vals["additional_slave_security_groups"] = cty.StringVal(p.AdditionalSlaveSecurityGroups)
}

func EncodeEmrCluster_KerberosAttributes(p *KerberosAttributes, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.KerberosAttributes {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_KerberosAttributes_AdDomainJoinPassword(v, ctyVal)
		EncodeEmrCluster_KerberosAttributes_AdDomainJoinUser(v, ctyVal)
		EncodeEmrCluster_KerberosAttributes_CrossRealmTrustPrincipalPassword(v, ctyVal)
		EncodeEmrCluster_KerberosAttributes_KdcAdminPassword(v, ctyVal)
		EncodeEmrCluster_KerberosAttributes_Realm(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["kerberos_attributes"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_KerberosAttributes_AdDomainJoinPassword(p *KerberosAttributes, vals map[string]cty.Value) {
	vals["ad_domain_join_password"] = cty.StringVal(p.AdDomainJoinPassword)
}

func EncodeEmrCluster_KerberosAttributes_AdDomainJoinUser(p *KerberosAttributes, vals map[string]cty.Value) {
	vals["ad_domain_join_user"] = cty.StringVal(p.AdDomainJoinUser)
}

func EncodeEmrCluster_KerberosAttributes_CrossRealmTrustPrincipalPassword(p *KerberosAttributes, vals map[string]cty.Value) {
	vals["cross_realm_trust_principal_password"] = cty.StringVal(p.CrossRealmTrustPrincipalPassword)
}

func EncodeEmrCluster_KerberosAttributes_KdcAdminPassword(p *KerberosAttributes, vals map[string]cty.Value) {
	vals["kdc_admin_password"] = cty.StringVal(p.KdcAdminPassword)
}

func EncodeEmrCluster_KerberosAttributes_Realm(p *KerberosAttributes, vals map[string]cty.Value) {
	vals["realm"] = cty.StringVal(p.Realm)
}

func EncodeEmrCluster_MasterInstanceFleet(p *MasterInstanceFleet, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MasterInstanceFleet {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceFleet_Id(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_Name(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_ProvisionedOnDemandCapacity(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_ProvisionedSpotCapacity(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_TargetOnDemandCapacity(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_TargetSpotCapacity(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs(v.InstanceTypeConfigs, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications(v.LaunchSpecifications, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["master_instance_fleet"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceFleet_Id(p *MasterInstanceFleet, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrCluster_MasterInstanceFleet_Name(p *MasterInstanceFleet, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrCluster_MasterInstanceFleet_ProvisionedOnDemandCapacity(p *MasterInstanceFleet, vals map[string]cty.Value) {
	vals["provisioned_on_demand_capacity"] = cty.IntVal(p.ProvisionedOnDemandCapacity)
}

func EncodeEmrCluster_MasterInstanceFleet_ProvisionedSpotCapacity(p *MasterInstanceFleet, vals map[string]cty.Value) {
	vals["provisioned_spot_capacity"] = cty.IntVal(p.ProvisionedSpotCapacity)
}

func EncodeEmrCluster_MasterInstanceFleet_TargetOnDemandCapacity(p *MasterInstanceFleet, vals map[string]cty.Value) {
	vals["target_on_demand_capacity"] = cty.IntVal(p.TargetOnDemandCapacity)
}

func EncodeEmrCluster_MasterInstanceFleet_TargetSpotCapacity(p *MasterInstanceFleet, vals map[string]cty.Value) {
	vals["target_spot_capacity"] = cty.IntVal(p.TargetSpotCapacity)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.InstanceTypeConfigs {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_InstanceType(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_WeightedCapacity(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_BidPrice(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_Configurations(v.Configurations, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig(v.EbsConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["instance_type_configs"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_BidPriceAsPercentageOfOnDemandPrice(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price_as_percentage_of_on_demand_price"] = cty.IntVal(p.BidPriceAsPercentageOfOnDemandPrice)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_InstanceType(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_WeightedCapacity(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.IntVal(p.WeightedCapacity)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_BidPrice(p *InstanceTypeConfigs, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_Configurations(p *Configurations, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Configurations {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_Configurations_Classification(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_Configurations_Properties(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["configurations"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_Configurations_Classification(p *Configurations, vals map[string]cty.Value) {
	vals["classification"] = cty.StringVal(p.Classification)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_Configurations_Properties(p *Configurations, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Properties {
		mVals[key] = cty.StringVal(value)
	}
	vals["properties"] = cty.MapVal(mVals)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig(p *EbsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_VolumesPerInstance(p *EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.IntVal(p.VolumesPerInstance)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_Iops(p *EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_Size(p *EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeEmrCluster_MasterInstanceFleet_InstanceTypeConfigs_EbsConfig_Type(p *EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications(p *LaunchSpecifications, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchSpecifications {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_OnDemandSpecification(v.OnDemandSpecification, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification(v.SpotSpecification, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_specifications"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_OnDemandSpecification(p *OnDemandSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OnDemandSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["on_demand_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_OnDemandSpecification_AllocationStrategy(p *OnDemandSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification(p *SpotSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SpotSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(v, ctyVal)
		EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["spot_specification"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_BlockDurationMinutes(p *SpotSpecification, vals map[string]cty.Value) {
	vals["block_duration_minutes"] = cty.IntVal(p.BlockDurationMinutes)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutAction(p *SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_action"] = cty.StringVal(p.TimeoutAction)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_TimeoutDurationMinutes(p *SpotSpecification, vals map[string]cty.Value) {
	vals["timeout_duration_minutes"] = cty.IntVal(p.TimeoutDurationMinutes)
}

func EncodeEmrCluster_MasterInstanceFleet_LaunchSpecifications_SpotSpecification_AllocationStrategy(p *SpotSpecification, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeEmrCluster_MasterInstanceGroup(p *MasterInstanceGroup, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MasterInstanceGroup {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceGroup_BidPrice(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_Id(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_InstanceCount(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_InstanceType(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_Name(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_EbsConfig(v.EbsConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["master_instance_group"] = cty.ListVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceGroup_BidPrice(p *MasterInstanceGroup, vals map[string]cty.Value) {
	vals["bid_price"] = cty.StringVal(p.BidPrice)
}

func EncodeEmrCluster_MasterInstanceGroup_Id(p *MasterInstanceGroup, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrCluster_MasterInstanceGroup_InstanceCount(p *MasterInstanceGroup, vals map[string]cty.Value) {
	vals["instance_count"] = cty.IntVal(p.InstanceCount)
}

func EncodeEmrCluster_MasterInstanceGroup_InstanceType(p *MasterInstanceGroup, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeEmrCluster_MasterInstanceGroup_Name(p *MasterInstanceGroup, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrCluster_MasterInstanceGroup_EbsConfig(p *EbsConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeEmrCluster_MasterInstanceGroup_EbsConfig_Iops(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_EbsConfig_Size(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_EbsConfig_Type(v, ctyVal)
		EncodeEmrCluster_MasterInstanceGroup_EbsConfig_VolumesPerInstance(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_config"] = cty.SetVal(valsForCollection)
}

func EncodeEmrCluster_MasterInstanceGroup_EbsConfig_Iops(p *EbsConfig, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeEmrCluster_MasterInstanceGroup_EbsConfig_Size(p *EbsConfig, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeEmrCluster_MasterInstanceGroup_EbsConfig_Type(p *EbsConfig, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeEmrCluster_MasterInstanceGroup_EbsConfig_VolumesPerInstance(p *EbsConfig, vals map[string]cty.Value) {
	vals["volumes_per_instance"] = cty.IntVal(p.VolumesPerInstance)
}

func EncodeEmrCluster_ClusterState(p *EmrClusterObservation, vals map[string]cty.Value) {
	vals["cluster_state"] = cty.StringVal(p.ClusterState)
}

func EncodeEmrCluster_MasterPublicDns(p *EmrClusterObservation, vals map[string]cty.Value) {
	vals["master_public_dns"] = cty.StringVal(p.MasterPublicDns)
}

func EncodeEmrCluster_Arn(p *EmrClusterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}