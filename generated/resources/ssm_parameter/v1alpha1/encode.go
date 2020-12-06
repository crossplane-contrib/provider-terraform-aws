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

func EncodeSsmParameter(r SsmParameter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSsmParameter_AllowedPattern(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Arn(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Description(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_KeyId(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Overwrite(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Value(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_DataType(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Tier(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Type(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Version(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSsmParameter_AllowedPattern(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["allowed_pattern"] = cty.StringVal(p.AllowedPattern)
}

func EncodeSsmParameter_Arn(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeSsmParameter_Description(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmParameter_Id(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmParameter_KeyId(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeSsmParameter_Name(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmParameter_Overwrite(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["overwrite"] = cty.BoolVal(p.Overwrite)
}

func EncodeSsmParameter_Tags(p SsmParameterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSsmParameter_Value(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeSsmParameter_DataType(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["data_type"] = cty.StringVal(p.DataType)
}

func EncodeSsmParameter_Tier(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["tier"] = cty.StringVal(p.Tier)
}

func EncodeSsmParameter_Type(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeSsmParameter_Version(p SsmParameterObservation, vals map[string]cty.Value) {
	vals["version"] = cty.NumberIntVal(p.Version)
}