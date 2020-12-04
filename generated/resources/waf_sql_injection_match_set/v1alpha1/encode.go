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

func EncodeWafSqlInjectionMatchSet(r WafSqlInjectionMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafSqlInjectionMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafSqlInjectionMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples(r.Spec.ForProvider.SqlInjectionMatchTuples, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeWafSqlInjectionMatchSet_Id(p WafSqlInjectionMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafSqlInjectionMatchSet_Name(p WafSqlInjectionMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples(p SqlInjectionMatchTuples, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_TextTransformation(p, ctyVal)
	EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_FieldToMatch(p.FieldToMatch, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["sql_injection_match_tuples"] = cty.SetVal(valsForCollection)
}

func EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_TextTransformation(p SqlInjectionMatchTuples, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_FieldToMatch(p FieldToMatch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_FieldToMatch_Data(p, ctyVal)
	EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_FieldToMatch_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafSqlInjectionMatchSet_SqlInjectionMatchTuples_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}