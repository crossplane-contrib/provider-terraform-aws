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
	r, ok := mr.(*DxHostedTransitVirtualInterface)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxHostedTransitVirtualInterface.")
	}
	return EncodeDxHostedTransitVirtualInterface(*r), nil
}

func EncodeDxHostedTransitVirtualInterface(r DxHostedTransitVirtualInterface) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedTransitVirtualInterface_Mtu(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_AddressFamily(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_AmazonAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_BgpAuthKey(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_CustomerAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_OwnerAccountId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_Vlan(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxHostedTransitVirtualInterface_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_AmazonSideAsn(r.Status.AtProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterface_AwsDevice(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDxHostedTransitVirtualInterface_Mtu(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["mtu"] = cty.NumberIntVal(p.Mtu)
}

func EncodeDxHostedTransitVirtualInterface_Name(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxHostedTransitVirtualInterface_AddressFamily(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["address_family"] = cty.StringVal(p.AddressFamily)
}

func EncodeDxHostedTransitVirtualInterface_AmazonAddress(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["amazon_address"] = cty.StringVal(p.AmazonAddress)
}

func EncodeDxHostedTransitVirtualInterface_BgpAuthKey(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_auth_key"] = cty.StringVal(p.BgpAuthKey)
}

func EncodeDxHostedTransitVirtualInterface_CustomerAddress(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["customer_address"] = cty.StringVal(p.CustomerAddress)
}

func EncodeDxHostedTransitVirtualInterface_ConnectionId(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeDxHostedTransitVirtualInterface_OwnerAccountId(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["owner_account_id"] = cty.StringVal(p.OwnerAccountId)
}

func EncodeDxHostedTransitVirtualInterface_BgpAsn(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.NumberIntVal(p.BgpAsn)
}

func EncodeDxHostedTransitVirtualInterface_Vlan(p DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["vlan"] = cty.NumberIntVal(p.Vlan)
}

func EncodeDxHostedTransitVirtualInterface_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedTransitVirtualInterface_Timeouts_Delete(p, ctyVal)
	EncodeDxHostedTransitVirtualInterface_Timeouts_Update(p, ctyVal)
	EncodeDxHostedTransitVirtualInterface_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxHostedTransitVirtualInterface_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxHostedTransitVirtualInterface_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDxHostedTransitVirtualInterface_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxHostedTransitVirtualInterface_Arn(p DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxHostedTransitVirtualInterface_JumboFrameCapable(p DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}

func EncodeDxHostedTransitVirtualInterface_AmazonSideAsn(p DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}

func EncodeDxHostedTransitVirtualInterface_AwsDevice(p DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}