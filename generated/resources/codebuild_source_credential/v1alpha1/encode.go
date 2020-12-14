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
	r, ok := mr.(*CodebuildSourceCredential)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CodebuildSourceCredential.")
	}
	return EncodeCodebuildSourceCredential(*r), nil
}

func EncodeCodebuildSourceCredential(r CodebuildSourceCredential) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodebuildSourceCredential_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_ServerType(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_Token(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_UserName(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_AuthType(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCodebuildSourceCredential_Id(p CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCodebuildSourceCredential_ServerType(p CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	vals["server_type"] = cty.StringVal(p.ServerType)
}

func EncodeCodebuildSourceCredential_Token(p CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	vals["token"] = cty.StringVal(p.Token)
}

func EncodeCodebuildSourceCredential_UserName(p CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	vals["user_name"] = cty.StringVal(p.UserName)
}

func EncodeCodebuildSourceCredential_AuthType(p CodebuildSourceCredentialParameters, vals map[string]cty.Value) {
	vals["auth_type"] = cty.StringVal(p.AuthType)
}

func EncodeCodebuildSourceCredential_Arn(p CodebuildSourceCredentialObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}