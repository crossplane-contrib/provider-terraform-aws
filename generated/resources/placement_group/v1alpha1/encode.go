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
	r, ok := mr.(*PlacementGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a PlacementGroup.")
	}
	return EncodePlacementGroup(*r), nil
}

func EncodePlacementGroup(r PlacementGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodePlacementGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodePlacementGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodePlacementGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodePlacementGroup_Strategy(r.Spec.ForProvider, ctyVal)
	EncodePlacementGroup_Arn(r.Status.AtProvider, ctyVal)
	EncodePlacementGroup_PlacementGroupId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodePlacementGroup_Tags(p PlacementGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodePlacementGroup_Id(p PlacementGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodePlacementGroup_Name(p PlacementGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodePlacementGroup_Strategy(p PlacementGroupParameters, vals map[string]cty.Value) {
	vals["strategy"] = cty.StringVal(p.Strategy)
}

func EncodePlacementGroup_Arn(p PlacementGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodePlacementGroup_PlacementGroupId(p PlacementGroupObservation, vals map[string]cty.Value) {
	vals["placement_group_id"] = cty.StringVal(p.PlacementGroupId)
}