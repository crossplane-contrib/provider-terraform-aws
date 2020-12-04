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

func EncodeWafregionalRateBasedRule(r WafregionalRateBasedRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalRateBasedRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRateBasedRule_RateKey(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRateBasedRule_RateLimit(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRateBasedRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRateBasedRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRateBasedRule_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRateBasedRule_Predicate(r.Spec.ForProvider.Predicate, ctyVal)
	EncodeWafregionalRateBasedRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafregionalRateBasedRule_Name(p WafregionalRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalRateBasedRule_RateKey(p WafregionalRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["rate_key"] = cty.StringVal(p.RateKey)
}

func EncodeWafregionalRateBasedRule_RateLimit(p WafregionalRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["rate_limit"] = cty.NumberIntVal(p.RateLimit)
}

func EncodeWafregionalRateBasedRule_Tags(p WafregionalRateBasedRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafregionalRateBasedRule_Id(p WafregionalRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalRateBasedRule_MetricName(p WafregionalRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafregionalRateBasedRule_Predicate(p Predicate, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalRateBasedRule_Predicate_DataId(p, ctyVal)
	EncodeWafregionalRateBasedRule_Predicate_Negated(p, ctyVal)
	EncodeWafregionalRateBasedRule_Predicate_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["predicate"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalRateBasedRule_Predicate_DataId(p Predicate, vals map[string]cty.Value) {
	vals["data_id"] = cty.StringVal(p.DataId)
}

func EncodeWafregionalRateBasedRule_Predicate_Negated(p Predicate, vals map[string]cty.Value) {
	vals["negated"] = cty.BoolVal(p.Negated)
}

func EncodeWafregionalRateBasedRule_Predicate_Type(p Predicate, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalRateBasedRule_Arn(p WafregionalRateBasedRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}