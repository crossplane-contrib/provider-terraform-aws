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
	r, ok := mr.(*GuarddutyMember)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeGuarddutyMember(r, ctyValue)
}

func DecodeGuarddutyMember(prev *GuarddutyMember, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeGuarddutyMember_DisableEmailNotification(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyMember_Email(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyMember_InvitationMessage(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyMember_Invite(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyMember_AccountId(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyMember_DetectorId(&new.Spec.ForProvider, valMap)
	DecodeGuarddutyMember_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeGuarddutyMember_RelationshipStatus(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_DisableEmailNotification(p *GuarddutyMemberParameters, vals map[string]cty.Value) {
	p.DisableEmailNotification = ctwhy.ValueAsBool(vals["disable_email_notification"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_Email(p *GuarddutyMemberParameters, vals map[string]cty.Value) {
	p.Email = ctwhy.ValueAsString(vals["email"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_InvitationMessage(p *GuarddutyMemberParameters, vals map[string]cty.Value) {
	p.InvitationMessage = ctwhy.ValueAsString(vals["invitation_message"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_Invite(p *GuarddutyMemberParameters, vals map[string]cty.Value) {
	p.Invite = ctwhy.ValueAsBool(vals["invite"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_AccountId(p *GuarddutyMemberParameters, vals map[string]cty.Value) {
	p.AccountId = ctwhy.ValueAsString(vals["account_id"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_DetectorId(p *GuarddutyMemberParameters, vals map[string]cty.Value) {
	p.DetectorId = ctwhy.ValueAsString(vals["detector_id"])
}

//containerTypeDecodeTemplate
func DecodeGuarddutyMember_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeGuarddutyMember_Timeouts_Create(p, valMap)
	DecodeGuarddutyMember_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeGuarddutyMember_RelationshipStatus(p *GuarddutyMemberObservation, vals map[string]cty.Value) {
	p.RelationshipStatus = ctwhy.ValueAsString(vals["relationship_status"])
}