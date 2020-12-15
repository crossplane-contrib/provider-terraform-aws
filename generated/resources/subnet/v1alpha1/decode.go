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
	r, ok := mr.(*Subnet)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSubnet(r, ctyValue)
}

func DecodeSubnet(prev *Subnet, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSubnet_Tags(&new.Spec.ForProvider, valMap)
	DecodeSubnet_VpcId(&new.Spec.ForProvider, valMap)
	DecodeSubnet_CidrBlock(&new.Spec.ForProvider, valMap)
	DecodeSubnet_Ipv6CidrBlock(&new.Spec.ForProvider, valMap)
	DecodeSubnet_OutpostArn(&new.Spec.ForProvider, valMap)
	DecodeSubnet_AssignIpv6AddressOnCreation(&new.Spec.ForProvider, valMap)
	DecodeSubnet_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeSubnet_AvailabilityZoneId(&new.Spec.ForProvider, valMap)
	DecodeSubnet_MapPublicIpOnLaunch(&new.Spec.ForProvider, valMap)
	DecodeSubnet_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeSubnet_Ipv6CidrBlockAssociationId(&new.Status.AtProvider, valMap)
	DecodeSubnet_OwnerId(&new.Status.AtProvider, valMap)
	DecodeSubnet_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeSubnet_Tags(p *SubnetParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_VpcId(p *SubnetParameters, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_CidrBlock(p *SubnetParameters, vals map[string]cty.Value) {
	p.CidrBlock = ctwhy.ValueAsString(vals["cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_Ipv6CidrBlock(p *SubnetParameters, vals map[string]cty.Value) {
	p.Ipv6CidrBlock = ctwhy.ValueAsString(vals["ipv6_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_OutpostArn(p *SubnetParameters, vals map[string]cty.Value) {
	p.OutpostArn = ctwhy.ValueAsString(vals["outpost_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_AssignIpv6AddressOnCreation(p *SubnetParameters, vals map[string]cty.Value) {
	p.AssignIpv6AddressOnCreation = ctwhy.ValueAsBool(vals["assign_ipv6_address_on_creation"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_AvailabilityZone(p *SubnetParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_AvailabilityZoneId(p *SubnetParameters, vals map[string]cty.Value) {
	p.AvailabilityZoneId = ctwhy.ValueAsString(vals["availability_zone_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_MapPublicIpOnLaunch(p *SubnetParameters, vals map[string]cty.Value) {
	p.MapPublicIpOnLaunch = ctwhy.ValueAsBool(vals["map_public_ip_on_launch"])
}

//containerTypeDecodeTemplate
func DecodeSubnet_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeSubnet_Timeouts_Create(p, valMap)
	DecodeSubnet_Timeouts_Delete(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_Ipv6CidrBlockAssociationId(p *SubnetObservation, vals map[string]cty.Value) {
	p.Ipv6CidrBlockAssociationId = ctwhy.ValueAsString(vals["ipv6_cidr_block_association_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_OwnerId(p *SubnetObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSubnet_Arn(p *SubnetObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}