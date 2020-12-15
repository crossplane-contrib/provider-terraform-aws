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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*PlacementGroup)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodePlacementGroup(r, ctyValue)
}

func DecodePlacementGroup(prev *PlacementGroup, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodePlacementGroup_Strategy(&new.Spec.ForProvider, valMap)
	DecodePlacementGroup_Tags(&new.Spec.ForProvider, valMap)
	DecodePlacementGroup_Id(&new.Spec.ForProvider, valMap)
	DecodePlacementGroup_Name(&new.Spec.ForProvider, valMap)
	DecodePlacementGroup_PlacementGroupId(&new.Status.AtProvider, valMap)
	DecodePlacementGroup_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodePlacementGroup_Strategy(p *PlacementGroupParameters, vals map[string]cty.Value) {
	p.Strategy = ctwhy.ValueAsString(vals["strategy"])
}

//primitiveMapTypeDecodeTemplate
func DecodePlacementGroup_Tags(p *PlacementGroupParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodePlacementGroup_Id(p *PlacementGroupParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodePlacementGroup_Name(p *PlacementGroupParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodePlacementGroup_PlacementGroupId(p *PlacementGroupObservation, vals map[string]cty.Value) {
	p.PlacementGroupId = ctwhy.ValueAsString(vals["placement_group_id"])
}

//primitiveTypeDecodeTemplate
func DecodePlacementGroup_Arn(p *PlacementGroupObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}