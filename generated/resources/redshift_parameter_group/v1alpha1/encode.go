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

func EncodeRedshiftParameterGroup(r RedshiftParameterGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftParameterGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRedshiftParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeRedshiftParameterGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRedshiftParameterGroup_Description(p RedshiftParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeRedshiftParameterGroup_Family(p RedshiftParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeRedshiftParameterGroup_Id(p RedshiftParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRedshiftParameterGroup_Name(p RedshiftParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRedshiftParameterGroup_Tags(p RedshiftParameterGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRedshiftParameterGroup_Parameter(p Parameter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRedshiftParameterGroup_Parameter_Name(p, ctyVal)
	EncodeRedshiftParameterGroup_Parameter_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeRedshiftParameterGroup_Parameter_Name(p Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRedshiftParameterGroup_Parameter_Value(p Parameter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeRedshiftParameterGroup_Arn(p RedshiftParameterGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}