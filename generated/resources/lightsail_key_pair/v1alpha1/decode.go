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
	r, ok := mr.(*LightsailKeyPair)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeLightsailKeyPair(r, ctyValue)
}

func DecodeLightsailKeyPair(prev *LightsailKeyPair, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeLightsailKeyPair_PublicKey(&new.Spec.ForProvider, valMap)
	DecodeLightsailKeyPair_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeLightsailKeyPair_PgpKey(&new.Spec.ForProvider, valMap)
	DecodeLightsailKeyPair_Name(&new.Spec.ForProvider, valMap)
	DecodeLightsailKeyPair_PrivateKey(&new.Status.AtProvider, valMap)
	DecodeLightsailKeyPair_Fingerprint(&new.Status.AtProvider, valMap)
	DecodeLightsailKeyPair_Arn(&new.Status.AtProvider, valMap)
	DecodeLightsailKeyPair_EncryptedFingerprint(&new.Status.AtProvider, valMap)
	DecodeLightsailKeyPair_EncryptedPrivateKey(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_PublicKey(p *LightsailKeyPairParameters, vals map[string]cty.Value) {
	p.PublicKey = ctwhy.ValueAsString(vals["public_key"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_NamePrefix(p *LightsailKeyPairParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_PgpKey(p *LightsailKeyPairParameters, vals map[string]cty.Value) {
	p.PgpKey = ctwhy.ValueAsString(vals["pgp_key"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_Name(p *LightsailKeyPairParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_PrivateKey(p *LightsailKeyPairObservation, vals map[string]cty.Value) {
	p.PrivateKey = ctwhy.ValueAsString(vals["private_key"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_Fingerprint(p *LightsailKeyPairObservation, vals map[string]cty.Value) {
	p.Fingerprint = ctwhy.ValueAsString(vals["fingerprint"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_Arn(p *LightsailKeyPairObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_EncryptedFingerprint(p *LightsailKeyPairObservation, vals map[string]cty.Value) {
	p.EncryptedFingerprint = ctwhy.ValueAsString(vals["encrypted_fingerprint"])
}

//primitiveTypeDecodeTemplate
func DecodeLightsailKeyPair_EncryptedPrivateKey(p *LightsailKeyPairObservation, vals map[string]cty.Value) {
	p.EncryptedPrivateKey = ctwhy.ValueAsString(vals["encrypted_private_key"])
}