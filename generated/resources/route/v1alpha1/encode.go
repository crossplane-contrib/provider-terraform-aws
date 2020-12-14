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
	r, ok := mr.(*Route)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Route.")
	}
	return EncodeRoute(*r), nil
}

func EncodeRoute(r Route) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute_GatewayId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute_RouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_DestinationCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeRoute_NatGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_NetworkInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_TransitGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_EgressOnlyGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_InstanceId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_LocalGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_VpcPeeringConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeRoute_DestinationIpv6CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeRoute_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeRoute_DestinationPrefixListId(r.Status.AtProvider, ctyVal)
	EncodeRoute_State(r.Status.AtProvider, ctyVal)
	EncodeRoute_InstanceOwnerId(r.Status.AtProvider, ctyVal)
	EncodeRoute_Origin(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeRoute_GatewayId(p RouteParameters, vals map[string]cty.Value) {
	vals["gateway_id"] = cty.StringVal(p.GatewayId)
}

func EncodeRoute_Id(p RouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute_RouteTableId(p RouteParameters, vals map[string]cty.Value) {
	vals["route_table_id"] = cty.StringVal(p.RouteTableId)
}

func EncodeRoute_DestinationCidrBlock(p RouteParameters, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeRoute_NatGatewayId(p RouteParameters, vals map[string]cty.Value) {
	vals["nat_gateway_id"] = cty.StringVal(p.NatGatewayId)
}

func EncodeRoute_NetworkInterfaceId(p RouteParameters, vals map[string]cty.Value) {
	vals["network_interface_id"] = cty.StringVal(p.NetworkInterfaceId)
}

func EncodeRoute_TransitGatewayId(p RouteParameters, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}

func EncodeRoute_EgressOnlyGatewayId(p RouteParameters, vals map[string]cty.Value) {
	vals["egress_only_gateway_id"] = cty.StringVal(p.EgressOnlyGatewayId)
}

func EncodeRoute_InstanceId(p RouteParameters, vals map[string]cty.Value) {
	vals["instance_id"] = cty.StringVal(p.InstanceId)
}

func EncodeRoute_LocalGatewayId(p RouteParameters, vals map[string]cty.Value) {
	vals["local_gateway_id"] = cty.StringVal(p.LocalGatewayId)
}

func EncodeRoute_VpcPeeringConnectionId(p RouteParameters, vals map[string]cty.Value) {
	vals["vpc_peering_connection_id"] = cty.StringVal(p.VpcPeeringConnectionId)
}

func EncodeRoute_DestinationIpv6CidrBlock(p RouteParameters, vals map[string]cty.Value) {
	vals["destination_ipv6_cidr_block"] = cty.StringVal(p.DestinationIpv6CidrBlock)
}

func EncodeRoute_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeRoute_Timeouts_Create(p, ctyVal)
	EncodeRoute_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeRoute_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeRoute_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeRoute_DestinationPrefixListId(p RouteObservation, vals map[string]cty.Value) {
	vals["destination_prefix_list_id"] = cty.StringVal(p.DestinationPrefixListId)
}

func EncodeRoute_State(p RouteObservation, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeRoute_InstanceOwnerId(p RouteObservation, vals map[string]cty.Value) {
	vals["instance_owner_id"] = cty.StringVal(p.InstanceOwnerId)
}

func EncodeRoute_Origin(p RouteObservation, vals map[string]cty.Value) {
	vals["origin"] = cty.StringVal(p.Origin)
}