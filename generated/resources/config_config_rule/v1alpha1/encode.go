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

func EncodeConfigConfigRule(r ConfigConfigRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeConfigConfigRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigRule_InputParameters(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigRule_MaximumExecutionFrequency(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeConfigConfigRule_Scope(r.Spec.ForProvider.Scope, ctyVal)
	EncodeConfigConfigRule_Source(r.Spec.ForProvider.Source, ctyVal)
	EncodeConfigConfigRule_Arn(r.Status.AtProvider, ctyVal)
	EncodeConfigConfigRule_RuleId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeConfigConfigRule_Description(p ConfigConfigRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeConfigConfigRule_Id(p ConfigConfigRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeConfigConfigRule_InputParameters(p ConfigConfigRuleParameters, vals map[string]cty.Value) {
	vals["input_parameters"] = cty.StringVal(p.InputParameters)
}

func EncodeConfigConfigRule_MaximumExecutionFrequency(p ConfigConfigRuleParameters, vals map[string]cty.Value) {
	vals["maximum_execution_frequency"] = cty.StringVal(p.MaximumExecutionFrequency)
}

func EncodeConfigConfigRule_Name(p ConfigConfigRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeConfigConfigRule_Tags(p ConfigConfigRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeConfigConfigRule_Scope(p Scope, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeConfigConfigRule_Scope_ComplianceResourceId(p, ctyVal)
	EncodeConfigConfigRule_Scope_ComplianceResourceTypes(p, ctyVal)
	EncodeConfigConfigRule_Scope_TagKey(p, ctyVal)
	EncodeConfigConfigRule_Scope_TagValue(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["scope"] = cty.ListVal(valsForCollection)
}

func EncodeConfigConfigRule_Scope_ComplianceResourceId(p Scope, vals map[string]cty.Value) {
	vals["compliance_resource_id"] = cty.StringVal(p.ComplianceResourceId)
}

func EncodeConfigConfigRule_Scope_ComplianceResourceTypes(p Scope, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ComplianceResourceTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["compliance_resource_types"] = cty.SetVal(colVals)
}

func EncodeConfigConfigRule_Scope_TagKey(p Scope, vals map[string]cty.Value) {
	vals["tag_key"] = cty.StringVal(p.TagKey)
}

func EncodeConfigConfigRule_Scope_TagValue(p Scope, vals map[string]cty.Value) {
	vals["tag_value"] = cty.StringVal(p.TagValue)
}

func EncodeConfigConfigRule_Source(p Source, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeConfigConfigRule_Source_Owner(p, ctyVal)
	EncodeConfigConfigRule_Source_SourceIdentifier(p, ctyVal)
	EncodeConfigConfigRule_Source_SourceDetail(p.SourceDetail, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["source"] = cty.ListVal(valsForCollection)
}

func EncodeConfigConfigRule_Source_Owner(p Source, vals map[string]cty.Value) {
	vals["owner"] = cty.StringVal(p.Owner)
}

func EncodeConfigConfigRule_Source_SourceIdentifier(p Source, vals map[string]cty.Value) {
	vals["source_identifier"] = cty.StringVal(p.SourceIdentifier)
}

func EncodeConfigConfigRule_Source_SourceDetail(p []SourceDetail, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeConfigConfigRule_Source_SourceDetail_MessageType(v, ctyVal)
		EncodeConfigConfigRule_Source_SourceDetail_EventSource(v, ctyVal)
		EncodeConfigConfigRule_Source_SourceDetail_MaximumExecutionFrequency(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["source_detail"] = cty.SetVal(valsForCollection)
}

func EncodeConfigConfigRule_Source_SourceDetail_MessageType(p SourceDetail, vals map[string]cty.Value) {
	vals["message_type"] = cty.StringVal(p.MessageType)
}

func EncodeConfigConfigRule_Source_SourceDetail_EventSource(p SourceDetail, vals map[string]cty.Value) {
	vals["event_source"] = cty.StringVal(p.EventSource)
}

func EncodeConfigConfigRule_Source_SourceDetail_MaximumExecutionFrequency(p SourceDetail, vals map[string]cty.Value) {
	vals["maximum_execution_frequency"] = cty.StringVal(p.MaximumExecutionFrequency)
}

func EncodeConfigConfigRule_Arn(p ConfigConfigRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeConfigConfigRule_RuleId(p ConfigConfigRuleObservation, vals map[string]cty.Value) {
	vals["rule_id"] = cty.StringVal(p.RuleId)
}