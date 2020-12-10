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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*CognitoUserGroup)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CognitoUserGroup.")
	}
	return EncodeCognitoUserGroup(*r), nil
}

func EncodeCognitoUserGroup(r CognitoUserGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCognitoUserGroup_Description(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_Precedence(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_UserPoolId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeCognitoUserGroup_Description(p CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCognitoUserGroup_Id(p CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoUserGroup_Name(p CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCognitoUserGroup_Precedence(p CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["precedence"] = cty.NumberIntVal(p.Precedence)
}

func EncodeCognitoUserGroup_RoleArn(p CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCognitoUserGroup_UserPoolId(p CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}