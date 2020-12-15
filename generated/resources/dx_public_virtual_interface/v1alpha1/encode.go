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
	r, ok := mr.(*DxPublicVirtualInterface)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxPublicVirtualInterface.")
	}
	return EncodeDxPublicVirtualInterface(*r), nil
}

func EncodeDxPublicVirtualInterface(r DxPublicVirtualInterface) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxPublicVirtualInterface_Name(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_AddressFamily(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_AmazonAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_RouteFilterPrefixes(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_Vlan(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_BgpAsn(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_BgpAuthKey(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_CustomerAddress(r.Spec.ForProvider, ctyVal)
	EncodeDxPublicVirtualInterface_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxPublicVirtualInterface_Arn(r.Status.AtProvider, ctyVal)
	EncodeDxPublicVirtualInterface_AwsDevice(r.Status.AtProvider, ctyVal)
	EncodeDxPublicVirtualInterface_AmazonSideAsn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxPublicVirtualInterface_Name(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDxPublicVirtualInterface_AddressFamily(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["address_family"] = cty.StringVal(p.AddressFamily)
}

func EncodeDxPublicVirtualInterface_ConnectionId(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeDxPublicVirtualInterface_AmazonAddress(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["amazon_address"] = cty.StringVal(p.AmazonAddress)
}

func EncodeDxPublicVirtualInterface_Tags(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
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

func EncodeDxPublicVirtualInterface_RouteFilterPrefixes(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RouteFilterPrefixes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["route_filter_prefixes"] = cty.SetVal(colVals)
}

func EncodeDxPublicVirtualInterface_Vlan(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["vlan"] = cty.NumberIntVal(p.Vlan)
}

func EncodeDxPublicVirtualInterface_BgpAsn(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_asn"] = cty.NumberIntVal(p.BgpAsn)
}

func EncodeDxPublicVirtualInterface_BgpAuthKey(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["bgp_auth_key"] = cty.StringVal(p.BgpAuthKey)
}

func EncodeDxPublicVirtualInterface_CustomerAddress(p DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	vals["customer_address"] = cty.StringVal(p.CustomerAddress)
}

func EncodeDxPublicVirtualInterface_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxPublicVirtualInterface_Timeouts_Create(p, ctyVal)
	EncodeDxPublicVirtualInterface_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxPublicVirtualInterface_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxPublicVirtualInterface_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxPublicVirtualInterface_Arn(p DxPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDxPublicVirtualInterface_AwsDevice(p DxPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["aws_device"] = cty.StringVal(p.AwsDevice)
}

func EncodeDxPublicVirtualInterface_AmazonSideAsn(p DxPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.StringVal(p.AmazonSideAsn)
}