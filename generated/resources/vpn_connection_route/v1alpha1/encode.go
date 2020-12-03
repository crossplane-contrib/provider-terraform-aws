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

package v1alpha1func EncodeVpnConnectionRoute(r VpnConnectionRoute) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeVpnConnectionRoute_DestinationCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnectionRoute_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpnConnectionRoute_VpnConnectionId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeVpnConnectionRoute_DestinationCidrBlock(p *VpnConnectionRouteParameters, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeVpnConnectionRoute_Id(p *VpnConnectionRouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpnConnectionRoute_VpnConnectionId(p *VpnConnectionRouteParameters, vals map[string]cty.Value) {
	vals["vpn_connection_id"] = cty.StringVal(p.VpnConnectionId)
}