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
	r, ok := mr.(*DxHostedPublicVirtualInterface)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDxHostedPublicVirtualInterface(r, ctyValue)
}

func DecodeDxHostedPublicVirtualInterface(prev *DxHostedPublicVirtualInterface, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDxHostedPublicVirtualInterface_CustomerAddress(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_RouteFilterPrefixes(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_Vlan(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_AmazonAddress(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_BgpAsn(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_BgpAuthKey(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_ConnectionId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_AddressFamily(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_Name(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_OwnerAccountId(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_Id(&new.Spec.ForProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDxHostedPublicVirtualInterface_AmazonSideAsn(&new.Status.AtProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_AwsDevice(&new.Status.AtProvider, valMap)
	DecodeDxHostedPublicVirtualInterface_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeDxHostedPublicVirtualInterface_CustomerAddress(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.CustomerAddress = ctwhy.ValueAsString(vals["customer_address"])
}

func DecodeDxHostedPublicVirtualInterface_RouteFilterPrefixes(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["route_filter_prefixes"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.RouteFilterPrefixes = goVals
}

func DecodeDxHostedPublicVirtualInterface_Vlan(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Vlan = ctwhy.ValueAsInt64(vals["vlan"])
}

func DecodeDxHostedPublicVirtualInterface_AmazonAddress(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AmazonAddress = ctwhy.ValueAsString(vals["amazon_address"])
}

func DecodeDxHostedPublicVirtualInterface_BgpAsn(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAsn = ctwhy.ValueAsInt64(vals["bgp_asn"])
}

func DecodeDxHostedPublicVirtualInterface_BgpAuthKey(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.BgpAuthKey = ctwhy.ValueAsString(vals["bgp_auth_key"])
}

func DecodeDxHostedPublicVirtualInterface_ConnectionId(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.ConnectionId = ctwhy.ValueAsString(vals["connection_id"])
}

func DecodeDxHostedPublicVirtualInterface_AddressFamily(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.AddressFamily = ctwhy.ValueAsString(vals["address_family"])
}

func DecodeDxHostedPublicVirtualInterface_Name(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeDxHostedPublicVirtualInterface_OwnerAccountId(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.OwnerAccountId = ctwhy.ValueAsString(vals["owner_account_id"])
}

func DecodeDxHostedPublicVirtualInterface_Id(p *DxHostedPublicVirtualInterfaceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeDxHostedPublicVirtualInterface_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDxHostedPublicVirtualInterface_Timeouts_Create(p, valMap)
	DecodeDxHostedPublicVirtualInterface_Timeouts_Delete(p, valMap)
}

func DecodeDxHostedPublicVirtualInterface_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

func DecodeDxHostedPublicVirtualInterface_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

func DecodeDxHostedPublicVirtualInterface_AmazonSideAsn(p *DxHostedPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsString(vals["amazon_side_asn"])
}

func DecodeDxHostedPublicVirtualInterface_AwsDevice(p *DxHostedPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.AwsDevice = ctwhy.ValueAsString(vals["aws_device"])
}

func DecodeDxHostedPublicVirtualInterface_Arn(p *DxHostedPublicVirtualInterfaceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}