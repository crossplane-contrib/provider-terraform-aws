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

func EncodeRouteTableAssociation(r RouteTableAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRouteTableAssociation_GatewayId(r.Spec.ForProvider, ctyVal)
	EncodeRouteTableAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeRouteTableAssociation_RouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeRouteTableAssociation_SubnetId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeRouteTableAssociation_GatewayId(p RouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["gateway_id"] = cty.StringVal(p.GatewayId)
}

func EncodeRouteTableAssociation_Id(p RouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRouteTableAssociation_RouteTableId(p RouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["route_table_id"] = cty.StringVal(p.RouteTableId)
}

func EncodeRouteTableAssociation_SubnetId(p RouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}