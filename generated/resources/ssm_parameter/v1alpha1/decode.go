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
	r, ok := mr.(*SsmParameter)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSsmParameter(r, ctyValue)
}

func DecodeSsmParameter(prev *SsmParameter, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSsmParameter_Overwrite(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Tags(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Tier(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Arn(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_DataType(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Description(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_KeyId(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Name(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Type(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Value(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_AllowedPattern(&new.Spec.ForProvider, valMap)
	DecodeSsmParameter_Version(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Overwrite(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.Overwrite = ctwhy.ValueAsBool(vals["overwrite"])
}

//primitiveMapTypeDecodeTemplate
func DecodeSsmParameter_Tags(p *SsmParameterParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Tier(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.Tier = ctwhy.ValueAsString(vals["tier"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Arn(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_DataType(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.DataType = ctwhy.ValueAsString(vals["data_type"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Description(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_KeyId(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.KeyId = ctwhy.ValueAsString(vals["key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Name(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Type(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Value(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.Value = ctwhy.ValueAsString(vals["value"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_AllowedPattern(p *SsmParameterParameters, vals map[string]cty.Value) {
	p.AllowedPattern = ctwhy.ValueAsString(vals["allowed_pattern"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmParameter_Version(p *SsmParameterObservation, vals map[string]cty.Value) {
	p.Version = ctwhy.ValueAsInt64(vals["version"])
}