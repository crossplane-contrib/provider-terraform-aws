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

func EncodeOpsworksRailsAppLayer(r OpsworksRailsAppLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksRailsAppLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_RubygemsVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_RubyVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_AppServer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_ManageBundler(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_PassengerVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_BundlerVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksRailsAppLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksRailsAppLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksRailsAppLayer_AutoAssignElasticIps(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksRailsAppLayer_ElasticLoadBalancer(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksRailsAppLayer_InstanceShutdownTimeout(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksRailsAppLayer_SystemPackages(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksRailsAppLayer_AutoAssignPublicIps(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksRailsAppLayer_CustomConfigureRecipes(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksRailsAppLayer_CustomInstanceProfileArn(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksRailsAppLayer_CustomUndeployRecipes(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksRailsAppLayer_Tags(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksRailsAppLayer_InstallUpdatesOnBoot(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksRailsAppLayer_RubygemsVersion(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["rubygems_version"] = cty.StringVal(p.RubygemsVersion)
}

func EncodeOpsworksRailsAppLayer_UseEbsOptimizedInstances(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksRailsAppLayer_CustomJson(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksRailsAppLayer_CustomShutdownRecipes(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksRailsAppLayer_DrainElbOnShutdown(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksRailsAppLayer_Id(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksRailsAppLayer_RubyVersion(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["ruby_version"] = cty.StringVal(p.RubyVersion)
}

func EncodeOpsworksRailsAppLayer_AppServer(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["app_server"] = cty.StringVal(p.AppServer)
}

func EncodeOpsworksRailsAppLayer_CustomSecurityGroupIds(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksRailsAppLayer_StackId(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksRailsAppLayer_ManageBundler(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["manage_bundler"] = cty.BoolVal(p.ManageBundler)
}

func EncodeOpsworksRailsAppLayer_Name(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksRailsAppLayer_PassengerVersion(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["passenger_version"] = cty.StringVal(p.PassengerVersion)
}

func EncodeOpsworksRailsAppLayer_AutoHealing(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksRailsAppLayer_BundlerVersion(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	vals["bundler_version"] = cty.StringVal(p.BundlerVersion)
}

func EncodeOpsworksRailsAppLayer_CustomDeployRecipes(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksRailsAppLayer_CustomSetupRecipes(p OpsworksRailsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksRailsAppLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksRailsAppLayer_EbsVolume_Type(p, ctyVal)
	EncodeOpsworksRailsAppLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksRailsAppLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksRailsAppLayer_EbsVolume_MountPoint(p, ctyVal)
	EncodeOpsworksRailsAppLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksRailsAppLayer_EbsVolume_RaidLevel(p, ctyVal)
	EncodeOpsworksRailsAppLayer_EbsVolume_Size(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksRailsAppLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksRailsAppLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksRailsAppLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksRailsAppLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksRailsAppLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksRailsAppLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksRailsAppLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksRailsAppLayer_Arn(p OpsworksRailsAppLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}