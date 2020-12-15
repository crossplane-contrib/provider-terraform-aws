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
	r, ok := mr.(*Ec2TransitGatewayVpcAttachmentAccepter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeEc2TransitGatewayVpcAttachmentAccepter(r, ctyValue)
}

func DecodeEc2TransitGatewayVpcAttachmentAccepter(prev *Ec2TransitGatewayVpcAttachmentAccepter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayAttachmentId(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_Id(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_Tags(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTableAssociation(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTablePropagation(&new.Spec.ForProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_Ipv6Support(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_SubnetIds(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_VpcId(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_VpcOwnerId(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayId(&new.Status.AtProvider, valMap)
	DecodeEc2TransitGatewayVpcAttachmentAccepter_DnsSupport(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayAttachmentId(p *Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	p.TransitGatewayAttachmentId = ctwhy.ValueAsString(vals["transit_gateway_attachment_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_Id(p *Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_Tags(p *Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTableAssociation(p *Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	p.TransitGatewayDefaultRouteTableAssociation = ctwhy.ValueAsBool(vals["transit_gateway_default_route_table_association"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTablePropagation(p *Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	p.TransitGatewayDefaultRouteTablePropagation = ctwhy.ValueAsBool(vals["transit_gateway_default_route_table_propagation"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_Ipv6Support(p *Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	p.Ipv6Support = ctwhy.ValueAsString(vals["ipv6_support"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_SubnetIds(p *Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["subnet_ids"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.SubnetIds = goVals
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_VpcId(p *Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_VpcOwnerId(p *Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	p.VpcOwnerId = ctwhy.ValueAsString(vals["vpc_owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayId(p *Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	p.TransitGatewayId = ctwhy.ValueAsString(vals["transit_gateway_id"])
}

//primitiveTypeDecodeTemplate
func DecodeEc2TransitGatewayVpcAttachmentAccepter_DnsSupport(p *Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	p.DnsSupport = ctwhy.ValueAsString(vals["dns_support"])
}