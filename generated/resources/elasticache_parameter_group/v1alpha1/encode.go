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

package v1alpha1func EncodeElasticacheParameterGroup(r ElasticacheParameterGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeElasticacheParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeElasticacheParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeElasticacheParameterGroup_Description(p *ElasticacheParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeElasticacheParameterGroup_Family(p *ElasticacheParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeElasticacheParameterGroup_Id(p *ElasticacheParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeElasticacheParameterGroup_Name(p *ElasticacheParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticacheParameterGroup_Parameter(p *Parameter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameter {
		ctyVal = make(map[string]cty.Value)
		EncodeElasticacheParameterGroup_Parameter_Name(v, ctyVal)
		EncodeElasticacheParameterGroup_Parameter_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeElasticacheParameterGroup_Parameter_Name(p *Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeElasticacheParameterGroup_Parameter_Value(p *Parameter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}