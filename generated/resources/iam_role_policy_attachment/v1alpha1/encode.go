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

package v1alpha1func EncodeIamRolePolicyAttachment(r IamRolePolicyAttachment) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeIamRolePolicyAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeIamRolePolicyAttachment_PolicyArn(r.Spec.ForProvider, ctyVal)
	EncodeIamRolePolicyAttachment_Role(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeIamRolePolicyAttachment_Id(p *IamRolePolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeIamRolePolicyAttachment_PolicyArn(p *IamRolePolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["policy_arn"] = cty.StringVal(p.PolicyArn)
}

func EncodeIamRolePolicyAttachment_Role(p *IamRolePolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["role"] = cty.StringVal(p.Role)
}