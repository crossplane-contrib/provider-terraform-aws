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

func EncodeSubnet(r Subnet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSubnet_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_AvailabilityZoneId(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_Id(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_OutpostArn(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_AssignIpv6AddressOnCreation(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_Ipv6CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_MapPublicIpOnLaunch(r.Spec.ForProvider, ctyVal)
	EncodeSubnet_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeSubnet_Arn(r.Status.AtProvider, ctyVal)
	EncodeSubnet_Ipv6CidrBlockAssociationId(r.Status.AtProvider, ctyVal)
	EncodeSubnet_OwnerId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeSubnet_VpcId(p SubnetParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeSubnet_AvailabilityZoneId(p SubnetParameters, vals map[string]cty.Value) {
	vals["availability_zone_id"] = cty.StringVal(p.AvailabilityZoneId)
}

func EncodeSubnet_Id(p SubnetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSubnet_OutpostArn(p SubnetParameters, vals map[string]cty.Value) {
	vals["outpost_arn"] = cty.StringVal(p.OutpostArn)
}

func EncodeSubnet_Tags(p SubnetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSubnet_AssignIpv6AddressOnCreation(p SubnetParameters, vals map[string]cty.Value) {
	vals["assign_ipv6_address_on_creation"] = cty.BoolVal(p.AssignIpv6AddressOnCreation)
}

func EncodeSubnet_AvailabilityZone(p SubnetParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeSubnet_CidrBlock(p SubnetParameters, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeSubnet_Ipv6CidrBlock(p SubnetParameters, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeSubnet_MapPublicIpOnLaunch(p SubnetParameters, vals map[string]cty.Value) {
	vals["map_public_ip_on_launch"] = cty.BoolVal(p.MapPublicIpOnLaunch)
}

func EncodeSubnet_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeSubnet_Timeouts_Create(p, ctyVal)
	EncodeSubnet_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeSubnet_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeSubnet_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeSubnet_Arn(p SubnetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeSubnet_Ipv6CidrBlockAssociationId(p SubnetObservation, vals map[string]cty.Value) {
	vals["ipv6_cidr_block_association_id"] = cty.StringVal(p.Ipv6CidrBlockAssociationId)
}

func EncodeSubnet_OwnerId(p SubnetObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}