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
	r, ok := mr.(*CognitoIdentityPool)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CognitoIdentityPool.")
	}
	return EncodeCognitoIdentityPool(*r), nil
}

func EncodeCognitoIdentityPool(r CognitoIdentityPool) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCognitoIdentityPool_AllowUnauthenticatedIdentities(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_DeveloperProviderName(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_OpenidConnectProviderArns(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_SamlProviderArns(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_SupportedLoginProviders(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_IdentityPoolName(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityPool_CognitoIdentityProviders(r.Spec.ForProvider.CognitoIdentityProviders, ctyVal)
	EncodeCognitoIdentityPool_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCognitoIdentityPool_AllowUnauthenticatedIdentities(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	vals["allow_unauthenticated_identities"] = cty.BoolVal(p.AllowUnauthenticatedIdentities)
}

func EncodeCognitoIdentityPool_DeveloperProviderName(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	vals["developer_provider_name"] = cty.StringVal(p.DeveloperProviderName)
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

func EncodeCognitoIdentityPool_SupportedLoginProviders(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	if len(p.SupportedLoginProviders) == 0 {
		vals["supported_login_providers"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.SupportedLoginProviders {
		mVals[key] = cty.StringVal(value)
	}
	vals["supported_login_providers"] = cty.MapVal(mVals)
}

func EncodeCognitoIdentityPool_Tags(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCognitoIdentityPool_Id(p CognitoIdentityPoolParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
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