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

package v1alpha1func EncodeWafregionalRegexMatchSet(r WafregionalRegexMatchSet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafregionalRegexMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRegexMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalRegexMatchSet_RegexMatchTuple(r.Spec.ForProvider.RegexMatchTuple, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeWafregionalRegexMatchSet_Id(p *WafregionalRegexMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalRegexMatchSet_Name(p *WafregionalRegexMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalRegexMatchSet_RegexMatchTuple(p *RegexMatchTuple, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegexMatchTuple {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalRegexMatchSet_RegexMatchTuple_RegexPatternSetId(v, ctyVal)
		EncodeWafregionalRegexMatchSet_RegexMatchTuple_TextTransformation(v, ctyVal)
		EncodeWafregionalRegexMatchSet_RegexMatchTuple_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regex_match_tuple"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalRegexMatchSet_RegexMatchTuple_RegexPatternSetId(p *RegexMatchTuple, vals map[string]cty.Value) {
	vals["regex_pattern_set_id"] = cty.StringVal(p.RegexPatternSetId)
}

func EncodeWafregionalRegexMatchSet_RegexMatchTuple_TextTransformation(p *RegexMatchTuple, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafregionalRegexMatchSet_RegexMatchTuple_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalRegexMatchSet_RegexMatchTuple_FieldToMatch_Data(v, ctyVal)
		EncodeWafregionalRegexMatchSet_RegexMatchTuple_FieldToMatch_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalRegexMatchSet_RegexMatchTuple_FieldToMatch_Data(p *FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafregionalRegexMatchSet_RegexMatchTuple_FieldToMatch_Type(p *FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}