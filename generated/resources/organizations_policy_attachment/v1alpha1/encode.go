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

func EncodeOrganizationsPolicyAttachment(r OrganizationsPolicyAttachment) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOrganizationsPolicyAttachment_Id(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicyAttachment_PolicyId(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsPolicyAttachment_TargetId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeOrganizationsPolicyAttachment_Id(p OrganizationsPolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsPolicyAttachment_PolicyId(p OrganizationsPolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["policy_id"] = cty.StringVal(p.PolicyId)
}

func EncodeOrganizationsPolicyAttachment_TargetId(p OrganizationsPolicyAttachmentParameters, vals map[string]cty.Value) {
	vals["target_id"] = cty.StringVal(p.TargetId)
}