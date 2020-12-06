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
	"github.com/zclconf/go-cty/cty"
)

func EncodeEc2TransitGatewayRouteTableAssociation(r Ec2TransitGatewayRouteTableAssociation) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGatewayRouteTableAssociation_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTableAssociation_TransitGatewayAttachmentId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTableAssociation_TransitGatewayRouteTableId(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTableAssociation_ResourceId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGatewayRouteTableAssociation_ResourceType(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGatewayRouteTableAssociation_Id(p Ec2TransitGatewayRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGatewayRouteTableAssociation_TransitGatewayAttachmentId(p Ec2TransitGatewayRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["transit_gateway_attachment_id"] = cty.StringVal(p.TransitGatewayAttachmentId)
}

func EncodeEc2TransitGatewayRouteTableAssociation_TransitGatewayRouteTableId(p Ec2TransitGatewayRouteTableAssociationParameters, vals map[string]cty.Value) {
	vals["transit_gateway_route_table_id"] = cty.StringVal(p.TransitGatewayRouteTableId)
}

func EncodeEc2TransitGatewayRouteTableAssociation_ResourceId(p Ec2TransitGatewayRouteTableAssociationObservation, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeEc2TransitGatewayRouteTableAssociation_ResourceType(p Ec2TransitGatewayRouteTableAssociationObservation, vals map[string]cty.Value) {
	vals["resource_type"] = cty.StringVal(p.ResourceType)
}