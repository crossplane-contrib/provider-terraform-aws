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

package v1alpha1func EncodeRdsClusterParameterGroup(r RdsClusterParameterGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRdsClusterParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterParameterGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterParameterGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRdsClusterParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeRdsClusterParameterGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeRdsClusterParameterGroup_Description(p *RdsClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeRdsClusterParameterGroup_Family(p *RdsClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeRdsClusterParameterGroup_Id(p *RdsClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRdsClusterParameterGroup_Name(p *RdsClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRdsClusterParameterGroup_NamePrefix(p *RdsClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeRdsClusterParameterGroup_Tags(p *RdsClusterParameterGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRdsClusterParameterGroup_Parameter(p *Parameter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameter {
		ctyVal = make(map[string]cty.Value)
		EncodeRdsClusterParameterGroup_Parameter_Name(v, ctyVal)
		EncodeRdsClusterParameterGroup_Parameter_Value(v, ctyVal)
		EncodeRdsClusterParameterGroup_Parameter_ApplyMethod(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeRdsClusterParameterGroup_Parameter_Name(p *Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRdsClusterParameterGroup_Parameter_Value(p *Parameter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeRdsClusterParameterGroup_Parameter_ApplyMethod(p *Parameter, vals map[string]cty.Value) {
	vals["apply_method"] = cty.StringVal(p.ApplyMethod)
}

func EncodeRdsClusterParameterGroup_Arn(p *RdsClusterParameterGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}