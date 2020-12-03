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

package v1alpha1func EncodeRoute53ResolverRule(r Route53ResolverRule) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRoute53ResolverRule_DomainName(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_ResolverEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_RuleType(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRule_TargetIp(r.Spec.ForProvider.TargetIp, ctyVal)
	EncodeRoute53ResolverRule_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRoute53ResolverRule_Arn(r.Status.AtProvider, ctyVal)
	EncodeRoute53ResolverRule_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeRoute53ResolverRule_ShareStatus(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeRoute53ResolverRule_DomainName(p *Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["domain_name"] = cty.StringVal(p.DomainName)
}

func EncodeRoute53ResolverRule_Id(p *Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53ResolverRule_ResolverEndpointId(p *Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["resolver_endpoint_id"] = cty.StringVal(p.ResolverEndpointId)
}

func EncodeRoute53ResolverRule_RuleType(p *Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["rule_type"] = cty.StringVal(p.RuleType)
}

func EncodeRoute53ResolverRule_Tags(p *Route53ResolverRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRoute53ResolverRule_Name(p *Route53ResolverRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53ResolverRule_TargetIp(p *TargetIp, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.TargetIp {
		ctyVal = make(map[string]cty.Value)
		EncodeRoute53ResolverRule_TargetIp_Ip(v, ctyVal)
		EncodeRoute53ResolverRule_TargetIp_Port(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["target_ip"] = cty.SetVal(valsForCollection)
}

func EncodeRoute53ResolverRule_TargetIp_Ip(p *TargetIp, vals map[string]cty.Value) {
	vals["ip"] = cty.StringVal(p.Ip)
}

func EncodeRoute53ResolverRule_TargetIp_Port(p *TargetIp, vals map[string]cty.Value) {
	vals["port"] = cty.IntVal(p.Port)
}

func EncodeRoute53ResolverRule_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeRoute53ResolverRule_Timeouts_Create(p, ctyVal)
	EncodeRoute53ResolverRule_Timeouts_Delete(p, ctyVal)
	EncodeRoute53ResolverRule_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverRule_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRoute53ResolverRule_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRoute53ResolverRule_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeRoute53ResolverRule_Arn(p *Route53ResolverRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeRoute53ResolverRule_OwnerId(p *Route53ResolverRuleObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeRoute53ResolverRule_ShareStatus(p *Route53ResolverRuleObservation, vals map[string]cty.Value) {
	vals["share_status"] = cty.StringVal(p.ShareStatus)
}