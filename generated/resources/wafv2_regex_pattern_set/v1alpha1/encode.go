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

package v1alpha1func EncodeWafv2RegexPatternSet(r Wafv2RegexPatternSet) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeWafv2RegexPatternSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RegexPatternSet_Scope(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RegexPatternSet_Tags(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RegexPatternSet_Description(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RegexPatternSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafv2RegexPatternSet_RegularExpression(r.Spec.ForProvider.RegularExpression, ctyVal)
	EncodeWafv2RegexPatternSet_LockToken(r.Status.AtProvider, ctyVal)
	EncodeWafv2RegexPatternSet_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeWafv2RegexPatternSet_Name(p *Wafv2RegexPatternSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafv2RegexPatternSet_Scope(p *Wafv2RegexPatternSetParameters, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeWafv2RegexPatternSet_Tags(p *Wafv2RegexPatternSetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeWafv2RegexPatternSet_Description(p *Wafv2RegexPatternSetParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeWafv2RegexPatternSet_Id(p *Wafv2RegexPatternSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafv2RegexPatternSet_RegularExpression(p *RegularExpression, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RegularExpression {
		ctyVal = make(map[string]cty.Value)
		EncodeWafv2RegexPatternSet_RegularExpression_RegexString(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["regular_expression"] = cty.SetVal(valsForCollection)
}

func EncodeWafv2RegexPatternSet_RegularExpression_RegexString(p *RegularExpression, vals map[string]cty.Value) {
	vals["regex_string"] = cty.StringVal(p.RegexString)
}

func EncodeWafv2RegexPatternSet_LockToken(p *Wafv2RegexPatternSetObservation, vals map[string]cty.Value) {
	vals["lock_token"] = cty.StringVal(p.LockToken)
}

func EncodeWafv2RegexPatternSet_Arn(p *Wafv2RegexPatternSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}