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
	r, ok := mr.(*KmsKey)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeKmsKey(r, ctyValue)
}

func DecodeKmsKey(prev *KmsKey, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeKmsKey_Tags(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_EnableKeyRotation(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_IsEnabled(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_KeyUsage(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_Policy(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_CustomerMasterKeySpec(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_DeletionWindowInDays(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_Description(&new.Spec.ForProvider, valMap)
	DecodeKmsKey_Arn(&new.Status.AtProvider, valMap)
	DecodeKmsKey_KeyId(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeKmsKey_Tags(p *KmsKeyParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_EnableKeyRotation(p *KmsKeyParameters, vals map[string]cty.Value) {
	p.EnableKeyRotation = ctwhy.ValueAsBool(vals["enable_key_rotation"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_IsEnabled(p *KmsKeyParameters, vals map[string]cty.Value) {
	p.IsEnabled = ctwhy.ValueAsBool(vals["is_enabled"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_KeyUsage(p *KmsKeyParameters, vals map[string]cty.Value) {
	p.KeyUsage = ctwhy.ValueAsString(vals["key_usage"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_Policy(p *KmsKeyParameters, vals map[string]cty.Value) {
	p.Policy = ctwhy.ValueAsString(vals["policy"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_CustomerMasterKeySpec(p *KmsKeyParameters, vals map[string]cty.Value) {
	p.CustomerMasterKeySpec = ctwhy.ValueAsString(vals["customer_master_key_spec"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_DeletionWindowInDays(p *KmsKeyParameters, vals map[string]cty.Value) {
	p.DeletionWindowInDays = ctwhy.ValueAsInt64(vals["deletion_window_in_days"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_Description(p *KmsKeyParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_Arn(p *KmsKeyObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsKey_KeyId(p *KmsKeyObservation, vals map[string]cty.Value) {
	p.KeyId = ctwhy.ValueAsString(vals["key_id"])
}