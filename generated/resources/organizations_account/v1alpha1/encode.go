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

func EncodeOrganizationsAccount(r OrganizationsAccount) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOrganizationsAccount_Name(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_ParentId(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_RoleName(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_Email(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_IamUserAccessToBilling(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_Id(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_Tags(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsAccount_Arn(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsAccount_JoinedMethod(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsAccount_Status(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsAccount_JoinedTimestamp(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOrganizationsAccount_Name(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsAccount_ParentId(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["parent_id"] = cty.StringVal(p.ParentId)
}

func EncodeOrganizationsAccount_RoleName(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["role_name"] = cty.StringVal(p.RoleName)
}

func EncodeOrganizationsAccount_Email(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeOrganizationsAccount_IamUserAccessToBilling(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["iam_user_access_to_billing"] = cty.StringVal(p.IamUserAccessToBilling)
}

func EncodeOrganizationsAccount_Id(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsAccount_Tags(p OrganizationsAccountParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeOrganizationsAccount_Arn(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsAccount_JoinedMethod(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["joined_method"] = cty.StringVal(p.JoinedMethod)
}

func EncodeOrganizationsAccount_Status(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeOrganizationsAccount_JoinedTimestamp(p OrganizationsAccountObservation, vals map[string]cty.Value) {
	vals["joined_timestamp"] = cty.StringVal(p.JoinedTimestamp)
}