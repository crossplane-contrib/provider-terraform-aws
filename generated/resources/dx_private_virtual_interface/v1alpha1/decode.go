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
	r, ok := mr.(*DxPrivateVirtualInterface)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxPrivateVirtualInterface(r, ctyValue)
}

func DecodeDxPrivateVirtualInterface(prev *DxPrivateVirtualInterface, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxPrivateVirtualInterface_Tags(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_BgpAsn(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_CustomerAddress(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_DxGatewayId(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_Id(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_Mtu(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_AddressFamily(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_BgpAuthKey(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_VpnGatewayId(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_AmazonAddress(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_ConnectionId(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_Name(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_Vlan(&new.Spec.ForProvider, valMap)
	DecodeDxPrivateVirtualInterface_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxPrivateVirtualInterface_AmazonSideAsn(&new.Status.AtProvider, valMap)
	DecodeDxPrivateVirtualInterface_AwsDevice(&new.Status.AtProvider, valMap)
	DecodeDxPrivateVirtualInterface_Arn(&new.Status.AtProvider, valMap)
	DecodeDxPrivateVirtualInterface_JumboFrameCapable(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Tags(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_BgpAsn(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAsn = ctwhy.ValueAsInt64(vals["bgp_asn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_CustomerAddress(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.CustomerAddress = ctwhy.ValueAsString(vals["customer_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_DxGatewayId(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.DxGatewayId = ctwhy.ValueAsString(vals["dx_gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Id(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Mtu(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Mtu = ctwhy.ValueAsInt64(vals["mtu"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_AddressFamily(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AddressFamily = ctwhy.ValueAsString(vals["address_family"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_BgpAuthKey(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAuthKey = ctwhy.ValueAsString(vals["bgp_auth_key"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_VpnGatewayId(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.VpnGatewayId = ctwhy.ValueAsString(vals["vpn_gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_AmazonAddress(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AmazonAddress = ctwhy.ValueAsString(vals["amazon_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_ConnectionId(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.ConnectionId = ctwhy.ValueAsString(vals["connection_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Name(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Vlan(p *DxPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Vlan = ctwhy.ValueAsInt64(vals["vlan"])
}

//containerTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxPrivateVirtualInterface_Timeouts_Create(p, valMap)
	DecodeDxPrivateVirtualInterface_Timeouts_Delete(p, valMap)
	DecodeDxPrivateVirtualInterface_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_AmazonSideAsn(p *DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsString(vals["amazon_side_asn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_AwsDevice(p *DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_Arn(p *DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPrivateVirtualInterface_JumboFrameCapable(p *DxPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.JumboFrameCapable = ctwhy.ValueAsBool(vals["jumbo_frame_capable"])
}