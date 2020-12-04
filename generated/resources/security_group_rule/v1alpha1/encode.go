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
	"github.com/zclconf/go-cty/cty"
)

func EncodeSecurityGroupRule(r SecurityGroupRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSecurityGroupRule_SourceSecurityGroupId(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_Type(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_CidrBlocks(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_FromPort(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_Self(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_SecurityGroupId(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_ToPort(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_Ipv6CidrBlocks(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_PrefixListIds(r.Spec.ForProvider, ctyVal)
	EncodeSecurityGroupRule_Protocol(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeSecurityGroupRule_SourceSecurityGroupId(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["source_security_group_id"] = cty.StringVal(p.SourceSecurityGroupId)
}

func EncodeSecurityGroupRule_Type(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeSecurityGroupRule_CidrBlocks(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CidrBlocks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cidr_blocks"] = cty.ListVal(colVals)
}

func EncodeSecurityGroupRule_FromPort(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["from_port"] = cty.NumberIntVal(p.FromPort)
}

func EncodeSecurityGroupRule_Id(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSecurityGroupRule_Self(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["self"] = cty.BoolVal(p.Self)
}

func EncodeSecurityGroupRule_SecurityGroupId(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["security_group_id"] = cty.StringVal(p.SecurityGroupId)
}

func EncodeSecurityGroupRule_ToPort(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["to_port"] = cty.NumberIntVal(p.ToPort)
}

func EncodeSecurityGroupRule_Description(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeSecurityGroupRule_Ipv6CidrBlocks(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Ipv6CidrBlocks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ipv6_cidr_blocks"] = cty.ListVal(colVals)
}

func EncodeSecurityGroupRule_PrefixListIds(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PrefixListIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["prefix_list_ids"] = cty.ListVal(colVals)
}

func EncodeSecurityGroupRule_Protocol(p SecurityGroupRuleParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}