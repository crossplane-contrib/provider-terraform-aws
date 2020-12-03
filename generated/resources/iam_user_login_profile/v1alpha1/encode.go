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

package v1alpha1func EncodeIamUserLoginProfile(r IamUserLoginProfile) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIamUserLoginProfile_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_PasswordLength(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_PasswordResetRequired(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_PgpKey(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_User(r.Spec.ForProvider, ctyVal)
	EncodeIamUserLoginProfile_KeyFingerprint(r.Status.AtProvider, ctyVal)
	EncodeIamUserLoginProfile_EncryptedPassword(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeIamUserLoginProfile_Id(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamUserLoginProfile_PasswordLength(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["password_length"] = cty.IntVal(p.PasswordLength)
}

func EncodeIamUserLoginProfile_PasswordResetRequired(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["password_reset_required"] = cty.BoolVal(p.PasswordResetRequired)
}

func EncodeIamUserLoginProfile_PgpKey(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["pgp_key"] = cty.StringVal(p.PgpKey)
}

func EncodeIamUserLoginProfile_User(p *IamUserLoginProfileParameters, vals map[string]cty.Value) {
	vals["user"] = cty.StringVal(p.User)
}

func EncodeIamUserLoginProfile_KeyFingerprint(p *IamUserLoginProfileObservation, vals map[string]cty.Value) {
	vals["key_fingerprint"] = cty.StringVal(p.KeyFingerprint)
}

func EncodeIamUserLoginProfile_EncryptedPassword(p *IamUserLoginProfileObservation, vals map[string]cty.Value) {
	vals["encrypted_password"] = cty.StringVal(p.EncryptedPassword)
}