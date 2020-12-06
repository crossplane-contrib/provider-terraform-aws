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

func EncodeIamPolicyAttachment(r IamPolicyAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeIamPolicyAttachment_Groups(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicyAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicyAttachment_Name(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicyAttachment_PolicyArn(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicyAttachment_Roles(r.Spec.ForProvider, ctyVal)
	EncodeIamPolicyAttachment_Users(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeIamPolicyAttachment_Groups(p IamPolicyAttachmentParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Groups {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["groups"] = cty.SetVal(colVals)
}

func EncodeIamPolicyAttachment_Id(p IamPolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamPolicyAttachment_Name(p IamPolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeIamPolicyAttachment_PolicyArn(p IamPolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["policy_arn"] = cty.StringVal(p.PolicyArn)
}

func EncodeIamPolicyAttachment_Roles(p IamPolicyAttachmentParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Roles {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["roles"] = cty.SetVal(colVals)
}

func EncodeIamPolicyAttachment_Users(p IamPolicyAttachmentParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Users {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["users"] = cty.SetVal(colVals)
}