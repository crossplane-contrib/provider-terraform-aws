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

func EncodeDxHostedTransitVirtualInterfaceAccepter(r DxHostedTransitVirtualInterfaceAccepter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDxHostedTransitVirtualInterfaceAccepter_DxGatewayId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Id(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_VirtualInterfaceId(r.Spec.ForProvider, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_DxGatewayId(p DxHostedTransitVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["dx_gateway_id"] = cty.StringVal(p.DxGatewayId)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Id(p DxHostedTransitVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Tags(p DxHostedTransitVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
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
	EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Create(p, ctyVal)
	EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDxHostedTransitVirtualInterfaceAccepter_Arn(p DxHostedTransitVirtualInterfaceAccepterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}