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
	r, ok := mr.(*ConfigOrganizationCustomRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeConfigOrganizationCustomRule(r, ctyValue)
}

func DecodeConfigOrganizationCustomRule(prev *ConfigOrganizationCustomRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeConfigOrganizationCustomRule_ExcludedAccounts(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_LambdaFunctionArn(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_ResourceIdScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_TagValueScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_Description(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_Id(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_InputParameters(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_MaximumExecutionFrequency(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_Name(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_ResourceTypesScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_TagKeyScope(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_TriggerTypes(&new.Spec.ForProvider, valMap)
	DecodeConfigOrganizationCustomRule_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeConfigOrganizationCustomRule_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveCollectionTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_ExcludedAccounts(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["excluded_accounts"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ExcludedAccounts = goVals
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_LambdaFunctionArn(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.LambdaFunctionArn = ctwhy.ValueAsString(vals["lambda_function_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_ResourceIdScope(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.ResourceIdScope = ctwhy.ValueAsString(vals["resource_id_scope"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_TagValueScope(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.TagValueScope = ctwhy.ValueAsString(vals["tag_value_scope"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Description(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Id(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_InputParameters(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.InputParameters = ctwhy.ValueAsString(vals["input_parameters"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_MaximumExecutionFrequency(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.MaximumExecutionFrequency = ctwhy.ValueAsString(vals["maximum_execution_frequency"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Name(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_ResourceTypesScope(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["resource_types_scope"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ResourceTypesScope = goVals
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_TagKeyScope(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	p.TagKeyScope = ctwhy.ValueAsString(vals["tag_key_scope"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_TriggerTypes(p *ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["trigger_types"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.TriggerTypes = goVals
}

//containerTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeConfigOrganizationCustomRule_Timeouts_Create(p, valMap)
	DecodeConfigOrganizationCustomRule_Timeouts_Delete(p, valMap)
	DecodeConfigOrganizationCustomRule_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeConfigOrganizationCustomRule_Arn(p *ConfigOrganizationCustomRuleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}