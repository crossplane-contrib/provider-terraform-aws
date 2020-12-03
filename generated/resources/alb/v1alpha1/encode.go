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

package v1alpha1func EncodeAlb(r Alb) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAlb_DropInvalidHeaderFields(r.Spec.ForProvider, ctyVal)
	EncodeAlb_IdleTimeout(r.Spec.ForProvider, ctyVal)
	EncodeAlb_Internal(r.Spec.ForProvider, ctyVal)
	EncodeAlb_IpAddressType(r.Spec.ForProvider, ctyVal)
	EncodeAlb_LoadBalancerType(r.Spec.ForProvider, ctyVal)
	EncodeAlb_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeAlb_Subnets(r.Spec.ForProvider, ctyVal)
	EncodeAlb_EnableCrossZoneLoadBalancing(r.Spec.ForProvider, ctyVal)
	EncodeAlb_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAlb_CustomerOwnedIpv4Pool(r.Spec.ForProvider, ctyVal)
	EncodeAlb_EnableDeletionProtection(r.Spec.ForProvider, ctyVal)
	EncodeAlb_EnableHttp2(r.Spec.ForProvider, ctyVal)
	EncodeAlb_Name(r.Spec.ForProvider, ctyVal)
	EncodeAlb_Id(r.Spec.ForProvider, ctyVal)
	EncodeAlb_SecurityGroups(r.Spec.ForProvider, ctyVal)
	EncodeAlb_AccessLogs(r.Spec.ForProvider.AccessLogs, ctyVal)
	EncodeAlb_SubnetMapping(r.Spec.ForProvider.SubnetMapping, ctyVal)
	EncodeAlb_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeAlb_DnsName(r.Status.AtProvider, ctyVal)
	EncodeAlb_VpcId(r.Status.AtProvider, ctyVal)
	EncodeAlb_ZoneId(r.Status.AtProvider, ctyVal)
	EncodeAlb_Arn(r.Status.AtProvider, ctyVal)
	EncodeAlb_ArnSuffix(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAlb_DropInvalidHeaderFields(p *AlbParameters, vals map[string]cty.Value) {
	vals["drop_invalid_header_fields"] = cty.BoolVal(p.DropInvalidHeaderFields)
}

func EncodeAlb_IdleTimeout(p *AlbParameters, vals map[string]cty.Value) {
	vals["idle_timeout"] = cty.IntVal(p.IdleTimeout)
}

func EncodeAlb_Internal(p *AlbParameters, vals map[string]cty.Value) {
	vals["internal"] = cty.BoolVal(p.Internal)
}

func EncodeAlb_IpAddressType(p *AlbParameters, vals map[string]cty.Value) {
	vals["ip_address_type"] = cty.StringVal(p.IpAddressType)
}

func EncodeAlb_LoadBalancerType(p *AlbParameters, vals map[string]cty.Value) {
	vals["load_balancer_type"] = cty.StringVal(p.LoadBalancerType)
}

func EncodeAlb_NamePrefix(p *AlbParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeAlb_Subnets(p *AlbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Subnets {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnets"] = cty.SetVal(colVals)
}

func EncodeAlb_EnableCrossZoneLoadBalancing(p *AlbParameters, vals map[string]cty.Value) {
	vals["enable_cross_zone_load_balancing"] = cty.BoolVal(p.EnableCrossZoneLoadBalancing)
}

func EncodeAlb_Tags(p *AlbParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAlb_CustomerOwnedIpv4Pool(p *AlbParameters, vals map[string]cty.Value) {
	vals["customer_owned_ipv4_pool"] = cty.StringVal(p.CustomerOwnedIpv4Pool)
}

func EncodeAlb_EnableDeletionProtection(p *AlbParameters, vals map[string]cty.Value) {
	vals["enable_deletion_protection"] = cty.BoolVal(p.EnableDeletionProtection)
}

func EncodeAlb_EnableHttp2(p *AlbParameters, vals map[string]cty.Value) {
	vals["enable_http2"] = cty.BoolVal(p.EnableHttp2)
}

func EncodeAlb_Name(p *AlbParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAlb_Id(p *AlbParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAlb_SecurityGroups(p *AlbParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_groups"] = cty.SetVal(colVals)
}

func EncodeAlb_AccessLogs(p *AccessLogs, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AccessLogs {
		ctyVal = make(map[string]cty.Value)
		EncodeAlb_AccessLogs_Bucket(v, ctyVal)
		EncodeAlb_AccessLogs_Enabled(v, ctyVal)
		EncodeAlb_AccessLogs_Prefix(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["access_logs"] = cty.ListVal(valsForCollection)
}

func EncodeAlb_AccessLogs_Bucket(p *AccessLogs, vals map[string]cty.Value) {
	vals["bucket"] = cty.StringVal(p.Bucket)
}

func EncodeAlb_AccessLogs_Enabled(p *AccessLogs, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAlb_AccessLogs_Prefix(p *AccessLogs, vals map[string]cty.Value) {
	vals["prefix"] = cty.StringVal(p.Prefix)
}

func EncodeAlb_SubnetMapping(p *SubnetMapping, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SubnetMapping {
		ctyVal = make(map[string]cty.Value)
		EncodeAlb_SubnetMapping_AllocationId(v, ctyVal)
		EncodeAlb_SubnetMapping_OutpostId(v, ctyVal)
		EncodeAlb_SubnetMapping_PrivateIpv4Address(v, ctyVal)
		EncodeAlb_SubnetMapping_SubnetId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["subnet_mapping"] = cty.SetVal(valsForCollection)
}

func EncodeAlb_SubnetMapping_AllocationId(p *SubnetMapping, vals map[string]cty.Value) {
	vals["allocation_id"] = cty.StringVal(p.AllocationId)
}

func EncodeAlb_SubnetMapping_OutpostId(p *SubnetMapping, vals map[string]cty.Value) {
	vals["outpost_id"] = cty.StringVal(p.OutpostId)
}

func EncodeAlb_SubnetMapping_PrivateIpv4Address(p *SubnetMapping, vals map[string]cty.Value) {
	vals["private_ipv4_address"] = cty.StringVal(p.PrivateIpv4Address)
}

func EncodeAlb_SubnetMapping_SubnetId(p *SubnetMapping, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeAlb_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeAlb_Timeouts_Create(p, ctyVal)
	EncodeAlb_Timeouts_Delete(p, ctyVal)
	EncodeAlb_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeAlb_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeAlb_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeAlb_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeAlb_DnsName(p *AlbObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeAlb_VpcId(p *AlbObservation, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeAlb_ZoneId(p *AlbObservation, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}

func EncodeAlb_Arn(p *AlbObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAlb_ArnSuffix(p *AlbObservation, vals map[string]cty.Value) {
	vals["arn_suffix"] = cty.StringVal(p.ArnSuffix)
}