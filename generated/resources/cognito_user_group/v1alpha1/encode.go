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

package v1alpha1func EncodeCognitoUserGroup(r CognitoUserGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCognitoUserGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_Precedence(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_UserPoolId(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserGroup_Description(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeCognitoUserGroup_Id(p *CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoUserGroup_Name(p *CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCognitoUserGroup_Precedence(p *CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["precedence"] = cty.IntVal(p.Precedence)
}

func EncodeCognitoUserGroup_RoleArn(p *CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCognitoUserGroup_UserPoolId(p *CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeCognitoUserGroup_Description(p *CognitoUserGroupParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}