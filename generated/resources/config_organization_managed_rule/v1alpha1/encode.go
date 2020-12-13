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
	r, ok := mr.(*ConfigOrganizationManagedRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ConfigOrganizationManagedRule.")
	}
	return EncodeConfigOrganizationManagedRule(*r), nil
}

func EncodeConfigOrganizationManagedRule(r ConfigOrganizationManagedRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigOrganizationManagedRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_ResourceTypesScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_TagKeyScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_ExcludedAccounts(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_InputParameters(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_MaximumExecutionFrequency(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_ResourceIdScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_RuleIdentifier(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_TagValueScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationManagedRule_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeConfigOrganizationManagedRule_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeConfigOrganizationManagedRule_Name(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeConfigOrganizationManagedRule_ResourceTypesScope(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ResourceTypesScope {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["resource_types_scope"] = cty.SetVal(colVals)
}

func EncodeConfigOrganizationManagedRule_TagKeyScope(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["tag_key_scope"] = cty.StringVal(p.TagKeyScope)
}

func EncodeConfigOrganizationManagedRule_Description(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeConfigOrganizationManagedRule_ExcludedAccounts(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ExcludedAccounts {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["excluded_accounts"] = cty.SetVal(colVals)
}

func EncodeConfigOrganizationManagedRule_InputParameters(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["input_parameters"] = cty.StringVal(p.InputParameters)
}

func EncodeConfigOrganizationManagedRule_MaximumExecutionFrequency(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["maximum_execution_frequency"] = cty.StringVal(p.MaximumExecutionFrequency)
}

func EncodeConfigOrganizationManagedRule_ResourceIdScope(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["resource_id_scope"] = cty.StringVal(p.ResourceIdScope)
}

func EncodeConfigOrganizationManagedRule_RuleIdentifier(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["rule_identifier"] = cty.StringVal(p.RuleIdentifier)
}

func EncodeConfigOrganizationManagedRule_TagValueScope(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["tag_value_scope"] = cty.StringVal(p.TagValueScope)
}

func EncodeConfigOrganizationManagedRule_Id(p ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigOrganizationManagedRule_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigOrganizationManagedRule_Timeouts_Create(p, ctyVal)
	EncodeConfigOrganizationManagedRule_Timeouts_Delete(p, ctyVal)
	EncodeConfigOrganizationManagedRule_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeConfigOrganizationManagedRule_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeConfigOrganizationManagedRule_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeConfigOrganizationManagedRule_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeConfigOrganizationManagedRule_Arn(p ConfigOrganizationManagedRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}