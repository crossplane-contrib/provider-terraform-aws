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
	r, ok := mr.(*SecurityhubMember)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSecurityhubMember(r, ctyValue)
}

func DecodeSecurityhubMember(prev *SecurityhubMember, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSecurityhubMember_AccountId(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubMember_Email(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubMember_Id(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubMember_Invite(&new.Spec.ForProvider, valMap)
	DecodeSecurityhubMember_MasterId(&new.Status.AtProvider, valMap)
	DecodeSecurityhubMember_MemberStatus(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubMember_AccountId(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	p.AccountId = ctwhy.ValueAsString(vals["account_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubMember_Email(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	p.Email = ctwhy.ValueAsString(vals["email"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubMember_Id(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubMember_Invite(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	p.Invite = ctwhy.ValueAsBool(vals["invite"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubMember_MasterId(p *SecurityhubMemberObservation, vals map[string]cty.Value) {
	p.MasterId = ctwhy.ValueAsString(vals["master_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityhubMember_MemberStatus(p *SecurityhubMemberObservation, vals map[string]cty.Value) {
	p.MemberStatus = ctwhy.ValueAsString(vals["member_status"])
}