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
	r, ok := mr.(*RouteTable)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a RouteTable.")
	}
	return EncodeRouteTable(*r), nil
}

func EncodeRouteTable(r RouteTable) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRouteTable_PropagatingVgws(r.Spec.ForProvider, ctyVal)
	EncodeRouteTable_Route(r.Spec.ForProvider.Route, ctyVal)
	EncodeRouteTable_Tags(r.Spec.ForProvider, ctyVal)
	EncodeRouteTable_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeRouteTable_Id(r.Spec.ForProvider, ctyVal)
	EncodeRouteTable_OwnerId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeRouteTable_PropagatingVgws(p RouteTableParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.PropagatingVgws {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["propagating_vgws"] = cty.SetVal(colVals)
}

func EncodeRouteTable_Route(p []Route, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeRouteTable_Route_VpcPeeringConnectionId(v, ctyVal)
		EncodeRouteTable_Route_CidrBlock(v, ctyVal)
		EncodeRouteTable_Route_GatewayId(v, ctyVal)
		EncodeRouteTable_Route_Ipv6CidrBlock(v, ctyVal)
		EncodeRouteTable_Route_LocalGatewayId(v, ctyVal)
		EncodeRouteTable_Route_NetworkInterfaceId(v, ctyVal)
		EncodeRouteTable_Route_NatGatewayId(v, ctyVal)
		EncodeRouteTable_Route_TransitGatewayId(v, ctyVal)
		EncodeRouteTable_Route_EgressOnlyGatewayId(v, ctyVal)
		EncodeRouteTable_Route_InstanceId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["route"] = cty.SetVal(valsForCollection)
}

func EncodeRouteTable_Route_VpcPeeringConnectionId(p Route, vals map[string]cty.Value) {
	vals["vpc_peering_connection_id"] = cty.StringVal(p.VpcPeeringConnectionId)
}

func EncodeRouteTable_Route_CidrBlock(p Route, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeRouteTable_Route_GatewayId(p Route, vals map[string]cty.Value) {
	vals["gateway_id"] = cty.StringVal(p.GatewayId)
}

func EncodeRouteTable_Route_Ipv6CidrBlock(p Route, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeRouteTable_Route_LocalGatewayId(p Route, vals map[string]cty.Value) {
	vals["local_gateway_id"] = cty.StringVal(p.LocalGatewayId)
}

func EncodeRouteTable_Route_NetworkInterfaceId(p Route, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeRouteTable_Route_NatGatewayId(p Route, vals map[string]cty.Value) {
	vals["nat_gateway_id"] = cty.StringVal(p.NatGatewayId)
}

func EncodeRouteTable_Route_TransitGatewayId(p Route, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}

func EncodeRouteTable_Route_EgressOnlyGatewayId(p Route, vals map[string]cty.Value) {
	vals["egress_only_gateway_id"] = cty.StringVal(p.EgressOnlyGatewayId)
}

func EncodeRouteTable_Route_InstanceId(p Route, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}

func EncodeRouteTable_Tags(p RouteTableParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeRouteTable_VpcId(p RouteTableParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeRouteTable_Id(p RouteTableParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRouteTable_OwnerId(p RouteTableObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}