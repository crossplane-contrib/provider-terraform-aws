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

func EncodeDirectoryServiceDirectory(r DirectoryServiceDirectory) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeDirectoryServiceDirectory_Description(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_EnableSso(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_ShortName(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Edition(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Password(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Size(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Type(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Alias(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Id(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_Name(r.Spec.ForProvider, ctyVal)
	EncodeDirectoryServiceDirectory_ConnectSettings(r.Spec.ForProvider.ConnectSettings, ctyVal)
	EncodeDirectoryServiceDirectory_VpcSettings(r.Spec.ForProvider.VpcSettings, ctyVal)
	EncodeDirectoryServiceDirectory_AccessUrl(r.Status.AtProvider, ctyVal)
	EncodeDirectoryServiceDirectory_DnsIpAddresses(r.Status.AtProvider, ctyVal)
	EncodeDirectoryServiceDirectory_SecurityGroupId(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeDirectoryServiceDirectory_Description(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDirectoryServiceDirectory_EnableSso(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["enable_sso"] = cty.BoolVal(p.EnableSso)
}

func EncodeDirectoryServiceDirectory_ShortName(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["short_name"] = cty.StringVal(p.ShortName)
}

func EncodeDirectoryServiceDirectory_Tags(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDirectoryServiceDirectory_Edition(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["edition"] = cty.StringVal(p.Edition)
}

func EncodeDirectoryServiceDirectory_Password(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["password"] = cty.StringVal(p.Password)
}

func EncodeDirectoryServiceDirectory_Size(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["size"] = cty.StringVal(p.Size)
}

func EncodeDirectoryServiceDirectory_Type(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeDirectoryServiceDirectory_Alias(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["alias"] = cty.StringVal(p.Alias)
}

func EncodeDirectoryServiceDirectory_Id(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDirectoryServiceDirectory_Name(p DirectoryServiceDirectoryParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDirectoryServiceDirectory_ConnectSettings(p ConnectSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDirectoryServiceDirectory_ConnectSettings_VpcId(p, ctyVal)
	EncodeDirectoryServiceDirectory_ConnectSettings_AvailabilityZones(p, ctyVal)
	EncodeDirectoryServiceDirectory_ConnectSettings_ConnectIps(p, ctyVal)
	EncodeDirectoryServiceDirectory_ConnectSettings_CustomerDnsIps(p, ctyVal)
	EncodeDirectoryServiceDirectory_ConnectSettings_CustomerUsername(p, ctyVal)
	EncodeDirectoryServiceDirectory_ConnectSettings_SubnetIds(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["connect_settings"] = cty.ListVal(valsForCollection)
}

func EncodeDirectoryServiceDirectory_ConnectSettings_VpcId(p ConnectSettings, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeDirectoryServiceDirectory_ConnectSettings_AvailabilityZones(p ConnectSettings, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeDirectoryServiceDirectory_ConnectSettings_ConnectIps(p ConnectSettings, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ConnectIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["connect_ips"] = cty.SetVal(colVals)
}

func EncodeDirectoryServiceDirectory_ConnectSettings_CustomerDnsIps(p ConnectSettings, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CustomerDnsIps {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["customer_dns_ips"] = cty.SetVal(colVals)
}

func EncodeDirectoryServiceDirectory_ConnectSettings_CustomerUsername(p ConnectSettings, vals map[string]cty.Value) {
	vals["customer_username"] = cty.StringVal(p.CustomerUsername)
}

func EncodeDirectoryServiceDirectory_ConnectSettings_SubnetIds(p ConnectSettings, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeDirectoryServiceDirectory_VpcSettings(p VpcSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeDirectoryServiceDirectory_VpcSettings_SubnetIds(p, ctyVal)
	EncodeDirectoryServiceDirectory_VpcSettings_VpcId(p, ctyVal)
	EncodeDirectoryServiceDirectory_VpcSettings_AvailabilityZones(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["vpc_settings"] = cty.ListVal(valsForCollection)
}

func EncodeDirectoryServiceDirectory_VpcSettings_SubnetIds(p VpcSettings, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["subnet_ids"] = cty.SetVal(colVals)
}

func EncodeDirectoryServiceDirectory_VpcSettings_VpcId(p VpcSettings, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeDirectoryServiceDirectory_VpcSettings_AvailabilityZones(p VpcSettings, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AvailabilityZones {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["availability_zones"] = cty.SetVal(colVals)
}

func EncodeDirectoryServiceDirectory_AccessUrl(p DirectoryServiceDirectoryObservation, vals map[string]cty.Value) {
	vals["access_url"] = cty.StringVal(p.AccessUrl)
}

func EncodeDirectoryServiceDirectory_DnsIpAddresses(p DirectoryServiceDirectoryObservation, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.DnsIpAddresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["dns_ip_addresses"] = cty.SetVal(colVals)
}

func EncodeDirectoryServiceDirectory_SecurityGroupId(p DirectoryServiceDirectoryObservation, vals map[string]cty.Value) {
	vals["security_group_id"] = cty.StringVal(p.SecurityGroupId)
}