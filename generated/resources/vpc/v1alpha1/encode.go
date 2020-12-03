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

package v1alpha1func EncodeVpc(r Vpc) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeVpc_EnableClassiclink(r.Spec.ForProvider, ctyVal)
	EncodeVpc_InstanceTenancy(r.Spec.ForProvider, ctyVal)
	EncodeVpc_CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeVpc_EnableClassiclinkDnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeVpc_EnableDnsSupport(r.Spec.ForProvider, ctyVal)
	EncodeVpc_AssignGeneratedIpv6CidrBlock(r.Spec.ForProvider, ctyVal)
	EncodeVpc_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpc_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpc_EnableDnsHostnames(r.Spec.ForProvider, ctyVal)
	EncodeVpc_Arn(r.Status.AtProvider, ctyVal)
	EncodeVpc_Ipv6CidrBlock(r.Status.AtProvider, ctyVal)
	EncodeVpc_MainRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeVpc_DefaultRouteTableId(r.Status.AtProvider, ctyVal)
	EncodeVpc_DefaultSecurityGroupId(r.Status.AtProvider, ctyVal)
	EncodeVpc_DhcpOptionsId(r.Status.AtProvider, ctyVal)
	EncodeVpc_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeVpc_DefaultNetworkAclId(r.Status.AtProvider, ctyVal)
	EncodeVpc_Ipv6AssociationId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeVpc_EnableClassiclink(p *VpcParameters, vals map[string]cty.Value) {
	vals["enable_classiclink"] = cty.BoolVal(p.EnableClassiclink)
}

func EncodeVpc_InstanceTenancy(p *VpcParameters, vals map[string]cty.Value) {
	vals["instance_tenancy"] = cty.StringVal(p.InstanceTenancy)
}

func EncodeVpc_CidrBlock(p *VpcParameters, vals map[string]cty.Value) {
	vals["cidr_block"] = cty.StringVal(p.CidrBlock)
}

func EncodeVpc_EnableClassiclinkDnsSupport(p *VpcParameters, vals map[string]cty.Value) {
	vals["enable_classiclink_dns_support"] = cty.BoolVal(p.EnableClassiclinkDnsSupport)
}

func EncodeVpc_EnableDnsSupport(p *VpcParameters, vals map[string]cty.Value) {
	vals["enable_dns_support"] = cty.BoolVal(p.EnableDnsSupport)
}

func EncodeVpc_AssignGeneratedIpv6CidrBlock(p *VpcParameters, vals map[string]cty.Value) {
	vals["assign_generated_ipv6_cidr_block"] = cty.BoolVal(p.AssignGeneratedIpv6CidrBlock)
}

func EncodeVpc_Id(p *VpcParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpc_Tags(p *VpcParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpc_EnableDnsHostnames(p *VpcParameters, vals map[string]cty.Value) {
	vals["enable_dns_hostnames"] = cty.BoolVal(p.EnableDnsHostnames)
}

func EncodeVpc_Arn(p *VpcObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeVpc_Ipv6CidrBlock(p *VpcObservation, vals map[string]cty.Value) {
	vals["ipv6_cidr_block"] = cty.StringVal(p.Ipv6CidrBlock)
}

func EncodeVpc_MainRouteTableId(p *VpcObservation, vals map[string]cty.Value) {
	vals["main_route_table_id"] = cty.StringVal(p.MainRouteTableId)
}

func EncodeVpc_DefaultRouteTableId(p *VpcObservation, vals map[string]cty.Value) {
	vals["default_route_table_id"] = cty.StringVal(p.DefaultRouteTableId)
}

func EncodeVpc_DefaultSecurityGroupId(p *VpcObservation, vals map[string]cty.Value) {
	vals["default_security_group_id"] = cty.StringVal(p.DefaultSecurityGroupId)
}

func EncodeVpc_DhcpOptionsId(p *VpcObservation, vals map[string]cty.Value) {
	vals["dhcp_options_id"] = cty.StringVal(p.DhcpOptionsId)
}

func EncodeVpc_OwnerId(p *VpcObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeVpc_DefaultNetworkAclId(p *VpcObservation, vals map[string]cty.Value) {
	vals["default_network_acl_id"] = cty.StringVal(p.DefaultNetworkAclId)
}

func EncodeVpc_Ipv6AssociationId(p *VpcObservation, vals map[string]cty.Value) {
	vals["ipv6_association_id"] = cty.StringVal(p.Ipv6AssociationId)
}