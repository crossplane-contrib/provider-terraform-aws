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
	
	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*Ec2TransitGatewayVpcAttachmentAccepter)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Ec2TransitGatewayVpcAttachmentAccepter.")
	}
	return EncodeEc2TransitGatewayVpcAttachmentAccepter(*r), nil
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter(r Ec2TransitGatewayVpcAttachmentAccepter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTableAssociation(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTablePropagation(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayAttachmentId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_VpcId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_DnsSupport(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_Ipv6Support(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_VpcOwnerId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayVpcAttachmentAccepter_SubnetIds(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_Tags(p Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTableAssociation(p Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	vals["transit_gateway_default_route_table_association"] = cty.BoolVal(p.TransitGatewayDefaultRouteTableAssociation)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayDefaultRouteTablePropagation(p Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	vals["transit_gateway_default_route_table_propagation"] = cty.BoolVal(p.TransitGatewayDefaultRouteTablePropagation)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayAttachmentId(p Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	vals["transit_gateway_attachment_id"] = cty.StringVal(p.TransitGatewayAttachmentId)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_Id(p Ec2TransitGatewayVpcAttachmentAccepterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_VpcId(p Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_DnsSupport(p Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["dns_support"] = cty.StringVal(p.DnsSupport)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_Ipv6Support(p Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["ipv6_support"] = cty.StringVal(p.Ipv6Support)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_TransitGatewayId(p Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["transit_gateway_id"] = cty.StringVal(p.TransitGatewayId)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_VpcOwnerId(p Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	vals["vpc_owner_id"] = cty.StringVal(p.VpcOwnerId)
}

func EncodeEc2TransitGatewayVpcAttachmentAccepter_SubnetIds(p Ec2TransitGatewayVpcAttachmentAccepterObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}