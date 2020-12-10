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
	"fmt"
	
	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*LaunchTemplate)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LaunchTemplate.")
	}
	return EncodeLaunchTemplate(*r), nil
}

func EncodeLaunchTemplate(r LaunchTemplate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_Description(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_ImageId(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_InstanceInitiatedShutdownBehavior(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_InstanceType(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_KeyName(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_RamDiskId(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_UpdateDefaultVersion(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_UserData(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_DisableApiTermination(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_EbsOptimized(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_KernelId(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_Name(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_DefaultVersion(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_Id(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_SecurityGroupNames(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_Tags(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeLaunchTemplate_HibernationOptions(r.Spec.ForProvider.HibernationOptions, ctyVal)
	EncodeLaunchTemplate_InstanceMarketOptions(r.Spec.ForProvider.InstanceMarketOptions, ctyVal)
	EncodeLaunchTemplate_CreditSpecification(r.Spec.ForProvider.CreditSpecification, ctyVal)
	EncodeLaunchTemplate_ElasticGpuSpecifications(r.Spec.ForProvider.ElasticGpuSpecifications, ctyVal)
	EncodeLaunchTemplate_ElasticInferenceAccelerator(r.Spec.ForProvider.ElasticInferenceAccelerator, ctyVal)
	EncodeLaunchTemplate_LicenseSpecification(r.Spec.ForProvider.LicenseSpecification, ctyVal)
	EncodeLaunchTemplate_CapacityReservationSpecification(r.Spec.ForProvider.CapacityReservationSpecification, ctyVal)
	EncodeLaunchTemplate_CpuOptions(r.Spec.ForProvider.CpuOptions, ctyVal)
	EncodeLaunchTemplate_MetadataOptions(r.Spec.ForProvider.MetadataOptions, ctyVal)
	EncodeLaunchTemplate_Monitoring(r.Spec.ForProvider.Monitoring, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings(r.Spec.ForProvider.BlockDeviceMappings, ctyVal)
	EncodeLaunchTemplate_IamInstanceProfile(r.Spec.ForProvider.IamInstanceProfile, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces(r.Spec.ForProvider.NetworkInterfaces, ctyVal)
	EncodeLaunchTemplate_Placement(r.Spec.ForProvider.Placement, ctyVal)
	EncodeLaunchTemplate_TagSpecifications(r.Spec.ForProvider.TagSpecifications, ctyVal)
	EncodeLaunchTemplate_Arn(r.Status.AtProvider, ctyVal)
	EncodeLaunchTemplate_LatestVersion(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLaunchTemplate_Description(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLaunchTemplate_ImageId(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["image_id"] = cty.StringVal(p.ImageId)
}

func EncodeLaunchTemplate_InstanceInitiatedShutdownBehavior(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["instance_initiated_shutdown_behavior"] = cty.StringVal(p.InstanceInitiatedShutdownBehavior)
}

func EncodeLaunchTemplate_InstanceType(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["instance_type"] = cty.StringVal(p.InstanceType)
}

func EncodeLaunchTemplate_KeyName(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["key_name"] = cty.StringVal(p.KeyName)
}

func EncodeLaunchTemplate_RamDiskId(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["ram_disk_id"] = cty.StringVal(p.RamDiskId)
}

func EncodeLaunchTemplate_UpdateDefaultVersion(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["update_default_version"] = cty.BoolVal(p.UpdateDefaultVersion)
}

func EncodeLaunchTemplate_UserData(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["user_data"] = cty.StringVal(p.UserData)
}

func EncodeLaunchTemplate_DisableApiTermination(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["disable_api_termination"] = cty.BoolVal(p.DisableApiTermination)
}

func EncodeLaunchTemplate_EbsOptimized(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["ebs_optimized"] = cty.StringVal(p.EbsOptimized)
}

func EncodeLaunchTemplate_KernelId(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["kernel_id"] = cty.StringVal(p.KernelId)
}

func EncodeLaunchTemplate_Name(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLaunchTemplate_DefaultVersion(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["default_version"] = cty.NumberIntVal(p.DefaultVersion)
}

func EncodeLaunchTemplate_Id(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLaunchTemplate_NamePrefix(p LaunchTemplateParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeLaunchTemplate_SecurityGroupNames(p LaunchTemplateParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupNames {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_names"] = cty.SetVal(colVals)
}

func EncodeLaunchTemplate_Tags(p LaunchTemplateParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeLaunchTemplate_VpcSecurityGroupIds(p LaunchTemplateParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeLaunchTemplate_HibernationOptions(p HibernationOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_HibernationOptions_Configured(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["hibernation_options"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_HibernationOptions_Configured(p HibernationOptions, vals map[string]cty.Value) {
	vals["configured"] = cty.BoolVal(p.Configured)
}

func EncodeLaunchTemplate_InstanceMarketOptions(p InstanceMarketOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_InstanceMarketOptions_MarketType(p, ctyVal)
	EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions(p.SpotOptions, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["instance_market_options"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_InstanceMarketOptions_MarketType(p InstanceMarketOptions, vals map[string]cty.Value) {
	vals["market_type"] = cty.StringVal(p.MarketType)
}

func EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions(p SpotOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_BlockDurationMinutes(p, ctyVal)
	EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_InstanceInterruptionBehavior(p, ctyVal)
	EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_MaxPrice(p, ctyVal)
	EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_SpotInstanceType(p, ctyVal)
	EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_ValidUntil(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["spot_options"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_BlockDurationMinutes(p SpotOptions, vals map[string]cty.Value) {
	vals["block_duration_minutes"] = cty.NumberIntVal(p.BlockDurationMinutes)
}

func EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_InstanceInterruptionBehavior(p SpotOptions, vals map[string]cty.Value) {
	vals["instance_interruption_behavior"] = cty.StringVal(p.InstanceInterruptionBehavior)
}

func EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_MaxPrice(p SpotOptions, vals map[string]cty.Value) {
	vals["max_price"] = cty.StringVal(p.MaxPrice)
}

func EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_SpotInstanceType(p SpotOptions, vals map[string]cty.Value) {
	vals["spot_instance_type"] = cty.StringVal(p.SpotInstanceType)
}

func EncodeLaunchTemplate_InstanceMarketOptions_SpotOptions_ValidUntil(p SpotOptions, vals map[string]cty.Value) {
	vals["valid_until"] = cty.StringVal(p.ValidUntil)
}

func EncodeLaunchTemplate_CreditSpecification(p CreditSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_CreditSpecification_CpuCredits(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["credit_specification"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_CreditSpecification_CpuCredits(p CreditSpecification, vals map[string]cty.Value) {
	vals["cpu_credits"] = cty.StringVal(p.CpuCredits)
}

func EncodeLaunchTemplate_ElasticGpuSpecifications(p ElasticGpuSpecifications, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_ElasticGpuSpecifications_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["elastic_gpu_specifications"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_ElasticGpuSpecifications_Type(p ElasticGpuSpecifications, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeLaunchTemplate_ElasticInferenceAccelerator(p ElasticInferenceAccelerator, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_ElasticInferenceAccelerator_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["elastic_inference_accelerator"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_ElasticInferenceAccelerator_Type(p ElasticInferenceAccelerator, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeLaunchTemplate_LicenseSpecification(p LicenseSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_LicenseSpecification_LicenseConfigurationArn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["license_specification"] = cty.SetVal(valsForCollection)
}

func EncodeLaunchTemplate_LicenseSpecification_LicenseConfigurationArn(p LicenseSpecification, vals map[string]cty.Value) {
	vals["license_configuration_arn"] = cty.StringVal(p.LicenseConfigurationArn)
}

func EncodeLaunchTemplate_CapacityReservationSpecification(p CapacityReservationSpecification, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_CapacityReservationSpecification_CapacityReservationPreference(p, ctyVal)
	EncodeLaunchTemplate_CapacityReservationSpecification_CapacityReservationTarget(p.CapacityReservationTarget, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["capacity_reservation_specification"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_CapacityReservationSpecification_CapacityReservationPreference(p CapacityReservationSpecification, vals map[string]cty.Value) {
	vals["capacity_reservation_preference"] = cty.StringVal(p.CapacityReservationPreference)
}

func EncodeLaunchTemplate_CapacityReservationSpecification_CapacityReservationTarget(p CapacityReservationTarget, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_CapacityReservationSpecification_CapacityReservationTarget_CapacityReservationId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["capacity_reservation_target"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_CapacityReservationSpecification_CapacityReservationTarget_CapacityReservationId(p CapacityReservationTarget, vals map[string]cty.Value) {
	vals["capacity_reservation_id"] = cty.StringVal(p.CapacityReservationId)
}

func EncodeLaunchTemplate_CpuOptions(p CpuOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_CpuOptions_CoreCount(p, ctyVal)
	EncodeLaunchTemplate_CpuOptions_ThreadsPerCore(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cpu_options"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_CpuOptions_CoreCount(p CpuOptions, vals map[string]cty.Value) {
	vals["core_count"] = cty.NumberIntVal(p.CoreCount)
}

func EncodeLaunchTemplate_CpuOptions_ThreadsPerCore(p CpuOptions, vals map[string]cty.Value) {
	vals["threads_per_core"] = cty.NumberIntVal(p.ThreadsPerCore)
}

func EncodeLaunchTemplate_MetadataOptions(p MetadataOptions, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_MetadataOptions_HttpEndpoint(p, ctyVal)
	EncodeLaunchTemplate_MetadataOptions_HttpPutResponseHopLimit(p, ctyVal)
	EncodeLaunchTemplate_MetadataOptions_HttpTokens(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["metadata_options"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_MetadataOptions_HttpEndpoint(p MetadataOptions, vals map[string]cty.Value) {
	vals["http_endpoint"] = cty.StringVal(p.HttpEndpoint)
}

func EncodeLaunchTemplate_MetadataOptions_HttpPutResponseHopLimit(p MetadataOptions, vals map[string]cty.Value) {
	vals["http_put_response_hop_limit"] = cty.NumberIntVal(p.HttpPutResponseHopLimit)
}

func EncodeLaunchTemplate_MetadataOptions_HttpTokens(p MetadataOptions, vals map[string]cty.Value) {
	vals["http_tokens"] = cty.StringVal(p.HttpTokens)
}

func EncodeLaunchTemplate_Monitoring(p Monitoring, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_Monitoring_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["monitoring"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_Monitoring_Enabled(p Monitoring, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeLaunchTemplate_BlockDeviceMappings(p BlockDeviceMappings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_BlockDeviceMappings_DeviceName(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_NoDevice(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_VirtualName(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs(p.Ebs, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["block_device_mappings"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_BlockDeviceMappings_DeviceName(p BlockDeviceMappings, vals map[string]cty.Value) {
	vals["device_name"] = cty.StringVal(p.DeviceName)
}

func EncodeLaunchTemplate_BlockDeviceMappings_NoDevice(p BlockDeviceMappings, vals map[string]cty.Value) {
	vals["no_device"] = cty.StringVal(p.NoDevice)
}

func EncodeLaunchTemplate_BlockDeviceMappings_VirtualName(p BlockDeviceMappings, vals map[string]cty.Value) {
	vals["virtual_name"] = cty.StringVal(p.VirtualName)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs(p Ebs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs_VolumeType(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs_DeleteOnTermination(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs_Encrypted(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs_Iops(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs_KmsKeyId(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs_SnapshotId(p, ctyVal)
	EncodeLaunchTemplate_BlockDeviceMappings_Ebs_VolumeSize(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs_VolumeType(p Ebs, vals map[string]cty.Value) {
	vals["volume_type"] = cty.StringVal(p.VolumeType)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs_DeleteOnTermination(p Ebs, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.StringVal(p.DeleteOnTermination)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs_Encrypted(p Ebs, vals map[string]cty.Value) {
	vals["encrypted"] = cty.StringVal(p.Encrypted)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs_Iops(p Ebs, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs_KmsKeyId(p Ebs, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs_SnapshotId(p Ebs, vals map[string]cty.Value) {
	vals["snapshot_id"] = cty.StringVal(p.SnapshotId)
}

func EncodeLaunchTemplate_BlockDeviceMappings_Ebs_VolumeSize(p Ebs, vals map[string]cty.Value) {
	vals["volume_size"] = cty.NumberIntVal(p.VolumeSize)
}

func EncodeLaunchTemplate_IamInstanceProfile(p IamInstanceProfile, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_IamInstanceProfile_Name(p, ctyVal)
	EncodeLaunchTemplate_IamInstanceProfile_Arn(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["iam_instance_profile"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_IamInstanceProfile_Name(p IamInstanceProfile, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLaunchTemplate_IamInstanceProfile_Arn(p IamInstanceProfile, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLaunchTemplate_NetworkInterfaces(p NetworkInterfaces, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_NetworkInterfaces_Ipv6Addresses(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_PrivateIpAddress(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_SubnetId(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_Ipv4AddressCount(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_Ipv6AddressCount(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_Description(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_DeviceIndex(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_Ipv4Addresses(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_NetworkInterfaceId(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_SecurityGroups(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_AssociatePublicIpAddress(p, ctyVal)
	EncodeLaunchTemplate_NetworkInterfaces_DeleteOnTermination(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["network_interfaces"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_NetworkInterfaces_Ipv6Addresses(p NetworkInterfaces, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ipv6Addresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ipv6_addresses"] = cty.SetVal(colVals)
}

func EncodeLaunchTemplate_NetworkInterfaces_PrivateIpAddress(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["private_ip_address"] = cty.StringVal(p.PrivateIpAddress)
}

func EncodeLaunchTemplate_NetworkInterfaces_SubnetId(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeLaunchTemplate_NetworkInterfaces_Ipv4AddressCount(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["ipv4_address_count"] = cty.NumberIntVal(p.Ipv4AddressCount)
}

func EncodeLaunchTemplate_NetworkInterfaces_Ipv6AddressCount(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["ipv6_address_count"] = cty.NumberIntVal(p.Ipv6AddressCount)
}

func EncodeLaunchTemplate_NetworkInterfaces_Description(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLaunchTemplate_NetworkInterfaces_DeviceIndex(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["device_index"] = cty.NumberIntVal(p.DeviceIndex)
}

func EncodeLaunchTemplate_NetworkInterfaces_Ipv4Addresses(p NetworkInterfaces, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ipv4Addresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ipv4_addresses"] = cty.SetVal(colVals)
}

func EncodeLaunchTemplate_NetworkInterfaces_NetworkInterfaceId(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeLaunchTemplate_NetworkInterfaces_SecurityGroups(p NetworkInterfaces, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeLaunchTemplate_NetworkInterfaces_AssociatePublicIpAddress(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["associate_public_ip_address"] = cty.StringVal(p.AssociatePublicIpAddress)
}

func EncodeLaunchTemplate_NetworkInterfaces_DeleteOnTermination(p NetworkInterfaces, vals map[string]cty.Value) {
	vals["delete_on_termination"] = cty.StringVal(p.DeleteOnTermination)
}

func EncodeLaunchTemplate_Placement(p Placement, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_Placement_Affinity(p, ctyVal)
	EncodeLaunchTemplate_Placement_AvailabilityZone(p, ctyVal)
	EncodeLaunchTemplate_Placement_GroupName(p, ctyVal)
	EncodeLaunchTemplate_Placement_HostId(p, ctyVal)
	EncodeLaunchTemplate_Placement_PartitionNumber(p, ctyVal)
	EncodeLaunchTemplate_Placement_SpreadDomain(p, ctyVal)
	EncodeLaunchTemplate_Placement_Tenancy(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["placement"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_Placement_Affinity(p Placement, vals map[string]cty.Value) {
	vals["affinity"] = cty.StringVal(p.Affinity)
}

func EncodeLaunchTemplate_Placement_AvailabilityZone(p Placement, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeLaunchTemplate_Placement_GroupName(p Placement, vals map[string]cty.Value) {
	vals["group_name"] = cty.StringVal(p.GroupName)
}

func EncodeLaunchTemplate_Placement_HostId(p Placement, vals map[string]cty.Value) {
	vals["host_id"] = cty.StringVal(p.HostId)
}

func EncodeLaunchTemplate_Placement_PartitionNumber(p Placement, vals map[string]cty.Value) {
	vals["partition_number"] = cty.NumberIntVal(p.PartitionNumber)
}

func EncodeLaunchTemplate_Placement_SpreadDomain(p Placement, vals map[string]cty.Value) {
	vals["spread_domain"] = cty.StringVal(p.SpreadDomain)
}

func EncodeLaunchTemplate_Placement_Tenancy(p Placement, vals map[string]cty.Value) {
	vals["tenancy"] = cty.StringVal(p.Tenancy)
}

func EncodeLaunchTemplate_TagSpecifications(p TagSpecifications, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLaunchTemplate_TagSpecifications_ResourceType(p, ctyVal)
	EncodeLaunchTemplate_TagSpecifications_Tags(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tag_specifications"] = cty.ListVal(valsForCollection)
}

func EncodeLaunchTemplate_TagSpecifications_ResourceType(p TagSpecifications, vals map[string]cty.Value) {
	vals["resource_type"] = cty.StringVal(p.ResourceType)
}

func EncodeLaunchTemplate_TagSpecifications_Tags(p TagSpecifications, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeLaunchTemplate_Arn(p LaunchTemplateObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLaunchTemplate_LatestVersion(p LaunchTemplateObservation, vals map[string]cty.Value) {
	vals["latest_version"] = cty.NumberIntVal(p.LatestVersion)
}