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
	r, ok := mr.(*DefaultVpc)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeDefaultVpc(r, ctyValue)
}

func DecodeDefaultVpc(prev *DefaultVpc, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeDefaultVpc_EnableDnsSupport(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpc_EnableClassiclinkDnsSupport(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpc_EnableDnsHostnames(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpc_EnableClassiclink(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpc_Tags(&new.Spec.ForProvider, valMap)
	DecodeDefaultVpc_CidrBlock(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_DefaultSecurityGroupId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_Ipv6AssociationId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_MainRouteTableId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_DefaultRouteTableId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_OwnerId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_Arn(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_DefaultNetworkAclId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_InstanceTenancy(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_AssignGeneratedIpv6CidrBlock(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_DhcpOptionsId(&new.Status.AtProvider, valMap)
	DecodeDefaultVpc_Ipv6CidrBlock(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_EnableDnsSupport(p *DefaultVpcParameters, vals map[string]cty.Value) {
	p.EnableDnsSupport = ctwhy.ValueAsBool(vals["enable_dns_support"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_EnableClassiclinkDnsSupport(p *DefaultVpcParameters, vals map[string]cty.Value) {
	p.EnableClassiclinkDnsSupport = ctwhy.ValueAsBool(vals["enable_classiclink_dns_support"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_EnableDnsHostnames(p *DefaultVpcParameters, vals map[string]cty.Value) {
	p.EnableDnsHostnames = ctwhy.ValueAsBool(vals["enable_dns_hostnames"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_EnableClassiclink(p *DefaultVpcParameters, vals map[string]cty.Value) {
	p.EnableClassiclink = ctwhy.ValueAsBool(vals["enable_classiclink"])
}

//primitiveMapTypeDecodeTemplate
func DecodeDefaultVpc_Tags(p *DefaultVpcParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_CidrBlock(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.CidrBlock = ctwhy.ValueAsString(vals["cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_DefaultSecurityGroupId(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.DefaultSecurityGroupId = ctwhy.ValueAsString(vals["default_security_group_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_Ipv6AssociationId(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.Ipv6AssociationId = ctwhy.ValueAsString(vals["ipv6_association_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_MainRouteTableId(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.MainRouteTableId = ctwhy.ValueAsString(vals["main_route_table_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_DefaultRouteTableId(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.DefaultRouteTableId = ctwhy.ValueAsString(vals["default_route_table_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_OwnerId(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.OwnerId = ctwhy.ValueAsString(vals["owner_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_Arn(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_DefaultNetworkAclId(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.DefaultNetworkAclId = ctwhy.ValueAsString(vals["default_network_acl_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_InstanceTenancy(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.InstanceTenancy = ctwhy.ValueAsString(vals["instance_tenancy"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_AssignGeneratedIpv6CidrBlock(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.AssignGeneratedIpv6CidrBlock = ctwhy.ValueAsBool(vals["assign_generated_ipv6_cidr_block"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_DhcpOptionsId(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.DhcpOptionsId = ctwhy.ValueAsString(vals["dhcp_options_id"])
}

//primitiveTypeDecodeTemplate
func DecodeDefaultVpc_Ipv6CidrBlock(p *DefaultVpcObservation, vals map[string]cty.Value) {
	p.Ipv6CidrBlock = ctwhy.ValueAsString(vals["ipv6_cidr_block"])
}