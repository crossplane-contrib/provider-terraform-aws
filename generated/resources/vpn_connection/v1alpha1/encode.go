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

func EncodeVpnConnection(r VpnConnection) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpnConnection_CustomerGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_StaticRoutesOnly(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_TransitGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Tunnel2PresharedKey(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Type(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_VpnGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Tunnel1InsideCidr(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Tunnel1PresharedKey(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Tunnel2InsideCidr(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnection_Tunnel1VgwInsideAddress(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Tunnel2Address(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Tunnel2BgpAsn(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Tunnel2BgpHoldtime(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Routes(r.Status.AtProvider.Routes, ctyVal)
	EncodeVpnConnection_Tunnel1Address(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Tunnel1BgpAsn(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_VgwTelemetry(r.Status.AtProvider.VgwTelemetry, ctyVal)
	EncodeVpnConnection_Tunnel2VgwInsideAddress(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Arn(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_CustomerGatewayConfiguration(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_TransitGatewayAttachmentId(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Tunnel1BgpHoldtime(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Tunnel1CgwInsideAddress(r.Status.AtProvider, ctyVal)
	EncodeVpnConnection_Tunnel2CgwInsideAddress(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeVpnConnection_CustomerGatewayId(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["customer_gateway_id"] = cty.StringVal(p.CustomerGatewayId)
}

func EncodeVpnConnection_Id(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpnConnection_StaticRoutesOnly(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["static_routes_only"] = cty.BoolVal(p.StaticRoutesOnly)
}

func EncodeVpnConnection_TransitGatewayId(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}

func EncodeVpnConnection_Tunnel2PresharedKey(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["tunnel2_preshared_key"] = cty.StringVal(p.Tunnel2PresharedKey)
}

func EncodeVpnConnection_Type(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeVpnConnection_VpnGatewayId(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["vpn_gateway_id"] = cty.StringVal(p.VpnGatewayId)
}

func EncodeVpnConnection_Tunnel1InsideCidr(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["tunnel1_inside_cidr"] = cty.StringVal(p.Tunnel1InsideCidr)
}

func EncodeVpnConnection_Tunnel1PresharedKey(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["tunnel1_preshared_key"] = cty.StringVal(p.Tunnel1PresharedKey)
}

func EncodeVpnConnection_Tunnel2InsideCidr(p VpnConnectionParameters, vals map[string]cty.Value) {
	vals["tunnel2_inside_cidr"] = cty.StringVal(p.Tunnel2InsideCidr)
}

func EncodeVpnConnection_Tags(p VpnConnectionParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpnConnection_Tunnel1VgwInsideAddress(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel1_vgw_inside_address"] = cty.StringVal(p.Tunnel1VgwInsideAddress)
}

func EncodeVpnConnection_Tunnel2Address(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel2_address"] = cty.StringVal(p.Tunnel2Address)
}

func EncodeVpnConnection_Tunnel2BgpAsn(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel2_bgp_asn"] = cty.StringVal(p.Tunnel2BgpAsn)
}

func EncodeVpnConnection_Tunnel2BgpHoldtime(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel2_bgp_holdtime"] = cty.NumberIntVal(p.Tunnel2BgpHoldtime)
}

func EncodeVpnConnection_Routes(p []Routes, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeVpnConnection_Routes_DestinationCidrBlock(v, ctyVal)
		EncodeVpnConnection_Routes_Source(v, ctyVal)
		EncodeVpnConnection_Routes_State(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["routes"] = cty.SetVal(valsForCollection)
}

func EncodeVpnConnection_Routes_DestinationCidrBlock(p Routes, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeVpnConnection_Routes_Source(p Routes, vals map[string]cty.Value) {
	vals["source"] = cty.StringVal(p.Source)
}

func EncodeVpnConnection_Routes_State(p Routes, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeVpnConnection_Tunnel1Address(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel1_address"] = cty.StringVal(p.Tunnel1Address)
}

func EncodeVpnConnection_Tunnel1BgpAsn(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel1_bgp_asn"] = cty.StringVal(p.Tunnel1BgpAsn)
}

func EncodeVpnConnection_VgwTelemetry(p []VgwTelemetry, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeVpnConnection_VgwTelemetry_LastStatusChange(v, ctyVal)
		EncodeVpnConnection_VgwTelemetry_OutsideIpAddress(v, ctyVal)
		EncodeVpnConnection_VgwTelemetry_Status(v, ctyVal)
		EncodeVpnConnection_VgwTelemetry_StatusMessage(v, ctyVal)
		EncodeVpnConnection_VgwTelemetry_AcceptedRouteCount(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["vgw_telemetry"] = cty.SetVal(valsForCollection)
}

func EncodeVpnConnection_VgwTelemetry_LastStatusChange(p VgwTelemetry, vals map[string]cty.Value) {
	vals["last_status_change"] = cty.StringVal(p.LastStatusChange)
}

func EncodeVpnConnection_VgwTelemetry_OutsideIpAddress(p VgwTelemetry, vals map[string]cty.Value) {
	vals["outside_ip_address"] = cty.StringVal(p.OutsideIpAddress)
}

func EncodeVpnConnection_VgwTelemetry_Status(p VgwTelemetry, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeVpnConnection_VgwTelemetry_StatusMessage(p VgwTelemetry, vals map[string]cty.Value) {
	vals["status_message"] = cty.StringVal(p.StatusMessage)
}

func EncodeVpnConnection_VgwTelemetry_AcceptedRouteCount(p VgwTelemetry, vals map[string]cty.Value) {
	vals["accepted_route_count"] = cty.NumberIntVal(p.AcceptedRouteCount)
}

func EncodeVpnConnection_Tunnel2VgwInsideAddress(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel2_vgw_inside_address"] = cty.StringVal(p.Tunnel2VgwInsideAddress)
}

func EncodeVpnConnection_Arn(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeVpnConnection_CustomerGatewayConfiguration(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["customer_gateway_configuration"] = cty.StringVal(p.CustomerGatewayConfiguration)
}

func EncodeVpnConnection_TransitGatewayAttachmentId(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["transit_gateway_attachment_id"] = cty.StringVal(p.TransitGatewayAttachmentId)
}

func EncodeVpnConnection_Tunnel1BgpHoldtime(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel1_bgp_holdtime"] = cty.NumberIntVal(p.Tunnel1BgpHoldtime)
}

func EncodeVpnConnection_Tunnel1CgwInsideAddress(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel1_cgw_inside_address"] = cty.StringVal(p.Tunnel1CgwInsideAddress)
}

func EncodeVpnConnection_Tunnel2CgwInsideAddress(p VpnConnectionObservation, vals map[string]cty.Value) {
	vals["tunnel2_cgw_inside_address"] = cty.StringVal(p.Tunnel2CgwInsideAddress)
}