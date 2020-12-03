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

package v1alpha1func EncodeVpnGatewayRoutePropagation(r VpnGatewayRoutePropagation) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeVpnGatewayRoutePropagation_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpnGatewayRoutePropagation_RouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeVpnGatewayRoutePropagation_VpnGatewayId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeVpnGatewayRoutePropagation_Id(p *VpnGatewayRoutePropagationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpnGatewayRoutePropagation_RouteTableId(p *VpnGatewayRoutePropagationParameters, vals map[string]cty.Value) {
	vals["route_table_id"] = cty.StringVal(p.RouteTableId)
}

func EncodeVpnGatewayRoutePropagation_VpnGatewayId(p *VpnGatewayRoutePropagationParameters, vals map[string]cty.Value) {
	vals["vpn_gateway_id"] = cty.StringVal(p.VpnGatewayId)
}