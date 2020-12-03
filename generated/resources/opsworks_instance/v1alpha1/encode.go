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

package v1alpha1func EncodeOpsworksInstance(r OpsworksInstance) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeOpsworksInstance_AmiId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_CreatedAt(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_Os(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_SshHostRsaKeyFingerprint(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_State(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_ReportedOsName(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_VirtualizationType(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_AgentVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_DeleteEbs(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_EcsClusterArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_Hostname(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_LayerIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_ReportedOsFamily(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_SshKeyName(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_DeleteEip(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_InfrastructureClass(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_InstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_PublicIp(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_ElasticIp(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_Platform(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_RootDeviceType(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_Status(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_Architecture(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_AutoScalingType(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_PrivateDns(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_RootDeviceVolumeId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_RegisteredBy(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_SshHostDsaKeyFingerprint(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_LastServiceErrorId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_PublicDns(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_ReportedOsVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_PrivateIp(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_ReportedAgentVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_Tenancy(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksInstance_EbsBlockDevice(r.Spec.ForProvider.EbsBlockDevice, ctyVal)
	EncodeOpsworksInstance_EphemeralBlockDevice(r.Spec.ForProvider.EphemeralBlockDevice, ctyVal)
	EncodeOpsworksInstance_RootBlockDevice(r.Spec.ForProvider.RootBlockDevice, ctyVal)
	EncodeOpsworksInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeOpsworksInstance_Ec2InstanceId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeOpsworksInstance_AmiId(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["ami_id"] = cty.StringVal(p.AmiId)
}

func EncodeOpsworksInstance_CreatedAt(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["created_at"] = cty.StringVal(p.CreatedAt)
}

func EncodeOpsworksInstance_Os(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["os"] = cty.StringVal(p.Os)
}

func EncodeOpsworksInstance_SshHostRsaKeyFingerprint(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["ssh_host_rsa_key_fingerprint"] = cty.StringVal(p.SshHostRsaKeyFingerprint)
}

func EncodeOpsworksInstance_State(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeOpsworksInstance_ReportedOsName(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["reported_os_name"] = cty.StringVal(p.ReportedOsName)
}

func EncodeOpsworksInstance_VirtualizationType(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["virtualization_type"] = cty.StringVal(p.VirtualizationType)
}

func EncodeOpsworksInstance_AgentVersion(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["agent_version"] = cty.StringVal(p.AgentVersion)
}

func EncodeOpsworksInstance_DeleteEbs(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["delete_ebs"] = cty.BoolVal(p.DeleteEbs)
}

func EncodeOpsworksInstance_EcsClusterArn(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["ecs_cluster_arn"] = cty.StringVal(p.EcsClusterArn)
}

func EncodeOpsworksInstance_Hostname(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["hostname"] = cty.StringVal(p.Hostname)
}

func EncodeOpsworksInstance_LayerIds(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LayerIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["layer_ids"] = cty.ListVal(colVals)
}

func EncodeOpsworksInstance_ReportedOsFamily(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["reported_os_family"] = cty.StringVal(p.ReportedOsFamily)
}

func EncodeOpsworksInstance_SshKeyName(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["ssh_key_name"] = cty.StringVal(p.SshKeyName)
}

func EncodeOpsworksInstance_StackId(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksInstance_AvailabilityZone(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeOpsworksInstance_DeleteEip(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["delete_eip"] = cty.BoolVal(p.DeleteEip)
}

func EncodeOpsworksInstance_InfrastructureClass(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["infrastructure_class"] = cty.StringVal(p.InfrastructureClass)
}

func EncodeOpsworksInstance_InstanceProfileArn(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["instance_profile_arn"] = cty.StringVal(p.InstanceProfileArn)
}

func EncodeOpsworksInstance_PublicIp(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["public_ip"] = cty.StringVal(p.PublicIp)
}

func EncodeOpsworksInstance_SecurityGroupIds(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.ListVal(colVals)
}

func EncodeOpsworksInstance_ElasticIp(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["elastic_ip"] = cty.StringVal(p.ElasticIp)
}

func EncodeOpsworksInstance_InstallUpdatesOnBoot(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksInstance_Platform(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["platform"] = cty.StringVal(p.Platform)
}

func EncodeOpsworksInstance_RootDeviceType(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["root_device_type"] = cty.StringVal(p.RootDeviceType)
}

func EncodeOpsworksInstance_Status(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeOpsworksInstance_SubnetId(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeOpsworksInstance_Architecture(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["architecture"] = cty.StringVal(p.Architecture)
}

func EncodeOpsworksInstance_AutoScalingType(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["auto_scaling_type"] = cty.StringVal(p.AutoScalingType)
}

func EncodeOpsworksInstance_EbsOptimized(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeOpsworksInstance_Id(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksInstance_PrivateDns(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["private_dns"] = cty.StringVal(p.PrivateDns)
}

func EncodeOpsworksInstance_RootDeviceVolumeId(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["root_device_volume_id"] = cty.StringVal(p.RootDeviceVolumeId)
}

func EncodeOpsworksInstance_RegisteredBy(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["registered_by"] = cty.StringVal(p.RegisteredBy)
}

func EncodeOpsworksInstance_SshHostDsaKeyFingerprint(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["ssh_host_dsa_key_fingerprint"] = cty.StringVal(p.SshHostDsaKeyFingerprint)
}

func EncodeOpsworksInstance_LastServiceErrorId(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["last_service_error_id"] = cty.StringVal(p.LastServiceErrorId)
}

func EncodeOpsworksInstance_PublicDns(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["public_dns"] = cty.StringVal(p.PublicDns)
}

func EncodeOpsworksInstance_ReportedOsVersion(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["reported_os_version"] = cty.StringVal(p.ReportedOsVersion)
}

func EncodeOpsworksInstance_InstanceType(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeOpsworksInstance_PrivateIp(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["private_ip"] = cty.StringVal(p.PrivateIp)
}

func EncodeOpsworksInstance_ReportedAgentVersion(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["reported_agent_version"] = cty.StringVal(p.ReportedAgentVersion)
}

func EncodeOpsworksInstance_Tenancy(p *OpsworksInstanceParameters, vals map[string]cty.Value) {
	vals["tenancy"] = cty.StringVal(p.Tenancy)
}

func EncodeOpsworksInstance_EbsBlockDevice(p *EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeOpsworksInstance_EbsBlockDevice_SnapshotId(v, ctyVal)
		EncodeOpsworksInstance_EbsBlockDevice_VolumeSize(v, ctyVal)
		EncodeOpsworksInstance_EbsBlockDevice_VolumeType(v, ctyVal)
		EncodeOpsworksInstance_EbsBlockDevice_DeleteOnTermination(v, ctyVal)
		EncodeOpsworksInstance_EbsBlockDevice_DeviceName(v, ctyVal)
		EncodeOpsworksInstance_EbsBlockDevice_Iops(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksInstance_EbsBlockDevice_SnapshotId(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeOpsworksInstance_EbsBlockDevice_VolumeSize(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeOpsworksInstance_EbsBlockDevice_VolumeType(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeOpsworksInstance_EbsBlockDevice_DeleteOnTermination(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeOpsworksInstance_EbsBlockDevice_DeviceName(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeOpsworksInstance_EbsBlockDevice_Iops(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeOpsworksInstance_EphemeralBlockDevice(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EphemeralBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeOpsworksInstance_EphemeralBlockDevice_DeviceName(v, ctyVal)
		EncodeOpsworksInstance_EphemeralBlockDevice_VirtualName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksInstance_EphemeralBlockDevice_DeviceName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeOpsworksInstance_EphemeralBlockDevice_VirtualName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeOpsworksInstance_RootBlockDevice(p *RootBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RootBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeOpsworksInstance_RootBlockDevice_DeleteOnTermination(v, ctyVal)
		EncodeOpsworksInstance_RootBlockDevice_Iops(v, ctyVal)
		EncodeOpsworksInstance_RootBlockDevice_VolumeSize(v, ctyVal)
		EncodeOpsworksInstance_RootBlockDevice_VolumeType(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["root_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksInstance_RootBlockDevice_DeleteOnTermination(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeOpsworksInstance_RootBlockDevice_Iops(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeOpsworksInstance_RootBlockDevice_VolumeSize(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeOpsworksInstance_RootBlockDevice_VolumeType(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeOpsworksInstance_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeOpsworksInstance_Timeouts_Create(p, ctyVal)
	EncodeOpsworksInstance_Timeouts_Delete(p, ctyVal)
	EncodeOpsworksInstance_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeOpsworksInstance_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeOpsworksInstance_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeOpsworksInstance_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeOpsworksInstance_Ec2InstanceId(p *OpsworksInstanceObservation, vals map[string]cty.Value) {
	vals["ec2_instance_id"] = cty.StringVal(p.Ec2InstanceId)
}