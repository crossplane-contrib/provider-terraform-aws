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
	r, ok := mr.(*Ec2TransitGateway)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2TransitGateway(r, ctyValue)
}

func DecodeEc2TransitGateway(prev *Ec2TransitGateway, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2TransitGateway_AmazonSideAsn(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_AutoAcceptSharedAttachments(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_DefaultRouteTablePropagation(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_DnsSupport(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_Tags(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_DefaultRouteTableAssociation(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_Description(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_Id(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_VpnEcmpSupport(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGateway_Arn(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGateway_OwnerId(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGateway_AssociationDefaultRouteTableId(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGateway_PropagationDefaultRouteTableId(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeEc2TransitGateway_AmazonSideAsn(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.AmazonSideAsn = ctwhy.ValueAsInt64(vals["amazon_side_asn"])
}

func DecodeEc2TransitGateway_AutoAcceptSharedAttachments(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.AutoAcceptSharedAttachments = ctwhy.ValueAsString(vals["auto_accept_shared_attachments"])
}

func DecodeEc2TransitGateway_DefaultRouteTablePropagation(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.DefaultRouteTablePropagation = ctwhy.ValueAsString(vals["default_route_table_propagation"])
}

func DecodeEc2TransitGateway_DnsSupport(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.DnsSupport = ctwhy.ValueAsString(vals["dns_support"])
}

func DecodeEc2TransitGateway_Tags(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeEc2TransitGateway_DefaultRouteTableAssociation(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.DefaultRouteTableAssociation = ctwhy.ValueAsString(vals["default_route_table_association"])
}

func DecodeEc2TransitGateway_Description(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeEc2TransitGateway_Id(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeEc2TransitGateway_VpnEcmpSupport(p *Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	p.VpnEcmpSupport = ctwhy.ValueAsString(vals["vpn_ecmp_support"])
}

func DecodeEc2TransitGateway_Arn(p *Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

func DecodeEc2TransitGateway_OwnerId(p *Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

func DecodeEc2TransitGateway_AssociationDefaultRouteTableId(p *Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	p.AssociationDefaultRouteTableId = ctwhy.ValueAsString(vals["association_default_route_table_id"])
}

func DecodeEc2TransitGateway_PropagationDefaultRouteTableId(p *Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	p.PropagationDefaultRouteTableId = ctwhy.ValueAsString(vals["propagation_default_route_table_id"])
}