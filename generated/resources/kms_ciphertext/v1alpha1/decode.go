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
	r, ok := mr.(*KmsCiphertext)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeKmsCiphertext(r, ctyValue)
}

func DecodeKmsCiphertext(prev *KmsCiphertext, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeKmsCiphertext_Context(&new.Spec.ForProvider, valMap)
	DecodeKmsCiphertext_KeyId(&new.Spec.ForProvider, valMap)
	DecodeKmsCiphertext_Plaintext(&new.Spec.ForProvider, valMap)
	DecodeKmsCiphertext_CiphertextBlob(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeKmsCiphertext_Context(p *KmsCiphertextParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["context"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Context = vMap
}

//primitiveTypeDecodeTemplate
func DecodeKmsCiphertext_KeyId(p *KmsCiphertextParameters, vals map[string]cty.Value) {
	p.KeyId = ctwhy.ValueAsString(vals["key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsCiphertext_Plaintext(p *KmsCiphertextParameters, vals map[string]cty.Value) {
	p.Plaintext = ctwhy.ValueAsString(vals["plaintext"])
}

//primitiveTypeDecodeTemplate
func DecodeKmsCiphertext_CiphertextBlob(p *KmsCiphertextObservation, vals map[string]cty.Value) {
	p.CiphertextBlob = ctwhy.ValueAsString(vals["ciphertext_blob"])
}