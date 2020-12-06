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

func EncodeDocdbClusterParameterGroup(r DocdbClusterParameterGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbClusterParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeDocdbClusterParameterGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDocdbClusterParameterGroup_Id(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDocdbClusterParameterGroup_Name(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDocdbClusterParameterGroup_NamePrefix(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeDocdbClusterParameterGroup_Tags(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDocdbClusterParameterGroup_Description(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDocdbClusterParameterGroup_Family(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeDocdbClusterParameterGroup_Parameter(p Parameter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbClusterParameterGroup_Parameter_ApplyMethod(p, ctyVal)
	EncodeDocdbClusterParameterGroup_Parameter_Name(p, ctyVal)
	EncodeDocdbClusterParameterGroup_Parameter_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeDocdbClusterParameterGroup_Parameter_ApplyMethod(p Parameter, vals map[string]cty.Value) {
	vals["apply_method"] = cty.StringVal(p.ApplyMethod)
}

func EncodeDocdbClusterParameterGroup_Parameter_Name(p Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDocdbClusterParameterGroup_Parameter_Value(p Parameter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeDocdbClusterParameterGroup_Arn(p DocdbClusterParameterGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}