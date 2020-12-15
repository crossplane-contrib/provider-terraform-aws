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
	r, ok := mr.(*Cloud9EnvironmentEc2)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloud9EnvironmentEc2(r, ctyValue)
}

func DecodeCloud9EnvironmentEc2(prev *Cloud9EnvironmentEc2, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloud9EnvironmentEc2_InstanceType(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_OwnerArn(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_Tags(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_Description(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_Id(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_AutomaticStopTimeMinutes(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_Name(&new.Spec.ForProvider, valMap)
	DecodeCloud9EnvironmentEc2_Type(&new.Status.AtProvider, valMap)
	DecodeCloud9EnvironmentEc2_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_InstanceType(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	p.InstanceType = ctwhy.ValueAsString(vals["instance_type"])
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_OwnerArn(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	p.OwnerArn = ctwhy.ValueAsString(vals["owner_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_SubnetId(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_Tags(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_Description(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_Id(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_AutomaticStopTimeMinutes(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	p.AutomaticStopTimeMinutes = ctwhy.ValueAsInt64(vals["automatic_stop_time_minutes"])
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_Name(p *Cloud9EnvironmentEc2Parameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_Type(p *Cloud9EnvironmentEc2Observation, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeCloud9EnvironmentEc2_Arn(p *Cloud9EnvironmentEc2Observation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}