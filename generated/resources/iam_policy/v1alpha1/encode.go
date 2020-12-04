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

func EncodeIamPolicy(r IamPolicy) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamPolicy_Description(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicy_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicy_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicy_Path(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicy_Policy(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicy_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeIamPolicy_Description(p IamPolicyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeIamPolicy_Id(p IamPolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamPolicy_Name(p IamPolicyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamPolicy_NamePrefix(p IamPolicyParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeIamPolicy_Path(p IamPolicyParameters, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeIamPolicy_Policy(p IamPolicyParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeIamPolicy_Arn(p IamPolicyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}