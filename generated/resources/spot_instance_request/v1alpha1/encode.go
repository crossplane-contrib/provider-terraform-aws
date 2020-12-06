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

func EncodeSpotInstanceRequest(r SpotInstanceRequest) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_SecondaryPrivateIps(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_SpotType(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_UserDataBase64(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_CpuCoreCount(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_GetPasswordData(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Ipv6AddressCount(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_VolumeTags(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Hibernation(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Tenancy(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_UserData(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_CpuThreadsPerCore(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Id(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_PlacementGroup(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_ValidUntil(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_WaitForFulfillment(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_SourceDestCheck(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Ami(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_AssociatePublicIpAddress(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Ipv6Addresses(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_DisableApiTermination(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_PrivateIp(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_SpotPrice(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_InstanceInterruptionBehaviour(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Monitoring(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_ValidFrom(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_BlockDurationMinutes(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_HostId(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_IamInstanceProfile(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_KeyName(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_LaunchGroup(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_InstanceInitiatedShutdownBehavior(r.Spec.ForProvider, ctyVal)
	EncodeSpotInstanceRequest_NetworkInterface(r.Spec.ForProvider.NetworkInterface, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice(r.Spec.ForProvider.RootBlockDevice, ctyVal)
	EncodeSpotInstanceRequest_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeSpotInstanceRequest_CreditSpecification(r.Spec.ForProvider.CreditSpecification, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice(r.Spec.ForProvider.EbsBlockDevice, ctyVal)
	EncodeSpotInstanceRequest_EphemeralBlockDevice(r.Spec.ForProvider.EphemeralBlockDevice, ctyVal)
	EncodeSpotInstanceRequest_MetadataOptions(r.Spec.ForProvider.MetadataOptions, ctyVal)
	EncodeSpotInstanceRequest_PrimaryNetworkInterfaceId(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_PrivateDns(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_SpotInstanceId(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_PublicIp(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_Arn(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_PasswordData(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_OutpostArn(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_SpotBidStatus(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_InstanceState(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_SpotRequestState(r.Status.AtProvider, ctyVal)
	EncodeSpotInstanceRequest_PublicDns(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSpotInstanceRequest_SecondaryPrivateIps(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecondaryPrivateIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["secondary_private_ips"] = cty.SetVal(colVals)
}

func EncodeSpotInstanceRequest_SpotType(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["spot_type"] = cty.StringVal(p.SpotType)
}

func EncodeSpotInstanceRequest_UserDataBase64(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["user_data_base64"] = cty.StringVal(p.UserDataBase64)
}

func EncodeSpotInstanceRequest_CpuCoreCount(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["cpu_core_count"] = cty.NumberIntVal(p.CpuCoreCount)
}

func EncodeSpotInstanceRequest_GetPasswordData(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["get_password_data"] = cty.BoolVal(p.GetPasswordData)
}

func EncodeSpotInstanceRequest_Ipv6AddressCount(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["ipv6_address_count"] = cty.NumberIntVal(p.Ipv6AddressCount)
}

func EncodeSpotInstanceRequest_VolumeTags(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.VolumeTags {
		mVals[key] = cty.StringVal(value)
	}
	vals["volume_tags"] = cty.MapVal(mVals)
}

func EncodeSpotInstanceRequest_Hibernation(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["hibernation"] = cty.BoolVal(p.Hibernation)
}

func EncodeSpotInstanceRequest_Tenancy(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["tenancy"] = cty.StringVal(p.Tenancy)
}

func EncodeSpotInstanceRequest_UserData(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["user_data"] = cty.StringVal(p.UserData)
}

func EncodeSpotInstanceRequest_CpuThreadsPerCore(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["cpu_threads_per_core"] = cty.NumberIntVal(p.CpuThreadsPerCore)
}

func EncodeSpotInstanceRequest_EbsOptimized(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeSpotInstanceRequest_Id(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSpotInstanceRequest_PlacementGroup(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["placement_group"] = cty.StringVal(p.PlacementGroup)
}

func EncodeSpotInstanceRequest_ValidUntil(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["valid_until"] = cty.StringVal(p.ValidUntil)
}

func EncodeSpotInstanceRequest_VpcSecurityGroupIds(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeSpotInstanceRequest_WaitForFulfillment(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["wait_for_fulfillment"] = cty.BoolVal(p.WaitForFulfillment)
}

func EncodeSpotInstanceRequest_SecurityGroups(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeSpotInstanceRequest_SourceDestCheck(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["source_dest_check"] = cty.BoolVal(p.SourceDestCheck)
}

func EncodeSpotInstanceRequest_Ami(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["ami"] = cty.StringVal(p.Ami)
}

func EncodeSpotInstanceRequest_AssociatePublicIpAddress(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["associate_public_ip_address"] = cty.BoolVal(p.AssociatePublicIpAddress)
}

func EncodeSpotInstanceRequest_AvailabilityZone(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeSpotInstanceRequest_Ipv6Addresses(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ipv6Addresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ipv6_addresses"] = cty.ListVal(colVals)
}

func EncodeSpotInstanceRequest_DisableApiTermination(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["disable_api_termination"] = cty.BoolVal(p.DisableApiTermination)
}

func EncodeSpotInstanceRequest_InstanceType(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeSpotInstanceRequest_PrivateIp(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["private_ip"] = cty.StringVal(p.PrivateIp)
}

func EncodeSpotInstanceRequest_SpotPrice(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["spot_price"] = cty.StringVal(p.SpotPrice)
}

func EncodeSpotInstanceRequest_InstanceInterruptionBehaviour(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["instance_interruption_behaviour"] = cty.StringVal(p.InstanceInterruptionBehaviour)
}

func EncodeSpotInstanceRequest_Monitoring(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["monitoring"] = cty.BoolVal(p.Monitoring)
}

func EncodeSpotInstanceRequest_Tags(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSpotInstanceRequest_SubnetId(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeSpotInstanceRequest_ValidFrom(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["valid_from"] = cty.StringVal(p.ValidFrom)
}

func EncodeSpotInstanceRequest_BlockDurationMinutes(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["block_duration_minutes"] = cty.NumberIntVal(p.BlockDurationMinutes)
}

func EncodeSpotInstanceRequest_HostId(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["host_id"] = cty.StringVal(p.HostId)
}

func EncodeSpotInstanceRequest_IamInstanceProfile(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["iam_instance_profile"] = cty.StringVal(p.IamInstanceProfile)
}

func EncodeSpotInstanceRequest_KeyName(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["key_name"] = cty.StringVal(p.KeyName)
}

func EncodeSpotInstanceRequest_LaunchGroup(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["launch_group"] = cty.StringVal(p.LaunchGroup)
}

func EncodeSpotInstanceRequest_InstanceInitiatedShutdownBehavior(p SpotInstanceRequestParameters, vals map[string]cty.Value) {
	vals["instance_initiated_shutdown_behavior"] = cty.StringVal(p.InstanceInitiatedShutdownBehavior)
}

func EncodeSpotInstanceRequest_NetworkInterface(p NetworkInterface, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_NetworkInterface_DeleteOnTermination(p, ctyVal)
	EncodeSpotInstanceRequest_NetworkInterface_DeviceIndex(p, ctyVal)
	EncodeSpotInstanceRequest_NetworkInterface_NetworkInterfaceId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["network_interface"] = cty.SetVal(valsForCollection)
}

func EncodeSpotInstanceRequest_NetworkInterface_DeleteOnTermination(p NetworkInterface, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeSpotInstanceRequest_NetworkInterface_DeviceIndex(p NetworkInterface, vals map[string]cty.Value) {
	vals["device_index"] = cty.NumberIntVal(p.DeviceIndex)
}

func EncodeSpotInstanceRequest_NetworkInterface_NetworkInterfaceId(p NetworkInterface, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeSpotInstanceRequest_RootBlockDevice(p RootBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_RootBlockDevice_Encrypted(p, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice_Iops(p, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice_KmsKeyId(p, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice_VolumeId(p, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice_VolumeSize(p, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice_VolumeType(p, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice_DeleteOnTermination(p, ctyVal)
	EncodeSpotInstanceRequest_RootBlockDevice_DeviceName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["root_block_device"] = cty.ListVal(valsForCollection)
}

func EncodeSpotInstanceRequest_RootBlockDevice_Encrypted(p RootBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeSpotInstanceRequest_RootBlockDevice_Iops(p RootBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeSpotInstanceRequest_RootBlockDevice_KmsKeyId(p RootBlockDevice, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeSpotInstanceRequest_RootBlockDevice_VolumeId(p RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeSpotInstanceRequest_RootBlockDevice_VolumeSize(p RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeSpotInstanceRequest_RootBlockDevice_VolumeType(p RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeSpotInstanceRequest_RootBlockDevice_DeleteOnTermination(p RootBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeSpotInstanceRequest_RootBlockDevice_DeviceName(p RootBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeSpotInstanceRequest_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_Timeouts_Create(p, ctyVal)
	EncodeSpotInstanceRequest_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeSpotInstanceRequest_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeSpotInstanceRequest_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeSpotInstanceRequest_CreditSpecification(p CreditSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_CreditSpecification_CpuCredits(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["credit_specification"] = cty.ListVal(valsForCollection)
}

func EncodeSpotInstanceRequest_CreditSpecification_CpuCredits(p CreditSpecification, vals map[string]cty.Value) {
	vals["cpu_credits"] = cty.StringVal(p.CpuCredits)
}

func EncodeSpotInstanceRequest_EbsBlockDevice(p EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_EbsBlockDevice_Encrypted(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_Iops(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_SnapshotId(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_VolumeSize(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_VolumeType(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_DeleteOnTermination(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_DeviceName(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_KmsKeyId(p, ctyVal)
	EncodeSpotInstanceRequest_EbsBlockDevice_VolumeId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_Encrypted(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_Iops(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_SnapshotId(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_VolumeSize(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_VolumeType(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_DeleteOnTermination(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_DeviceName(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_KmsKeyId(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeSpotInstanceRequest_EbsBlockDevice_VolumeId(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_id"] = cty.StringVal(p.VolumeId)
}

func EncodeSpotInstanceRequest_EphemeralBlockDevice(p EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_EphemeralBlockDevice_DeviceName(p, ctyVal)
	EncodeSpotInstanceRequest_EphemeralBlockDevice_NoDevice(p, ctyVal)
	EncodeSpotInstanceRequest_EphemeralBlockDevice_VirtualName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeSpotInstanceRequest_EphemeralBlockDevice_DeviceName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeSpotInstanceRequest_EphemeralBlockDevice_NoDevice(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["no_device"] = cty.BoolVal(p.NoDevice)
}

func EncodeSpotInstanceRequest_EphemeralBlockDevice_VirtualName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeSpotInstanceRequest_MetadataOptions(p MetadataOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeSpotInstanceRequest_MetadataOptions_HttpEndpoint(p, ctyVal)
	EncodeSpotInstanceRequest_MetadataOptions_HttpPutResponseHopLimit(p, ctyVal)
	EncodeSpotInstanceRequest_MetadataOptions_HttpTokens(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["metadata_options"] = cty.ListVal(valsForCollection)
}

func EncodeSpotInstanceRequest_MetadataOptions_HttpEndpoint(p MetadataOptions, vals map[string]cty.Value) {
	vals["http_endpoint"] = cty.StringVal(p.HttpEndpoint)
}

func EncodeSpotInstanceRequest_MetadataOptions_HttpPutResponseHopLimit(p MetadataOptions, vals map[string]cty.Value) {
	vals["http_put_response_hop_limit"] = cty.NumberIntVal(p.HttpPutResponseHopLimit)
}

func EncodeSpotInstanceRequest_MetadataOptions_HttpTokens(p MetadataOptions, vals map[string]cty.Value) {
	vals["http_tokens"] = cty.StringVal(p.HttpTokens)
}

func EncodeSpotInstanceRequest_PrimaryNetworkInterfaceId(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["primary_network_interface_id"] = cty.StringVal(p.PrimaryNetworkInterfaceId)
}

func EncodeSpotInstanceRequest_PrivateDns(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["private_dns"] = cty.StringVal(p.PrivateDns)
}

func EncodeSpotInstanceRequest_SpotInstanceId(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["spot_instance_id"] = cty.StringVal(p.SpotInstanceId)
}

func EncodeSpotInstanceRequest_PublicIp(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["public_ip"] = cty.StringVal(p.PublicIp)
}

func EncodeSpotInstanceRequest_Arn(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeSpotInstanceRequest_PasswordData(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["password_data"] = cty.StringVal(p.PasswordData)
}

func EncodeSpotInstanceRequest_OutpostArn(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["outpost_arn"] = cty.StringVal(p.OutpostArn)
}

func EncodeSpotInstanceRequest_SpotBidStatus(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["spot_bid_status"] = cty.StringVal(p.SpotBidStatus)
}

func EncodeSpotInstanceRequest_InstanceState(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["instance_state"] = cty.StringVal(p.InstanceState)
}

func EncodeSpotInstanceRequest_SpotRequestState(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["spot_request_state"] = cty.StringVal(p.SpotRequestState)
}

func EncodeSpotInstanceRequest_PublicDns(p SpotInstanceRequestObservation, vals map[string]cty.Value) {
	vals["public_dns"] = cty.StringVal(p.PublicDns)
}