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

func EncodeWafXssMatchSet(r WafXssMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafXssMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafXssMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafXssMatchSet_XssMatchTuples(r.Spec.ForProvider.XssMatchTuples, ctyVal)
	EncodeWafXssMatchSet_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafXssMatchSet_Name(p WafXssMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafXssMatchSet_Id(p WafXssMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafXssMatchSet_XssMatchTuples(p XssMatchTuples, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafXssMatchSet_XssMatchTuples_TextTransformation(p, ctyVal)
	EncodeWafXssMatchSet_XssMatchTuples_FieldToMatch(p.FieldToMatch, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["xss_match_tuples"] = cty.SetVal(valsForCollection)
}

func EncodeWafXssMatchSet_XssMatchTuples_TextTransformation(p XssMatchTuples, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafXssMatchSet_XssMatchTuples_FieldToMatch(p FieldToMatch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafXssMatchSet_XssMatchTuples_FieldToMatch_Data(p, ctyVal)
	EncodeWafXssMatchSet_XssMatchTuples_FieldToMatch_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafXssMatchSet_XssMatchTuples_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafXssMatchSet_XssMatchTuples_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafXssMatchSet_Arn(p WafXssMatchSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}