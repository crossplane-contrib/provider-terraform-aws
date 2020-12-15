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
	r, ok := mr.(*KmsExternalKey)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeKmsExternalKey(r, ctyValue)
}

func DecodeKmsExternalKey(prev *KmsExternalKey, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeKmsExternalKey_Policy(&new.Spec.ForProvider, valMap)
	DecodeKmsExternalKey_ValidTo(&new.Spec.ForProvider, valMap)
	DecodeKmsExternalKey_Enabled(&new.Spec.ForProvider, valMap)
	DecodeKmsExternalKey_KeyMaterialBase64(&new.Spec.ForProvider, valMap)
	DecodeKmsExternalKey_DeletionWindowInDays(&new.Spec.ForProvider, valMap)
	DecodeKmsExternalKey_Description(&new.Spec.ForProvider, valMap)
	DecodeKmsExternalKey_Tags(&new.Spec.ForProvider, valMap)
	DecodeKmsExternalKey_KeyUsage(&new.Status.AtProvider, valMap)
	DecodeKmsExternalKey_Arn(&new.Status.AtProvider, valMap)
	DecodeKmsExternalKey_KeyState(&new.Status.AtProvider, valMap)
	DecodeKmsExternalKey_ExpirationModel(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_Policy(p *KmsExternalKeyParameters, vals map[string]cty.Value) {
	p.Policy = ctwhy.ValueAsString(vals["policy"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_ValidTo(p *KmsExternalKeyParameters, vals map[string]cty.Value) {
	p.ValidTo = ctwhy.ValueAsString(vals["valid_to"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_Enabled(p *KmsExternalKeyParameters, vals map[string]cty.Value) {
	p.Enabled = ctwhy.ValueAsBool(vals["enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_KeyMaterialBase64(p *KmsExternalKeyParameters, vals map[string]cty.Value) {
	p.KeyMaterialBase64 = ctwhy.ValueAsString(vals["key_material_base64"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_DeletionWindowInDays(p *KmsExternalKeyParameters, vals map[string]cty.Value) {
	p.DeletionWindowInDays = ctwhy.ValueAsInt64(vals["deletion_window_in_days"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_Description(p *KmsExternalKeyParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveMapTypeDecodeTemplate
func DecodeKmsExternalKey_Tags(p *KmsExternalKeyParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_KeyUsage(p *KmsExternalKeyObservation, vals map[string]cty.Value) {
	p.KeyUsage = ctwhy.ValueAsString(vals["key_usage"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_Arn(p *KmsExternalKeyObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_KeyState(p *KmsExternalKeyObservation, vals map[string]cty.Value) {
	p.KeyState = ctwhy.ValueAsString(vals["key_state"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsExternalKey_ExpirationModel(p *KmsExternalKeyObservation, vals map[string]cty.Value) {
	p.ExpirationModel = ctwhy.ValueAsString(vals["expiration_model"])
}