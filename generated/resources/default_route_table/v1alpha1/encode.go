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

func EncodeDefaultRouteTable(r DefaultRouteTable) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDefaultRouteTable_PropagatingVgws(r.Spec.ForProvider, ctyVal)
	EncodeDefaultRouteTable_Route(r.Spec.ForProvider.Route, ctyVal)
	EncodeDefaultRouteTable_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDefaultRouteTable_DefaultRouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeDefaultRouteTable_Id(r.Spec.ForProvider, ctyVal)
	EncodeDefaultRouteTable_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeDefaultRouteTable_VpcId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDefaultRouteTable_PropagatingVgws(p DefaultRouteTableParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PropagatingVgws {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["propagating_vgws"] = cty.SetVal(colVals)
}

func EncodeDefaultRouteTable_Route(p []Route, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeDefaultRouteTable_Route_EgressOnlyGatewayId(v, ctyVal)
		EncodeDefaultRouteTable_Route_InstanceId(v, ctyVal)
		EncodeDefaultRouteTable_Route_NetworkInterfaceId(v, ctyVal)
		EncodeDefaultRouteTable_Route_GatewayId(v, ctyVal)
		EncodeDefaultRouteTable_Route_Ipv6CidrBlock(v, ctyVal)
		EncodeDefaultRouteTable_Route_TransitGatewayId(v, ctyVal)
		EncodeDefaultRouteTable_Route_VpcPeeringConnectionId(v, ctyVal)
		EncodeDefaultRouteTable_Route_CidrBlock(v, ctyVal)
		EncodeDefaultRouteTable_Route_NatGatewayId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["route"] = cty.SetVal(valsForCollection)
}

func EncodeDefaultRouteTable_Route_EgressOnlyGatewayId(p Route, vals map[string]cty.Value) {
	vals["egress_only_gateway_id"] = cty.StringVal(p.EgressOnlyGatewayId)
}

func EncodeDefaultRouteTable_Route_InstanceId(p Route, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}

func EncodeDefaultRouteTable_Route_NetworkInterfaceId(p Route, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeDefaultRouteTable_Route_GatewayId(p Route, vals map[string]cty.Value) {
	vals["gateway_id"] = cty.StringVal(p.GatewayId)
}

func EncodeDefaultRouteTable_Route_Ipv6CidrBlock(p Route, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeDefaultRouteTable_Route_TransitGatewayId(p Route, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}

func EncodeDefaultRouteTable_Route_VpcPeeringConnectionId(p Route, vals map[string]cty.Value) {
	vals["vpc_peering_connection_id"] = cty.StringVal(p.VpcPeeringConnectionId)
}

func EncodeDefaultRouteTable_Route_CidrBlock(p Route, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeDefaultRouteTable_Route_NatGatewayId(p Route, vals map[string]cty.Value) {
	vals["nat_gateway_id"] = cty.StringVal(p.NatGatewayId)
}

func EncodeDefaultRouteTable_Tags(p DefaultRouteTableParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDefaultRouteTable_DefaultRouteTableId(p DefaultRouteTableParameters, vals map[string]cty.Value) {
	vals["default_route_table_id"] = cty.StringVal(p.DefaultRouteTableId)
}

func EncodeDefaultRouteTable_Id(p DefaultRouteTableParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDefaultRouteTable_OwnerId(p DefaultRouteTableObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeDefaultRouteTable_VpcId(p DefaultRouteTableObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}