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

package v1alpha1func EncodeNeptuneParameterGroup(r NeptuneParameterGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeNeptuneParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeNeptuneParameterGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeNeptuneParameterGroup_Name(p *NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNeptuneParameterGroup_Tags(p *NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeNeptuneParameterGroup_Description(p *NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeNeptuneParameterGroup_Family(p *NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeNeptuneParameterGroup_Id(p *NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNeptuneParameterGroup_Parameter(p *Parameter, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Parameter {
		ctyVal = make(map[string]cty.Value)
		EncodeNeptuneParameterGroup_Parameter_ApplyMethod(v, ctyVal)
		EncodeNeptuneParameterGroup_Parameter_Name(v, ctyVal)
		EncodeNeptuneParameterGroup_Parameter_Value(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeNeptuneParameterGroup_Parameter_ApplyMethod(p *Parameter, vals map[string]cty.Value) {
	vals["apply_method"] = cty.StringVal(p.ApplyMethod)
}

func EncodeNeptuneParameterGroup_Parameter_Name(p *Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNeptuneParameterGroup_Parameter_Value(p *Parameter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeNeptuneParameterGroup_Arn(p *NeptuneParameterGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}