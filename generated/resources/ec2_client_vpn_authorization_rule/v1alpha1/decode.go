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
	r, ok := mr.(*Ec2ClientVpnAuthorizationRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2ClientVpnAuthorizationRule(r, ctyValue)
}

func DecodeEc2ClientVpnAuthorizationRule(prev *Ec2ClientVpnAuthorizationRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2ClientVpnAuthorizationRule_AccessGroupId(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnAuthorizationRule_AuthorizeAllGroups(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnAuthorizationRule_ClientVpnEndpointId(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnAuthorizationRule_Description(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnAuthorizationRule_Id(&new.Spec.ForProvider, valMap)
	DecodeEc2ClientVpnAuthorizationRule_TargetNetworkCidr(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnAuthorizationRule_AccessGroupId(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	p.AccessGroupId = ctwhy.ValueAsString(vals["access_group_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnAuthorizationRule_AuthorizeAllGroups(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	p.AuthorizeAllGroups = ctwhy.ValueAsBool(vals["authorize_all_groups"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnAuthorizationRule_ClientVpnEndpointId(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	p.ClientVpnEndpointId = ctwhy.ValueAsString(vals["client_vpn_endpoint_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnAuthorizationRule_Description(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnAuthorizationRule_Id(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2ClientVpnAuthorizationRule_TargetNetworkCidr(p *Ec2ClientVpnAuthorizationRuleParameters, vals map[string]cty.Value) {
	p.TargetNetworkCidr = ctwhy.ValueAsString(vals["target_network_cidr"])
}