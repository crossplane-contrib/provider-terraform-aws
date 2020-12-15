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
	r, ok := mr.(*Vpc)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Vpc.")
	}
	return EncodeVpc(*r), nil
}

func EncodeVpc(r Vpc) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeVpc_CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeVpc_InstanceTenancy(r.Spec.ForProvider, ctyVal)
	EncodeVpc_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpc_AssignGeneratedIpv6CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeVpc_EnableClassiclink(r.Spec.ForProvider, ctyVal)
	EncodeVpc_EnableDnsHostnames(r.Spec.ForProvider, ctyVal)
	EncodeVpc_EnableDnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeVpc_EnableClassiclinkDnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeVpc_Arn(r.Status.AtProvider, ctyVal)
	EncodeVpc_DhcpOptionsId(r.Status.AtProvider, ctyVal)
	EncodeVpc_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeVpc_DefaultNetworkAclId(r.Status.AtProvider, ctyVal)
	EncodeVpc_Ipv6CidrBlock(r.Status.AtProvider, ctyVal)
	EncodeVpc_DefaultRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeVpc_Ipv6AssociationId(r.Status.AtProvider, ctyVal)
	EncodeVpc_MainRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeVpc_DefaultSecurityGroupId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeVpc_CidrBlock(p VpcParameters, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeVpc_InstanceTenancy(p VpcParameters, vals map[string]cty.Value) {
	vals["instance_tenancy"] = cty.StringVal(p.InstanceTenancy)
}

func EncodeVpc_Tags(p VpcParameters, vals map[string]cty.Value) {
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

func EncodeVpc_AssignGeneratedIpv6CidrBlock(p VpcParameters, vals map[string]cty.Value) {
	vals["assign_generated_ipv6_cidr_block"] = cty.BoolVal(p.AssignGeneratedIpv6CidrBlock)
}

func EncodeVpc_EnableClassiclink(p VpcParameters, vals map[string]cty.Value) {
	vals["enable_classiclink"] = cty.BoolVal(p.EnableClassiclink)
}

func EncodeVpc_EnableDnsHostnames(p VpcParameters, vals map[string]cty.Value) {
	vals["enable_dns_hostnames"] = cty.BoolVal(p.EnableDnsHostnames)
}

func EncodeVpc_EnableDnsSupport(p VpcParameters, vals map[string]cty.Value) {
	vals["enable_dns_support"] = cty.BoolVal(p.EnableDnsSupport)
}

func EncodeVpc_EnableClassiclinkDnsSupport(p VpcParameters, vals map[string]cty.Value) {
	vals["enable_classiclink_dns_support"] = cty.BoolVal(p.EnableClassiclinkDnsSupport)
}

func EncodeVpc_Arn(p VpcObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeVpc_DhcpOptionsId(p VpcObservation, vals map[string]cty.Value) {
	vals["dhcp_options_id"] = cty.StringVal(p.DhcpOptionsId)
}

func EncodeVpc_OwnerId(p VpcObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeVpc_DefaultNetworkAclId(p VpcObservation, vals map[string]cty.Value) {
	vals["default_network_acl_id"] = cty.StringVal(p.DefaultNetworkAclId)
}

func EncodeVpc_Ipv6CidrBlock(p VpcObservation, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeVpc_DefaultRouteTableId(p VpcObservation, vals map[string]cty.Value) {
	vals["default_route_table_id"] = cty.StringVal(p.DefaultRouteTableId)
}

func EncodeVpc_Ipv6AssociationId(p VpcObservation, vals map[string]cty.Value) {
	vals["ipv6_association_id"] = cty.StringVal(p.Ipv6AssociationId)
}

func EncodeVpc_MainRouteTableId(p VpcObservation, vals map[string]cty.Value) {
	vals["main_route_table_id"] = cty.StringVal(p.MainRouteTableId)
}

func EncodeVpc_DefaultSecurityGroupId(p VpcObservation, vals map[string]cty.Value) {
	vals["default_security_group_id"] = cty.StringVal(p.DefaultSecurityGroupId)
}