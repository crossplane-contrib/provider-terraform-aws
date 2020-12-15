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
	r, ok := mr.(*OrganizationsAccount)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeOrganizationsAccount(r, ctyValue)
}

func DecodeOrganizationsAccount(prev *OrganizationsAccount, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeOrganizationsAccount_ParentId(&new.Spec.ForProvider, valMap)
	DecodeOrganizationsAccount_Email(&new.Spec.ForProvider, valMap)
	DecodeOrganizationsAccount_IamUserAccessToBilling(&new.Spec.ForProvider, valMap)
	DecodeOrganizationsAccount_Name(&new.Spec.ForProvider, valMap)
	DecodeOrganizationsAccount_RoleName(&new.Spec.ForProvider, valMap)
	DecodeOrganizationsAccount_Tags(&new.Spec.ForProvider, valMap)
	DecodeOrganizationsAccount_Arn(&new.Status.AtProvider, valMap)
	DecodeOrganizationsAccount_JoinedMethod(&new.Status.AtProvider, valMap)
	DecodeOrganizationsAccount_JoinedTimestamp(&new.Status.AtProvider, valMap)
	DecodeOrganizationsAccount_Status(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_ParentId(p *OrganizationsAccountParameters, vals map[string]cty.Value) {
	p.ParentId = ctwhy.ValueAsString(vals["parent_id"])
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_Email(p *OrganizationsAccountParameters, vals map[string]cty.Value) {
	p.Email = ctwhy.ValueAsString(vals["email"])
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_IamUserAccessToBilling(p *OrganizationsAccountParameters, vals map[string]cty.Value) {
	p.IamUserAccessToBilling = ctwhy.ValueAsString(vals["iam_user_access_to_billing"])
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_Name(p *OrganizationsAccountParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_RoleName(p *OrganizationsAccountParameters, vals map[string]cty.Value) {
	p.RoleName = ctwhy.ValueAsString(vals["role_name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeOrganizationsAccount_Tags(p *OrganizationsAccountParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_Arn(p *OrganizationsAccountObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_JoinedMethod(p *OrganizationsAccountObservation, vals map[string]cty.Value) {
	p.JoinedMethod = ctwhy.ValueAsString(vals["joined_method"])
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_JoinedTimestamp(p *OrganizationsAccountObservation, vals map[string]cty.Value) {
	p.JoinedTimestamp = ctwhy.ValueAsString(vals["joined_timestamp"])
}

//primitiveTypeDecodeTemplate
func DecodeOrganizationsAccount_Status(p *OrganizationsAccountObservation, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}