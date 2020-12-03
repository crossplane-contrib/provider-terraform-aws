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

package v1alpha1func EncodeWafregionalSqlInjectionMatchSet(r WafregionalSqlInjectionMatchSet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafregionalSqlInjectionMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalSqlInjectionMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple(r.Spec.ForProvider.SqlInjectionMatchTuple, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeWafregionalSqlInjectionMatchSet_Name(p *WafregionalSqlInjectionMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalSqlInjectionMatchSet_Id(p *WafregionalSqlInjectionMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple(p *SqlInjectionMatchTuple, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SqlInjectionMatchTuple {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_TextTransformation(v, ctyVal)
		EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sql_injection_match_tuple"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_TextTransformation(p *SqlInjectionMatchTuple, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_FieldToMatch_Data(v, ctyVal)
		EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_FieldToMatch_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_FieldToMatch_Data(p *FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafregionalSqlInjectionMatchSet_SqlInjectionMatchTuple_FieldToMatch_Type(p *FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}