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
	r, ok := mr.(*DefaultSubnet)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DefaultSubnet.")
	}
	return EncodeDefaultSubnet(*r), nil
}

func EncodeDefaultSubnet(r DefaultSubnet) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDefaultSubnet_MapPublicIpOnLaunch(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSubnet_OutpostArn(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSubnet_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSubnet_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSubnet_Id(r.Spec.ForProvider, ctyVal)
	EncodeDefaultSubnet_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDefaultSubnet_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeDefaultSubnet_Arn(r.Status.AtProvider, ctyVal)
	EncodeDefaultSubnet_AssignIpv6AddressOnCreation(r.Status.AtProvider, ctyVal)
	EncodeDefaultSubnet_CidrBlock(r.Status.AtProvider, ctyVal)
	EncodeDefaultSubnet_Ipv6CidrBlock(r.Status.AtProvider, ctyVal)
	EncodeDefaultSubnet_Ipv6CidrBlockAssociationId(r.Status.AtProvider, ctyVal)
	EncodeDefaultSubnet_AvailabilityZoneId(r.Status.AtProvider, ctyVal)
	EncodeDefaultSubnet_VpcId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDefaultSubnet_MapPublicIpOnLaunch(p DefaultSubnetParameters, vals map[string]cty.Value) {
	vals["map_public_ip_on_launch"] = cty.BoolVal(p.MapPublicIpOnLaunch)
}

func EncodeDefaultSubnet_OutpostArn(p DefaultSubnetParameters, vals map[string]cty.Value) {
	vals["outpost_arn"] = cty.StringVal(p.OutpostArn)
}

func EncodeDefaultSubnet_Tags(p DefaultSubnetParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDefaultSubnet_AvailabilityZone(p DefaultSubnetParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeDefaultSubnet_Id(p DefaultSubnetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDefaultSubnet_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeDefaultSubnet_Timeouts_Create(p, ctyVal)
	EncodeDefaultSubnet_Timeouts_Delete(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDefaultSubnet_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDefaultSubnet_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDefaultSubnet_OwnerId(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeDefaultSubnet_Arn(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDefaultSubnet_AssignIpv6AddressOnCreation(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["assign_ipv6_address_on_creation"] = cty.BoolVal(p.AssignIpv6AddressOnCreation)
}

func EncodeDefaultSubnet_CidrBlock(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeDefaultSubnet_Ipv6CidrBlock(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeDefaultSubnet_Ipv6CidrBlockAssociationId(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["ipv6_cidr_block_association_id"] = cty.StringVal(p.Ipv6CidrBlockAssociationId)
}

func EncodeDefaultSubnet_AvailabilityZoneId(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["availability_zone_id"] = cty.StringVal(p.AvailabilityZoneId)
}

func EncodeDefaultSubnet_VpcId(p DefaultSubnetObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}