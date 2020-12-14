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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*DxTransitVirtualInterface)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxTransitVirtualInterface(r, ctyValue)
}

func DecodeDxTransitVirtualInterface(prev *DxTransitVirtualInterface, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxTransitVirtualInterface_AddressFamily(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_BgpAsn(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_CustomerAddress(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_Mtu(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_Id(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_BgpAuthKey(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_Name(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_Tags(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_Vlan(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_AmazonAddress(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_ConnectionId(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_DxGatewayId(&new.Spec.ForProvider, valMap)
	DecodeDxTransitVirtualInterface_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxTransitVirtualInterface_Arn(&new.Status.AtProvider, valMap)
	DecodeDxTransitVirtualInterface_JumboFrameCapable(&new.Status.AtProvider, valMap)
	DecodeDxTransitVirtualInterface_AmazonSideAsn(&new.Status.AtProvider, valMap)
	DecodeDxTransitVirtualInterface_AwsDevice(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDxTransitVirtualInterface_AddressFamily(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AddressFamily = ctwhy.ValueAsString(vals["address_family"])
}

func DecodeDxTransitVirtualInterface_BgpAsn(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAsn = ctwhy.ValueAsInt64(vals["bgp_asn"])
}

func DecodeDxTransitVirtualInterface_CustomerAddress(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.CustomerAddress = ctwhy.ValueAsString(vals["customer_address"])
}

func DecodeDxTransitVirtualInterface_Mtu(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Mtu = ctwhy.ValueAsInt64(vals["mtu"])
}

func DecodeDxTransitVirtualInterface_Id(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDxTransitVirtualInterface_BgpAuthKey(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAuthKey = ctwhy.ValueAsString(vals["bgp_auth_key"])
}

func DecodeDxTransitVirtualInterface_Name(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeDxTransitVirtualInterface_Tags(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDxTransitVirtualInterface_Vlan(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Vlan = ctwhy.ValueAsInt64(vals["vlan"])
}

func DecodeDxTransitVirtualInterface_AmazonAddress(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AmazonAddress = ctwhy.ValueAsString(vals["amazon_address"])
}

func DecodeDxTransitVirtualInterface_ConnectionId(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.ConnectionId = ctwhy.ValueAsString(vals["connection_id"])
}

func DecodeDxTransitVirtualInterface_DxGatewayId(p *DxTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.DxGatewayId = ctwhy.ValueAsString(vals["dx_gateway_id"])
}

func DecodeDxTransitVirtualInterface_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxTransitVirtualInterface_Timeouts_Update(p, valMap)
	DecodeDxTransitVirtualInterface_Timeouts_Create(p, valMap)
	DecodeDxTransitVirtualInterface_Timeouts_Delete(p, valMap)
}

func DecodeDxTransitVirtualInterface_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeDxTransitVirtualInterface_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDxTransitVirtualInterface_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDxTransitVirtualInterface_Arn(p *DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeDxTransitVirtualInterface_JumboFrameCapable(p *DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.JumboFrameCapable = ctwhy.ValueAsBool(vals["jumbo_frame_capable"])
}

func DecodeDxTransitVirtualInterface_AmazonSideAsn(p *DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsString(vals["amazon_side_asn"])
}

func DecodeDxTransitVirtualInterface_AwsDevice(p *DxTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}