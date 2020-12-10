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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*NeptuneParameterGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NeptuneParameterGroup.")
	}
	return EncodeNeptuneParameterGroup(*r), nil
}

func EncodeNeptuneParameterGroup(r NeptuneParameterGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeNeptuneParameterGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeNeptuneParameterGroup_Description(p NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeNeptuneParameterGroup_Family(p NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
}

func EncodeNeptuneParameterGroup_Id(p NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNeptuneParameterGroup_Name(p NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNeptuneParameterGroup_Tags(p NeptuneParameterGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeNeptuneParameterGroup_Parameter(p Parameter, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneParameterGroup_Parameter_ApplyMethod(p, ctyVal)
	EncodeNeptuneParameterGroup_Parameter_Name(p, ctyVal)
	EncodeNeptuneParameterGroup_Parameter_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["parameter"] = cty.SetVal(valsForCollection)
}

func EncodeNeptuneParameterGroup_Parameter_ApplyMethod(p Parameter, vals map[string]cty.Value) {
	vals["apply_method"] = cty.StringVal(p.ApplyMethod)
}

func EncodeNeptuneParameterGroup_Parameter_Name(p Parameter, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNeptuneParameterGroup_Parameter_Value(p Parameter, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeNeptuneParameterGroup_Arn(p NeptuneParameterGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}