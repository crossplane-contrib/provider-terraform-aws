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

package v1alpha1func EncodeEc2ClientVpnAuthorizationRule(r Ec2ClientVpnAuthorizationRule) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeEc2ClientVpnAuthorizationRule_AccessGroupId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_AuthorizeAllGroups(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_ClientVpnEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_TargetNetworkCidr(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeEc2ClientVpnAuthorizationRule_AccessGroupId(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["access_group_id"] = cty.StringVal(p.AccessGroupId)
}

func EncodeEc2ClientVpnAuthorizationRule_AuthorizeAllGroups(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["authorize_all_groups"] = cty.BoolVal(p.AuthorizeAllGroups)
}

func EncodeEc2ClientVpnAuthorizationRule_ClientVpnEndpointId(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["client_vpn_endpoint_id"] = cty.StringVal(p.ClientVpnEndpointId)
}

func EncodeEc2ClientVpnAuthorizationRule_Description(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2ClientVpnAuthorizationRule_Id(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2ClientVpnAuthorizationRule_TargetNetworkCidr(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["target_network_cidr"] = cty.StringVal(p.TargetNetworkCidr)
}