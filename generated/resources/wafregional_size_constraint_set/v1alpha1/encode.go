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

package v1alpha1func EncodeWafregionalSizeConstraintSet(r WafregionalSizeConstraintSet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafregionalSizeConstraintSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalSizeConstraintSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalSizeConstraintSet_SizeConstraints(r.Spec.ForProvider.SizeConstraints, ctyVal)
	EncodeWafregionalSizeConstraintSet_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeWafregionalSizeConstraintSet_Id(p *WafregionalSizeConstraintSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalSizeConstraintSet_Name(p *WafregionalSizeConstraintSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalSizeConstraintSet_SizeConstraints(p *SizeConstraints, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SizeConstraints {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalSizeConstraintSet_SizeConstraints_Size(v, ctyVal)
		EncodeWafregionalSizeConstraintSet_SizeConstraints_TextTransformation(v, ctyVal)
		EncodeWafregionalSizeConstraintSet_SizeConstraints_ComparisonOperator(v, ctyVal)
		EncodeWafregionalSizeConstraintSet_SizeConstraints_FieldToMatch(v.FieldToMatch, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["size_constraints"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalSizeConstraintSet_SizeConstraints_Size(p *SizeConstraints, vals map[string]cty.Value) {
	vals["size"] = cty.IntVal(p.Size)
}

func EncodeWafregionalSizeConstraintSet_SizeConstraints_TextTransformation(p *SizeConstraints, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafregionalSizeConstraintSet_SizeConstraints_ComparisonOperator(p *SizeConstraints, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafregionalSizeConstraintSet_SizeConstraints_FieldToMatch(p *FieldToMatch, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.FieldToMatch {
		ctyVal = make(map[string]cty.Value)
		EncodeWafregionalSizeConstraintSet_SizeConstraints_FieldToMatch_Data(v, ctyVal)
		EncodeWafregionalSizeConstraintSet_SizeConstraints_FieldToMatch_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalSizeConstraintSet_SizeConstraints_FieldToMatch_Data(p *FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafregionalSizeConstraintSet_SizeConstraints_FieldToMatch_Type(p *FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalSizeConstraintSet_Arn(p *WafregionalSizeConstraintSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}