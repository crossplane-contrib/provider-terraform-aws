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

func EncodeLaunchConfiguration(r LaunchConfiguration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchConfiguration_AssociatePublicIpAddress(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_KeyName(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_VpcClassicLinkSecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_IamInstanceProfile(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_SpotPrice(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_EnableMonitoring(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_ImageId(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_UserData(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_UserDataBase64(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_PlacementTenancy(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_VpcClassicLinkId(r.Spec.ForProvider, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice(r.Spec.ForProvider.EbsBlockDevice, ctyVal)
	EncodeLaunchConfiguration_EphemeralBlockDevice(r.Spec.ForProvider.EphemeralBlockDevice, ctyVal)
	EncodeLaunchConfiguration_RootBlockDevice(r.Spec.ForProvider.RootBlockDevice, ctyVal)
	EncodeLaunchConfiguration_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLaunchConfiguration_AssociatePublicIpAddress(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["associate_public_ip_address"] = cty.BoolVal(p.AssociatePublicIpAddress)
}

func EncodeLaunchConfiguration_KeyName(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["key_name"] = cty.StringVal(p.KeyName)
}

func EncodeLaunchConfiguration_VpcClassicLinkSecurityGroups(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcClassicLinkSecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_classic_link_security_groups"] = cty.SetVal(colVals)
}

func EncodeLaunchConfiguration_EbsOptimized(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.BoolVal(p.EbsOptimized)
}

func EncodeLaunchConfiguration_IamInstanceProfile(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["iam_instance_profile"] = cty.StringVal(p.IamInstanceProfile)
}

func EncodeLaunchConfiguration_NamePrefix(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeLaunchConfiguration_SpotPrice(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["spot_price"] = cty.StringVal(p.SpotPrice)
}

func EncodeLaunchConfiguration_EnableMonitoring(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["enable_monitoring"] = cty.BoolVal(p.EnableMonitoring)
}

func EncodeLaunchConfiguration_ImageId(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["image_id"] = cty.StringVal(p.ImageId)
}

func EncodeLaunchConfiguration_Name(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLaunchConfiguration_SecurityGroups(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeLaunchConfiguration_UserData(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["user_data"] = cty.StringVal(p.UserData)
}

func EncodeLaunchConfiguration_UserDataBase64(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["user_data_base64"] = cty.StringVal(p.UserDataBase64)
}

func EncodeLaunchConfiguration_Id(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLaunchConfiguration_InstanceType(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeLaunchConfiguration_PlacementTenancy(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["placement_tenancy"] = cty.StringVal(p.PlacementTenancy)
}

func EncodeLaunchConfiguration_VpcClassicLinkId(p LaunchConfigurationParameters, vals map[string]cty.Value) {
	vals["vpc_classic_link_id"] = cty.StringVal(p.VpcClassicLinkId)
}

func EncodeLaunchConfiguration_EbsBlockDevice(p EbsBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchConfiguration_EbsBlockDevice_DeviceName(p, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice_Encrypted(p, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice_Iops(p, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice_NoDevice(p, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice_SnapshotId(p, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice_VolumeSize(p, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice_VolumeType(p, ctyVal)
	EncodeLaunchConfiguration_EbsBlockDevice_DeleteOnTermination(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeLaunchConfiguration_EbsBlockDevice_DeviceName(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeLaunchConfiguration_EbsBlockDevice_Encrypted(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeLaunchConfiguration_EbsBlockDevice_Iops(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeLaunchConfiguration_EbsBlockDevice_NoDevice(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["no_device"] = cty.BoolVal(p.NoDevice)
}

func EncodeLaunchConfiguration_EbsBlockDevice_SnapshotId(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeLaunchConfiguration_EbsBlockDevice_VolumeSize(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeLaunchConfiguration_EbsBlockDevice_VolumeType(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeLaunchConfiguration_EbsBlockDevice_DeleteOnTermination(p EbsBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeLaunchConfiguration_EphemeralBlockDevice(p EphemeralBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchConfiguration_EphemeralBlockDevice_DeviceName(p, ctyVal)
	EncodeLaunchConfiguration_EphemeralBlockDevice_VirtualName(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ephemeral_block_device"] = cty.SetVal(valsForCollection)
}

func EncodeLaunchConfiguration_EphemeralBlockDevice_DeviceName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeLaunchConfiguration_EphemeralBlockDevice_VirtualName(p EphemeralBlockDevice, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeLaunchConfiguration_RootBlockDevice(p RootBlockDevice, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchConfiguration_RootBlockDevice_VolumeType(p, ctyVal)
	EncodeLaunchConfiguration_RootBlockDevice_DeleteOnTermination(p, ctyVal)
	EncodeLaunchConfiguration_RootBlockDevice_Encrypted(p, ctyVal)
	EncodeLaunchConfiguration_RootBlockDevice_Iops(p, ctyVal)
	EncodeLaunchConfiguration_RootBlockDevice_VolumeSize(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["root_block_device"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchConfiguration_RootBlockDevice_VolumeType(p RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeLaunchConfiguration_RootBlockDevice_DeleteOnTermination(p RootBlockDevice, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.BoolVal(p.DeleteOnTermination)
}

func EncodeLaunchConfiguration_RootBlockDevice_Encrypted(p RootBlockDevice, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeLaunchConfiguration_RootBlockDevice_Iops(p RootBlockDevice, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeLaunchConfiguration_RootBlockDevice_VolumeSize(p RootBlockDevice, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeLaunchConfiguration_Arn(p LaunchConfigurationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}