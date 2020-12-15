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
	r, ok := mr.(*LightsailKeyPair)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a LightsailKeyPair.")
	}
	return EncodeLightsailKeyPair(*r), nil
}

func EncodeLightsailKeyPair(r LightsailKeyPair) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLightsailKeyPair_Name(r.Spec.ForProvider, ctyVal)
	EncodeLightsailKeyPair_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeLightsailKeyPair_PgpKey(r.Spec.ForProvider, ctyVal)
	EncodeLightsailKeyPair_PublicKey(r.Spec.ForProvider, ctyVal)
	EncodeLightsailKeyPair_Arn(r.Status.AtProvider, ctyVal)
	EncodeLightsailKeyPair_EncryptedFingerprint(r.Status.AtProvider, ctyVal)
	EncodeLightsailKeyPair_EncryptedPrivateKey(r.Status.AtProvider, ctyVal)
	EncodeLightsailKeyPair_PrivateKey(r.Status.AtProvider, ctyVal)
	EncodeLightsailKeyPair_Fingerprint(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeLightsailKeyPair_Name(p LightsailKeyPairParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLightsailKeyPair_NamePrefix(p LightsailKeyPairParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeLightsailKeyPair_PgpKey(p LightsailKeyPairParameters, vals map[string]cty.Value) {
	vals["pgp_key"] = cty.StringVal(p.PgpKey)
}

func EncodeLightsailKeyPair_PublicKey(p LightsailKeyPairParameters, vals map[string]cty.Value) {
	vals["public_key"] = cty.StringVal(p.PublicKey)
}

func EncodeLightsailKeyPair_Arn(p LightsailKeyPairObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLightsailKeyPair_EncryptedFingerprint(p LightsailKeyPairObservation, vals map[string]cty.Value) {
	vals["encrypted_fingerprint"] = cty.StringVal(p.EncryptedFingerprint)
}

func EncodeLightsailKeyPair_EncryptedPrivateKey(p LightsailKeyPairObservation, vals map[string]cty.Value) {
	vals["encrypted_private_key"] = cty.StringVal(p.EncryptedPrivateKey)
}

func EncodeLightsailKeyPair_PrivateKey(p LightsailKeyPairObservation, vals map[string]cty.Value) {
	vals["private_key"] = cty.StringVal(p.PrivateKey)
}

func EncodeLightsailKeyPair_Fingerprint(p LightsailKeyPairObservation, vals map[string]cty.Value) {
	vals["fingerprint"] = cty.StringVal(p.Fingerprint)
}