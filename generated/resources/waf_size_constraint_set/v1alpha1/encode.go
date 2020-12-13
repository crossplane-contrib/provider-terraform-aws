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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*WafSizeConstraintSet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafSizeConstraintSet.")
	}
	return EncodeWafSizeConstraintSet(*r), nil
}

func EncodeWafSizeConstraintSet(r WafSizeConstraintSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafSizeConstraintSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafSizeConstraintSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafSizeConstraintSet_SizeConstraints(r.Spec.ForProvider.SizeConstraints, ctyVal)
	EncodeWafSizeConstraintSet_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeWafSizeConstraintSet_Id(p WafSizeConstraintSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafSizeConstraintSet_Name(p WafSizeConstraintSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafSizeConstraintSet_SizeConstraints(p SizeConstraints, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafSizeConstraintSet_SizeConstraints_ComparisonOperator(p, ctyVal)
	EncodeWafSizeConstraintSet_SizeConstraints_Size(p, ctyVal)
	EncodeWafSizeConstraintSet_SizeConstraints_TextTransformation(p, ctyVal)
	EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch(p.FieldToMatch, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["size_constraints"] = cty.SetVal(valsForCollection)
}

func EncodeWafSizeConstraintSet_SizeConstraints_ComparisonOperator(p SizeConstraints, vals map[string]cty.Value) {
	vals["comparison_operator"] = cty.StringVal(p.ComparisonOperator)
}

func EncodeWafSizeConstraintSet_SizeConstraints_Size(p SizeConstraints, vals map[string]cty.Value) {
	vals["size"] = cty.NumberIntVal(p.Size)
}

func EncodeWafSizeConstraintSet_SizeConstraints_TextTransformation(p SizeConstraints, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch(p FieldToMatch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Data(p, ctyVal)
	EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafSizeConstraintSet_SizeConstraints_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafSizeConstraintSet_Arn(p WafSizeConstraintSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}