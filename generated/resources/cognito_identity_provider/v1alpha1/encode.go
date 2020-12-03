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

package v1alpha1func EncodeCognitoIdentityProvider(r CognitoIdentityProvider) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeCognitoIdentityProvider_AttributeMapping(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityProvider_Id(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityProvider_IdpIdentifiers(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityProvider_ProviderDetails(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityProvider_ProviderName(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityProvider_ProviderType(r.Spec.ForProvider, ctyVal)
	EncodeCognitoIdentityProvider_UserPoolId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeCognitoIdentityProvider_AttributeMapping(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.AttributeMapping {
		mVals[key] = cty.StringVal(value)
	}
	vals["attribute_mapping"] = cty.MapVal(mVals)
}

func EncodeCognitoIdentityProvider_Id(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCognitoIdentityProvider_IdpIdentifiers(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IdpIdentifiers {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["idp_identifiers"] = cty.ListVal(colVals)
}

func EncodeCognitoIdentityProvider_ProviderDetails(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ProviderDetails {
		mVals[key] = cty.StringVal(value)
	}
	vals["provider_details"] = cty.MapVal(mVals)
}

func EncodeCognitoIdentityProvider_ProviderName(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	vals["provider_name"] = cty.StringVal(p.ProviderName)
}

func EncodeCognitoIdentityProvider_ProviderType(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	vals["provider_type"] = cty.StringVal(p.ProviderType)
}

func EncodeCognitoIdentityProvider_UserPoolId(p *CognitoIdentityProviderParameters, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}