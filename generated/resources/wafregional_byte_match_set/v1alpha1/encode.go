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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*WafregionalByteMatchSet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a WafregionalByteMatchSet.")
	}
	return EncodeWafregionalByteMatchSet(*r), nil
}

func EncodeWafregionalByteMatchSet(r WafregionalByteMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalByteMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalByteMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalByteMatchSet_ByteMatchTuples(r.Spec.ForProvider.ByteMatchTuples, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeWafregionalByteMatchSet_Id(p WafregionalByteMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalByteMatchSet_Name(p WafregionalByteMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalByteMatchSet_ByteMatchTuples(p ByteMatchTuples, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalByteMatchSet_ByteMatchTuples_PositionalConstraint(p, ctyVal)
	EncodeWafregionalByteMatchSet_ByteMatchTuples_TargetString(p, ctyVal)
	EncodeWafregionalByteMatchSet_ByteMatchTuples_TextTransformation(p, ctyVal)
	EncodeWafregionalByteMatchSet_ByteMatchTuples_FieldToMatch(p.FieldToMatch, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["byte_match_tuples"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalByteMatchSet_ByteMatchTuples_PositionalConstraint(p ByteMatchTuples, vals map[string]cty.Value) {
	vals["positional_constraint"] = cty.StringVal(p.PositionalConstraint)
}

func EncodeWafregionalByteMatchSet_ByteMatchTuples_TargetString(p ByteMatchTuples, vals map[string]cty.Value) {
	vals["target_string"] = cty.StringVal(p.TargetString)
}

func EncodeWafregionalByteMatchSet_ByteMatchTuples_TextTransformation(p ByteMatchTuples, vals map[string]cty.Value) {
	vals["text_transformation"] = cty.StringVal(p.TextTransformation)
}

func EncodeWafregionalByteMatchSet_ByteMatchTuples_FieldToMatch(p FieldToMatch, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalByteMatchSet_ByteMatchTuples_FieldToMatch_Data(p, ctyVal)
	EncodeWafregionalByteMatchSet_ByteMatchTuples_FieldToMatch_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["field_to_match"] = cty.ListVal(valsForCollection)
}

func EncodeWafregionalByteMatchSet_ByteMatchTuples_FieldToMatch_Data(p FieldToMatch, vals map[string]cty.Value) {
	vals["data"] = cty.StringVal(p.Data)
}

func EncodeWafregionalByteMatchSet_ByteMatchTuples_FieldToMatch_Type(p FieldToMatch, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}