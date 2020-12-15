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
	r, ok := mr.(*SecurityGroupRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSecurityGroupRule(r, ctyValue)
}

func DecodeSecurityGroupRule(prev *SecurityGroupRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSecurityGroupRule_FromPort(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_Ipv6CidrBlocks(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_PrefixListIds(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_Protocol(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_SecurityGroupId(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_ToPort(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_Type(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_Description(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_Id(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_Self(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_SourceSecurityGroupId(&new.Spec.ForProvider, valMap)
	DecodeSecurityGroupRule_CidrBlocks(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_FromPort(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.FromPort = ctwhy.ValueAsInt64(vals["from_port"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeSecurityGroupRule_Ipv6CidrBlocks(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["ipv6_cidr_blocks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Ipv6CidrBlocks = goVals
}

//primitiveCollectionTypeDecodeTemplate
func DecodeSecurityGroupRule_PrefixListIds(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["prefix_list_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.PrefixListIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_Protocol(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.Protocol = ctwhy.ValueAsString(vals["protocol"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_SecurityGroupId(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.SecurityGroupId = ctwhy.ValueAsString(vals["security_group_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_ToPort(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.ToPort = ctwhy.ValueAsInt64(vals["to_port"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_Type(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_Description(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_Id(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_Self(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.Self = ctwhy.ValueAsBool(vals["self"])
}

//primitiveTypeDecodeTemplate
func DecodeSecurityGroupRule_SourceSecurityGroupId(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	p.SourceSecurityGroupId = ctwhy.ValueAsString(vals["source_security_group_id"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeSecurityGroupRule_CidrBlocks(p *SecurityGroupRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsList(vals["cidr_blocks"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.CidrBlocks = goVals
}