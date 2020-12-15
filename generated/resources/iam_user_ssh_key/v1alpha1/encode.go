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
	r, ok := mr.(*IamUserSshKey)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a IamUserSshKey.")
	}
	return EncodeIamUserSshKey(*r), nil
}

func EncodeIamUserSshKey(r IamUserSshKey) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamUserSshKey_PublicKey(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_Status(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_Username(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_Encoding(r.Spec.ForProvider, ctyVal)
	EncodeIamUserSshKey_SshPublicKeyId(r.Status.AtProvider, ctyVal)
	EncodeIamUserSshKey_Fingerprint(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
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

func EncodeIamUserSshKey_SshPublicKeyId(p IamUserSshKeyObservation, vals map[string]cty.Value) {
	vals["ssh_public_key_id"] = cty.StringVal(p.SshPublicKeyId)
}

func EncodeIamUserSshKey_Fingerprint(p IamUserSshKeyObservation, vals map[string]cty.Value) {
	vals["fingerprint"] = cty.StringVal(p.Fingerprint)
}