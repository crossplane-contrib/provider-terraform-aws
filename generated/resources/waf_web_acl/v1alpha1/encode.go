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

func EncodeWafWebAcl(r WafWebAcl) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafWebAcl_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeWafWebAcl_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafWebAcl_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafWebAcl_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafWebAcl_DefaultAction(r.Spec.ForProvider.DefaultAction, ctyVal)
	EncodeWafWebAcl_LoggingConfiguration(r.Spec.ForProvider.LoggingConfiguration, ctyVal)
	EncodeWafWebAcl_Rules(r.Spec.ForProvider.Rules, ctyVal)
	EncodeWafWebAcl_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafWebAcl_MetricName(p WafWebAclParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafWebAcl_Name(p WafWebAclParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafWebAcl_Tags(p WafWebAclParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafWebAcl_Id(p WafWebAclParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafWebAcl_DefaultAction(p DefaultAction, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafWebAcl_DefaultAction_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["default_action"] = cty.ListVal(valsForCollection)
}

func EncodeWafWebAcl_DefaultAction_Type(p DefaultAction, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafWebAcl_LoggingConfiguration(p LoggingConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafWebAcl_LoggingConfiguration_LogDestination(p, ctyVal)
	EncodeWafWebAcl_LoggingConfiguration_RedactedFields(p.RedactedFields, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["logging_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeWafWebAcl_LoggingConfiguration_LogDestination(p LoggingConfiguration, vals map[string]cty.Value) {
	vals["log_destination"] = cty.StringVal(p.LogDestination)
}

func EncodeWafWebAcl_LoggingConfiguration_RedactedFields(p RedactedFields, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch(p.FieldToMatch, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["redacted_fields"] = cty.ListVal(valsForCollection)
}

func EncodeWafWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch(p []FieldToMatch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeWafWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Data(v, ctyVal)
		EncodeWafWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.SetVal(valsForCollection)
}

func EncodeWafWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafWebAcl_Rules(p Rules, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafWebAcl_Rules_Type(p, ctyVal)
	EncodeWafWebAcl_Rules_Priority(p, ctyVal)
	EncodeWafWebAcl_Rules_RuleId(p, ctyVal)
	EncodeWafWebAcl_Rules_Action(p.Action, ctyVal)
	EncodeWafWebAcl_Rules_OverrideAction(p.OverrideAction, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["rules"] = cty.SetVal(valsForCollection)
}

func EncodeWafWebAcl_Rules_Type(p Rules, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafWebAcl_Rules_Priority(p Rules, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeWafWebAcl_Rules_RuleId(p Rules, vals map[string]cty.Value) {
	vals["rule_id"] = cty.StringVal(p.RuleId)
}

func EncodeWafWebAcl_Rules_Action(p Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafWebAcl_Rules_Action_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeWafWebAcl_Rules_Action_Type(p Action, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafWebAcl_Rules_OverrideAction(p OverrideAction, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafWebAcl_Rules_OverrideAction_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["override_action"] = cty.ListVal(valsForCollection)
}

func EncodeWafWebAcl_Rules_OverrideAction_Type(p OverrideAction, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafWebAcl_Arn(p WafWebAclObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}