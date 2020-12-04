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

func EncodeLb(r Lb) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLb_EnableCrossZoneLoadBalancing(r.Spec.ForProvider, ctyVal)
	EncodeLb_EnableHttp2(r.Spec.ForProvider, ctyVal)
	EncodeLb_Internal(r.Spec.ForProvider, ctyVal)
	EncodeLb_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeLb_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeLb_Tags(r.Spec.ForProvider, ctyVal)
	EncodeLb_Name(r.Spec.ForProvider, ctyVal)
	EncodeLb_DropInvalidHeaderFields(r.Spec.ForProvider, ctyVal)
	EncodeLb_EnableDeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeLb_Id(r.Spec.ForProvider, ctyVal)
	EncodeLb_IdleTimeout(r.Spec.ForProvider, ctyVal)
	EncodeLb_IpAddressType(r.Spec.ForProvider, ctyVal)
	EncodeLb_LoadBalancerType(r.Spec.ForProvider, ctyVal)
	EncodeLb_CustomerOwnedIpv4Pool(r.Spec.ForProvider, ctyVal)
	EncodeLb_Subnets(r.Spec.ForProvider, ctyVal)
	EncodeLb_AccessLogs(r.Spec.ForProvider.AccessLogs, ctyVal)
	EncodeLb_SubnetMapping(r.Spec.ForProvider.SubnetMapping, ctyVal)
	EncodeLb_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeLb_VpcId(r.Status.AtProvider, ctyVal)
	EncodeLb_ZoneId(r.Status.AtProvider, ctyVal)
	EncodeLb_ArnSuffix(r.Status.AtProvider, ctyVal)
	EncodeLb_DnsName(r.Status.AtProvider, ctyVal)
	EncodeLb_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLb_EnableCrossZoneLoadBalancing(p LbParameters, vals map[string]cty.Value) {
	vals["enable_cross_zone_load_balancing"] = cty.BoolVal(p.EnableCrossZoneLoadBalancing)
}

func EncodeLb_EnableHttp2(p LbParameters, vals map[string]cty.Value) {
	vals["enable_http2"] = cty.BoolVal(p.EnableHttp2)
}

func EncodeLb_Internal(p LbParameters, vals map[string]cty.Value) {
	vals["internal"] = cty.BoolVal(p.Internal)
}

func EncodeLb_NamePrefix(p LbParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeLb_SecurityGroups(p LbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeLb_Tags(p LbParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeLb_Name(p LbParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLb_DropInvalidHeaderFields(p LbParameters, vals map[string]cty.Value) {
	vals["drop_invalid_header_fields"] = cty.BoolVal(p.DropInvalidHeaderFields)
}

func EncodeLb_EnableDeletionProtection(p LbParameters, vals map[string]cty.Value) {
	vals["enable_deletion_protection"] = cty.BoolVal(p.EnableDeletionProtection)
}

func EncodeLb_Id(p LbParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLb_IdleTimeout(p LbParameters, vals map[string]cty.Value) {
	vals["idle_timeout"] = cty.NumberIntVal(p.IdleTimeout)
}

func EncodeLb_IpAddressType(p LbParameters, vals map[string]cty.Value) {
	vals["ip_address_type"] = cty.StringVal(p.IpAddressType)
}

func EncodeLb_LoadBalancerType(p LbParameters, vals map[string]cty.Value) {
	vals["load_balancer_type"] = cty.StringVal(p.LoadBalancerType)
}

func EncodeLb_CustomerOwnedIpv4Pool(p LbParameters, vals map[string]cty.Value) {
	vals["customer_owned_ipv4_pool"] = cty.StringVal(p.CustomerOwnedIpv4Pool)
}

func EncodeLb_Subnets(p LbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}

func EncodeLb_AccessLogs(p AccessLogs, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLb_AccessLogs_Enabled(p, ctyVal)
	EncodeLb_AccessLogs_Prefix(p, ctyVal)
	EncodeLb_AccessLogs_Bucket(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["access_logs"] = cty.ListVal(valsForCollection)
}

func EncodeLb_AccessLogs_Enabled(p AccessLogs, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeLb_AccessLogs_Prefix(p AccessLogs, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeLb_AccessLogs_Bucket(p AccessLogs, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeLb_SubnetMapping(p SubnetMapping, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLb_SubnetMapping_AllocationId(p, ctyVal)
	EncodeLb_SubnetMapping_OutpostId(p, ctyVal)
	EncodeLb_SubnetMapping_PrivateIpv4Address(p, ctyVal)
	EncodeLb_SubnetMapping_SubnetId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["subnet_mapping"] = cty.SetVal(valsForCollection)
}

func EncodeLb_SubnetMapping_AllocationId(p SubnetMapping, vals map[string]cty.Value) {
	vals["allocation_id"] = cty.StringVal(p.AllocationId)
}

func EncodeLb_SubnetMapping_OutpostId(p SubnetMapping, vals map[string]cty.Value) {
	vals["outpost_id"] = cty.StringVal(p.OutpostId)
}

func EncodeLb_SubnetMapping_PrivateIpv4Address(p SubnetMapping, vals map[string]cty.Value) {
	vals["private_ipv4_address"] = cty.StringVal(p.PrivateIpv4Address)
}

func EncodeLb_SubnetMapping_SubnetId(p SubnetMapping, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeLb_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeLb_Timeouts_Delete(p, ctyVal)
	EncodeLb_Timeouts_Update(p, ctyVal)
	EncodeLb_Timeouts_Create(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeLb_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeLb_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeLb_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeLb_VpcId(p LbObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeLb_ZoneId(p LbObservation, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeLb_ArnSuffix(p LbObservation, vals map[string]cty.Value) {
	vals["arn_suffix"] = cty.StringVal(p.ArnSuffix)
}

func EncodeLb_DnsName(p LbObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeLb_Arn(p LbObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}