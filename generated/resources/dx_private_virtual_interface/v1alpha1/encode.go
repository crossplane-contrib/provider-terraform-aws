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

func EncodeDxPrivateVirtualInterface(r DxPrivateVirtualInterface) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxPrivateVirtualInterface_AmazonAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_Vlan(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_AddressFamily(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_DxGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_Mtu(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_VpnGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_BgpAuthKey(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_CustomerAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxPrivateVirtualInterface_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_AmazonSideAsn(r.Status.AtProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_AwsDevice(r.Status.AtProvider, ctyVal)
	EncodeDxPrivateVirtualInterface_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxPrivateVirtualInterface_AmazonAddress(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["amazon_address"] = cty.StringVal(p.AmazonAddress)
}

func EncodeDxPrivateVirtualInterface_BgpAsn(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.NumberIntVal(p.BgpAsn)
}

func EncodeDxPrivateVirtualInterface_ConnectionId(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeDxPrivateVirtualInterface_Id(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxPrivateVirtualInterface_Name(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxPrivateVirtualInterface_Vlan(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["vlan"] = cty.NumberIntVal(p.Vlan)
}

func EncodeDxPrivateVirtualInterface_AddressFamily(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["address_family"] = cty.StringVal(p.AddressFamily)
}

func EncodeDxPrivateVirtualInterface_DxGatewayId(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["dx_gateway_id"] = cty.StringVal(p.DxGatewayId)
}

func EncodeDxPrivateVirtualInterface_Mtu(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["mtu"] = cty.NumberIntVal(p.Mtu)
}

func EncodeDxPrivateVirtualInterface_VpnGatewayId(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["vpn_gateway_id"] = cty.StringVal(p.VpnGatewayId)
}

func EncodeDxPrivateVirtualInterface_BgpAuthKey(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_auth_key"] = cty.StringVal(p.BgpAuthKey)
}

func EncodeDxPrivateVirtualInterface_CustomerAddress(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["customer_address"] = cty.StringVal(p.CustomerAddress)
}

func EncodeDxPrivateVirtualInterface_Tags(p DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDxPrivateVirtualInterface_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxPrivateVirtualInterface_Timeouts_Create(p, ctyVal)
	EncodeDxPrivateVirtualInterface_Timeouts_Delete(p, ctyVal)
	EncodeDxPrivateVirtualInterface_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxPrivateVirtualInterface_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxPrivateVirtualInterface_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxPrivateVirtualInterface_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDxPrivateVirtualInterface_Arn(p DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxPrivateVirtualInterface_AmazonSideAsn(p DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}

func EncodeDxPrivateVirtualInterface_AwsDevice(p DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}

func EncodeDxPrivateVirtualInterface_JumboFrameCapable(p DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}