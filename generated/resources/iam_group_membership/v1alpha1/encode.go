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

package v1alpha1func EncodeIamGroupMembership(r IamGroupMembership) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIamGroupMembership_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamGroupMembership_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamGroupMembership_Users(r.Spec.ForProvider, ctyVal)
	EncodeIamGroupMembership_Group(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeIamGroupMembership_Id(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamGroupMembership_Name(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamGroupMembership_Users(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Users {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["users"] = cty.SetVal(colVals)
}

func EncodeIamGroupMembership_Group(p *IamGroupMembershipParameters, vals map[string]cty.Value) {
	vals["group"] = cty.StringVal(p.Group)
}