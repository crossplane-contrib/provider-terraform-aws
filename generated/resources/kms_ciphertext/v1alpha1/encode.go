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

package v1alpha1func EncodeKmsCiphertext(r KmsCiphertext) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeKmsCiphertext_Plaintext(r.Spec.ForProvider, ctyVal)
	EncodeKmsCiphertext_Context(r.Spec.ForProvider, ctyVal)
	EncodeKmsCiphertext_Id(r.Spec.ForProvider, ctyVal)
	EncodeKmsCiphertext_KeyId(r.Spec.ForProvider, ctyVal)
	EncodeKmsCiphertext_CiphertextBlob(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeKmsCiphertext_Plaintext(p *KmsCiphertextParameters, vals map[string]cty.Value) {
	vals["plaintext"] = cty.StringVal(p.Plaintext)
}

func EncodeKmsCiphertext_Context(p *KmsCiphertextParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Context {
		mVals[key] = cty.StringVal(value)
	}
	vals["context"] = cty.MapVal(mVals)
}

func EncodeKmsCiphertext_Id(p *KmsCiphertextParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeKmsCiphertext_KeyId(p *KmsCiphertextParameters, vals map[string]cty.Value) {
	vals["key_id"] = cty.StringVal(p.KeyId)
}

func EncodeKmsCiphertext_CiphertextBlob(p *KmsCiphertextObservation, vals map[string]cty.Value) {
	vals["ciphertext_blob"] = cty.StringVal(p.CiphertextBlob)
}