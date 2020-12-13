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
	r, ok := mr.(*DocdbClusterParameterGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DocdbClusterParameterGroup.")
	}
	return EncodeDocdbClusterParameterGroup(*r), nil
}

func EncodeDocdbClusterParameterGroup(r DocdbClusterParameterGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDocdbClusterParameterGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Family(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDocdbClusterParameterGroup_Parameter(r.Spec.ForProvider.Parameter, ctyVal)
	EncodeDocdbClusterParameterGroup_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDocdbClusterParameterGroup_Description(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDocdbClusterParameterGroup_Family(p DocdbClusterParameterGroupParameters, vals map[string]cty.Value) {
	vals["family"] = cty.StringVal(p.Family)
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