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

func EncodeGlobalacceleratorAccelerator(r GlobalacceleratorAccelerator) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlobalacceleratorAccelerator_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_IpAddressType(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Attributes(r.Spec.ForProvider.Attributes, ctyVal)
	EncodeGlobalacceleratorAccelerator_DnsName(r.Status.AtProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_HostedZoneId(r.Status.AtProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_IpSets(r.Status.AtProvider.IpSets, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGlobalacceleratorAccelerator_Enabled(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeGlobalacceleratorAccelerator_Id(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlobalacceleratorAccelerator_IpAddressType(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["ip_address_type"] = cty.StringVal(p.IpAddressType)
}

func EncodeGlobalacceleratorAccelerator_Name(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlobalacceleratorAccelerator_Tags(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGlobalacceleratorAccelerator_Attributes(p Attributes, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsEnabled(p, ctyVal)
	EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Bucket(p, ctyVal)
	EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Prefix(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["attributes"] = cty.ListVal(valsForCollection)
}

func EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsEnabled(p Attributes, vals map[string]cty.Value) {
	vals["flow_logs_enabled"] = cty.BoolVal(p.FlowLogsEnabled)
}

func EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Bucket(p Attributes, vals map[string]cty.Value) {
	vals["flow_logs_s3_bucket"] = cty.StringVal(p.FlowLogsS3Bucket)
}

func EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Prefix(p Attributes, vals map[string]cty.Value) {
	vals["flow_logs_s3_prefix"] = cty.StringVal(p.FlowLogsS3Prefix)
}

func EncodeGlobalacceleratorAccelerator_DnsName(p GlobalacceleratorAcceleratorObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeGlobalacceleratorAccelerator_HostedZoneId(p GlobalacceleratorAcceleratorObservation, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}

func EncodeGlobalacceleratorAccelerator_IpSets(p []IpSets, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGlobalacceleratorAccelerator_IpSets_IpAddresses(v, ctyVal)
		EncodeGlobalacceleratorAccelerator_IpSets_IpFamily(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["ip_sets"] = cty.ListVal(valsForCollection)
}

func EncodeGlobalacceleratorAccelerator_IpSets_IpAddresses(p IpSets, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IpAddresses {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["ip_addresses"] = cty.ListVal(colVals)
}

func EncodeGlobalacceleratorAccelerator_IpSets_IpFamily(p IpSets, vals map[string]cty.Value) {
	vals["ip_family"] = cty.StringVal(p.IpFamily)
}