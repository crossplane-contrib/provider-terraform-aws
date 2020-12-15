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
	r, ok := mr.(*SsmActivation)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSsmActivation(r, ctyValue)
}

func DecodeSsmActivation(prev *SsmActivation, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSsmActivation_RegistrationLimit(&new.Spec.ForProvider, valMap)
	DecodeSsmActivation_IamRole(&new.Spec.ForProvider, valMap)
	DecodeSsmActivation_Id(&new.Spec.ForProvider, valMap)
	DecodeSsmActivation_Name(&new.Spec.ForProvider, valMap)
	DecodeSsmActivation_Tags(&new.Spec.ForProvider, valMap)
	DecodeSsmActivation_Description(&new.Spec.ForProvider, valMap)
	DecodeSsmActivation_ExpirationDate(&new.Spec.ForProvider, valMap)
	DecodeSsmActivation_Expired(&new.Status.AtProvider, valMap)
	DecodeSsmActivation_RegistrationCount(&new.Status.AtProvider, valMap)
	DecodeSsmActivation_ActivationCode(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_RegistrationLimit(p *SsmActivationParameters, vals map[string]cty.Value) {
	p.RegistrationLimit = ctwhy.ValueAsInt64(vals["registration_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_IamRole(p *SsmActivationParameters, vals map[string]cty.Value) {
	p.IamRole = ctwhy.ValueAsString(vals["iam_role"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_Id(p *SsmActivationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_Name(p *SsmActivationParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeSsmActivation_Tags(p *SsmActivationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_Description(p *SsmActivationParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_ExpirationDate(p *SsmActivationParameters, vals map[string]cty.Value) {
	p.ExpirationDate = ctwhy.ValueAsString(vals["expiration_date"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_Expired(p *SsmActivationObservation, vals map[string]cty.Value) {
	p.Expired = ctwhy.ValueAsBool(vals["expired"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_RegistrationCount(p *SsmActivationObservation, vals map[string]cty.Value) {
	p.RegistrationCount = ctwhy.ValueAsInt64(vals["registration_count"])
}

//primitiveTypeDecodeTemplate
func DecodeSsmActivation_ActivationCode(p *SsmActivationObservation, vals map[string]cty.Value) {
	p.ActivationCode = ctwhy.ValueAsString(vals["activation_code"])
}