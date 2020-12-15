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
	r, ok := mr.(*DxHostedTransitVirtualInterfaceAccepter)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DxHostedTransitVirtualInterfaceAccepter.")
	}
	return EncodeDxHostedTransitVirtualInterfaceAccepter(*r), nil
}

func EncodeDxHostedTransitVirtualInterfaceAccepter(r DxHostedTransitVirtualInterfaceAccepter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedTransitVirtualInterfaceAccepter_DxGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_VirtualInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_DxGatewayId(p DxHostedTransitVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["dx_gateway_id"] = cty.StringVal(p.DxGatewayId)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Tags(p DxHostedTransitVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
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

func EncodeDxHostedTransitVirtualInterfaceAccepter_VirtualInterfaceId(p DxHostedTransitVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["virtual_interface_id"] = cty.StringVal(p.VirtualInterfaceId)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Delete(p, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Arn(p DxHostedTransitVirtualInterfaceAccepterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}