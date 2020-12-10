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
	r, ok := mr.(*OpsworksNodejsAppLayer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OpsworksNodejsAppLayer.")
	}
	return EncodeOpsworksNodejsAppLayer(*r), nil
}

func EncodeOpsworksNodejsAppLayer(r OpsworksNodejsAppLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksNodejsAppLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_NodejsVersion(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksNodejsAppLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksNodejsAppLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksNodejsAppLayer_CustomDeployRecipes(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksNodejsAppLayer_CustomSecurityGroupIds(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksNodejsAppLayer_CustomUndeployRecipes(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksNodejsAppLayer_Name(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksNodejsAppLayer_NodejsVersion(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["nodejs_version"] = cty.StringVal(p.NodejsVersion)
}

func EncodeOpsworksNodejsAppLayer_Tags(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksNodejsAppLayer_AutoAssignPublicIps(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksNodejsAppLayer_CustomShutdownRecipes(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksNodejsAppLayer_InstanceShutdownTimeout(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksNodejsAppLayer_StackId(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksNodejsAppLayer_UseEbsOptimizedInstances(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksNodejsAppLayer_CustomSetupRecipes(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksNodejsAppLayer_ElasticLoadBalancer(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksNodejsAppLayer_Id(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksNodejsAppLayer_AutoAssignElasticIps(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksNodejsAppLayer_AutoHealing(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksNodejsAppLayer_CustomConfigureRecipes(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksNodejsAppLayer_CustomInstanceProfileArn(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksNodejsAppLayer_CustomJson(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksNodejsAppLayer_DrainElbOnShutdown(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksNodejsAppLayer_InstallUpdatesOnBoot(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksNodejsAppLayer_SystemPackages(p OpsworksNodejsAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksNodejsAppLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksNodejsAppLayer_EbsVolume_RaidLevel(p, ctyVal)
	EncodeOpsworksNodejsAppLayer_EbsVolume_Size(p, ctyVal)
	EncodeOpsworksNodejsAppLayer_EbsVolume_Type(p, ctyVal)
	EncodeOpsworksNodejsAppLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksNodejsAppLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksNodejsAppLayer_EbsVolume_MountPoint(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksNodejsAppLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksNodejsAppLayer_Arn(p OpsworksNodejsAppLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}