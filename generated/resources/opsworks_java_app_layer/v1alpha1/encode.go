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

func EncodeOpsworksJavaAppLayer(r OpsworksJavaAppLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksJavaAppLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_JvmOptions(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_AppServer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_JvmVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_JvmType(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_AppServerVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksJavaAppLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksJavaAppLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksJavaAppLayer_AutoHealing(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksJavaAppLayer_CustomConfigureRecipes(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksJavaAppLayer_JvmOptions(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["jvm_options"] = cty.StringVal(p.JvmOptions)
}

func EncodeOpsworksJavaAppLayer_Id(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksJavaAppLayer_InstallUpdatesOnBoot(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksJavaAppLayer_StackId(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksJavaAppLayer_AppServer(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["app_server"] = cty.StringVal(p.AppServer)
}

func EncodeOpsworksJavaAppLayer_AutoAssignPublicIps(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksJavaAppLayer_CustomJson(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksJavaAppLayer_DrainElbOnShutdown(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksJavaAppLayer_UseEbsOptimizedInstances(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksJavaAppLayer_CustomSecurityGroupIds(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksJavaAppLayer_Tags(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksJavaAppLayer_CustomInstanceProfileArn(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksJavaAppLayer_JvmVersion(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["jvm_version"] = cty.StringVal(p.JvmVersion)
}

func EncodeOpsworksJavaAppLayer_AutoAssignElasticIps(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksJavaAppLayer_CustomShutdownRecipes(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksJavaAppLayer_CustomUndeployRecipes(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksJavaAppLayer_JvmType(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["jvm_type"] = cty.StringVal(p.JvmType)
}

func EncodeOpsworksJavaAppLayer_AppServerVersion(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["app_server_version"] = cty.StringVal(p.AppServerVersion)
}

func EncodeOpsworksJavaAppLayer_CustomSetupRecipes(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksJavaAppLayer_InstanceShutdownTimeout(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksJavaAppLayer_Name(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksJavaAppLayer_SystemPackages(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksJavaAppLayer_CustomDeployRecipes(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksJavaAppLayer_ElasticLoadBalancer(p OpsworksJavaAppLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksJavaAppLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksJavaAppLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksJavaAppLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksJavaAppLayer_EbsVolume_MountPoint(p, ctyVal)
	EncodeOpsworksJavaAppLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksJavaAppLayer_EbsVolume_RaidLevel(p, ctyVal)
	EncodeOpsworksJavaAppLayer_EbsVolume_Size(p, ctyVal)
	EncodeOpsworksJavaAppLayer_EbsVolume_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksJavaAppLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksJavaAppLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksJavaAppLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksJavaAppLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksJavaAppLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksJavaAppLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksJavaAppLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksJavaAppLayer_Arn(p OpsworksJavaAppLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}