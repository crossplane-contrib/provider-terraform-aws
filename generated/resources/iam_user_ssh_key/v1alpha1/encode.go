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
	"github.com/zclconf/go-cty/cty"
)

func EncodeIamUserSshKey(r IamUserSshKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamUserSshKey_PublicKey(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_Status(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_Username(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_Encoding(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_SshPublicKeyId(r.Status.AtProvider, ctyVal)
	EncodeIamUserSshKey_Fingerprint(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamUserSshKey_PublicKey(p IamUserSshKeyParameters, vals map[string]cty.Value) {
	vals["public_key"] = cty.StringVal(p.PublicKey)
}

func EncodeIamUserSshKey_Status(p IamUserSshKeyParameters, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeIamUserSshKey_Username(p IamUserSshKeyParameters, vals map[string]cty.Value) {
	vals["username"] = cty.StringVal(p.Username)
}

func EncodeIamUserSshKey_Encoding(p IamUserSshKeyParameters, vals map[string]cty.Value) {
	vals["encoding"] = cty.StringVal(p.Encoding)
}

func EncodeIamUserSshKey_Id(p IamUserSshKeyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamUserSshKey_SshPublicKeyId(p IamUserSshKeyObservation, vals map[string]cty.Value) {
	vals["ssh_public_key_id"] = cty.StringVal(p.SshPublicKeyId)
}

func EncodeIamUserSshKey_Fingerprint(p IamUserSshKeyObservation, vals map[string]cty.Value) {
	vals["fingerprint"] = cty.StringVal(p.Fingerprint)
}