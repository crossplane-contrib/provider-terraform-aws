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

package v1alpha1func EncodeInstance(r Instance) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeInstance_VolumeTags(r.Spec.ForProvider, ctyVal)
	EncodeInstance_CpuCoreCount(r.Spec.ForProvider, ctyVal)
	EncodeInstance_SourceDestCheck(r.Spec.ForProvider, ctyVal)
	EncodeInstance_Monitoring(r.Spec.ForProvider, ctyVal)
	EncodeInstance_PrivateIp(r.Spec.ForProvider, ctyVal)
	EncodeInstance_UserData(r.Spec.ForProvider, ctyVal)
	EncodeInstance_UserDataBase64(r.Spec.ForProvider, ctyVal)
	EncodeInstance_Ami(r.Spec.ForProvider, ctyVal)
	EncodeInstance_IamInstanceProfile(r.Spec.ForProvider, ctyVal)
	EncodeInstance_KeyName(r.Spec.ForProvider, ctyVal)
	EncodeInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeInstance_Hibernation(r.Spec.ForProvider, ctyVal)
	EncodeInstance_Ipv6AddressCount(r.Spec.ForProvider, ctyVal)
	EncodeInstance_SecondaryPrivateIps(r.Spec.ForProvider, ctyVal)
	EncodeInstance_DisableApiTermination(r.Spec.ForProvider, ctyVal)
	EncodeInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeInstance_InstanceInitiatedShutdownBehavior(r.Spec.ForProvider, ctyVal)
	EncodeInstance_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeInstance_Tenancy(r.Spec.ForProvider, ctyVal)
	EncodeInstance_GetPasswordData(r.Spec.ForProvider, ctyVal)
	EncodeInstance_HostId(r.Spec.ForProvider, ctyVal)
	EncodeInstance_AssociatePublicIpAddress(r.Spec.ForProvider, ctyVal)
	EncodeInstance_PlacementGroup(r.Spec.ForProvider, ctyVal)
	EncodeInstance_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeInstance_CpuThreadsPerCore(r.Spec.ForProvider, ctyVal)
	EncodeInstance_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeInstance_Ipv6Addresses(r.Spec.ForProvider, ctyVal)
	EncodeInstance_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeInstance_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeInstance_CreditSpecification(r.Spec.ForProvider.CreditSpecification, ctyVal)
	EncodeInstance_EbsBlockDevice(r.Spec.ForProvider.EbsBlockDevice, ctyVal)
	EncodeInstance_EphemeralBlockDevice(r.Spec.ForProvider.EphemeralBlockDevice, ctyVal)
	EncodeInstance_MetadataOptions(r.Spec.ForProvider.MetadataOptions, ctyVal)
	EncodeInstance_NetworkInterface(r.Spec.ForProvider.NetworkInterface, ctyVal)
	EncodeInstance_RootBlockDevice(r.Spec.ForProvider.RootBlockDevice, ctyVal)
	EncodeInstance_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeInstance_PasswordData(r.Status.AtProvider, ctyVal)
	EncodeInstance_InstanceState(r.Status.AtProvider, ctyVal)
	EncodeInstance_OutpostArn(r.Status.AtProvider, ctyVal)
	EncodeInstance_Arn(r.Status.AtProvider, ctyVal)
	EncodeInstance_PublicDns(r.Status.AtProvider, ctyVal)
	EncodeInstance_PrimaryNetworkInterfaceId(r.Status.AtProvider, ctyVal)
	EncodeInstance_PrivateDns(r.Status.AtProvider, ctyVal)
	EncodeInstance_PublicIp(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeInstance_Tags(p *InstanceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeInstance_VolumeTags(p *InstanceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.VolumeTags {
		mVals[key] = cty.StringVal(value)
	}
	vals["volume_tags"] = cty.MapVal(mVals)
}

func EncodeInstance_CpuCoreCount(p *InstanceParameters, vals map[string]cty.Value) {
	vals["cpu_core_count"] = cty.IntVal(p.CpuCoreCount)
}

func EncodeInstance_SourceDestCheck(p *InstanceParameters, vals map[string]cty.Value) {
	vals["source_dest_check"] = cty.BoolVal(p.SourceDestCheck)
}

func EncodeInstance_Monitoring(p *InstanceParameters, vals map[string]cty.Value) {
	vals["monitoring"] = cty.BoolVal(p.Monitoring)
}

func EncodeInstance_PrivateIp(p *InstanceParameters, vals map[string]cty.Value) {
	vals["private_ip"] = cty.StringVal(p.PrivateIp)
}

func EncodeInstance_UserData(p *InstanceParameters, vals map[string]cty.Value) {
	vals["user_data"] = cty.StringVal(p.UserData)
}

func EncodeInstance_UserDataBase64(p *InstanceParameters, vals map[string]cty.Value) {
	vals["user_data_base64"] = cty.StringVal(p.UserDataBase64)
}

func EncodeInstance_Ami(p *InstanceParameters, vals map[string]cty.Value) {
	vals["ami"] = cty.StringVal(p.Ami)
}

func EncodeInstance_IamInstanceProfile(p *InstanceParameters, vals map[string]cty.Value) {
	vals["iam_instance_profile"] = cty.StringVal(p.IamInstanceProfile)
}

func EncodeInstance_KeyName(p *InstanceParameters, vals map[string]cty.Value) {
	vals["key_name"] = cty.StringVal(p.KeyName)
}

func EncodeInstance_AvailabilityZone(p *InstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeInstance_Hibernation(p *InstanceParameters, vals map[string]cty.Value) {
	vals["hibernation"] = cty.BoolVal(p.Hibernation)
}

func EncodeInstance_Ipv6AddressCount(p *InstanceParameters, vals map[string]cty.Value) {
	vals["ipv6_address_count"] = cty.IntVal(p.Ipv6AddressCount)
}

func EncodeInstance_SecondaryPrivateIps(p *InstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecondaryPrivateIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["secondary_private_ips"] = cty.SetVal(colVals)
}

func EncodeInstance_DisableApiTermination(p *InstanceParameters, vals map[string]cty.Value) {
	vals["disable_api_termination"] = cty.BoolVal(p.DisableApiTermination)
}

func EncodeInstance_Id(p *InstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeInstance_InstanceInitiatedShutdownBehavior(p *InstanceParameters, vals map[string]cty.Value) {
	vals["instance_initiated_shutdown_behavior"] = cty.StringVal(p.InstanceInitiatedShutdownBehavior)
}

func EncodeInstance_InstanceType(p *InstanceParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeInstance_Tenancy(p *InstanceParameters, vals map[string]cty.Value) {
	vals["tenancy"] = cty.StringVal(p.Tenancy)
}

func EncodeInstance_GetPasswordData(p *InstanceParameters, vals map[string]cty.Value) {
	vals["get_password_data"] = cty.BoolVal(p.GetPasswordData)
}

func EncodeInstance_HostId(p *InstanceParameters, vals map[string]cty.Value) {
	vals["host_id"] = cty.StringVal(p.HostId)
}

func EncodeInstance_AssociatePublicIpAddress(p *InstanceParameters, vals map[string]cty.Value) {
	vals["associate_public_ip_address"] = cty.BoolVal(p.AssociatePublicIpAddress)
}

func EncodeInstance_PlacementGroup(p *InstanceParameters, vals map[string]cty.Value) {
	vals["placement_group"] = cty.StringVal(p.PlacementGroup)
}

func EncodeInstance_SubnetId(p *InstanceParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeInstance_CpuThreadsPerCore(p *InstanceParameters, vals map[string]cty.Value) {
	vals["cpu_threads_per_core"] = cty.IntVal(p.CpuThreadsPerCore)
}

func EncodeInstance_EbsOptimized(p *InstanceParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeInstance_Ipv6Addresses(p *InstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ipv6Addresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ipv6_addresses"] = cty.ListVal(colVals)
}

func EncodeInstance_VpcSecurityGroupIds(p *InstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeInstance_SecurityGroups(p *InstanceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeInstance_CreditSpecification(p *CreditSpecification, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CreditSpecification {
		ctyVal = make(map[string]cty.Value)
		EncodeInstance_CreditSpecification_CpuCredits(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["credit_specification"] = cty.ListVal(valsForCollection)
}

func EncodeInstance_CreditSpecification_CpuCredits(p *CreditSpecification, vals map[string]cty.Value) {
	vals["cpu_credits"] = cty.StringVal(p.CpuCredits)
}

func EncodeInstance_EbsBlockDevice(p *EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeInstance_EbsBlockDevice_KmsKeyId(v, ctyVal)
		EncodeInstance_EbsBlockDevice_SnapshotId(v, ctyVal)
		EncodeInstance_EbsBlockDevice_VolumeId(v, ctyVal)
		EncodeInstance_EbsBlockDevice_VolumeType(v, ctyVal)
		EncodeInstance_EbsBlockDevice_Encrypted(v, ctyVal)
		EncodeInstance_EbsBlockDevice_DeviceName(v, ctyVal)
		EncodeInstance_EbsBlockDevice_Iops(v, ctyVal)
		EncodeInstance_EbsBlockDevice_VolumeSize(v, ctyVal)
		EncodeInstance_EbsBlockDevice_DeleteOnTermination(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeInstance_EbsBlockDevice_KmsKeyId(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeInstance_EbsBlockDevice_SnapshotId(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeInstance_EbsBlockDevice_VolumeId(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeInstance_EbsBlockDevice_VolumeType(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeInstance_EbsBlockDevice_Encrypted(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeInstance_EbsBlockDevice_DeviceName(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeInstance_EbsBlockDevice_Iops(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeInstance_EbsBlockDevice_VolumeSize(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeInstance_EbsBlockDevice_DeleteOnTermination(p *EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeInstance_EphemeralBlockDevice(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EphemeralBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeInstance_EphemeralBlockDevice_DeviceName(v, ctyVal)
		EncodeInstance_EphemeralBlockDevice_NoDevice(v, ctyVal)
		EncodeInstance_EphemeralBlockDevice_VirtualName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeInstance_EphemeralBlockDevice_DeviceName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeInstance_EphemeralBlockDevice_NoDevice(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["no_device"] = cty.BoolVal(p.NoDevice)
}

func EncodeInstance_EphemeralBlockDevice_VirtualName(p *EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeInstance_MetadataOptions(p *MetadataOptions, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.MetadataOptions {
		ctyVal = make(map[string]cty.Value)
		EncodeInstance_MetadataOptions_HttpEndpoint(v, ctyVal)
		EncodeInstance_MetadataOptions_HttpPutResponseHopLimit(v, ctyVal)
		EncodeInstance_MetadataOptions_HttpTokens(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["metadata_options"] = cty.ListVal(valsForCollection)
}

func EncodeInstance_MetadataOptions_HttpEndpoint(p *MetadataOptions, vals map[string]cty.Value) {
	vals["http_endpoint"] = cty.StringVal(p.HttpEndpoint)
}

func EncodeInstance_MetadataOptions_HttpPutResponseHopLimit(p *MetadataOptions, vals map[string]cty.Value) {
	vals["http_put_response_hop_limit"] = cty.IntVal(p.HttpPutResponseHopLimit)
}

func EncodeInstance_MetadataOptions_HttpTokens(p *MetadataOptions, vals map[string]cty.Value) {
	vals["http_tokens"] = cty.StringVal(p.HttpTokens)
}

func EncodeInstance_NetworkInterface(p *NetworkInterface, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.NetworkInterface {
		ctyVal = make(map[string]cty.Value)
		EncodeInstance_NetworkInterface_DeleteOnTermination(v, ctyVal)
		EncodeInstance_NetworkInterface_DeviceIndex(v, ctyVal)
		EncodeInstance_NetworkInterface_NetworkInterfaceId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["network_interface"] = cty.SetVal(valsForCollection)
}

func EncodeInstance_NetworkInterface_DeleteOnTermination(p *NetworkInterface, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeInstance_NetworkInterface_DeviceIndex(p *NetworkInterface, vals map[string]cty.Value) {
	vals["device_index"] = cty.IntVal(p.DeviceIndex)
}

func EncodeInstance_NetworkInterface_NetworkInterfaceId(p *NetworkInterface, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeInstance_RootBlockDevice(p *RootBlockDevice, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RootBlockDevice {
		ctyVal = make(map[string]cty.Value)
		EncodeInstance_RootBlockDevice_VolumeType(v, ctyVal)
		EncodeInstance_RootBlockDevice_DeleteOnTermination(v, ctyVal)
		EncodeInstance_RootBlockDevice_DeviceName(v, ctyVal)
		EncodeInstance_RootBlockDevice_Encrypted(v, ctyVal)
		EncodeInstance_RootBlockDevice_Iops(v, ctyVal)
		EncodeInstance_RootBlockDevice_KmsKeyId(v, ctyVal)
		EncodeInstance_RootBlockDevice_VolumeId(v, ctyVal)
		EncodeInstance_RootBlockDevice_VolumeSize(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["root_block_device"] = cty.ListVal(valsForCollection)
}

func EncodeInstance_RootBlockDevice_VolumeType(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeInstance_RootBlockDevice_DeleteOnTermination(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeInstance_RootBlockDevice_DeviceName(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeInstance_RootBlockDevice_Encrypted(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeInstance_RootBlockDevice_Iops(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeInstance_RootBlockDevice_KmsKeyId(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeInstance_RootBlockDevice_VolumeId(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeInstance_RootBlockDevice_VolumeSize(p *RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.IntVal(p.VolumeSize)
}

func EncodeInstance_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeInstance_Timeouts_Create(p, ctyVal)
	EncodeInstance_Timeouts_Delete(p, ctyVal)
	EncodeInstance_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeInstance_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeInstance_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeInstance_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeInstance_PasswordData(p *InstanceObservation, vals map[string]cty.Value) {
	vals["password_data"] = cty.StringVal(p.PasswordData)
}

func EncodeInstance_InstanceState(p *InstanceObservation, vals map[string]cty.Value) {
	vals["instance_state"] = cty.StringVal(p.InstanceState)
}

func EncodeInstance_OutpostArn(p *InstanceObservation, vals map[string]cty.Value) {
	vals["outpost_arn"] = cty.StringVal(p.OutpostArn)
}

func EncodeInstance_Arn(p *InstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeInstance_PublicDns(p *InstanceObservation, vals map[string]cty.Value) {
	vals["public_dns"] = cty.StringVal(p.PublicDns)
}

func EncodeInstance_PrimaryNetworkInterfaceId(p *InstanceObservation, vals map[string]cty.Value) {
	vals["primary_network_interface_id"] = cty.StringVal(p.PrimaryNetworkInterfaceId)
}

func EncodeInstance_PrivateDns(p *InstanceObservation, vals map[string]cty.Value) {
	vals["private_dns"] = cty.StringVal(p.PrivateDns)
}

func EncodeInstance_PublicIp(p *InstanceObservation, vals map[string]cty.Value) {
	vals["public_ip"] = cty.StringVal(p.PublicIp)
}