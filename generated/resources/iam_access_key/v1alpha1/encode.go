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
	r, ok := mr.(*IamAccessKey)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamAccessKey.")
	}
	return EncodeIamAccessKey(*r), nil
}

func EncodeIamAccessKey(r IamAccessKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamAccessKey_PgpKey(r.Spec.ForProvider, ctyVal)
	EncodeIamAccessKey_Status(r.Spec.ForProvider, ctyVal)
	EncodeIamAccessKey_User(r.Spec.ForProvider, ctyVal)
	EncodeIamAccessKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamAccessKey_KeyFingerprint(r.Status.AtProvider, ctyVal)
	EncodeIamAccessKey_Secret(r.Status.AtProvider, ctyVal)
	EncodeIamAccessKey_SesSmtpPasswordV4(r.Status.AtProvider, ctyVal)
	EncodeIamAccessKey_EncryptedSecret(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeIamAccessKey_PgpKey(p IamAccessKeyParameters, vals map[string]cty.Value) {
	vals["pgp_key"] = cty.StringVal(p.PgpKey)
}

func EncodeIamAccessKey_Status(p IamAccessKeyParameters, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeIamAccessKey_User(p IamAccessKeyParameters, vals map[string]cty.Value) {
	vals["user"] = cty.StringVal(p.User)
}

func EncodeIamAccessKey_Id(p IamAccessKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamAccessKey_KeyFingerprint(p IamAccessKeyObservation, vals map[string]cty.Value) {
	vals["key_fingerprint"] = cty.StringVal(p.KeyFingerprint)
}

func EncodeIamAccessKey_Secret(p IamAccessKeyObservation, vals map[string]cty.Value) {
	vals["secret"] = cty.StringVal(p.Secret)
}

func EncodeIamAccessKey_SesSmtpPasswordV4(p IamAccessKeyObservation, vals map[string]cty.Value) {
	vals["ses_smtp_password_v4"] = cty.StringVal(p.SesSmtpPasswordV4)
}

func EncodeIamAccessKey_EncryptedSecret(p IamAccessKeyObservation, vals map[string]cty.Value) {
	vals["encrypted_secret"] = cty.StringVal(p.EncryptedSecret)
}