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
	r, ok := mr.(*IotRoleAlias)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeIotRoleAlias(r, ctyValue)
}

func DecodeIotRoleAlias(prev *IotRoleAlias, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeIotRoleAlias_Alias(&new.Spec.ForProvider, valMap)
	DecodeIotRoleAlias_CredentialDuration(&new.Spec.ForProvider, valMap)
	DecodeIotRoleAlias_Id(&new.Spec.ForProvider, valMap)
	DecodeIotRoleAlias_RoleArn(&new.Spec.ForProvider, valMap)
	DecodeIotRoleAlias_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeIotRoleAlias_Alias(p *IotRoleAliasParameters, vals map[string]cty.Value) {
	p.Alias = ctwhy.ValueAsString(vals["alias"])
}

//primitiveTypeDecodeTemplate
func DecodeIotRoleAlias_CredentialDuration(p *IotRoleAliasParameters, vals map[string]cty.Value) {
	p.CredentialDuration = ctwhy.ValueAsInt64(vals["credential_duration"])
}

//primitiveTypeDecodeTemplate
func DecodeIotRoleAlias_Id(p *IotRoleAliasParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeIotRoleAlias_RoleArn(p *IotRoleAliasParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeIotRoleAlias_Arn(p *IotRoleAliasObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}