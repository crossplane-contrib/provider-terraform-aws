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
	r, ok := mr.(*Route53ResolverRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53ResolverRule.")
	}
	return EncodeRoute53ResolverRule(*r), nil
}

func EncodeRoute53ResolverRule(r Route53ResolverRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_RuleType(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_ResolverEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_TargetIp(r.Spec.ForProvider.TargetIp, ctyVal)
	EncodeRoute53ResolverRule_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRoute53ResolverRule_ShareStatus(r.Status.AtProvider, ctyVal)
	EncodeRoute53ResolverRule_Arn(r.Status.AtProvider, ctyVal)
	EncodeRoute53ResolverRule_OwnerId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverRule_Tags(p Route53ResolverRuleParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRoute53ResolverRule_RuleType(p Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["rule_type"] = cty.StringVal(p.RuleType)
}

func EncodeRoute53ResolverRule_DomainName(p Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeRoute53ResolverRule_Id(p Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53ResolverRule_Name(p Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53ResolverRule_ResolverEndpointId(p Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["resolver_endpoint_id"] = cty.StringVal(p.ResolverEndpointId)
}

func EncodeRoute53ResolverRule_TargetIp(p TargetIp, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverRule_TargetIp_Ip(p, ctyVal)
	EncodeRoute53ResolverRule_TargetIp_Port(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["target_ip"] = cty.SetVal(valsForCollection)
}

func EncodeRoute53ResolverRule_TargetIp_Ip(p TargetIp, vals map[string]cty.Value) {
	vals["ip"] = cty.StringVal(p.Ip)
}

func EncodeRoute53ResolverRule_TargetIp_Port(p TargetIp, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeRoute53ResolverRule_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverRule_Timeouts_Create(p, ctyVal)
	EncodeRoute53ResolverRule_Timeouts_Delete(p, ctyVal)
	EncodeRoute53ResolverRule_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverRule_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRoute53ResolverRule_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRoute53ResolverRule_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeRoute53ResolverRule_ShareStatus(p Route53ResolverRuleObservation, vals map[string]cty.Value) {
	vals["share_status"] = cty.StringVal(p.ShareStatus)
}

func EncodeRoute53ResolverRule_Arn(p Route53ResolverRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeRoute53ResolverRule_OwnerId(p Route53ResolverRuleObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}