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

package v1alpha1func EncodeWafregionalWebAcl(r WafregionalWebAcl) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafregionalWebAcl_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalWebAcl_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalWebAcl_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalWebAcl_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalWebAcl_DefaultAction(r.Spec.ForProvider.DefaultAction, ctyVal)
	EncodeWafregionalWebAcl_LoggingConfiguration(r.Spec.ForProvider.LoggingConfiguration, ctyVal)
	EncodeWafregionalWebAcl_Rule(r.Spec.ForProvider.Rule, ctyVal)
	EncodeWafregionalWebAcl_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeWafregionalWebAcl_MetricName(p *WafregionalWebAclParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafregionalWebAcl_Name(p *WafregionalWebAclParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalWebAcl_Tags(p *WafregionalWebAclParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafregionalWebAcl_Id(p *WafregionalWebAclParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalWebAcl_DefaultAction(p *DefaultAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DefaultAction {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalWebAcl_DefaultAction_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["default_action"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalWebAcl_DefaultAction_Type(p *DefaultAction, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalWebAcl_LoggingConfiguration(p *LoggingConfiguration, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LoggingConfiguration {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalWebAcl_LoggingConfiguration_LogDestination(v, ctyVal)
		EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields(v.RedactedFields, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["logging_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalWebAcl_LoggingConfiguration_LogDestination(p *LoggingConfiguration, vals map[string]cty.Value) {
	vals["log_destination"] = cty.StringVal(p.LogDestination)
}

func EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields(p *RedactedFields, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RedactedFields {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["redacted_fields"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Data(v, ctyVal)
		EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Data(p *FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafregionalWebAcl_LoggingConfiguration_RedactedFields_FieldToMatch_Type(p *FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalWebAcl_Rule(p *Rule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Rule {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalWebAcl_Rule_Priority(v, ctyVal)
		EncodeWafregionalWebAcl_Rule_RuleId(v, ctyVal)
		EncodeWafregionalWebAcl_Rule_Type(v, ctyVal)
		EncodeWafregionalWebAcl_Rule_OverrideAction(v.OverrideAction, ctyVal)
		EncodeWafregionalWebAcl_Rule_Action(v.Action, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["rule"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalWebAcl_Rule_Priority(p *Rule, vals map[string]cty.Value) {
	vals["priority"] = cty.IntVal(p.Priority)
}

func EncodeWafregionalWebAcl_Rule_RuleId(p *Rule, vals map[string]cty.Value) {
	vals["rule_id"] = cty.StringVal(p.RuleId)
}

func EncodeWafregionalWebAcl_Rule_Type(p *Rule, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalWebAcl_Rule_OverrideAction(p *OverrideAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OverrideAction {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalWebAcl_Rule_OverrideAction_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["override_action"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalWebAcl_Rule_OverrideAction_Type(p *OverrideAction, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalWebAcl_Rule_Action(p *Action, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Action {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalWebAcl_Rule_Action_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalWebAcl_Rule_Action_Type(p *Action, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalWebAcl_Arn(p *WafregionalWebAclObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}