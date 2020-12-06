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

func EncodeCognitoIdentityPool(r CognitoIdentityPool) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCognitoIdentityPool_SupportedLoginProviders(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_AllowUnauthenticatedIdentities(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_OpenidConnectProviderArns(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_SamlProviderArns(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_DeveloperProviderName(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_IdentityPoolName(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_CognitoIdentityProviders(r.Spec.ForProvider.CognitoIdentityProviders, ctyVal)
	EncodeCognitoIdentityPool_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCognitoIdentityPool_SupportedLoginProviders(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.SupportedLoginProviders {
		mVals[key] = cty.StringVal(value)
	}
	vals["supported_login_providers"] = cty.MapVal(mVals)
}

func EncodeCognitoIdentityPool_AllowUnauthenticatedIdentities(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	vals["allow_unauthenticated_identities"] = cty.BoolVal(p.AllowUnauthenticatedIdentities)
}

func EncodeCognitoIdentityPool_Id(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoIdentityPool_OpenidConnectProviderArns(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.OpenidConnectProviderArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["openid_connect_provider_arns"] = cty.SetVal(colVals)
}

func EncodeCognitoIdentityPool_SamlProviderArns(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SamlProviderArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["saml_provider_arns"] = cty.ListVal(colVals)
}

func EncodeCognitoIdentityPool_Tags(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCognitoIdentityPool_DeveloperProviderName(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	vals["developer_provider_name"] = cty.StringVal(p.DeveloperProviderName)
}

func EncodeCognitoIdentityPool_IdentityPoolName(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	vals["identity_pool_name"] = cty.StringVal(p.IdentityPoolName)
}

func EncodeCognitoIdentityPool_CognitoIdentityProviders(p CognitoIdentityProviders, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCognitoIdentityPool_CognitoIdentityProviders_ClientId(p, ctyVal)
	EncodeCognitoIdentityPool_CognitoIdentityProviders_ProviderName(p, ctyVal)
	EncodeCognitoIdentityPool_CognitoIdentityProviders_ServerSideTokenCheck(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cognito_identity_providers"] = cty.SetVal(valsForCollection)
}

func EncodeCognitoIdentityPool_CognitoIdentityProviders_ClientId(p CognitoIdentityProviders, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeCognitoIdentityPool_CognitoIdentityProviders_ProviderName(p CognitoIdentityProviders, vals map[string]cty.Value) {
	vals["provider_name"] = cty.StringVal(p.ProviderName)
}

func EncodeCognitoIdentityPool_CognitoIdentityProviders_ServerSideTokenCheck(p CognitoIdentityProviders, vals map[string]cty.Value) {
	vals["server_side_token_check"] = cty.BoolVal(p.ServerSideTokenCheck)
}

func EncodeCognitoIdentityPool_Arn(p CognitoIdentityPoolObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}