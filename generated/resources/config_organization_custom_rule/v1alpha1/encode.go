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

func EncodeConfigOrganizationCustomRule(r ConfigOrganizationCustomRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigOrganizationCustomRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_MaximumExecutionFrequency(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_TagKeyScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_TriggerTypes(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_ResourceIdScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_ResourceTypesScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_TagValueScope(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_ExcludedAccounts(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_InputParameters(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_LambdaFunctionArn(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeConfigOrganizationCustomRule_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeConfigOrganizationCustomRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeConfigOrganizationCustomRule_Description(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeConfigOrganizationCustomRule_MaximumExecutionFrequency(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["maximum_execution_frequency"] = cty.StringVal(p.MaximumExecutionFrequency)
}

func EncodeConfigOrganizationCustomRule_TagKeyScope(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["tag_key_scope"] = cty.StringVal(p.TagKeyScope)
}

func EncodeConfigOrganizationCustomRule_TriggerTypes(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.TriggerTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["trigger_types"] = cty.SetVal(colVals)
}

func EncodeConfigOrganizationCustomRule_ResourceIdScope(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["resource_id_scope"] = cty.StringVal(p.ResourceIdScope)
}

func EncodeConfigOrganizationCustomRule_ResourceTypesScope(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ResourceTypesScope {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["resource_types_scope"] = cty.SetVal(colVals)
}

func EncodeConfigOrganizationCustomRule_TagValueScope(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["tag_value_scope"] = cty.StringVal(p.TagValueScope)
}

func EncodeConfigOrganizationCustomRule_ExcludedAccounts(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ExcludedAccounts {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["excluded_accounts"] = cty.SetVal(colVals)
}

func EncodeConfigOrganizationCustomRule_Id(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigOrganizationCustomRule_InputParameters(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["input_parameters"] = cty.StringVal(p.InputParameters)
}

func EncodeConfigOrganizationCustomRule_LambdaFunctionArn(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["lambda_function_arn"] = cty.StringVal(p.LambdaFunctionArn)
}

func EncodeConfigOrganizationCustomRule_Name(p ConfigOrganizationCustomRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeConfigOrganizationCustomRule_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigOrganizationCustomRule_Timeouts_Create(p, ctyVal)
	EncodeConfigOrganizationCustomRule_Timeouts_Delete(p, ctyVal)
	EncodeConfigOrganizationCustomRule_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeConfigOrganizationCustomRule_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeConfigOrganizationCustomRule_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeConfigOrganizationCustomRule_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeConfigOrganizationCustomRule_Arn(p ConfigOrganizationCustomRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}