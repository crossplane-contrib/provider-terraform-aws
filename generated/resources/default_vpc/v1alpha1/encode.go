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
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*DefaultVpc)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a DefaultVpc.")
	}
	return EncodeDefaultVpc(*r), nil
}

func EncodeDefaultVpc(r DefaultVpc) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDefaultVpc_EnableDnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpc_EnableClassiclinkDnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpc_EnableDnsHostnames(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpc_EnableClassiclink(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpc_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDefaultVpc_CidrBlock(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_DefaultSecurityGroupId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_Ipv6AssociationId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_MainRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_DefaultRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_Arn(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_DefaultNetworkAclId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_InstanceTenancy(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_AssignGeneratedIpv6CidrBlock(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_DhcpOptionsId(r.Status.AtProvider, ctyVal)
	EncodeDefaultVpc_Ipv6CidrBlock(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeDefaultVpc_EnableDnsSupport(p DefaultVpcParameters, vals map[string]cty.Value) {
	vals["enable_dns_support"] = cty.BoolVal(p.EnableDnsSupport)
}

func EncodeDefaultVpc_EnableClassiclinkDnsSupport(p DefaultVpcParameters, vals map[string]cty.Value) {
	vals["enable_classiclink_dns_support"] = cty.BoolVal(p.EnableClassiclinkDnsSupport)
}

func EncodeDefaultVpc_EnableDnsHostnames(p DefaultVpcParameters, vals map[string]cty.Value) {
	vals["enable_dns_hostnames"] = cty.BoolVal(p.EnableDnsHostnames)
}

func EncodeDefaultVpc_EnableClassiclink(p DefaultVpcParameters, vals map[string]cty.Value) {
	vals["enable_classiclink"] = cty.BoolVal(p.EnableClassiclink)
}

func EncodeDefaultVpc_Tags(p DefaultVpcParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDefaultVpc_CidrBlock(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeDefaultVpc_DefaultSecurityGroupId(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["default_security_group_id"] = cty.StringVal(p.DefaultSecurityGroupId)
}

func EncodeDefaultVpc_Ipv6AssociationId(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["ipv6_association_id"] = cty.StringVal(p.Ipv6AssociationId)
}

func EncodeDefaultVpc_MainRouteTableId(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["main_route_table_id"] = cty.StringVal(p.MainRouteTableId)
}

func EncodeDefaultVpc_DefaultRouteTableId(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["default_route_table_id"] = cty.StringVal(p.DefaultRouteTableId)
}

func EncodeDefaultVpc_OwnerId(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeDefaultVpc_Arn(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDefaultVpc_DefaultNetworkAclId(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["default_network_acl_id"] = cty.StringVal(p.DefaultNetworkAclId)
}

func EncodeDefaultVpc_InstanceTenancy(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["instance_tenancy"] = cty.StringVal(p.InstanceTenancy)
}

func EncodeDefaultVpc_AssignGeneratedIpv6CidrBlock(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["assign_generated_ipv6_cidr_block"] = cty.BoolVal(p.AssignGeneratedIpv6CidrBlock)
}

func EncodeDefaultVpc_DhcpOptionsId(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["dhcp_options_id"] = cty.StringVal(p.DhcpOptionsId)
}

func EncodeDefaultVpc_Ipv6CidrBlock(p DefaultVpcObservation, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}