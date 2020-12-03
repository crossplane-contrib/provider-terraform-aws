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

package v1alpha1func EncodeIamRolePolicy(r IamRolePolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIamRolePolicy_Role(r.Spec.ForProvider, ctyVal)
	EncodeIamRolePolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamRolePolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamRolePolicy_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamRolePolicy_Policy(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeIamRolePolicy_Role(p *IamRolePolicyParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}

func EncodeIamRolePolicy_Id(p *IamRolePolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamRolePolicy_Name(p *IamRolePolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamRolePolicy_NamePrefix(p *IamRolePolicyParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamRolePolicy_Policy(p *IamRolePolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}