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

func EncodeRoute53VpcAssociationAuthorization(r Route53VpcAssociationAuthorization) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute53VpcAssociationAuthorization_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53VpcAssociationAuthorization_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeRoute53VpcAssociationAuthorization_VpcRegion(r.Spec.ForProvider, ctyVal)
	EncodeRoute53VpcAssociationAuthorization_ZoneId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeRoute53VpcAssociationAuthorization_Id(p Route53VpcAssociationAuthorizationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53VpcAssociationAuthorization_VpcId(p Route53VpcAssociationAuthorizationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeRoute53VpcAssociationAuthorization_VpcRegion(p Route53VpcAssociationAuthorizationParameters, vals map[string]cty.Value) {
	vals["vpc_region"] = cty.StringVal(p.VpcRegion)
}

func EncodeRoute53VpcAssociationAuthorization_ZoneId(p Route53VpcAssociationAuthorizationParameters, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}