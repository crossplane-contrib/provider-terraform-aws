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
	r, ok := mr.(*DxHostedPrivateVirtualInterfaceAccepter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxHostedPrivateVirtualInterfaceAccepter(r, ctyValue)
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter(prev *DxHostedPrivateVirtualInterfaceAccepter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxHostedPrivateVirtualInterfaceAccepter_VirtualInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterfaceAccepter_VpnGatewayId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterfaceAccepter_DxGatewayId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterfaceAccepter_Id(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterfaceAccepter_Tags(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxHostedPrivateVirtualInterfaceAccepter_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_VirtualInterfaceId(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	p.VirtualInterfaceId = ctwhy.ValueAsString(vals["virtual_interface_id"])
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_VpnGatewayId(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	p.VpnGatewayId = ctwhy.ValueAsString(vals["vpn_gateway_id"])
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_DxGatewayId(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	p.DxGatewayId = ctwhy.ValueAsString(vals["dx_gateway_id"])
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_Id(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_Tags(p *DxHostedPrivateVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Create(p, valMap)
	DecodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Delete(p, valMap)
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDxHostedPrivateVirtualInterfaceAccepter_Arn(p *DxHostedPrivateVirtualInterfaceAccepterObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}