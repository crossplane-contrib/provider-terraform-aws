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

package v1alpha1func EncodeWafByteMatchSet(r WafByteMatchSet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafByteMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafByteMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafByteMatchSet_ByteMatchTuples(r.Spec.ForProvider.ByteMatchTuples, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeWafByteMatchSet_Name(p *WafByteMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafByteMatchSet_Id(p *WafByteMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafByteMatchSet_ByteMatchTuples(p *ByteMatchTuples, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ByteMatchTuples {
		ctyVal = make(map[string]cty.Value)
		EncodeWafByteMatchSet_ByteMatchTuples_TargetString(v, ctyVal)
		EncodeWafByteMatchSet_ByteMatchTuples_TextTransformation(v, ctyVal)
		EncodeWafByteMatchSet_ByteMatchTuples_PositionalConstraint(v, ctyVal)
		EncodeWafByteMatchSet_ByteMatchTuples_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["byte_match_tuples"] = cty.SetVal(valsForCollection)
}

func EncodeWafByteMatchSet_ByteMatchTuples_TargetString(p *ByteMatchTuples, vals map[string]cty.Value) {
	vals["target_string"] = cty.StringVal(p.TargetString)
}

func EncodeWafByteMatchSet_ByteMatchTuples_TextTransformation(p *ByteMatchTuples, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafByteMatchSet_ByteMatchTuples_PositionalConstraint(p *ByteMatchTuples, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafByteMatchSet_ByteMatchTuples_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafByteMatchSet_ByteMatchTuples_FieldToMatch_Data(v, ctyVal)
		EncodeWafByteMatchSet_ByteMatchTuples_FieldToMatch_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafByteMatchSet_ByteMatchTuples_FieldToMatch_Data(p *FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafByteMatchSet_ByteMatchTuples_FieldToMatch_Type(p *FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}