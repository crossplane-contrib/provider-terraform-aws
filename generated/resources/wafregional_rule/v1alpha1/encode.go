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
	r, ok := mr.(*WafregionalRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafregionalRule.")
	}
	return EncodeWafregionalRule(*r), nil
}

func EncodeWafregionalRule(r WafregionalRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRule_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRule_Predicate(r.Spec.ForProvider.Predicate, ctyVal)
	EncodeWafregionalRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafregionalRule_Id(p WafregionalRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalRule_MetricName(p WafregionalRuleParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafregionalRule_Name(p WafregionalRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalRule_Tags(p WafregionalRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafregionalRule_Predicate(p Predicate, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalRule_Predicate_DataId(p, ctyVal)
	EncodeWafregionalRule_Predicate_Negated(p, ctyVal)
	EncodeWafregionalRule_Predicate_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["predicate"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalRule_Predicate_DataId(p Predicate, vals map[string]cty.Value) {
	vals["data_id"] = cty.StringVal(p.DataId)
}

func EncodeWafregionalRule_Predicate_Negated(p Predicate, vals map[string]cty.Value) {
	vals["negated"] = cty.BoolVal(p.Negated)
}

func EncodeWafregionalRule_Predicate_Type(p Predicate, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalRule_Arn(p WafregionalRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}