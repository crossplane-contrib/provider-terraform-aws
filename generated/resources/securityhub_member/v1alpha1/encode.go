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

package v1alpha1func EncodeSecurityhubMember(r SecurityhubMember) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSecurityhubMember_Invite(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubMember_AccountId(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubMember_Email(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubMember_Id(r.Spec.ForProvider, ctyVal)
	EncodeSecurityhubMember_MasterId(r.Status.AtProvider, ctyVal)
	EncodeSecurityhubMember_MemberStatus(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSecurityhubMember_Invite(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	vals["invite"] = cty.BoolVal(p.Invite)
}

func EncodeSecurityhubMember_AccountId(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeSecurityhubMember_Email(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeSecurityhubMember_Id(p *SecurityhubMemberParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSecurityhubMember_MasterId(p *SecurityhubMemberObservation, vals map[string]cty.Value) {
	vals["master_id"] = cty.StringVal(p.MasterId)
}

func EncodeSecurityhubMember_MemberStatus(p *SecurityhubMemberObservation, vals map[string]cty.Value) {
	vals["member_status"] = cty.StringVal(p.MemberStatus)
}