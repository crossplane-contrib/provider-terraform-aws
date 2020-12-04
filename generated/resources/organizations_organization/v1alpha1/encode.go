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

func EncodeOrganizationsOrganization(r OrganizationsOrganization) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeOrganizationsOrganization_FeatureSet(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganization_Id(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganization_AwsServiceAccessPrincipals(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganization_EnabledPolicyTypes(r.Spec.ForProvider, ctyVal)
	EncodeOrganizationsOrganization_MasterAccountId(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsOrganization_Arn(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsOrganization_MasterAccountArn(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsOrganization_MasterAccountEmail(r.Status.AtProvider, ctyVal)
	EncodeOrganizationsOrganization_NonMasterAccounts(r.Status.AtProvider.NonMasterAccounts, ctyVal)
	EncodeOrganizationsOrganization_Roots(r.Status.AtProvider.Roots, ctyVal)
	EncodeOrganizationsOrganization_Accounts(r.Status.AtProvider.Accounts, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeOrganizationsOrganization_FeatureSet(p OrganizationsOrganizationParameters, vals map[string]cty.Value) {
	vals["feature_set"] = cty.StringVal(p.FeatureSet)
}

func EncodeOrganizationsOrganization_Id(p OrganizationsOrganizationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganization_AwsServiceAccessPrincipals(p OrganizationsOrganizationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AwsServiceAccessPrincipals {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["aws_service_access_principals"] = cty.SetVal(colVals)
}

func EncodeOrganizationsOrganization_EnabledPolicyTypes(p OrganizationsOrganizationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.EnabledPolicyTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["enabled_policy_types"] = cty.SetVal(colVals)
}

func EncodeOrganizationsOrganization_MasterAccountId(p OrganizationsOrganizationObservation, vals map[string]cty.Value) {
	vals["master_account_id"] = cty.StringVal(p.MasterAccountId)
}

func EncodeOrganizationsOrganization_Arn(p OrganizationsOrganizationObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsOrganization_MasterAccountArn(p OrganizationsOrganizationObservation, vals map[string]cty.Value) {
	vals["master_account_arn"] = cty.StringVal(p.MasterAccountArn)
}

func EncodeOrganizationsOrganization_MasterAccountEmail(p OrganizationsOrganizationObservation, vals map[string]cty.Value) {
	vals["master_account_email"] = cty.StringVal(p.MasterAccountEmail)
}

func EncodeOrganizationsOrganization_NonMasterAccounts(p []NonMasterAccounts, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeOrganizationsOrganization_NonMasterAccounts_Id(v, ctyVal)
		EncodeOrganizationsOrganization_NonMasterAccounts_Name(v, ctyVal)
		EncodeOrganizationsOrganization_NonMasterAccounts_Status(v, ctyVal)
		EncodeOrganizationsOrganization_NonMasterAccounts_Arn(v, ctyVal)
		EncodeOrganizationsOrganization_NonMasterAccounts_Email(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["non_master_accounts"] = cty.ListVal(valsForCollection)
}

func EncodeOrganizationsOrganization_NonMasterAccounts_Id(p NonMasterAccounts, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganization_NonMasterAccounts_Name(p NonMasterAccounts, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsOrganization_NonMasterAccounts_Status(p NonMasterAccounts, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeOrganizationsOrganization_NonMasterAccounts_Arn(p NonMasterAccounts, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsOrganization_NonMasterAccounts_Email(p NonMasterAccounts, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeOrganizationsOrganization_Roots(p []Roots, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeOrganizationsOrganization_Roots_Name(v, ctyVal)
		EncodeOrganizationsOrganization_Roots_PolicyTypes(v.PolicyTypes, ctyVal)
		EncodeOrganizationsOrganization_Roots_Arn(v, ctyVal)
		EncodeOrganizationsOrganization_Roots_Id(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["roots"] = cty.ListVal(valsForCollection)
}

func EncodeOrganizationsOrganization_Roots_Name(p Roots, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsOrganization_Roots_PolicyTypes(p []PolicyTypes, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeOrganizationsOrganization_Roots_PolicyTypes_Type(v, ctyVal)
		EncodeOrganizationsOrganization_Roots_PolicyTypes_Status(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["policy_types"] = cty.ListVal(valsForCollection)
}

func EncodeOrganizationsOrganization_Roots_PolicyTypes_Type(p PolicyTypes, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeOrganizationsOrganization_Roots_PolicyTypes_Status(p PolicyTypes, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeOrganizationsOrganization_Roots_Arn(p Roots, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsOrganization_Roots_Id(p Roots, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganization_Accounts(p []Accounts, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeOrganizationsOrganization_Accounts_Arn(v, ctyVal)
		EncodeOrganizationsOrganization_Accounts_Email(v, ctyVal)
		EncodeOrganizationsOrganization_Accounts_Id(v, ctyVal)
		EncodeOrganizationsOrganization_Accounts_Name(v, ctyVal)
		EncodeOrganizationsOrganization_Accounts_Status(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["accounts"] = cty.ListVal(valsForCollection)
}

func EncodeOrganizationsOrganization_Accounts_Arn(p Accounts, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeOrganizationsOrganization_Accounts_Email(p Accounts, vals map[string]cty.Value) {
	vals["email"] = cty.StringVal(p.Email)
}

func EncodeOrganizationsOrganization_Accounts_Id(p Accounts, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeOrganizationsOrganization_Accounts_Name(p Accounts, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeOrganizationsOrganization_Accounts_Status(p Accounts, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}