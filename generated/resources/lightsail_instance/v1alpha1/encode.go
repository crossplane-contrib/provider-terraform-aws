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
	r, ok := mr.(*LightsailInstance)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LightsailInstance.")
	}
	return EncodeLightsailInstance(*r), nil
}

func EncodeLightsailInstance(r LightsailInstance) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLightsailInstance_AvailabilityZone(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_KeyPairName(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_Tags(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_BlueprintId(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_UserData(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_BundleId(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_Id(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_Name(r.Spec.ForProvider, ctyVal)
	EncodeLightsailInstance_CreatedAt(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_IsStaticIp(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_Arn(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_CpuCount(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_Ipv6Address(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_PrivateIpAddress(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_Username(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_PublicIpAddress(r.Status.AtProvider, ctyVal)
	EncodeLightsailInstance_RamSize(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeLightsailInstance_AvailabilityZone(p LightsailInstanceParameters, vals map[string]cty.Value) {
	vals["availability_zone"] = cty.StringVal(p.AvailabilityZone)
}

func EncodeLightsailInstance_KeyPairName(p LightsailInstanceParameters, vals map[string]cty.Value) {
	vals["key_pair_name"] = cty.StringVal(p.KeyPairName)
}

func EncodeLightsailInstance_Tags(p LightsailInstanceParameters, vals map[string]cty.Value) {
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

func EncodeLightsailInstance_BlueprintId(p LightsailInstanceParameters, vals map[string]cty.Value) {
	vals["blueprint_id"] = cty.StringVal(p.BlueprintId)
}

func EncodeLightsailInstance_UserData(p LightsailInstanceParameters, vals map[string]cty.Value) {
	vals["user_data"] = cty.StringVal(p.UserData)
}

func EncodeLightsailInstance_BundleId(p LightsailInstanceParameters, vals map[string]cty.Value) {
	vals["bundle_id"] = cty.StringVal(p.BundleId)
}

func EncodeLightsailInstance_Id(p LightsailInstanceParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLightsailInstance_Name(p LightsailInstanceParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLightsailInstance_CreatedAt(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["created_at"] = cty.StringVal(p.CreatedAt)
}

func EncodeLightsailInstance_IsStaticIp(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["is_static_ip"] = cty.BoolVal(p.IsStaticIp)
}

func EncodeLightsailInstance_Arn(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLightsailInstance_CpuCount(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["cpu_count"] = cty.NumberIntVal(p.CpuCount)
}

func EncodeLightsailInstance_Ipv6Address(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["ipv6_address"] = cty.StringVal(p.Ipv6Address)
}

func EncodeLightsailInstance_PrivateIpAddress(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["private_ip_address"] = cty.StringVal(p.PrivateIpAddress)
}

func EncodeLightsailInstance_Username(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeLightsailInstance_PublicIpAddress(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["public_ip_address"] = cty.StringVal(p.PublicIpAddress)
}

func EncodeLightsailInstance_RamSize(p LightsailInstanceObservation, vals map[string]cty.Value) {
	vals["ram_size"] = cty.NumberIntVal(p.RamSize)
}