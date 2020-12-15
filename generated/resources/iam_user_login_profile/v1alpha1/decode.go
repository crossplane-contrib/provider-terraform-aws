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
	r, ok := mr.(*IamUserLoginProfile)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamUserLoginProfile(r, ctyValue)
}

func DecodeIamUserLoginProfile(prev *IamUserLoginProfile, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamUserLoginProfile_Id(&new.Spec.ForProvider, valMap)
	DecodeIamUserLoginProfile_PasswordLength(&new.Spec.ForProvider, valMap)
	DecodeIamUserLoginProfile_PasswordResetRequired(&new.Spec.ForProvider, valMap)
	DecodeIamUserLoginProfile_PgpKey(&new.Spec.ForProvider, valMap)
	DecodeIamUserLoginProfile_User(&new.Spec.ForProvider, valMap)
	DecodeIamUserLoginProfile_EncryptedPassword(&new.Status.AtProvider, valMap)
	DecodeIamUserLoginProfile_KeyFingerprint(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamUserLoginProfile_Id(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserLoginProfile_PasswordLength(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	p.PasswordLength = ctwhy.ValueAsInt64(vals["password_length"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserLoginProfile_PasswordResetRequired(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	p.PasswordResetRequired = ctwhy.ValueAsBool(vals["password_reset_required"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserLoginProfile_PgpKey(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	p.PgpKey = ctwhy.ValueAsString(vals["pgp_key"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserLoginProfile_User(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	p.User = ctwhy.ValueAsString(vals["user"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserLoginProfile_EncryptedPassword(p *IamUserLoginProfileObservation, vals map[string]cty.Value) {
	p.EncryptedPassword = ctwhy.ValueAsString(vals["encrypted_password"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserLoginProfile_KeyFingerprint(p *IamUserLoginProfileObservation, vals map[string]cty.Value) {
	p.KeyFingerprint = ctwhy.ValueAsString(vals["key_fingerprint"])
}