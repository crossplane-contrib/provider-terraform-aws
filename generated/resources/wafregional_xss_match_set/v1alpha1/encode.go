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

func EncodeWafregionalXssMatchSet(r WafregionalXssMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalXssMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalXssMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalXssMatchSet_XssMatchTuple(r.Spec.ForProvider.XssMatchTuple, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeWafregionalXssMatchSet_Id(p WafregionalXssMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalXssMatchSet_Name(p WafregionalXssMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalXssMatchSet_XssMatchTuple(p XssMatchTuple, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalXssMatchSet_XssMatchTuple_TextTransformation(p, ctyVal)
	EncodeWafregionalXssMatchSet_XssMatchTuple_FieldToMatch(p.FieldToMatch, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["xss_match_tuple"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalXssMatchSet_XssMatchTuple_TextTransformation(p XssMatchTuple, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafregionalXssMatchSet_XssMatchTuple_FieldToMatch(p FieldToMatch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalXssMatchSet_XssMatchTuple_FieldToMatch_Data(p, ctyVal)
	EncodeWafregionalXssMatchSet_XssMatchTuple_FieldToMatch_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalXssMatchSet_XssMatchTuple_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafregionalXssMatchSet_XssMatchTuple_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}