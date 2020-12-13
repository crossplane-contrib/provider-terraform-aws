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
	"fmt"

	"github.com/zclconf/go-cty/cty"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*CognitoUserPoolClient)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CognitoUserPoolClient.")
	}
	return EncodeCognitoUserPoolClient(*r), nil
}

func EncodeCognitoUserPoolClient(r CognitoUserPoolClient) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCognitoUserPoolClient_DefaultRedirectUri(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_RefreshTokenValidity(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_UserPoolId(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_WriteAttributes(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_AllowedOauthScopes(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_ExplicitAuthFlows(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_GenerateSecret(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_PreventUserExistenceErrors(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_ReadAttributes(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_AllowedOauthFlows(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_SupportedIdentityProviders(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_AllowedOauthFlowsUserPoolClient(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_CallbackUrls(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_LogoutUrls(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_Name(r.Spec.ForProvider, ctyVal)
	EncodeCognitoUserPoolClient_AnalyticsConfiguration(r.Spec.ForProvider.AnalyticsConfiguration, ctyVal)
	EncodeCognitoUserPoolClient_ClientSecret(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCognitoUserPoolClient_DefaultRedirectUri(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["default_redirect_uri"] = cty.StringVal(p.DefaultRedirectUri)
}

func EncodeCognitoUserPoolClient_Id(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoUserPoolClient_RefreshTokenValidity(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["refresh_token_validity"] = cty.NumberIntVal(p.RefreshTokenValidity)
}

func EncodeCognitoUserPoolClient_UserPoolId(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeCognitoUserPoolClient_WriteAttributes(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.WriteAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["write_attributes"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_AllowedOauthScopes(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedOauthScopes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_oauth_scopes"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_ExplicitAuthFlows(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ExplicitAuthFlows {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["explicit_auth_flows"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_GenerateSecret(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["generate_secret"] = cty.BoolVal(p.GenerateSecret)
}

func EncodeCognitoUserPoolClient_PreventUserExistenceErrors(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["prevent_user_existence_errors"] = cty.StringVal(p.PreventUserExistenceErrors)
}

func EncodeCognitoUserPoolClient_ReadAttributes(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ReadAttributes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["read_attributes"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_AllowedOauthFlows(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowedOauthFlows {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allowed_oauth_flows"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_SupportedIdentityProviders(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SupportedIdentityProviders {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["supported_identity_providers"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_AllowedOauthFlowsUserPoolClient(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["allowed_oauth_flows_user_pool_client"] = cty.BoolVal(p.AllowedOauthFlowsUserPoolClient)
}

func EncodeCognitoUserPoolClient_CallbackUrls(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CallbackUrls {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["callback_urls"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_LogoutUrls(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.LogoutUrls {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["logout_urls"] = cty.SetVal(colVals)
}

func EncodeCognitoUserPoolClient_Name(p CognitoUserPoolClientParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCognitoUserPoolClient_AnalyticsConfiguration(p AnalyticsConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeCognitoUserPoolClient_AnalyticsConfiguration_ApplicationId(p, ctyVal)
	EncodeCognitoUserPoolClient_AnalyticsConfiguration_ExternalId(p, ctyVal)
	EncodeCognitoUserPoolClient_AnalyticsConfiguration_RoleArn(p, ctyVal)
	EncodeCognitoUserPoolClient_AnalyticsConfiguration_UserDataShared(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["analytics_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeCognitoUserPoolClient_AnalyticsConfiguration_ApplicationId(p AnalyticsConfiguration, vals map[string]cty.Value) {
	vals["application_id"] = cty.StringVal(p.ApplicationId)
}

func EncodeCognitoUserPoolClient_AnalyticsConfiguration_ExternalId(p AnalyticsConfiguration, vals map[string]cty.Value) {
	vals["external_id"] = cty.StringVal(p.ExternalId)
}

func EncodeCognitoUserPoolClient_AnalyticsConfiguration_RoleArn(p AnalyticsConfiguration, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCognitoUserPoolClient_AnalyticsConfiguration_UserDataShared(p AnalyticsConfiguration, vals map[string]cty.Value) {
	vals["user_data_shared"] = cty.BoolVal(p.UserDataShared)
}

func EncodeCognitoUserPoolClient_ClientSecret(p CognitoUserPoolClientObservation, vals map[string]cty.Value) {
	vals["client_secret"] = cty.StringVal(p.ClientSecret)
}