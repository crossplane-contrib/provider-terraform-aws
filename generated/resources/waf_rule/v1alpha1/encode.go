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
	r, ok := mr.(*WafRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafRule.")
	}
	return EncodeWafRule(*r), nil
}

func EncodeWafRule(r WafRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafRule_MetricName(r.Spec.ForProvider, ctyVal)
	EncodeWafRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafRule_Predicates(r.Spec.ForProvider.Predicates, ctyVal)
	EncodeWafRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafRule_MetricName(p WafRuleParameters, vals map[string]cty.Value) {
	vals["metric_name"] = cty.StringVal(p.MetricName)
}

func EncodeWafRule_Name(p WafRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafRule_Tags(p WafRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafRule_Id(p WafRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafRule_Predicates(p Predicates, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafRule_Predicates_DataId(p, ctyVal)
	EncodeWafRule_Predicates_Negated(p, ctyVal)
	EncodeWafRule_Predicates_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["predicates"] = cty.SetVal(valsForCollection)
}

func EncodeWafRule_Predicates_DataId(p Predicates, vals map[string]cty.Value) {
	vals["data_id"] = cty.StringVal(p.DataId)
}

func EncodeWafRule_Predicates_Negated(p Predicates, vals map[string]cty.Value) {
	vals["negated"] = cty.BoolVal(p.Negated)
}

func EncodeWafRule_Predicates_Type(p Predicates, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafRule_Arn(p WafRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}