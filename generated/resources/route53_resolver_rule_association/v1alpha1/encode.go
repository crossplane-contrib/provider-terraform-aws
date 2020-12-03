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

package v1alpha1func EncodeRoute53ResolverRuleAssociation(r Route53ResolverRuleAssociation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRoute53ResolverRuleAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRuleAssociation_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRuleAssociation_ResolverRuleId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRuleAssociation_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRuleAssociation_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeRoute53ResolverRuleAssociation_Id(p *Route53ResolverRuleAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53ResolverRuleAssociation_Name(p *Route53ResolverRuleAssociationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53ResolverRuleAssociation_ResolverRuleId(p *Route53ResolverRuleAssociationParameters, vals map[string]cty.Value) {
	vals["resolver_rule_id"] = cty.StringVal(p.ResolverRuleId)
}

func EncodeRoute53ResolverRuleAssociation_VpcId(p *Route53ResolverRuleAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeRoute53ResolverRuleAssociation_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeRoute53ResolverRuleAssociation_Timeouts_Create(p, ctyVal)
	EncodeRoute53ResolverRuleAssociation_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverRuleAssociation_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRoute53ResolverRuleAssociation_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}