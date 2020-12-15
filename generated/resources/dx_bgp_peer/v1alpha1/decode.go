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
	r, ok := mr.(*DxBgpPeer)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxBgpPeer(r, ctyValue)
}

func DecodeDxBgpPeer(prev *DxBgpPeer, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxBgpPeer_AddressFamily(&new.Spec.ForProvider, valMap)
	DecodeDxBgpPeer_BgpAuthKey(&new.Spec.ForProvider, valMap)
	DecodeDxBgpPeer_CustomerAddress(&new.Spec.ForProvider, valMap)
	DecodeDxBgpPeer_AmazonAddress(&new.Spec.ForProvider, valMap)
	DecodeDxBgpPeer_BgpAsn(&new.Spec.ForProvider, valMap)
	DecodeDxBgpPeer_VirtualInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeDxBgpPeer_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxBgpPeer_AwsDevice(&new.Status.AtProvider, valMap)
	DecodeDxBgpPeer_BgpPeerId(&new.Status.AtProvider, valMap)
	DecodeDxBgpPeer_BgpStatus(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_AddressFamily(p *DxBgpPeerParameters, vals map[string]cty.Value) {
	p.AddressFamily = ctwhy.ValueAsString(vals["address_family"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_BgpAuthKey(p *DxBgpPeerParameters, vals map[string]cty.Value) {
	p.BgpAuthKey = ctwhy.ValueAsString(vals["bgp_auth_key"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_CustomerAddress(p *DxBgpPeerParameters, vals map[string]cty.Value) {
	p.CustomerAddress = ctwhy.ValueAsString(vals["customer_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_AmazonAddress(p *DxBgpPeerParameters, vals map[string]cty.Value) {
	p.AmazonAddress = ctwhy.ValueAsString(vals["amazon_address"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_BgpAsn(p *DxBgpPeerParameters, vals map[string]cty.Value) {
	p.BgpAsn = ctwhy.ValueAsInt64(vals["bgp_asn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_VirtualInterfaceId(p *DxBgpPeerParameters, vals map[string]cty.Value) {
	p.VirtualInterfaceId = ctwhy.ValueAsString(vals["virtual_interface_id"])
}

//containerTypeDecodeTemplate
func DecodeDxBgpPeer_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxBgpPeer_Timeouts_Create(p, valMap)
	DecodeDxBgpPeer_Timeouts_Delete(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_AwsDevice(p *DxBgpPeerObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_BgpPeerId(p *DxBgpPeerObservation, vals map[string]cty.Value) {
	p.BgpPeerId = ctwhy.ValueAsString(vals["bgp_peer_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxBgpPeer_BgpStatus(p *DxBgpPeerObservation, vals map[string]cty.Value) {
	p.BgpStatus = ctwhy.ValueAsString(vals["bgp_status"])
}