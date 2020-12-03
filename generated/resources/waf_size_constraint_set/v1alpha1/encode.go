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

package v1alpha1func EncodeWafSizeConstraintSet(r WafSizeConstraintSet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafSizeConstraintSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafSizeConstraintSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafSizeConstraintSet_SizeConstraints(r.Spec.ForProvider.SizeConstraints, ctyVal)
	EncodeWafSizeConstraintSet_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeWafSizeConstraintSet_Id(p *WafSizeConstraintSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafSizeConstraintSet_Name(p *WafSizeConstraintSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafSizeConstraintSet_SizeConstraints(p *SizeConstraints, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraints {
		ctyVal = make(map[string]cty.Value)
		EncodeWafSizeConstraintSet_SizeConstraints_TextTransformation(v, ctyVal)
		EncodeWafSizeConstraintSet_SizeConstraints_ComparisonOperator(v, ctyVal)
		EncodeWafSizeConstraintSet_SizeConstraints_Size(v, ctyVal)
		EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraints"] = cty.SetVal(valsForCollection)
}

func EncodeWafSizeConstraintSet_SizeConstraints_TextTransformation(p *SizeConstraints, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafSizeConstraintSet_SizeConstraints_ComparisonOperator(p *SizeConstraints, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafSizeConstraintSet_SizeConstraints_Size(p *SizeConstraints, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Data(v, ctyVal)
		EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Data(p *FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Type(p *FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafSizeConstraintSet_Arn(p *WafSizeConstraintSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}