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

func EncodeOpsworksStaticWebLayer(r OpsworksStaticWebLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksStaticWebLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksStaticWebLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksStaticWebLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksStaticWebLayer_AutoAssignElasticIps(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksStaticWebLayer_AutoAssignPublicIps(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksStaticWebLayer_CustomDeployRecipes(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksStaticWebLayer_CustomInstanceProfileArn(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksStaticWebLayer_CustomJson(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksStaticWebLayer_CustomSecurityGroupIds(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksStaticWebLayer_CustomUndeployRecipes(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksStaticWebLayer_DrainElbOnShutdown(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksStaticWebLayer_InstanceShutdownTimeout(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksStaticWebLayer_Name(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksStaticWebLayer_Tags(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksStaticWebLayer_CustomSetupRecipes(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksStaticWebLayer_ElasticLoadBalancer(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksStaticWebLayer_Id(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksStaticWebLayer_SystemPackages(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksStaticWebLayer_AutoHealing(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksStaticWebLayer_CustomConfigureRecipes(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksStaticWebLayer_CustomShutdownRecipes(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksStaticWebLayer_InstallUpdatesOnBoot(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksStaticWebLayer_UseEbsOptimizedInstances(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksStaticWebLayer_StackId(p OpsworksStaticWebLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksStaticWebLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksStaticWebLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksStaticWebLayer_EbsVolume_RaidLevel(p, ctyVal)
	EncodeOpsworksStaticWebLayer_EbsVolume_Size(p, ctyVal)
	EncodeOpsworksStaticWebLayer_EbsVolume_Type(p, ctyVal)
	EncodeOpsworksStaticWebLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksStaticWebLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksStaticWebLayer_EbsVolume_MountPoint(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksStaticWebLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksStaticWebLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksStaticWebLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksStaticWebLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksStaticWebLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksStaticWebLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksStaticWebLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksStaticWebLayer_Arn(p OpsworksStaticWebLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}