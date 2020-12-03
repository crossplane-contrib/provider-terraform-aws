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

package v1alpha1func EncodeOpsworksPhpAppLayer(r OpsworksPhpAppLayer) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeOpsworksPhpAppLayer_CustomSetupRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_InstanceShutdownTimeout(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_Name(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_StackId(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_AutoHealing(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_CustomInstanceProfileArn(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_CustomShutdownRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_CustomUndeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_ElasticLoadBalancer(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_Id(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_InstallUpdatesOnBoot(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_DrainElbOnShutdown(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_SystemPackages(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_UseEbsOptimizedInstances(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_AutoAssignElasticIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_AutoAssignPublicIps(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_CustomConfigureRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_CustomDeployRecipes(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_CustomJson(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_CustomSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeOpsworksPhpAppLayer_EbsVolume(r.Spec.ForProvider.EbsVolume, ctyVal)
	EncodeOpsworksPhpAppLayer_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeOpsworksPhpAppLayer_CustomSetupRecipes(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSetupRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_setup_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksPhpAppLayer_InstanceShutdownTimeout(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["instance_shutdown_timeout"] = cty.IntVal(p.InstanceShutdownTimeout)
}

func EncodeOpsworksPhpAppLayer_Name(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOpsworksPhpAppLayer_StackId(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["stack_id"] = cty.StringVal(p.StackId)
}

func EncodeOpsworksPhpAppLayer_Tags(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOpsworksPhpAppLayer_AutoHealing(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_healing"] = cty.BoolVal(p.AutoHealing)
}

func EncodeOpsworksPhpAppLayer_CustomInstanceProfileArn(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_instance_profile_arn"] = cty.StringVal(p.CustomInstanceProfileArn)
}

func EncodeOpsworksPhpAppLayer_CustomShutdownRecipes(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomShutdownRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_shutdown_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksPhpAppLayer_CustomUndeployRecipes(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomUndeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_undeploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksPhpAppLayer_ElasticLoadBalancer(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["elastic_load_balancer"] = cty.StringVal(p.ElasticLoadBalancer)
}

func EncodeOpsworksPhpAppLayer_Id(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOpsworksPhpAppLayer_InstallUpdatesOnBoot(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["install_updates_on_boot"] = cty.BoolVal(p.InstallUpdatesOnBoot)
}

func EncodeOpsworksPhpAppLayer_DrainElbOnShutdown(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["drain_elb_on_shutdown"] = cty.BoolVal(p.DrainElbOnShutdown)
}

func EncodeOpsworksPhpAppLayer_SystemPackages(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SystemPackages {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["system_packages"] = cty.SetVal(colVals)
}

func EncodeOpsworksPhpAppLayer_UseEbsOptimizedInstances(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["use_ebs_optimized_instances"] = cty.BoolVal(p.UseEbsOptimizedInstances)
}

func EncodeOpsworksPhpAppLayer_AutoAssignElasticIps(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_elastic_ips"] = cty.BoolVal(p.AutoAssignElasticIps)
}

func EncodeOpsworksPhpAppLayer_AutoAssignPublicIps(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["auto_assign_public_ips"] = cty.BoolVal(p.AutoAssignPublicIps)
}

func EncodeOpsworksPhpAppLayer_CustomConfigureRecipes(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomConfigureRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_configure_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksPhpAppLayer_CustomDeployRecipes(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomDeployRecipes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_deploy_recipes"] = cty.ListVal(colVals)
}

func EncodeOpsworksPhpAppLayer_CustomJson(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	vals["custom_json"] = cty.StringVal(p.CustomJson)
}

func EncodeOpsworksPhpAppLayer_CustomSecurityGroupIds(p *OpsworksPhpAppLayerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["custom_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeOpsworksPhpAppLayer_EbsVolume(p *EbsVolume, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.EbsVolume {
		ctyVal = make(map[string]cty.Value)
		EncodeOpsworksPhpAppLayer_EbsVolume_Type(v, ctyVal)
		EncodeOpsworksPhpAppLayer_EbsVolume_Encrypted(v, ctyVal)
		EncodeOpsworksPhpAppLayer_EbsVolume_Iops(v, ctyVal)
		EncodeOpsworksPhpAppLayer_EbsVolume_MountPoint(v, ctyVal)
		EncodeOpsworksPhpAppLayer_EbsVolume_NumberOfDisks(v, ctyVal)
		EncodeOpsworksPhpAppLayer_EbsVolume_RaidLevel(v, ctyVal)
		EncodeOpsworksPhpAppLayer_EbsVolume_Size(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ebs_volume"] = cty.SetVal(valsForCollection)
}

func EncodeOpsworksPhpAppLayer_EbsVolume_Type(p *EbsVolume, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOpsworksPhpAppLayer_EbsVolume_Encrypted(p *EbsVolume, vals map[string]cty.Value) {
	vals["encrypted"] = cty.BoolVal(p.Encrypted)
}

func EncodeOpsworksPhpAppLayer_EbsVolume_Iops(p *EbsVolume, vals map[string]cty.Value) {
	vals["iops"] = cty.IntVal(p.Iops)
}

func EncodeOpsworksPhpAppLayer_EbsVolume_MountPoint(p *EbsVolume, vals map[string]cty.Value) {
	vals["mount_point"] = cty.StringVal(p.MountPoint)
}

func EncodeOpsworksPhpAppLayer_EbsVolume_NumberOfDisks(p *EbsVolume, vals map[string]cty.Value) {
	vals["number_of_disks"] = cty.IntVal(p.NumberOfDisks)
}

func EncodeOpsworksPhpAppLayer_EbsVolume_RaidLevel(p *EbsVolume, vals map[string]cty.Value) {
	vals["raid_level"] = cty.StringVal(p.RaidLevel)
}

func EncodeOpsworksPhpAppLayer_EbsVolume_Size(p *EbsVolume, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeOpsworksPhpAppLayer_Arn(p *OpsworksPhpAppLayerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}