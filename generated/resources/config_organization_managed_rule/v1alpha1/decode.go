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
	r, ok := mr.(*ConfigOrganizationManagedRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeConfigOrganizationManagedRule(r, ctyValue)
}

func DecodeConfigOrganizationManagedRule(prev *ConfigOrganizationManagedRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeConfigOrganizationManagedRule_ResourceIdScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_ResourceTypesScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_RuleIdentifier(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_ExcludedAccounts(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_InputParameters(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_MaximumExecutionFrequency(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_Name(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_TagKeyScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_TagValueScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_Description(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationManagedRule_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeConfigOrganizationManagedRule_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_ResourceIdScope(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.ResourceIdScope = ctwhy.ValueAsString(vals["resource_id_scope"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_ResourceTypesScope(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["resource_types_scope"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ResourceTypesScope = goVals
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_RuleIdentifier(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.RuleIdentifier = ctwhy.ValueAsString(vals["rule_identifier"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_ExcludedAccounts(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["excluded_accounts"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ExcludedAccounts = goVals
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_InputParameters(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.InputParameters = ctwhy.ValueAsString(vals["input_parameters"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_MaximumExecutionFrequency(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.MaximumExecutionFrequency = ctwhy.ValueAsString(vals["maximum_execution_frequency"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_Name(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_TagKeyScope(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.TagKeyScope = ctwhy.ValueAsString(vals["tag_key_scope"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_TagValueScope(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.TagValueScope = ctwhy.ValueAsString(vals["tag_value_scope"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_Description(p *ConfigOrganizationManagedRuleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//containerTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeConfigOrganizationManagedRule_Timeouts_Create(p, valMap)
	DecodeConfigOrganizationManagedRule_Timeouts_Delete(p, valMap)
	DecodeConfigOrganizationManagedRule_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationManagedRule_Arn(p *ConfigOrganizationManagedRuleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}