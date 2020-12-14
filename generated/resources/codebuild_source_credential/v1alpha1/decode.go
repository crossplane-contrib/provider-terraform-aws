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
	r, ok := mr.(*CodebuildSourceCredential)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCodebuildSourceCredential(r, ctyValue)
}

func DecodeCodebuildSourceCredential(prev *CodebuildSourceCredential, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCodebuildSourceCredential_Id(&new.Spec.ForProvider, valMap)
	DecodeCodebuildSourceCredential_ServerType(&new.Spec.ForProvider, valMap)
	DecodeCodebuildSourceCredential_Token(&new.Spec.ForProvider, valMap)
	DecodeCodebuildSourceCredential_UserName(&new.Spec.ForProvider, valMap)
	DecodeCodebuildSourceCredential_AuthType(&new.Spec.ForProvider, valMap)
	DecodeCodebuildSourceCredential_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeCodebuildSourceCredential_Id(p *CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeCodebuildSourceCredential_ServerType(p *CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	p.ServerType = ctwhy.ValueAsString(vals["server_type"])
}

func DecodeCodebuildSourceCredential_Token(p *CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	p.Token = ctwhy.ValueAsString(vals["token"])
}

func DecodeCodebuildSourceCredential_UserName(p *CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	p.UserName = ctwhy.ValueAsString(vals["user_name"])
}

func DecodeCodebuildSourceCredential_AuthType(p *CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	p.AuthType = ctwhy.ValueAsString(vals["auth_type"])
}

func DecodeCodebuildSourceCredential_Arn(p *CodebuildSourceCredentialObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}