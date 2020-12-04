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

func EncodeIamUserPolicy(r IamUserPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamUserPolicy_User(r.Spec.ForProvider, ctyVal)
	EncodeIamUserPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamUserPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamUserPolicy_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamUserPolicy_Policy(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeIamUserPolicy_User(p IamUserPolicyParameters, vals map[string]cty.Value) {
	vals["user"] = cty.StringVal(p.User)
}

func EncodeIamUserPolicy_Id(p IamUserPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamUserPolicy_Name(p IamUserPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamUserPolicy_NamePrefix(p IamUserPolicyParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamUserPolicy_Policy(p IamUserPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}