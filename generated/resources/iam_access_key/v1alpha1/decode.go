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
	r, ok := mr.(*IamAccessKey)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamAccessKey(r, ctyValue)
}

func DecodeIamAccessKey(prev *IamAccessKey, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamAccessKey_PgpKey(&new.Spec.ForProvider, valMap)
	DecodeIamAccessKey_Status(&new.Spec.ForProvider, valMap)
	DecodeIamAccessKey_User(&new.Spec.ForProvider, valMap)
	DecodeIamAccessKey_KeyFingerprint(&new.Status.AtProvider, valMap)
	DecodeIamAccessKey_Secret(&new.Status.AtProvider, valMap)
	DecodeIamAccessKey_SesSmtpPasswordV4(&new.Status.AtProvider, valMap)
	DecodeIamAccessKey_EncryptedSecret(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamAccessKey_PgpKey(p *IamAccessKeyParameters, vals map[string]cty.Value) {
	p.PgpKey = ctwhy.ValueAsString(vals["pgp_key"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccessKey_Status(p *IamAccessKeyParameters, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccessKey_User(p *IamAccessKeyParameters, vals map[string]cty.Value) {
	p.User = ctwhy.ValueAsString(vals["user"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccessKey_KeyFingerprint(p *IamAccessKeyObservation, vals map[string]cty.Value) {
	p.KeyFingerprint = ctwhy.ValueAsString(vals["key_fingerprint"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccessKey_Secret(p *IamAccessKeyObservation, vals map[string]cty.Value) {
	p.Secret = ctwhy.ValueAsString(vals["secret"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccessKey_SesSmtpPasswordV4(p *IamAccessKeyObservation, vals map[string]cty.Value) {
	p.SesSmtpPasswordV4 = ctwhy.ValueAsString(vals["ses_smtp_password_v4"])
}

//primitiveTypeDecodeTemplate
func DecodeIamAccessKey_EncryptedSecret(p *IamAccessKeyObservation, vals map[string]cty.Value) {
	p.EncryptedSecret = ctwhy.ValueAsString(vals["encrypted_secret"])
}