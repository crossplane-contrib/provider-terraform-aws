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
	r, ok := mr.(*DxTransitVirtualInterface)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxTransitVirtualInterface.")
	}
	return EncodeDxTransitVirtualInterface(*r), nil
}

func EncodeDxTransitVirtualInterface(r DxTransitVirtualInterface) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxTransitVirtualInterface_AddressFamily(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_AmazonAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_Mtu(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_Vlan(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_DxGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_BgpAuthKey(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_CustomerAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxTransitVirtualInterface_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxTransitVirtualInterface_AmazonSideAsn(r.Status.AtProvider, ctyVal)
	EncodeDxTransitVirtualInterface_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxTransitVirtualInterface_AwsDevice(r.Status.AtProvider, ctyVal)
	EncodeDxTransitVirtualInterface_JumboFrameCapable(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxTransitVirtualInterface_AddressFamily(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["address_family"] = cty.StringVal(p.AddressFamily)
}

func EncodeDxTransitVirtualInterface_AmazonAddress(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["amazon_address"] = cty.StringVal(p.AmazonAddress)
}

func EncodeDxTransitVirtualInterface_Mtu(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["mtu"] = cty.NumberIntVal(p.Mtu)
}

func EncodeDxTransitVirtualInterface_Vlan(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["vlan"] = cty.NumberIntVal(p.Vlan)
}

func EncodeDxTransitVirtualInterface_BgpAsn(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.NumberIntVal(p.BgpAsn)
}

func EncodeDxTransitVirtualInterface_ConnectionId(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeDxTransitVirtualInterface_Tags(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDxTransitVirtualInterface_DxGatewayId(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["dx_gateway_id"] = cty.StringVal(p.DxGatewayId)
}

func EncodeDxTransitVirtualInterface_Name(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxTransitVirtualInterface_BgpAuthKey(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_auth_key"] = cty.StringVal(p.BgpAuthKey)
}

func EncodeDxTransitVirtualInterface_CustomerAddress(p DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["customer_address"] = cty.StringVal(p.CustomerAddress)
}

func EncodeDxTransitVirtualInterface_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxTransitVirtualInterface_Timeouts_Update(p, ctyVal)
	EncodeDxTransitVirtualInterface_Timeouts_Create(p, ctyVal)
	EncodeDxTransitVirtualInterface_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxTransitVirtualInterface_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDxTransitVirtualInterface_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxTransitVirtualInterface_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxTransitVirtualInterface_AmazonSideAsn(p DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}

func EncodeDxTransitVirtualInterface_Arn(p DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxTransitVirtualInterface_AwsDevice(p DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}

func EncodeDxTransitVirtualInterface_JumboFrameCapable(p DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["jumbo_frame_capable"] = cty.BoolVal(p.JumboFrameCapable)
}