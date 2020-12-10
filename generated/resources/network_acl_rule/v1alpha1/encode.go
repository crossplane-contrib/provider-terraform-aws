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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*NetworkAclRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a NetworkAclRule.")
	}
	return EncodeNetworkAclRule(*r), nil
}

func EncodeNetworkAclRule(r NetworkAclRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeNetworkAclRule_IcmpCode(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_Ipv6CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_RuleAction(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_ToPort(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_Egress(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_FromPort(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_IcmpType(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_NetworkAclId(r.Spec.ForProvider, ctyVal)
	EncodeNetworkAclRule_RuleNumber(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeNetworkAclRule_IcmpCode(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["icmp_code"] = cty.StringVal(p.IcmpCode)
}

func EncodeNetworkAclRule_Id(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeNetworkAclRule_Ipv6CidrBlock(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeNetworkAclRule_Protocol(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeNetworkAclRule_RuleAction(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["rule_action"] = cty.StringVal(p.RuleAction)
}

func EncodeNetworkAclRule_ToPort(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["to_port"] = cty.NumberIntVal(p.ToPort)
}

func EncodeNetworkAclRule_CidrBlock(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeNetworkAclRule_Egress(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["egress"] = cty.BoolVal(p.Egress)
}

func EncodeNetworkAclRule_FromPort(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["from_port"] = cty.NumberIntVal(p.FromPort)
}

func EncodeNetworkAclRule_IcmpType(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["icmp_type"] = cty.StringVal(p.IcmpType)
}

func EncodeNetworkAclRule_NetworkAclId(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["network_acl_id"] = cty.StringVal(p.NetworkAclId)
}

func EncodeNetworkAclRule_RuleNumber(p NetworkAclRuleParameters, vals map[string]cty.Value) {
	vals["rule_number"] = cty.NumberIntVal(p.RuleNumber)
}