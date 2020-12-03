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

package v1alpha1func EncodeSpotFleetRequest(r SpotFleetRequest) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSpotFleetRequest_FleetType(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_InstancePoolsToUseCount(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_ValidFrom(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_WaitForFulfillment(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_AllocationStrategy(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_TargetCapacity(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_ValidUntil(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_ExcessCapacityTerminationPolicy(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_Id(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_InstanceInterruptionBehaviour(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_ReplaceUnhealthyInstances(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_TargetGroupArns(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_TerminateInstancesWithExpiration(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_IamFleetRole(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_LoadBalancers(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_SpotPrice(r.Spec.ForProvider, ctyVal)
	EncodeSpotFleetRequest_LaunchSpecification(r.Spec.ForProvider.LaunchSpecification, ctyVal)
	EncodeSpotFleetRequest_LaunchTemplateConfig(r.Spec.ForProvider.LaunchTemplateConfig, ctyVal)
	EncodeSpotFleetRequest_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeSpotFleetRequest_ClientToken(r.Status.AtProvider, ctyVal)
	EncodeSpotFleetRequest_SpotRequestState(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSpotFleetRequest_FleetType(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["fleet_type"] = cty.StringVal(p.FleetType)
}

func EncodeSpotFleetRequest_InstancePoolsToUseCount(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["instance_pools_to_use_count"] = cty.IntVal(p.InstancePoolsToUseCount)
}

func EncodeSpotFleetRequest_Tags(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSpotFleetRequest_ValidFrom(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["valid_from"] = cty.StringVal(p.ValidFrom)
}

func EncodeSpotFleetRequest_WaitForFulfillment(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["wait_for_fulfillment"] = cty.BoolVal(p.WaitForFulfillment)
}

func EncodeSpotFleetRequest_AllocationStrategy(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["allocation_strategy"] = cty.StringVal(p.AllocationStrategy)
}

func EncodeSpotFleetRequest_TargetCapacity(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["target_capacity"] = cty.IntVal(p.TargetCapacity)
}

func EncodeSpotFleetRequest_ValidUntil(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["valid_until"] = cty.StringVal(p.ValidUntil)
}

func EncodeSpotFleetRequest_ExcessCapacityTerminationPolicy(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["excess_capacity_termination_policy"] = cty.StringVal(p.ExcessCapacityTerminationPolicy)
}

func EncodeSpotFleetRequest_Id(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSpotFleetRequest_InstanceInterruptionBehaviour(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["instance_interruption_behaviour"] = cty.StringVal(p.InstanceInterruptionBehaviour)
}

func EncodeSpotFleetRequest_ReplaceUnhealthyInstances(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["replace_unhealthy_instances"] = cty.BoolVal(p.ReplaceUnhealthyInstances)
}

func EncodeSpotFleetRequest_TargetGroupArns(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TargetGroupArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["target_group_arns"] = cty.SetVal(colVals)
}

func EncodeSpotFleetRequest_TerminateInstancesWithExpiration(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["terminate_instances_with_expiration"] = cty.BoolVal(p.TerminateInstancesWithExpiration)
}

func EncodeSpotFleetRequest_IamFleetRole(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["iam_fleet_role"] = cty.StringVal(p.IamFleetRole)
}

func EncodeSpotFleetRequest_LoadBalancers(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LoadBalancers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["load_balancers"] = cty.SetVal(colVals)
}

func EncodeSpotFleetRequest_SpotPrice(p *SpotFleetRequestParameters, vals map[string]cty.Value) {
	vals["spot_price"] = cty.StringVal(p.SpotPrice)
}

func EncodeSpotFleetRequest_LaunchSpecification(p *LaunchSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeSpotFleetRequest_LaunchSpecification_IamInstanceProfileArn(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_SpotPrice(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_UserData(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_IamInstanceProfile(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_Monitoring(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_AvailabilityZone(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsOptimized(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_PlacementGroup(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_PlacementTenancy(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_SubnetId(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_Tags(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_Ami(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_AssociatePublicIpAddress(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_VpcSecurityGroupIds(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_WeightedCapacity(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_InstanceType(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_KeyName(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice(v.EbsBlockDevice, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EphemeralBlockDevice(v.EphemeralBlockDevice, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice(v.RootBlockDevice, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_specification"] = cty.SetVal(valsForCollection)
}

func EncodeSpotFleetRequest_LaunchSpecification_IamInstanceProfileArn(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["iam_instance_profile_arn"] = cty.StringVal(p.IamInstanceProfileArn)
}

func EncodeSpotFleetRequest_LaunchSpecification_SpotPrice(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["spot_price"] = cty.StringVal(p.SpotPrice)
}

func EncodeSpotFleetRequest_LaunchSpecification_UserData(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["user_data"] = cty.StringVal(p.UserData)
}

func EncodeSpotFleetRequest_LaunchSpecification_IamInstanceProfile(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["iam_instance_profile"] = cty.StringVal(p.IamInstanceProfile)
}

func EncodeSpotFleetRequest_LaunchSpecification_Monitoring(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["monitoring"] = cty.BoolVal(p.Monitoring)
}

func EncodeSpotFleetRequest_LaunchSpecification_AvailabilityZone(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsOptimized(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeSpotFleetRequest_LaunchSpecification_PlacementGroup(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["placement_group"] = cty.StringVal(p.PlacementGroup)
}

func EncodeSpotFleetRequest_LaunchSpecification_PlacementTenancy(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["placement_tenancy"] = cty.StringVal(p.PlacementTenancy)
}

func EncodeSpotFleetRequest_LaunchSpecification_SubnetId(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeSpotFleetRequest_LaunchSpecification_Tags(p *LaunchSpecification, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSpotFleetRequest_LaunchSpecification_Ami(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["ami"] = cty.StringVal(p.Ami)
}

func EncodeSpotFleetRequest_LaunchSpecification_AssociatePublicIpAddress(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["associate_public_ip_address"] = cty.BoolVal(p.AssociatePublicIpAddress)
}

func EncodeSpotFleetRequest_LaunchSpecification_VpcSecurityGroupIds(p *LaunchSpecification, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeSpotFleetRequest_LaunchSpecification_WeightedCapacity(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.StringVal(p.WeightedCapacity)
}

func EncodeSpotFleetRequest_LaunchSpecification_InstanceType(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeSpotFleetRequest_LaunchSpecification_KeyName(p *LaunchSpecification, vals map[string]cty.Value) {
	vals["key_name"] = cty.StringVal(p.KeyName)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice(p *EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_KmsKeyId(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_SnapshotId(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_VolumeSize(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_VolumeType(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_DeleteOnTermination(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_DeviceName(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_Encrypted(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_Iops(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_KmsKeyId(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_SnapshotId(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_VolumeSize(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_VolumeType(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_DeleteOnTermination(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_DeviceName(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_Encrypted(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeSpotFleetRequest_LaunchSpecification_EbsBlockDevice_Iops(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeSpotFleetRequest_LaunchSpecification_EphemeralBlockDevice(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EphemeralBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeSpotFleetRequest_LaunchSpecification_EphemeralBlockDevice_DeviceName(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_EphemeralBlockDevice_VirtualName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeSpotFleetRequest_LaunchSpecification_EphemeralBlockDevice_DeviceName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeSpotFleetRequest_LaunchSpecification_EphemeralBlockDevice_VirtualName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice(p *RootBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RootBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_Encrypted(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_Iops(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_KmsKeyId(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_VolumeSize(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_VolumeType(v, ctyVal)
		EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_DeleteOnTermination(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["root_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_Encrypted(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_Iops(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_KmsKeyId(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_VolumeSize(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_VolumeType(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeSpotFleetRequest_LaunchSpecification_RootBlockDevice_DeleteOnTermination(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig(p *LaunchTemplateConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchTemplateConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides(v.Overrides, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification(v.LaunchTemplateSpecification, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_template_config"] = cty.SetVal(valsForCollection)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides(p *Overrides, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Overrides {
		ctyVal = make(map[string]cty.Value)
		EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_SubnetId(v, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_WeightedCapacity(v, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_AvailabilityZone(v, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_InstanceType(v, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_Priority(v, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_SpotPrice(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["overrides"] = cty.SetVal(valsForCollection)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_SubnetId(p *Overrides, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_WeightedCapacity(p *Overrides, vals map[string]cty.Value) {
	vals["weighted_capacity"] = cty.IntVal(p.WeightedCapacity)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_AvailabilityZone(p *Overrides, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_InstanceType(p *Overrides, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_Priority(p *Overrides, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_Overrides_SpotPrice(p *Overrides, vals map[string]cty.Value) {
	vals["spot_price"] = cty.StringVal(p.SpotPrice)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LaunchTemplateSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification_Id(v, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification_Name(v, ctyVal)
		EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification_Version(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["launch_template_specification"] = cty.ListVal(valsForCollection)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification_Id(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification_Name(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSpotFleetRequest_LaunchTemplateConfig_LaunchTemplateSpecification_Version(p *LaunchTemplateSpecification, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeSpotFleetRequest_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeSpotFleetRequest_Timeouts_Delete(p, ctyVal)
	EncodeSpotFleetRequest_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeSpotFleetRequest_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeSpotFleetRequest_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeSpotFleetRequest_ClientToken(p *SpotFleetRequestObservation, vals map[string]cty.Value) {
	vals["client_token"] = cty.StringVal(p.ClientToken)
}

func EncodeSpotFleetRequest_SpotRequestState(p *SpotFleetRequestObservation, vals map[string]cty.Value) {
	vals["spot_request_state"] = cty.StringVal(p.SpotRequestState)
}