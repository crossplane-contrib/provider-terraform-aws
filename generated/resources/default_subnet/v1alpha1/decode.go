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
	r, ok := mr.(*DefaultSubnet)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDefaultSubnet(r, ctyValue)
}

func DecodeDefaultSubnet(prev *DefaultSubnet, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDefaultSubnet_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeDefaultSubnet_Id(&new.Spec.ForProvider, valMap)
	DecodeDefaultSubnet_MapPublicIpOnLaunch(&new.Spec.ForProvider, valMap)
	DecodeDefaultSubnet_OutpostArn(&new.Spec.ForProvider, valMap)
	DecodeDefaultSubnet_Tags(&new.Spec.ForProvider, valMap)
	DecodeDefaultSubnet_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeDefaultSubnet_AvailabilityZoneId(&new.Status.AtProvider, valMap)
	DecodeDefaultSubnet_Ipv6CidrBlock(&new.Status.AtProvider, valMap)
	DecodeDefaultSubnet_OwnerId(&new.Status.AtProvider, valMap)
	DecodeDefaultSubnet_Arn(&new.Status.AtProvider, valMap)
	DecodeDefaultSubnet_AssignIpv6AddressOnCreation(&new.Status.AtProvider, valMap)
	DecodeDefaultSubnet_CidrBlock(&new.Status.AtProvider, valMap)
	DecodeDefaultSubnet_Ipv6CidrBlockAssociationId(&new.Status.AtProvider, valMap)
	DecodeDefaultSubnet_VpcId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_AvailabilityZone(p *DefaultSubnetParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_Id(p *DefaultSubnetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_MapPublicIpOnLaunch(p *DefaultSubnetParameters, vals map[string]cty.Value) {
	p.MapPublicIpOnLaunch = ctwhy.ValueAsBool(vals["map_public_ip_on_launch"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_OutpostArn(p *DefaultSubnetParameters, vals map[string]cty.Value) {
	p.OutpostArn = ctwhy.ValueAsString(vals["outpost_arn"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDefaultSubnet_Tags(p *DefaultSubnetParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//containerTypeDecodeTemplate
func DecodeDefaultSubnet_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeDefaultSubnet_Timeouts_Create(p, valMap)
	DecodeDefaultSubnet_Timeouts_Delete(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	p.Create = ctwhy.ValueAsString(vals["create"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	p.Delete = ctwhy.ValueAsString(vals["delete"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_AvailabilityZoneId(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.AvailabilityZoneId = ctwhy.ValueAsString(vals["availability_zone_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_Ipv6CidrBlock(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.Ipv6CidrBlock = ctwhy.ValueAsString(vals["ipv6_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_OwnerId(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_Arn(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_AssignIpv6AddressOnCreation(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.AssignIpv6AddressOnCreation = ctwhy.ValueAsBool(vals["assign_ipv6_address_on_creation"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_CidrBlock(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.CidrBlock = ctwhy.ValueAsString(vals["cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_Ipv6CidrBlockAssociationId(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.Ipv6CidrBlockAssociationId = ctwhy.ValueAsString(vals["ipv6_cidr_block_association_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultSubnet_VpcId(p *DefaultSubnetObservation, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}