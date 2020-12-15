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
	r, ok := mr.(*NetworkAclRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeNetworkAclRule(r, ctyValue)
}

func DecodeNetworkAclRule(prev *NetworkAclRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeNetworkAclRule_IcmpCode(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_Ipv6CidrBlock(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_NetworkAclId(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_Protocol(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_ToPort(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_Egress(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_FromPort(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_IcmpType(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_Id(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_RuleAction(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_RuleNumber(&new.Spec.ForProvider, valMap)
	DecodeNetworkAclRule_CidrBlock(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_IcmpCode(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.IcmpCode = ctwhy.ValueAsString(vals["icmp_code"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_Ipv6CidrBlock(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.Ipv6CidrBlock = ctwhy.ValueAsString(vals["ipv6_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_NetworkAclId(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.NetworkAclId = ctwhy.ValueAsString(vals["network_acl_id"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_Protocol(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.Protocol = ctwhy.ValueAsString(vals["protocol"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_ToPort(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.ToPort = ctwhy.ValueAsInt64(vals["to_port"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_Egress(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.Egress = ctwhy.ValueAsBool(vals["egress"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_FromPort(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.FromPort = ctwhy.ValueAsInt64(vals["from_port"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_IcmpType(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.IcmpType = ctwhy.ValueAsString(vals["icmp_type"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_Id(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_RuleAction(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.RuleAction = ctwhy.ValueAsString(vals["rule_action"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_RuleNumber(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.RuleNumber = ctwhy.ValueAsInt64(vals["rule_number"])
}

//primitiveTypeDecodeTemplate
func DecodeNetworkAclRule_CidrBlock(p *NetworkAclRuleParameters, vals map[string]cty.Value) {
	p.CidrBlock = ctwhy.ValueAsString(vals["cidr_block"])
}