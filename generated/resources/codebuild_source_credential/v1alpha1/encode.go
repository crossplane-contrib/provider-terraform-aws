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

func EncodeCodebuildSourceCredential(r CodebuildSourceCredential) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCodebuildSourceCredential_Id(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_ServerType(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_Token(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_UserName(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_AuthType(r.Spec.ForProvider, ctyVal)
	EncodeCodebuildSourceCredential_Arn(r.Status.AtProvider, ctyVal)
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