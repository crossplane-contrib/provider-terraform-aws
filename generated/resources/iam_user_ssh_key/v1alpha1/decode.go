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
	r, ok := mr.(*IamUserSshKey)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIamUserSshKey(r, ctyValue)
}

func DecodeIamUserSshKey(prev *IamUserSshKey, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIamUserSshKey_PublicKey(&new.Spec.ForProvider, valMap)
	DecodeIamUserSshKey_Status(&new.Spec.ForProvider, valMap)
	DecodeIamUserSshKey_Username(&new.Spec.ForProvider, valMap)
	DecodeIamUserSshKey_Encoding(&new.Spec.ForProvider, valMap)
	DecodeIamUserSshKey_SshPublicKeyId(&new.Status.AtProvider, valMap)
	DecodeIamUserSshKey_Fingerprint(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIamUserSshKey_PublicKey(p *IamUserSshKeyParameters, vals map[string]cty.Value) {
	p.PublicKey = ctwhy.ValueAsString(vals["public_key"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserSshKey_Status(p *IamUserSshKeyParameters, vals map[string]cty.Value) {
	p.Status = ctwhy.ValueAsString(vals["status"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserSshKey_Username(p *IamUserSshKeyParameters, vals map[string]cty.Value) {
	p.Username = ctwhy.ValueAsString(vals["username"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserSshKey_Encoding(p *IamUserSshKeyParameters, vals map[string]cty.Value) {
	p.Encoding = ctwhy.ValueAsString(vals["encoding"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserSshKey_SshPublicKeyId(p *IamUserSshKeyObservation, vals map[string]cty.Value) {
	p.SshPublicKeyId = ctwhy.ValueAsString(vals["ssh_public_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeIamUserSshKey_Fingerprint(p *IamUserSshKeyObservation, vals map[string]cty.Value) {
	p.Fingerprint = ctwhy.ValueAsString(vals["fingerprint"])
}