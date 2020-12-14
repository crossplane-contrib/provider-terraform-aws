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
	r, ok := mr.(*GuarddutyMember)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GuarddutyMember.")
	}
	return EncodeGuarddutyMember(*r), nil
}

func EncodeGuarddutyMember(r GuarddutyMember) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyMember_Invite(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyMember_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyMember_DetectorId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyMember_DisableEmailNotification(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyMember_Email(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyMember_Id(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyMember_InvitationMessage(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyMember_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeGuarddutyMember_RelationshipStatus(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyMember_Invite(p GuarddutyMemberParameters, vals map[string]cty.Value) {
	vals["invite"] = cty.BoolVal(p.Invite)
}

func EncodeGuarddutyMember_AccountId(p GuarddutyMemberParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeGuarddutyMember_DetectorId(p GuarddutyMemberParameters, vals map[string]cty.Value) {
	vals["detector_id"] = cty.StringVal(p.DetectorId)
}

func EncodeGuarddutyMember_DisableEmailNotification(p GuarddutyMemberParameters, vals map[string]cty.Value) {
	vals["disable_email_notification"] = cty.BoolVal(p.DisableEmailNotification)
}

func EncodeGuarddutyMember_Email(p GuarddutyMemberParameters, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeGuarddutyMember_Id(p GuarddutyMemberParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGuarddutyMember_InvitationMessage(p GuarddutyMemberParameters, vals map[string]cty.Value) {
	vals["invitation_message"] = cty.StringVal(p.InvitationMessage)
}

func EncodeGuarddutyMember_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyMember_Timeouts_Update(p, ctyVal)
	EncodeGuarddutyMember_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyMember_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeGuarddutyMember_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeGuarddutyMember_RelationshipStatus(p GuarddutyMemberObservation, vals map[string]cty.Value) {
	vals["relationship_status"] = cty.StringVal(p.RelationshipStatus)
}