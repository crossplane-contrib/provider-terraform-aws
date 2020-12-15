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
	r, ok := mr.(*Route53ResolverRuleAssociation)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route53ResolverRuleAssociation.")
	}
	return EncodeRoute53ResolverRuleAssociation(*r), nil
}

func EncodeRoute53ResolverRuleAssociation(r Route53ResolverRuleAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverRuleAssociation_Name(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRuleAssociation_ResolverRuleId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRuleAssociation_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53ResolverRuleAssociation_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverRuleAssociation_Name(p Route53ResolverRuleAssociationParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeRoute53ResolverRuleAssociation_ResolverRuleId(p Route53ResolverRuleAssociationParameters, vals map[string]cty.Value) {
	vals["resolver_rule_id"] = cty.StringVal(p.ResolverRuleId)
}

func EncodeRoute53ResolverRuleAssociation_VpcId(p Route53ResolverRuleAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeRoute53ResolverRuleAssociation_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53ResolverRuleAssociation_Timeouts_Create(p, ctyVal)
	EncodeRoute53ResolverRuleAssociation_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRoute53ResolverRuleAssociation_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRoute53ResolverRuleAssociation_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}