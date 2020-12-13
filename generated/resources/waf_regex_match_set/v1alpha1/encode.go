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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*WafRegexMatchSet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafRegexMatchSet.")
	}
	return EncodeWafRegexMatchSet(*r), nil
}

func EncodeWafRegexMatchSet(r WafRegexMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafRegexMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafRegexMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafRegexMatchSet_RegexMatchTuple(r.Spec.ForProvider.RegexMatchTuple, ctyVal)
	EncodeWafRegexMatchSet_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeWafRegexMatchSet_Id(p WafRegexMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafRegexMatchSet_Name(p WafRegexMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafRegexMatchSet_RegexMatchTuple(p RegexMatchTuple, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafRegexMatchSet_RegexMatchTuple_TextTransformation(p, ctyVal)
	EncodeWafRegexMatchSet_RegexMatchTuple_RegexPatternSetId(p, ctyVal)
	EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch(p.FieldToMatch, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["regex_match_tuple"] = cty.SetVal(valsForCollection)
}

func EncodeWafRegexMatchSet_RegexMatchTuple_TextTransformation(p RegexMatchTuple, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafRegexMatchSet_RegexMatchTuple_RegexPatternSetId(p RegexMatchTuple, vals map[string]cty.Value) {
	vals["regex_pattern_set_id"] = cty.StringVal(p.RegexPatternSetId)
}

func EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch(p FieldToMatch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Type(p, ctyVal)
	EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Data(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafRegexMatchSet_Arn(p WafRegexMatchSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}