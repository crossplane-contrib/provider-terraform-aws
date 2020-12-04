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

func EncodeWafRegexMatchSet(r WafRegexMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafRegexMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafRegexMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafRegexMatchSet_RegexMatchTuple(r.Spec.ForProvider.RegexMatchTuple, ctyVal)
	EncodeWafRegexMatchSet_Arn(r.Status.AtProvider, ctyVal)
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
	EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Data(p, ctyVal)
	EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafRegexMatchSet_RegexMatchTuple_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafRegexMatchSet_Arn(p WafRegexMatchSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}