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

func EncodeDxHostedPublicVirtualInterface(r DxHostedPublicVirtualInterface) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedPublicVirtualInterface_RouteFilterPrefixes(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_AmazonAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_BgpAuthKey(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_Vlan(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_CustomerAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_OwnerAccountId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_AddressFamily(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxHostedPublicVirtualInterface_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_AwsDevice(r.Status.AtProvider, ctyVal)
	EncodeDxHostedPublicVirtualInterface_AmazonSideAsn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxHostedPublicVirtualInterface_RouteFilterPrefixes(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RouteFilterPrefixes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["route_filter_prefixes"] = cty.SetVal(colVals)
}

func EncodeDxHostedPublicVirtualInterface_AmazonAddress(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["amazon_address"] = cty.StringVal(p.AmazonAddress)
}

func EncodeDxHostedPublicVirtualInterface_BgpAuthKey(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_auth_key"] = cty.StringVal(p.BgpAuthKey)
}

func EncodeDxHostedPublicVirtualInterface_Name(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxHostedPublicVirtualInterface_Vlan(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["vlan"] = cty.NumberIntVal(p.Vlan)
}

func EncodeDxHostedPublicVirtualInterface_CustomerAddress(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["customer_address"] = cty.StringVal(p.CustomerAddress)
}

func EncodeDxHostedPublicVirtualInterface_Id(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxHostedPublicVirtualInterface_OwnerAccountId(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["owner_account_id"] = cty.StringVal(p.OwnerAccountId)
}

func EncodeDxHostedPublicVirtualInterface_BgpAsn(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.NumberIntVal(p.BgpAsn)
}

func EncodeDxHostedPublicVirtualInterface_ConnectionId(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeDxHostedPublicVirtualInterface_AddressFamily(p DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["address_family"] = cty.StringVal(p.AddressFamily)
}

func EncodeDxHostedPublicVirtualInterface_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedPublicVirtualInterface_Timeouts_Create(p, ctyVal)
	EncodeDxHostedPublicVirtualInterface_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxHostedPublicVirtualInterface_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxHostedPublicVirtualInterface_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxHostedPublicVirtualInterface_Arn(p DxHostedPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxHostedPublicVirtualInterface_AwsDevice(p DxHostedPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}

func EncodeDxHostedPublicVirtualInterface_AmazonSideAsn(p DxHostedPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}