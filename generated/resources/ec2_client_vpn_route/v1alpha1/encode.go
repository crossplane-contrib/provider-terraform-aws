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

func EncodeEc2ClientVpnRoute(r Ec2ClientVpnRoute) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeEc2ClientVpnRoute_ClientVpnEndpointId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Description(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_DestinationCidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Id(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_TargetVpcSubnetId(r.Spec.ForProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Origin(r.Status.AtProvider, ctyVal)
	EncodeEc2ClientVpnRoute_Type(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeEc2ClientVpnRoute_ClientVpnEndpointId(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["client_vpn_endpoint_id"] = cty.StringVal(p.ClientVpnEndpointId)
}

func EncodeEc2ClientVpnRoute_Description(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeEc2ClientVpnRoute_DestinationCidrBlock(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["destination_cidr_block"] = cty.StringVal(p.DestinationCidrBlock)
}

func EncodeEc2ClientVpnRoute_Id(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeEc2ClientVpnRoute_TargetVpcSubnetId(p Ec2ClientVpnRouteParameters, vals map[string]cty.Value) {
	vals["target_vpc_subnet_id"] = cty.StringVal(p.TargetVpcSubnetId)
}

func EncodeEc2ClientVpnRoute_Origin(p Ec2ClientVpnRouteObservation, vals map[string]cty.Value) {
	vals["origin"] = cty.StringVal(p.Origin)
}

func EncodeEc2ClientVpnRoute_Type(p Ec2ClientVpnRouteObservation, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}