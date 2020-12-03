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

package v1alpha1func EncodeOpsworksCustomLayer(r OpsworksCustomLayer) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeOpsworksCustomLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_ShortName(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksCustomLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksCustomLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeOpsworksCustomLayer_CustomDeployRecipes(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksCustomLayer_Id(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksCustomLayer_InstanceShutdownTimeout(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.IntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksCustomLayer_AutoAssignElasticIps(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksCustomLayer_AutoAssignPublicIps(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksCustomLayer_AutoHealing(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksCustomLayer_CustomJson(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksCustomLayer_CustomShutdownRecipes(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksCustomLayer_CustomUndeployRecipes(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksCustomLayer_Name(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksCustomLayer_StackId(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksCustomLayer_Tags(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksCustomLayer_CustomConfigureRecipes(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksCustomLayer_CustomSecurityGroupIds(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksCustomLayer_ElasticLoadBalancer(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksCustomLayer_InstallUpdatesOnBoot(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksCustomLayer_ShortName(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["short_name"] = cty.StringVal(p.ShortName)
}

func EncodeOpsworksCustomLayer_UseEbsOptimizedInstances(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksCustomLayer_CustomInstanceProfileArn(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksCustomLayer_CustomSetupRecipes(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksCustomLayer_DrainElbOnShutdown(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksCustomLayer_SystemPackages(p *OpsworksCustomLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksCustomLayer_EbsVolume(p *EbsVolume, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsVolume {
		ctyVal = make(map[string]cty.Value)
		EncodeOpsworksCustomLayer_EbsVolume_Iops(v, ctyVal)
		EncodeOpsworksCustomLayer_EbsVolume_MountPoint(v, ctyVal)
		EncodeOpsworksCustomLayer_EbsVolume_NumberOfDisks(v, ctyVal)
		EncodeOpsworksCustomLayer_EbsVolume_RaidLevel(v, ctyVal)
		EncodeOpsworksCustomLayer_EbsVolume_Size(v, ctyVal)
		EncodeOpsworksCustomLayer_EbsVolume_Type(v, ctyVal)
		EncodeOpsworksCustomLayer_EbsVolume_Encrypted(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksCustomLayer_EbsVolume_Iops(p *EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeOpsworksCustomLayer_EbsVolume_MountPoint(p *EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksCustomLayer_EbsVolume_NumberOfDisks(p *EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.IntVal(p.NumberOfDisks)
}

func EncodeOpsworksCustomLayer_EbsVolume_RaidLevel(p *EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksCustomLayer_EbsVolume_Size(p *EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeOpsworksCustomLayer_EbsVolume_Type(p *EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksCustomLayer_EbsVolume_Encrypted(p *EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksCustomLayer_Arn(p *OpsworksCustomLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}