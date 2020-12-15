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
	r, ok := mr.(*KmsCiphertext)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a KmsCiphertext.")
	}
	return EncodeKmsCiphertext(*r), nil
}

func EncodeKmsCiphertext(r KmsCiphertext) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeKmsCiphertext_Context(r.Spec.ForProvider, ctyVal)
	EncodeKmsCiphertext_KeyId(r.Spec.ForProvider, ctyVal)
	EncodeKmsCiphertext_Plaintext(r.Spec.ForProvider, ctyVal)
	EncodeKmsCiphertext_CiphertextBlob(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeKmsCiphertext_Context(p KmsCiphertextParameters, vals map[string]cty.Value) {
	if len(p.Context) == 0 {
		vals["context"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Context {
		mVals[key] = cty.StringVal(value)
	}
	vals["context"] = cty.MapVal(mVals)
}

func EncodeKmsCiphertext_KeyId(p KmsCiphertextParameters, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeKmsCiphertext_Plaintext(p KmsCiphertextParameters, vals map[string]cty.Value) {
	vals["plaintext"] = cty.StringVal(p.Plaintext)
}

func EncodeKmsCiphertext_CiphertextBlob(p KmsCiphertextObservation, vals map[string]cty.Value) {
	vals["ciphertext_blob"] = cty.StringVal(p.CiphertextBlob)
}