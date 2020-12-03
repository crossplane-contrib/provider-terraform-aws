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

package v1alpha1func EncodeApiGatewayAuthorizer(r ApiGatewayAuthorizer) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayAuthorizer_AuthorizerCredentials(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_AuthorizerUri(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_IdentityValidationExpression(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_Name(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_IdentitySource(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_ProviderArns(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_Type(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayAuthorizer_AuthorizerCredentials(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_credentials"] = cty.StringVal(p.AuthorizerCredentials)
}

func EncodeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_result_ttl_in_seconds"] = cty.IntVal(p.AuthorizerResultTtlInSeconds)
}

func EncodeApiGatewayAuthorizer_AuthorizerUri(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_uri"] = cty.StringVal(p.AuthorizerUri)
}

func EncodeApiGatewayAuthorizer_Id(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayAuthorizer_IdentityValidationExpression(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["identity_validation_expression"] = cty.StringVal(p.IdentityValidationExpression)
}

func EncodeApiGatewayAuthorizer_Name(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayAuthorizer_RestApiId(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayAuthorizer_IdentitySource(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["identity_source"] = cty.StringVal(p.IdentitySource)
}

func EncodeApiGatewayAuthorizer_ProviderArns(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ProviderArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["provider_arns"] = cty.SetVal(colVals)
}

func EncodeApiGatewayAuthorizer_Type(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}