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

func EncodeDxBgpPeer(r DxBgpPeer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxBgpPeer_CustomerAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxBgpPeer_VirtualInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeDxBgpPeer_AddressFamily(r.Spec.ForProvider, ctyVal)
	EncodeDxBgpPeer_AmazonAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxBgpPeer_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxBgpPeer_BgpAuthKey(r.Spec.ForProvider, ctyVal)
	EncodeDxBgpPeer_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxBgpPeer_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxBgpPeer_AwsDevice(r.Status.AtProvider, ctyVal)
	EncodeDxBgpPeer_BgpPeerId(r.Status.AtProvider, ctyVal)
	EncodeDxBgpPeer_BgpStatus(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxBgpPeer_CustomerAddress(p DxBgpPeerParameters, vals map[string]cty.Value) {
	vals["customer_address"] = cty.StringVal(p.CustomerAddress)
}

func EncodeDxBgpPeer_VirtualInterfaceId(p DxBgpPeerParameters, vals map[string]cty.Value) {
	vals["virtual_interface_id"] = cty.StringVal(p.VirtualInterfaceId)
}

func EncodeDxBgpPeer_AddressFamily(p DxBgpPeerParameters, vals map[string]cty.Value) {
	vals["address_family"] = cty.StringVal(p.AddressFamily)
}

func EncodeDxBgpPeer_AmazonAddress(p DxBgpPeerParameters, vals map[string]cty.Value) {
	vals["amazon_address"] = cty.StringVal(p.AmazonAddress)
}

func EncodeDxBgpPeer_BgpAsn(p DxBgpPeerParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.NumberIntVal(p.BgpAsn)
}

func EncodeDxBgpPeer_BgpAuthKey(p DxBgpPeerParameters, vals map[string]cty.Value) {
	vals["bgp_auth_key"] = cty.StringVal(p.BgpAuthKey)
}

func EncodeDxBgpPeer_Id(p DxBgpPeerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxBgpPeer_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxBgpPeer_Timeouts_Create(p, ctyVal)
	EncodeDxBgpPeer_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxBgpPeer_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxBgpPeer_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxBgpPeer_AwsDevice(p DxBgpPeerObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}

func EncodeDxBgpPeer_BgpPeerId(p DxBgpPeerObservation, vals map[string]cty.Value) {
	vals["bgp_peer_id"] = cty.StringVal(p.BgpPeerId)
}

func EncodeDxBgpPeer_BgpStatus(p DxBgpPeerObservation, vals map[string]cty.Value) {
	vals["bgp_status"] = cty.StringVal(p.BgpStatus)
}