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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*Route)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeRoute(r, ctyValue)
}

func DecodeRoute(prev *Route, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeRoute_NetworkInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeRoute_VpcPeeringConnectionId(&new.Spec.ForProvider, valMap)
	DecodeRoute_InstanceId(&new.Spec.ForProvider, valMap)
	DecodeRoute_NatGatewayId(&new.Spec.ForProvider, valMap)
	DecodeRoute_DestinationCidrBlock(&new.Spec.ForProvider, valMap)
	DecodeRoute_EgressOnlyGatewayId(&new.Spec.ForProvider, valMap)
	DecodeRoute_GatewayId(&new.Spec.ForProvider, valMap)
	DecodeRoute_RouteTableId(&new.Spec.ForProvider, valMap)
	DecodeRoute_TransitGatewayId(&new.Spec.ForProvider, valMap)
	DecodeRoute_DestinationIpv6CidrBlock(&new.Spec.ForProvider, valMap)
	DecodeRoute_LocalGatewayId(&new.Spec.ForProvider, valMap)
	DecodeRoute_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeRoute_State(&new.Status.AtProvider, valMap)
	DecodeRoute_Origin(&new.Status.AtProvider, valMap)
	DecodeRoute_DestinationPrefixListId(&new.Status.AtProvider, valMap)
	DecodeRoute_InstanceOwnerId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeRoute_NetworkInterfaceId(p *RouteParameters, vals map[string]cty.Value) {
	p.NetworkInterfaceId = ctwhy.ValueAsString(vals["network_interface_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_VpcPeeringConnectionId(p *RouteParameters, vals map[string]cty.Value) {
	p.VpcPeeringConnectionId = ctwhy.ValueAsString(vals["vpc_peering_connection_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_InstanceId(p *RouteParameters, vals map[string]cty.Value) {
	p.InstanceId = ctwhy.ValueAsString(vals["instance_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_NatGatewayId(p *RouteParameters, vals map[string]cty.Value) {
	p.NatGatewayId = ctwhy.ValueAsString(vals["nat_gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_DestinationCidrBlock(p *RouteParameters, vals map[string]cty.Value) {
	p.DestinationCidrBlock = ctwhy.ValueAsString(vals["destination_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_EgressOnlyGatewayId(p *RouteParameters, vals map[string]cty.Value) {
	p.EgressOnlyGatewayId = ctwhy.ValueAsString(vals["egress_only_gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_GatewayId(p *RouteParameters, vals map[string]cty.Value) {
	p.GatewayId = ctwhy.ValueAsString(vals["gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_RouteTableId(p *RouteParameters, vals map[string]cty.Value) {
	p.RouteTableId = ctwhy.ValueAsString(vals["route_table_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_TransitGatewayId(p *RouteParameters, vals map[string]cty.Value) {
	p.TransitGatewayId = ctwhy.ValueAsString(vals["transit_gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_DestinationIpv6CidrBlock(p *RouteParameters, vals map[string]cty.Value) {
	p.DestinationIpv6CidrBlock = ctwhy.ValueAsString(vals["destination_ipv6_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_LocalGatewayId(p *RouteParameters, vals map[string]cty.Value) {
	p.LocalGatewayId = ctwhy.ValueAsString(vals["local_gateway_id"])
}

//containerTypeDecodeTemplate
func DecodeRoute_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeRoute_Timeouts_Create(p, valMap)
	DecodeRoute_Timeouts_Delete(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeRoute_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_State(p *RouteObservation, vals map[string]cty.Value) {
	p.State = ctwhy.ValueAsString(vals["state"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_Origin(p *RouteObservation, vals map[string]cty.Value) {
	p.Origin = ctwhy.ValueAsString(vals["origin"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_DestinationPrefixListId(p *RouteObservation, vals map[string]cty.Value) {
	p.DestinationPrefixListId = ctwhy.ValueAsString(vals["destination_prefix_list_id"])
}

//primitiveTypeDecodeTemplate
func DecodeRoute_InstanceOwnerId(p *RouteObservation, vals map[string]cty.Value) {
	p.InstanceOwnerId = ctwhy.ValueAsString(vals["instance_owner_id"])
}