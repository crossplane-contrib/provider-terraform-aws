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
	r, ok := mr.(*NeptuneSubnetGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NeptuneSubnetGroup.")
	}
	return EncodeNeptuneSubnetGroup(*r), nil
}

func EncodeNeptuneSubnetGroup(r NeptuneSubnetGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNeptuneSubnetGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeNeptuneSubnetGroup_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeNeptuneSubnetGroup_Description(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeNeptuneSubnetGroup_Name(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeNeptuneSubnetGroup_NamePrefix(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeNeptuneSubnetGroup_SubnetIds(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeNeptuneSubnetGroup_Tags(p NeptuneSubnetGroupParameters, vals map[string]cty.Value) {
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

func EncodeNeptuneSubnetGroup_Arn(p NeptuneSubnetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}