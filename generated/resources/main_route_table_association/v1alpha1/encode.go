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

func EncodeMainRouteTableAssociation(r MainRouteTableAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeMainRouteTableAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeMainRouteTableAssociation_RouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeMainRouteTableAssociation_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeMainRouteTableAssociation_OriginalRouteTableId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeMainRouteTableAssociation_Id(p MainRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeMainRouteTableAssociation_RouteTableId(p MainRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["route_table_id"] = cty.StringVal(p.RouteTableId)
}

func EncodeMainRouteTableAssociation_VpcId(p MainRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeMainRouteTableAssociation_OriginalRouteTableId(p MainRouteTableAssociationObservation, vals map[string]cty.Value) {
	vals["original_route_table_id"] = cty.StringVal(p.OriginalRouteTableId)
}