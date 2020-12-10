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
	r, ok := mr.(*OpsworksMysqlLayer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OpsworksMysqlLayer.")
	}
	return EncodeOpsworksMysqlLayer(*r), nil
}

func EncodeOpsworksMysqlLayer(r OpsworksMysqlLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksMysqlLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_RootPassword(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_RootPasswordOnAllInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksMysqlLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksMysqlLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksMysqlLayer_SystemPackages(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksMysqlLayer_UseEbsOptimizedInstances(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksMysqlLayer_CustomConfigureRecipes(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMysqlLayer_CustomSetupRecipes(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMysqlLayer_Name(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksMysqlLayer_RootPassword(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["root_password"] = cty.StringVal(p.RootPassword)
}

func EncodeOpsworksMysqlLayer_StackId(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksMysqlLayer_AutoAssignPublicIps(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksMysqlLayer_Id(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksMysqlLayer_InstanceShutdownTimeout(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksMysqlLayer_ElasticLoadBalancer(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksMysqlLayer_InstallUpdatesOnBoot(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksMysqlLayer_RootPasswordOnAllInstances(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["root_password_on_all_instances"] = cty.BoolVal(p.RootPasswordOnAllInstances)
}

func EncodeOpsworksMysqlLayer_CustomUndeployRecipes(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMysqlLayer_DrainElbOnShutdown(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksMysqlLayer_CustomInstanceProfileArn(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksMysqlLayer_CustomJson(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksMysqlLayer_CustomSecurityGroupIds(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksMysqlLayer_CustomShutdownRecipes(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMysqlLayer_Tags(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksMysqlLayer_AutoAssignElasticIps(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksMysqlLayer_AutoHealing(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksMysqlLayer_CustomDeployRecipes(p OpsworksMysqlLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksMysqlLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksMysqlLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksMysqlLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksMysqlLayer_EbsVolume_MountPoint(p, ctyVal)
	EncodeOpsworksMysqlLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksMysqlLayer_EbsVolume_RaidLevel(p, ctyVal)
	EncodeOpsworksMysqlLayer_EbsVolume_Size(p, ctyVal)
	EncodeOpsworksMysqlLayer_EbsVolume_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksMysqlLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksMysqlLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksMysqlLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksMysqlLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksMysqlLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksMysqlLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksMysqlLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksMysqlLayer_Arn(p OpsworksMysqlLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}