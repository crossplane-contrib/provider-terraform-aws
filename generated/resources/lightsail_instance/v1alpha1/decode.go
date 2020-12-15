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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*LightsailInstance)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLightsailInstance(r, ctyValue)
}

func DecodeLightsailInstance(prev *LightsailInstance, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLightsailInstance_BundleId(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_UserData(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_AvailabilityZone(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_BlueprintId(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_Name(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_Id(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_KeyPairName(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_Tags(&new.Spec.ForProvider, valMap)
	DecodeLightsailInstance_CreatedAt(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_PublicIpAddress(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_RamSize(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_Username(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_CpuCount(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_IsStaticIp(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_Arn(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_Ipv6Address(&new.Status.AtProvider, valMap)
	DecodeLightsailInstance_PrivateIpAddress(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_BundleId(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	p.BundleId = ctwhy.ValueAsString(vals["bundle_id"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_UserData(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	p.UserData = ctwhy.ValueAsString(vals["user_data"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_AvailabilityZone(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	p.AvailabilityZone = ctwhy.ValueAsString(vals["availability_zone"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_BlueprintId(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	p.BlueprintId = ctwhy.ValueAsString(vals["blueprint_id"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_Name(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_Id(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_KeyPairName(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	p.KeyPairName = ctwhy.ValueAsString(vals["key_pair_name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeLightsailInstance_Tags(p *LightsailInstanceParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_CreatedAt(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.CreatedAt = ctwhy.ValueAsString(vals["created_at"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_PublicIpAddress(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.PublicIpAddress = ctwhy.ValueAsString(vals["public_ip_address"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_RamSize(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.RamSize = ctwhy.ValueAsInt64(vals["ram_size"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_Username(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.Username = ctwhy.ValueAsString(vals["username"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_CpuCount(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.CpuCount = ctwhy.ValueAsInt64(vals["cpu_count"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_IsStaticIp(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.IsStaticIp = ctwhy.ValueAsBool(vals["is_static_ip"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_Arn(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_Ipv6Address(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.Ipv6Address = ctwhy.ValueAsString(vals["ipv6_address"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailInstance_PrivateIpAddress(p *LightsailInstanceObservation, vals map[string]cty.Value) {
	p.PrivateIpAddress = ctwhy.ValueAsString(vals["private_ip_address"])
}