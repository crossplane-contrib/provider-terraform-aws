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

package v1alpha1func EncodeDxHostedPrivateVirtualInterfaceAccepter(r DxHostedPrivateVirtualInterfaceAccepter) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_DxGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_VirtualInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_VpnGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_DxGatewayId(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["dx_gateway_id"] = cty.StringVal(p.DxGatewayId)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_Id(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_Tags(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_VirtualInterfaceId(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["virtual_interface_id"] = cty.StringVal(p.VirtualInterfaceId)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_VpnGatewayId(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["vpn_gateway_id"] = cty.StringVal(p.VpnGatewayId)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Create(p, ctyVal)
	EncodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxHostedPrivateVirtualInterfaceAccepter_Arn(p *DxHostedPrivateVirtualInterfaceAccepterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}