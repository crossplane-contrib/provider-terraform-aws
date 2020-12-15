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
	r, ok := mr.(*Ec2ClientVpnAuthorizationRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2ClientVpnAuthorizationRule.")
	}
	return EncodeEc2ClientVpnAuthorizationRule(*r), nil
}

func EncodeEc2ClientVpnAuthorizationRule(r Ec2ClientVpnAuthorizationRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2ClientVpnAuthorizationRule_TargetNetworkCidr(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_AccessGroupId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_AuthorizeAllGroups(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_ClientVpnEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnAuthorizationRule_Description(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2ClientVpnAuthorizationRule_TargetNetworkCidr(p Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["target_network_cidr"] = cty.StringVal(p.TargetNetworkCidr)
}

func EncodeEc2ClientVpnAuthorizationRule_AccessGroupId(p Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["access_group_id"] = cty.StringVal(p.AccessGroupId)
}

func EncodeEc2ClientVpnAuthorizationRule_AuthorizeAllGroups(p Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["authorize_all_groups"] = cty.BoolVal(p.AuthorizeAllGroups)
}

func EncodeEc2ClientVpnAuthorizationRule_ClientVpnEndpointId(p Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["client_vpn_endpoint_id"] = cty.StringVal(p.ClientVpnEndpointId)
}

func EncodeEc2ClientVpnAuthorizationRule_Description(p Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}