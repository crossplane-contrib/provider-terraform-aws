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
	r, ok := mr.(*IamGroupMembership)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamGroupMembership(r, ctyValue)
}

func DecodeIamGroupMembership(prev *IamGroupMembership, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamGroupMembership_Name(&new.Spec.ForProvider, valMap)
	DecodeIamGroupMembership_Users(&new.Spec.ForProvider, valMap)
	DecodeIamGroupMembership_Group(&new.Spec.ForProvider, valMap)
	DecodeIamGroupMembership_Id(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamGroupMembership_Name(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeIamGroupMembership_Users(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["users"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Users = goVals
}

//primitiveTypeDecodeTemplate
func DecodeIamGroupMembership_Group(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	p.Group = ctwhy.ValueAsString(vals["group"])
}

//primitiveTypeDecodeTemplate
func DecodeIamGroupMembership_Id(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}