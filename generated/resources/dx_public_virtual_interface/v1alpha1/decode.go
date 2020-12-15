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
	r, ok := mr.(*DxPublicVirtualInterface)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxPublicVirtualInterface(r, ctyValue)
}

func DecodeDxPublicVirtualInterface(prev *DxPublicVirtualInterface, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxPublicVirtualInterface_Name(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_AddressFamily(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_ConnectionId(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_AmazonAddress(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_Tags(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_RouteFilterPrefixes(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_Vlan(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_BgpAsn(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_BgpAuthKey(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_CustomerAddress(&new.Spec.ForProvider, valMap)
	DecodeDxPublicVirtualInterface_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxPublicVirtualInterface_Arn(&new.Status.AtProvider, valMap)
	DecodeDxPublicVirtualInterface_AwsDevice(&new.Status.AtProvider, valMap)
	DecodeDxPublicVirtualInterface_AmazonSideAsn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_Name(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_AddressFamily(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AddressFamily = ctwhy.ValueAsString(vals["address_family"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_ConnectionId(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.ConnectionId = ctwhy.ValueAsString(vals["connection_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_AmazonAddress(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AmazonAddress = ctwhy.ValueAsString(vals["amazon_address"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_Tags(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveCollectionTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_RouteFilterPrefixes(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["route_filter_prefixes"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.RouteFilterPrefixes = goVals
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_Vlan(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Vlan = ctwhy.ValueAsInt64(vals["vlan"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_BgpAsn(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAsn = ctwhy.ValueAsInt64(vals["bgp_asn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_BgpAuthKey(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAuthKey = ctwhy.ValueAsString(vals["bgp_auth_key"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_CustomerAddress(p *DxPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.CustomerAddress = ctwhy.ValueAsString(vals["customer_address"])
}

//containerTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxPublicVirtualInterface_Timeouts_Create(p, valMap)
	DecodeDxPublicVirtualInterface_Timeouts_Delete(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_Arn(p *DxPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_AwsDevice(p *DxPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}

//primitiveTypeDecodeTemplate
func DecodeDxPublicVirtualInterface_AmazonSideAsn(p *DxPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsString(vals["amazon_side_asn"])
}