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

func EncodeOpsworksMemcachedLayer(r OpsworksMemcachedLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksMemcachedLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_AllocatedMemory(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMemcachedLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksMemcachedLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksMemcachedLayer_Id(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksMemcachedLayer_DrainElbOnShutdown(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksMemcachedLayer_CustomConfigureRecipes(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMemcachedLayer_CustomSetupRecipes(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMemcachedLayer_CustomUndeployRecipes(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMemcachedLayer_InstallUpdatesOnBoot(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksMemcachedLayer_Name(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksMemcachedLayer_StackId(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksMemcachedLayer_AutoHealing(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksMemcachedLayer_CustomSecurityGroupIds(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksMemcachedLayer_InstanceShutdownTimeout(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksMemcachedLayer_AutoAssignElasticIps(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksMemcachedLayer_AutoAssignPublicIps(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksMemcachedLayer_CustomDeployRecipes(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMemcachedLayer_CustomInstanceProfileArn(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksMemcachedLayer_CustomJson(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksMemcachedLayer_CustomShutdownRecipes(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMemcachedLayer_ElasticLoadBalancer(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksMemcachedLayer_SystemPackages(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksMemcachedLayer_AllocatedMemory(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["allocated_memory"] = cty.NumberIntVal(p.AllocatedMemory)
}

func EncodeOpsworksMemcachedLayer_UseEbsOptimizedInstances(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksMemcachedLayer_Tags(p OpsworksMemcachedLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksMemcachedLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksMemcachedLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksMemcachedLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksMemcachedLayer_EbsVolume_MountPoint(p, ctyVal)
	EncodeOpsworksMemcachedLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksMemcachedLayer_EbsVolume_RaidLevel(p, ctyVal)
	EncodeOpsworksMemcachedLayer_EbsVolume_Size(p, ctyVal)
	EncodeOpsworksMemcachedLayer_EbsVolume_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksMemcachedLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksMemcachedLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksMemcachedLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksMemcachedLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksMemcachedLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksMemcachedLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksMemcachedLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksMemcachedLayer_Arn(p OpsworksMemcachedLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}