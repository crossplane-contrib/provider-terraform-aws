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
	r, ok := mr.(*ResourcegroupsGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ResourcegroupsGroup.")
	}
	return EncodeResourcegroupsGroup(*r), nil
}

func EncodeResourcegroupsGroup(r ResourcegroupsGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeResourcegroupsGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeResourcegroupsGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeResourcegroupsGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeResourcegroupsGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeResourcegroupsGroup_ResourceQuery(r.Spec.ForProvider.ResourceQuery, ctyVal)
	EncodeResourcegroupsGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeResourcegroupsGroup_Id(p ResourcegroupsGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeResourcegroupsGroup_Name(p ResourcegroupsGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeResourcegroupsGroup_Tags(p ResourcegroupsGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeResourcegroupsGroup_Description(p ResourcegroupsGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeResourcegroupsGroup_ResourceQuery(p ResourceQuery, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeResourcegroupsGroup_ResourceQuery_Query(p, ctyVal)
	EncodeResourcegroupsGroup_ResourceQuery_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["resource_query"] = cty.ListVal(valsForCollection)
}

func EncodeResourcegroupsGroup_ResourceQuery_Query(p ResourceQuery, vals map[string]cty.Value) {
	vals["query"] = cty.StringVal(p.Query)
}

func EncodeResourcegroupsGroup_ResourceQuery_Type(p ResourceQuery, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeResourcegroupsGroup_Arn(p ResourcegroupsGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}