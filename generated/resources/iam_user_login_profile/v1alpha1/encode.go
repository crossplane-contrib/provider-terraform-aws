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
	r, ok := mr.(*IamUserLoginProfile)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamUserLoginProfile.")
	}
	return EncodeIamUserLoginProfile(*r), nil
}

func EncodeIamUserLoginProfile(r IamUserLoginProfile) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamUserLoginProfile_PasswordLength(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_PasswordResetRequired(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_PgpKey(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_User(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_EncryptedPassword(r.Status.AtProvider, ctyVal)
	EncodeIamUserLoginProfile_KeyFingerprint(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamUserLoginProfile_PasswordLength(p IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["password_length"] = cty.NumberIntVal(p.PasswordLength)
}

func EncodeIamUserLoginProfile_PasswordResetRequired(p IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["password_reset_required"] = cty.BoolVal(p.PasswordResetRequired)
}

func EncodeIamUserLoginProfile_PgpKey(p IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["pgp_key"] = cty.StringVal(p.PgpKey)
}

func EncodeIamUserLoginProfile_User(p IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["user"] = cty.StringVal(p.User)
}

func EncodeIamUserLoginProfile_EncryptedPassword(p IamUserLoginProfileObservation, vals map[string]cty.Value) {
	vals["encrypted_password"] = cty.StringVal(p.EncryptedPassword)
}

func EncodeIamUserLoginProfile_KeyFingerprint(p IamUserLoginProfileObservation, vals map[string]cty.Value) {
	vals["key_fingerprint"] = cty.StringVal(p.KeyFingerprint)
}