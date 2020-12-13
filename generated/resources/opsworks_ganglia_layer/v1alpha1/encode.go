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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*OpsworksGangliaLayer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a OpsworksGangliaLayer.")
	}
	return EncodeOpsworksGangliaLayer(*r), nil
}

func EncodeOpsworksGangliaLayer(r OpsworksGangliaLayer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksGangliaLayer_Url(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_Password(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_Username(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksGangliaLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksGangliaLayer_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeOpsworksGangliaLayer_Url(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["url"] = cty.StringVal(p.Url)
}

func EncodeOpsworksGangliaLayer_AutoHealing(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksGangliaLayer_InstanceShutdownTimeout(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.NumberIntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksGangliaLayer_Name(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksGangliaLayer_Password(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeOpsworksGangliaLayer_StackId(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksGangliaLayer_InstallUpdatesOnBoot(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksGangliaLayer_AutoAssignPublicIps(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksGangliaLayer_CustomSecurityGroupIds(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksGangliaLayer_CustomSetupRecipes(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksGangliaLayer_CustomUndeployRecipes(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksGangliaLayer_Id(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksGangliaLayer_UseEbsOptimizedInstances(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksGangliaLayer_Username(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeOpsworksGangliaLayer_CustomInstanceProfileArn(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksGangliaLayer_CustomShutdownRecipes(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksGangliaLayer_DrainElbOnShutdown(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksGangliaLayer_ElasticLoadBalancer(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksGangliaLayer_Tags(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksGangliaLayer_SystemPackages(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksGangliaLayer_AutoAssignElasticIps(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksGangliaLayer_CustomConfigureRecipes(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksGangliaLayer_CustomDeployRecipes(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksGangliaLayer_CustomJson(p OpsworksGangliaLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksGangliaLayer_EbsVolume(p EbsVolume, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeOpsworksGangliaLayer_EbsVolume_Size(p, ctyVal)
	EncodeOpsworksGangliaLayer_EbsVolume_Type(p, ctyVal)
	EncodeOpsworksGangliaLayer_EbsVolume_Encrypted(p, ctyVal)
	EncodeOpsworksGangliaLayer_EbsVolume_Iops(p, ctyVal)
	EncodeOpsworksGangliaLayer_EbsVolume_MountPoint(p, ctyVal)
	EncodeOpsworksGangliaLayer_EbsVolume_NumberOfDisks(p, ctyVal)
	EncodeOpsworksGangliaLayer_EbsVolume_RaidLevel(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksGangliaLayer_EbsVolume_Size(p EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeOpsworksGangliaLayer_EbsVolume_Type(p EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksGangliaLayer_EbsVolume_Encrypted(p EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksGangliaLayer_EbsVolume_Iops(p EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.NumberIntVal(p.Iops)
}

func EncodeOpsworksGangliaLayer_EbsVolume_MountPoint(p EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksGangliaLayer_EbsVolume_NumberOfDisks(p EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.NumberIntVal(p.NumberOfDisks)
}

func EncodeOpsworksGangliaLayer_EbsVolume_RaidLevel(p EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksGangliaLayer_Arn(p OpsworksGangliaLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}