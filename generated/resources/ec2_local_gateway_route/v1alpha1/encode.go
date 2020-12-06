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

func EncodeEc2LocalGatewayRoute(r Ec2LocalGatewayRoute) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2LocalGatewayRoute_DestinationCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeEc2LocalGatewayRoute_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2LocalGatewayRoute_LocalGatewayRouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeEc2LocalGatewayRoute_LocalGatewayVirtualInterfaceGroupId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeEc2LocalGatewayRoute_DestinationCidrBlock(p Ec2LocalGatewayRouteParameters, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeEc2LocalGatewayRoute_Id(p Ec2LocalGatewayRouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2LocalGatewayRoute_LocalGatewayRouteTableId(p Ec2LocalGatewayRouteParameters, vals map[string]cty.Value) {
	vals["local_gateway_route_table_id"] = cty.StringVal(p.LocalGatewayRouteTableId)
}

func EncodeEc2LocalGatewayRoute_LocalGatewayVirtualInterfaceGroupId(p Ec2LocalGatewayRouteParameters, vals map[string]cty.Value) {
	vals["local_gateway_virtual_interface_group_id"] = cty.StringVal(p.LocalGatewayVirtualInterfaceGroupId)
}