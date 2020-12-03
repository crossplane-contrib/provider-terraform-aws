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

package v1alpha1func EncodeConfigRemediationConfiguration(r ConfigRemediationConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeConfigRemediationConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigRemediationConfiguration_ResourceType(r.Spec.ForProvider, ctyVal)
	EncodeConfigRemediationConfiguration_TargetId(r.Spec.ForProvider, ctyVal)
	EncodeConfigRemediationConfiguration_TargetType(r.Spec.ForProvider, ctyVal)
	EncodeConfigRemediationConfiguration_TargetVersion(r.Spec.ForProvider, ctyVal)
	EncodeConfigRemediationConfiguration_ConfigRuleName(r.Spec.ForProvider, ctyVal)
	EncodeConfigRemediationConfiguration_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeConfigRemediationConfiguration_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeConfigRemediationConfiguration_Id(p *ConfigRemediationConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigRemediationConfiguration_ResourceType(p *ConfigRemediationConfigurationParameters, vals map[string]cty.Value) {
	vals["resource_type"] = cty.StringVal(p.ResourceType)
}

func EncodeConfigRemediationConfiguration_TargetId(p *ConfigRemediationConfigurationParameters, vals map[string]cty.Value) {
	vals["target_id"] = cty.StringVal(p.TargetId)
}

func EncodeConfigRemediationConfiguration_TargetType(p *ConfigRemediationConfigurationParameters, vals map[string]cty.Value) {
	vals["target_type"] = cty.StringVal(p.TargetType)
}

func EncodeConfigRemediationConfiguration_TargetVersion(p *ConfigRemediationConfigurationParameters, vals map[string]cty.Value) {
	vals["target_version"] = cty.StringVal(p.TargetVersion)
}

func EncodeConfigRemediationConfiguration_ConfigRuleName(p *ConfigRemediationConfigurationParameters, vals map[string]cty.Value) {
	vals["config_rule_name"] = cty.StringVal(p.ConfigRuleName)
}

func EncodeConfigRemediationConfiguration_Parameter(p *Parameter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameter {
		ctyVal = make(map[string]cty.Value)
		EncodeConfigRemediationConfiguration_Parameter_ResourceValue(v, ctyVal)
		EncodeConfigRemediationConfiguration_Parameter_StaticValue(v, ctyVal)
		EncodeConfigRemediationConfiguration_Parameter_Name(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeConfigRemediationConfiguration_Parameter_ResourceValue(p *Parameter, vals map[string]cty.Value) {
	vals["resource_value"] = cty.StringVal(p.ResourceValue)
}

func EncodeConfigRemediationConfiguration_Parameter_StaticValue(p *Parameter, vals map[string]cty.Value) {
	vals["static_value"] = cty.StringVal(p.StaticValue)
}

func EncodeConfigRemediationConfiguration_Parameter_Name(p *Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeConfigRemediationConfiguration_Arn(p *ConfigRemediationConfigurationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}