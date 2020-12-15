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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*SsmParameter)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SsmParameter.")
	}
	return EncodeSsmParameter(*r), nil
}

func EncodeSsmParameter(r SsmParameter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSsmParameter_Tier(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Type(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Value(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Description(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Name(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_DataType(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Id(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_KeyId(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Overwrite(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_AllowedPattern(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Arn(r.Spec.ForProvider, ctyVal)
	EncodeSsmParameter_Version(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeSsmParameter_Tier(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["tier"] = cty.StringVal(p.Tier)
}

func EncodeSsmParameter_Type(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeSsmParameter_Value(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeSsmParameter_Description(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSsmParameter_Name(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSsmParameter_DataType(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["data_type"] = cty.StringVal(p.DataType)
}

func EncodeSsmParameter_Id(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSsmParameter_KeyId(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeSsmParameter_Overwrite(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["overwrite"] = cty.BoolVal(p.Overwrite)
}

func EncodeSsmParameter_Tags(p SsmParameterParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSsmParameter_AllowedPattern(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["allowed_pattern"] = cty.StringVal(p.AllowedPattern)
}

func EncodeSsmParameter_Arn(p SsmParameterParameters, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeSsmParameter_Version(p SsmParameterObservation, vals map[string]cty.Value) {
	vals["version"] = cty.NumberIntVal(p.Version)
}