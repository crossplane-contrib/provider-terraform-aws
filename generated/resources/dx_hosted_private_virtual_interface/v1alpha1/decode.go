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
	r, ok := mr.(*DxHostedPrivateVirtualInterface)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxHostedPrivateVirtualInterface(r, ctyValue)
}

func DecodeDxHostedPrivateVirtualInterface(prev *DxHostedPrivateVirtualInterface, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxHostedPrivateVirtualInterface_CustomerAddress(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_AddressFamily(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_BgpAsn(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_Id(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_AmazonAddress(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_BgpAuthKey(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_Mtu(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_ConnectionId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_Name(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_OwnerAccountId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_Vlan(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxHostedPrivateVirtualInterface_Arn(&new.Status.AtProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_AmazonSideAsn(&new.Status.AtProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_AwsDevice(&new.Status.AtProvider, valMap)
	DecodeDxHostedPrivateVirtualInterface_JumboFrameCapable(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDxHostedPrivateVirtualInterface_CustomerAddress(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.CustomerAddress = ctwhy.ValueAsString(vals["customer_address"])
}

func DecodeDxHostedPrivateVirtualInterface_AddressFamily(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AddressFamily = ctwhy.ValueAsString(vals["address_family"])
}

func DecodeDxHostedPrivateVirtualInterface_BgpAsn(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAsn = ctwhy.ValueAsInt64(vals["bgp_asn"])
}

func DecodeDxHostedPrivateVirtualInterface_Id(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDxHostedPrivateVirtualInterface_AmazonAddress(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AmazonAddress = ctwhy.ValueAsString(vals["amazon_address"])
}

func DecodeDxHostedPrivateVirtualInterface_BgpAuthKey(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAuthKey = ctwhy.ValueAsString(vals["bgp_auth_key"])
}

func DecodeDxHostedPrivateVirtualInterface_Mtu(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Mtu = ctwhy.ValueAsInt64(vals["mtu"])
}

func DecodeDxHostedPrivateVirtualInterface_ConnectionId(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.ConnectionId = ctwhy.ValueAsString(vals["connection_id"])
}

func DecodeDxHostedPrivateVirtualInterface_Name(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeDxHostedPrivateVirtualInterface_OwnerAccountId(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.OwnerAccountId = ctwhy.ValueAsString(vals["owner_account_id"])
}

func DecodeDxHostedPrivateVirtualInterface_Vlan(p *DxHostedPrivateVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Vlan = ctwhy.ValueAsInt64(vals["vlan"])
}

func DecodeDxHostedPrivateVirtualInterface_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxHostedPrivateVirtualInterface_Timeouts_Update(p, valMap)
	DecodeDxHostedPrivateVirtualInterface_Timeouts_Create(p, valMap)
	DecodeDxHostedPrivateVirtualInterface_Timeouts_Delete(p, valMap)
}

func DecodeDxHostedPrivateVirtualInterface_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

func DecodeDxHostedPrivateVirtualInterface_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDxHostedPrivateVirtualInterface_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDxHostedPrivateVirtualInterface_Arn(p *DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeDxHostedPrivateVirtualInterface_AmazonSideAsn(p *DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsString(vals["amazon_side_asn"])
}

func DecodeDxHostedPrivateVirtualInterface_AwsDevice(p *DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}

func DecodeDxHostedPrivateVirtualInterface_JumboFrameCapable(p *DxHostedPrivateVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.JumboFrameCapable = ctwhy.ValueAsBool(vals["jumbo_frame_capable"])
}