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

package v1alpha1func EncodeEmrSecurityConfiguration(r EmrSecurityConfiguration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEmrSecurityConfiguration_Configuration(r.Spec.ForProvider, ctyVal)
	EncodeEmrSecurityConfiguration_Id(r.Spec.ForProvider, ctyVal)
	EncodeEmrSecurityConfiguration_Name(r.Spec.ForProvider, ctyVal)
	EncodeEmrSecurityConfiguration_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeEmrSecurityConfiguration_CreationDate(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeEmrSecurityConfiguration_Configuration(p *EmrSecurityConfigurationParameters, vals map[string]cty.Value) {
	vals["configuration"] = cty.StringVal(p.Configuration)
}

func EncodeEmrSecurityConfiguration_Id(p *EmrSecurityConfigurationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEmrSecurityConfiguration_Name(p *EmrSecurityConfigurationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeEmrSecurityConfiguration_NamePrefix(p *EmrSecurityConfigurationParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeEmrSecurityConfiguration_CreationDate(p *EmrSecurityConfigurationObservation, vals map[string]cty.Value) {
	vals["creation_date"] = cty.StringVal(p.CreationDate)
}