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
	r, ok := mr.(*DxHostedPublicVirtualInterfaceAccepter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxHostedPublicVirtualInterfaceAccepter(r, ctyValue)
}

func DecodeDxHostedPublicVirtualInterfaceAccepter(prev *DxHostedPublicVirtualInterfaceAccepter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxHostedPublicVirtualInterfaceAccepter_Id(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterfaceAccepter_Tags(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterfaceAccepter_VirtualInterfaceId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterfaceAccepter_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxHostedPublicVirtualInterfaceAccepter_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDxHostedPublicVirtualInterfaceAccepter_Id(p *DxHostedPublicVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDxHostedPublicVirtualInterfaceAccepter_Tags(p *DxHostedPublicVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeDxHostedPublicVirtualInterfaceAccepter_VirtualInterfaceId(p *DxHostedPublicVirtualInterfaceAccepterParameters, vals map[string]cty.Value) {
	p.VirtualInterfaceId = ctwhy.ValueAsString(vals["virtual_interface_id"])
}

func DecodeDxHostedPublicVirtualInterfaceAccepter_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxHostedPublicVirtualInterfaceAccepter_Timeouts_Create(p, valMap)
	DecodeDxHostedPublicVirtualInterfaceAccepter_Timeouts_Delete(p, valMap)
}

func DecodeDxHostedPublicVirtualInterfaceAccepter_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDxHostedPublicVirtualInterfaceAccepter_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDxHostedPublicVirtualInterfaceAccepter_Arn(p *DxHostedPublicVirtualInterfaceAccepterObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}