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
	r, ok := mr.(*GlobalacceleratorAccelerator)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a GlobalacceleratorAccelerator.")
	}
	return EncodeGlobalacceleratorAccelerator(*r), nil
}

func EncodeGlobalacceleratorAccelerator(r GlobalacceleratorAccelerator) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGlobalacceleratorAccelerator_Name(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Id(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_IpAddressType(r.Spec.ForProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_Attributes(r.Spec.ForProvider.Attributes, ctyVal)
	EncodeGlobalacceleratorAccelerator_IpSets(r.Status.AtProvider.IpSets, ctyVal)
	EncodeGlobalacceleratorAccelerator_DnsName(r.Status.AtProvider, ctyVal)
	EncodeGlobalacceleratorAccelerator_HostedZoneId(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeGlobalacceleratorAccelerator_Name(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGlobalacceleratorAccelerator_Tags(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
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

func EncodeGlobalacceleratorAccelerator_Enabled(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeGlobalacceleratorAccelerator_Id(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGlobalacceleratorAccelerator_IpAddressType(p GlobalacceleratorAcceleratorParameters, vals map[string]cty.Value) {
	vals["ip_address_type"] = cty.StringVal(p.IpAddressType)
}

func EncodeGlobalacceleratorAccelerator_Attributes(p Attributes, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Prefix(p, ctyVal)
	EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsEnabled(p, ctyVal)
	EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Bucket(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["attributes"] = cty.ListVal(valsForCollection)
}

func EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Prefix(p Attributes, vals map[string]cty.Value) {
	vals["flow_logs_s3_prefix"] = cty.StringVal(p.FlowLogsS3Prefix)
}

func EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsEnabled(p Attributes, vals map[string]cty.Value) {
	vals["flow_logs_enabled"] = cty.BoolVal(p.FlowLogsEnabled)
}

func EncodeGlobalacceleratorAccelerator_Attributes_FlowLogsS3Bucket(p Attributes, vals map[string]cty.Value) {
	vals["flow_logs_s3_bucket"] = cty.StringVal(p.FlowLogsS3Bucket)
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

func EncodeGlobalacceleratorAccelerator_DnsName(p GlobalacceleratorAcceleratorObservation, vals map[string]cty.Value) {
	vals["dns_name"] = cty.StringVal(p.DnsName)
}

func EncodeGlobalacceleratorAccelerator_HostedZoneId(p GlobalacceleratorAcceleratorObservation, vals map[string]cty.Value) {
	vals["hosted_zone_id"] = cty.StringVal(p.HostedZoneId)
}