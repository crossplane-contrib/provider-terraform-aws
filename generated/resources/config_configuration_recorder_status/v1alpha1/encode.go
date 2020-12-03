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

package v1alpha1func EncodeConfigConfigurationRecorderStatus(r ConfigConfigurationRecorderStatus) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeConfigConfigurationRecorderStatus_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationRecorderStatus_IsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigurationRecorderStatus_Name(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeConfigConfigurationRecorderStatus_Id(p *ConfigConfigurationRecorderStatusParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigConfigurationRecorderStatus_IsEnabled(p *ConfigConfigurationRecorderStatusParameters, vals map[string]cty.Value) {
	vals["is_enabled"] = cty.BoolVal(p.IsEnabled)
}

func EncodeConfigConfigurationRecorderStatus_Name(p *ConfigConfigurationRecorderStatusParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}