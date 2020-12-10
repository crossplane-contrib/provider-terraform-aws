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
	r, ok := mr.(*WafRateBasedRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafRateBasedRule.")
	}
	return EncodeWafRateBasedRule(*r), nil
}

func EncodeWafRateBasedRule(r WafRateBasedRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafRateBasedRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafRateBasedRule_RateKey(r.Spec.ForProvider, ctyVal)
	EncodeWafRateBasedRule_RateLimit(r.Spec.ForProvider, ctyVal)
	EncodeWafRateBasedRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafRateBasedRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafRateBasedRule_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeWafRateBasedRule_Predicates(r.Spec.ForProvider.Predicates, ctyVal)
	EncodeWafRateBasedRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafRateBasedRule_Name(p WafRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafRateBasedRule_RateKey(p WafRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["rate_key"] = cty.StringVal(p.RateKey)
}

func EncodeWafRateBasedRule_RateLimit(p WafRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["rate_limit"] = cty.NumberIntVal(p.RateLimit)
}

func EncodeWafRateBasedRule_Tags(p WafRateBasedRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafRateBasedRule_Id(p WafRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafRateBasedRule_MetricName(p WafRateBasedRuleParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafRateBasedRule_Predicates(p Predicates, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafRateBasedRule_Predicates_Type(p, ctyVal)
	EncodeWafRateBasedRule_Predicates_DataId(p, ctyVal)
	EncodeWafRateBasedRule_Predicates_Negated(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["predicates"] = cty.SetVal(valsForCollection)
}

func EncodeWafRateBasedRule_Predicates_Type(p Predicates, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafRateBasedRule_Predicates_DataId(p Predicates, vals map[string]cty.Value) {
	vals["data_id"] = cty.StringVal(p.DataId)
}

func EncodeWafRateBasedRule_Predicates_Negated(p Predicates, vals map[string]cty.Value) {
	vals["negated"] = cty.BoolVal(p.Negated)
}

func EncodeWafRateBasedRule_Arn(p WafRateBasedRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}