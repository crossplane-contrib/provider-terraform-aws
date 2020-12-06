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
	"github.com/zclconf/go-cty/cty"
)

func EncodeDbParameterGroup(r DbParameterGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDbParameterGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeDbParameterGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDbParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeDbParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDbParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeDbParameterGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDbParameterGroup_NamePrefix(p DbParameterGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeDbParameterGroup_Tags(p DbParameterGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDbParameterGroup_Description(p DbParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDbParameterGroup_Family(p DbParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeDbParameterGroup_Id(p DbParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbParameterGroup_Name(p DbParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbParameterGroup_Parameter(p Parameter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDbParameterGroup_Parameter_Name(p, ctyVal)
	EncodeDbParameterGroup_Parameter_Value(p, ctyVal)
	EncodeDbParameterGroup_Parameter_ApplyMethod(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeDbParameterGroup_Parameter_Name(p Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbParameterGroup_Parameter_Value(p Parameter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeDbParameterGroup_Parameter_ApplyMethod(p Parameter, vals map[string]cty.Value) {
	vals["apply_method"] = cty.StringVal(p.ApplyMethod)
}

func EncodeDbParameterGroup_Arn(p DbParameterGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}