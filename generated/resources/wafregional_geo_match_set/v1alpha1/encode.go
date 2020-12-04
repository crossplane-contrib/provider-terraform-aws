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

func EncodeWafregionalGeoMatchSet(r WafregionalGeoMatchSet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalGeoMatchSet_Id(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalGeoMatchSet_Name(r.Spec.ForProvider, ctyVal)
	EncodeWafregionalGeoMatchSet_GeoMatchConstraint(r.Spec.ForProvider.GeoMatchConstraint, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeWafregionalGeoMatchSet_Id(p WafregionalGeoMatchSetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeWafregionalGeoMatchSet_Name(p WafregionalGeoMatchSetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeWafregionalGeoMatchSet_GeoMatchConstraint(p GeoMatchConstraint, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeWafregionalGeoMatchSet_GeoMatchConstraint_Type(p, ctyVal)
	EncodeWafregionalGeoMatchSet_GeoMatchConstraint_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["geo_match_constraint"] = cty.SetVal(valsForCollection)
}

func EncodeWafregionalGeoMatchSet_GeoMatchConstraint_Type(p GeoMatchConstraint, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeWafregionalGeoMatchSet_GeoMatchConstraint_Value(p GeoMatchConstraint, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}