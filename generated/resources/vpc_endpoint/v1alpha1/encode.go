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

package v1alpha1func EncodeVpcEndpoint(r VpcEndpoint) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeVpcEndpoint_Tags(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_AutoAccept(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_Id(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_SecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_VpcEndpointType(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_SubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_Policy(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_PrivateDnsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_RouteTableIds(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_ServiceName(r.Spec.ForProvider, ctyVal)
	EncodeVpcEndpoint_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeVpcEndpoint_DnsEntry(r.Status.AtProvider.DnsEntry, ctyVal)
	EncodeVpcEndpoint_CidrBlocks(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpoint_PrefixListId(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpoint_RequesterManaged(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpoint_Arn(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpoint_NetworkInterfaceIds(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpoint_OwnerId(r.Status.AtProvider, ctyVal)
	EncodeVpcEndpoint_State(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeVpcEndpoint_Tags(p *VpcEndpointParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeVpcEndpoint_AutoAccept(p *VpcEndpointParameters, vals map[string]cty.Value) {
	vals["auto_accept"] = cty.BoolVal(p.AutoAccept)
}

func EncodeVpcEndpoint_Id(p *VpcEndpointParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeVpcEndpoint_SecurityGroupIds(p *VpcEndpointParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["security_group_ids"] = cty.SetVal(colVals)
}

func EncodeVpcEndpoint_VpcEndpointType(p *VpcEndpointParameters, vals map[string]cty.Value) {
	vals["vpc_endpoint_type"] = cty.StringVal(p.VpcEndpointType)
}

func EncodeVpcEndpoint_VpcId(p *VpcEndpointParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeVpcEndpoint_SubnetIds(p *VpcEndpointParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeVpcEndpoint_Policy(p *VpcEndpointParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeVpcEndpoint_PrivateDnsEnabled(p *VpcEndpointParameters, vals map[string]cty.Value) {
	vals["private_dns_enabled"] = cty.BoolVal(p.PrivateDnsEnabled)
}

func EncodeVpcEndpoint_RouteTableIds(p *VpcEndpointParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RouteTableIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["route_table_ids"] = cty.SetVal(colVals)
}

func EncodeVpcEndpoint_ServiceName(p *VpcEndpointParameters, vals map[string]cty.Value) {
	vals["service_name"] = cty.StringVal(p.ServiceName)
}

func EncodeVpcEndpoint_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeVpcEndpoint_Timeouts_Create(p, ctyVal)
	EncodeVpcEndpoint_Timeouts_Delete(p, ctyVal)
	EncodeVpcEndpoint_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeVpcEndpoint_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeVpcEndpoint_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeVpcEndpoint_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeVpcEndpoint_DnsEntry(p *DnsEntry, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.DnsEntry {
		ctyVal = make(map[string]cty.Value)
		EncodeVpcEndpoint_DnsEntry_DnsName(v, ctyVal)
		EncodeVpcEndpoint_DnsEntry_HostedZoneId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["dns_entry"] = cty.ListVal(valsForCollection)
}

func EncodeVpcEndpoint_DnsEntry_DnsName(p *DnsEntry, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeVpcEndpoint_DnsEntry_HostedZoneId(p *DnsEntry, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeVpcEndpoint_CidrBlocks(p *VpcEndpointObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CidrBlocks {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cidr_blocks"] = cty.ListVal(colVals)
}

func EncodeVpcEndpoint_PrefixListId(p *VpcEndpointObservation, vals map[string]cty.Value) {
	vals["prefix_list_id"] = cty.StringVal(p.PrefixListId)
}

func EncodeVpcEndpoint_RequesterManaged(p *VpcEndpointObservation, vals map[string]cty.Value) {
	vals["requester_managed"] = cty.BoolVal(p.RequesterManaged)
}

func EncodeVpcEndpoint_Arn(p *VpcEndpointObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeVpcEndpoint_NetworkInterfaceIds(p *VpcEndpointObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NetworkInterfaceIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["network_interface_ids"] = cty.SetVal(colVals)
}

func EncodeVpcEndpoint_OwnerId(p *VpcEndpointObservation, vals map[string]cty.Value) {
	vals["owner_id"] = cty.StringVal(p.OwnerId)
}

func EncodeVpcEndpoint_State(p *VpcEndpointObservation, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}