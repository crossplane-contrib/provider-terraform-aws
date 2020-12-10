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
	r, ok := mr.(*OpsworksHaproxyLayer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OpsworksHaproxyLayer.")
	}
	return EncodeOpsworksHaproxyLayer(*r), nil
}

func EncodeOpsworksHaproxyLayer(r OpsworksHaproxyLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksHaproxyLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_HealthcheckUrl(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_StatsPassword(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_StatsUrl(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_StatsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_HealthcheckMethod(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_StatsUser(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksHaproxyLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksHaproxyLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksHaproxyLayer_UseEbsOptimizedInstances(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksHaproxyLayer_HealthcheckUrl(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["healthcheck_url"] = cty.StringVal(p.HealthcheckUrl)
}

func EncodeOpsworksHaproxyLayer_Tags(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksHaproxyLayer_CustomConfigureRecipes(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksHaproxyLayer_StackId(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksHaproxyLayer_StatsPassword(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["stats_password"] = cty.StringVal(p.StatsPassword)
}

func EncodeOpsworksHaproxyLayer_SystemPackages(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksHaproxyLayer_InstallUpdatesOnBoot(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksHaproxyLayer_InstanceShutdownTimeout(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksHaproxyLayer_StatsUrl(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["stats_url"] = cty.StringVal(p.StatsUrl)
}

func EncodeOpsworksHaproxyLayer_AutoAssignPublicIps(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksHaproxyLayer_CustomSetupRecipes(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksHaproxyLayer_StatsEnabled(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["stats_enabled"] = cty.BoolVal(p.StatsEnabled)
}

func EncodeOpsworksHaproxyLayer_CustomInstanceProfileArn(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksHaproxyLayer_ElasticLoadBalancer(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksHaproxyLayer_HealthcheckMethod(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["healthcheck_method"] = cty.StringVal(p.HealthcheckMethod)
}

func EncodeOpsworksHaproxyLayer_Name(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksHaproxyLayer_AutoHealing(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksHaproxyLayer_CustomShutdownRecipes(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksHaproxyLayer_CustomJson(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksHaproxyLayer_CustomUndeployRecipes(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksHaproxyLayer_DrainElbOnShutdown(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksHaproxyLayer_Id(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksHaproxyLayer_StatsUser(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["stats_user"] = cty.StringVal(p.StatsUser)
}

func EncodeOpsworksHaproxyLayer_AutoAssignElasticIps(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksHaproxyLayer_CustomDeployRecipes(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksHaproxyLayer_CustomSecurityGroupIds(p OpsworksHaproxyLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksHaproxyLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksHaproxyLayer_EbsVolume_Type(p, ctyVal)
	EncodeOpsworksHaproxyLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksHaproxyLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksHaproxyLayer_EbsVolume_MountPoint(p, ctyVal)
	EncodeOpsworksHaproxyLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksHaproxyLayer_EbsVolume_RaidLevel(p, ctyVal)
	EncodeOpsworksHaproxyLayer_EbsVolume_Size(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksHaproxyLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksHaproxyLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksHaproxyLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksHaproxyLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksHaproxyLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksHaproxyLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksHaproxyLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksHaproxyLayer_Arn(p OpsworksHaproxyLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}