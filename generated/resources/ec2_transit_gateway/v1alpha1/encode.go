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

func EncodeEc2TransitGateway(r Ec2TransitGateway) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2TransitGateway_DefaultRouteTableAssociation(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_DefaultRouteTablePropagation(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_DnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_Tags(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_AmazonSideAsn(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_AutoAcceptSharedAttachments(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_VpnEcmpSupport(r.Spec.ForProvider, ctyVal)
	EncodeEc2TransitGateway_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGateway_PropagationDefaultRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGateway_Arn(r.Status.AtProvider, ctyVal)
	EncodeEc2TransitGateway_AssociationDefaultRouteTableId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2TransitGateway_DefaultRouteTableAssociation(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["default_route_table_association"] = cty.StringVal(p.DefaultRouteTableAssociation)
}

func EncodeEc2TransitGateway_Id(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2TransitGateway_DefaultRouteTablePropagation(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["default_route_table_propagation"] = cty.StringVal(p.DefaultRouteTablePropagation)
}

func EncodeEc2TransitGateway_Description(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2TransitGateway_DnsSupport(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["dns_support"] = cty.StringVal(p.DnsSupport)
}

func EncodeEc2TransitGateway_Tags(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeEc2TransitGateway_AmazonSideAsn(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["amazon_side_asn"] = cty.NumberIntVal(p.AmazonSideAsn)
}

func EncodeEc2TransitGateway_AutoAcceptSharedAttachments(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["auto_accept_shared_attachments"] = cty.StringVal(p.AutoAcceptSharedAttachments)
}

func EncodeEc2TransitGateway_VpnEcmpSupport(p Ec2TransitGatewayParameters, vals map[string]cty.Value) {
	vals["vpn_ecmp_support"] = cty.StringVal(p.VpnEcmpSupport)
}

func EncodeEc2TransitGateway_OwnerId(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeEc2TransitGateway_PropagationDefaultRouteTableId(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["propagation_default_route_table_id"] = cty.StringVal(p.PropagationDefaultRouteTableId)
}

func EncodeEc2TransitGateway_Arn(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeEc2TransitGateway_AssociationDefaultRouteTableId(p Ec2TransitGatewayObservation, vals map[string]cty.Value) {
	vals["association_default_route_table_id"] = cty.StringVal(p.AssociationDefaultRouteTableId)
}