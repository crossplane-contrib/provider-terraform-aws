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

func EncodeWafGeoMatchSet(r WafGeoMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafGeoMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafGeoMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafGeoMatchSet_GeoMatchConstraint(r.Spec.ForProvider.GeoMatchConstraint, ctyVal)
	EncodeWafGeoMatchSet_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeWafGeoMatchSet_Id(p WafGeoMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafGeoMatchSet_Name(p WafGeoMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafGeoMatchSet_GeoMatchConstraint(p GeoMatchConstraint, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafGeoMatchSet_GeoMatchConstraint_Type(p, ctyVal)
	EncodeWafGeoMatchSet_GeoMatchConstraint_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["geo_match_constraint"] = cty.SetVal(valsForCollection)
}

func EncodeWafGeoMatchSet_GeoMatchConstraint_Type(p GeoMatchConstraint, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafGeoMatchSet_GeoMatchConstraint_Value(p GeoMatchConstraint, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeWafGeoMatchSet_Arn(p WafGeoMatchSetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}