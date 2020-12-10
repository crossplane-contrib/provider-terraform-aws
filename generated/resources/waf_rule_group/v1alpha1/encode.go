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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*WafRuleGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafRuleGroup.")
	}
	return EncodeWafRuleGroup(*r), nil
}

func EncodeWafRuleGroup(r WafRuleGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafRuleGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafRuleGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafRuleGroup_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeWafRuleGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafRuleGroup_ActivatedRule(r.Spec.ForProvider.ActivatedRule, ctyVal)
	EncodeWafRuleGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafRuleGroup_Tags(p WafRuleGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafRuleGroup_Id(p WafRuleGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafRuleGroup_MetricName(p WafRuleGroupParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafRuleGroup_Name(p WafRuleGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafRuleGroup_ActivatedRule(p ActivatedRule, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafRuleGroup_ActivatedRule_Priority(p, ctyVal)
	EncodeWafRuleGroup_ActivatedRule_RuleId(p, ctyVal)
	EncodeWafRuleGroup_ActivatedRule_Type(p, ctyVal)
	EncodeWafRuleGroup_ActivatedRule_Action(p.Action, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["activated_rule"] = cty.SetVal(valsForCollection)
}

func EncodeWafRuleGroup_ActivatedRule_Priority(p ActivatedRule, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeWafRuleGroup_ActivatedRule_RuleId(p ActivatedRule, vals map[string]cty.Value) {
	vals["rule_id"] = cty.StringVal(p.RuleId)
}

func EncodeWafRuleGroup_ActivatedRule_Type(p ActivatedRule, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafRuleGroup_ActivatedRule_Action(p Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafRuleGroup_ActivatedRule_Action_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeWafRuleGroup_ActivatedRule_Action_Type(p Action, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafRuleGroup_Arn(p WafRuleGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}