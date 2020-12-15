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
	r, ok := mr.(*DxHostedTransitVirtualInterface)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxHostedTransitVirtualInterface(r, ctyValue)
}

func DecodeDxHostedTransitVirtualInterface(prev *DxHostedTransitVirtualInterface, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxHostedTransitVirtualInterface_BgpAsn(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_CustomerAddress(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_Id(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_AmazonAddress(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_OwnerAccountId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_AddressFamily(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_BgpAuthKey(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_ConnectionId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_Mtu(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_Name(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_Vlan(&new.Spec.ForProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxHostedTransitVirtualInterface_AwsDevice(&new.Status.AtProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_JumboFrameCapable(&new.Status.AtProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_AmazonSideAsn(&new.Status.AtProvider, valMap)
	DecodeDxHostedTransitVirtualInterface_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_BgpAsn(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAsn = ctwhy.ValueAsInt64(vals["bgp_asn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_CustomerAddress(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.CustomerAddress = ctwhy.ValueAsString(vals["customer_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Id(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_AmazonAddress(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AmazonAddress = ctwhy.ValueAsString(vals["amazon_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_OwnerAccountId(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.OwnerAccountId = ctwhy.ValueAsString(vals["owner_account_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_AddressFamily(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AddressFamily = ctwhy.ValueAsString(vals["address_family"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_BgpAuthKey(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAuthKey = ctwhy.ValueAsString(vals["bgp_auth_key"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_ConnectionId(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.ConnectionId = ctwhy.ValueAsString(vals["connection_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Mtu(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Mtu = ctwhy.ValueAsInt64(vals["mtu"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Name(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Vlan(p *DxHostedTransitVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Vlan = ctwhy.ValueAsInt64(vals["vlan"])
}

//containerTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxHostedTransitVirtualInterface_Timeouts_Create(p, valMap)
	DecodeDxHostedTransitVirtualInterface_Timeouts_Delete(p, valMap)
	DecodeDxHostedTransitVirtualInterface_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_AwsDevice(p *DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_JumboFrameCapable(p *DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.JumboFrameCapable = ctwhy.ValueAsBool(vals["jumbo_frame_capable"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_AmazonSideAsn(p *DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsString(vals["amazon_side_asn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxHostedTransitVirtualInterface_Arn(p *DxHostedTransitVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}