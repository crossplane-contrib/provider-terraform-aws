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
	r, ok := mr.(*ApiGatewayAuthorizer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayAuthorizer.")
	}
	return EncodeApiGatewayAuthorizer(*r), nil
}

func EncodeApiGatewayAuthorizer(r ApiGatewayAuthorizer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayAuthorizer_AuthorizerCredentials(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_IdentitySource(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_Name(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_AuthorizerUri(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_IdentityValidationExpression(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_ProviderArns(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAuthorizer_Type(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayAuthorizer_AuthorizerCredentials(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_credentials"] = cty.StringVal(p.AuthorizerCredentials)
}

func EncodeApiGatewayAuthorizer_Id(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayAuthorizer_IdentitySource(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["identity_source"] = cty.StringVal(p.IdentitySource)
}

func EncodeApiGatewayAuthorizer_Name(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayAuthorizer_RestApiId(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_result_ttl_in_seconds"] = cty.NumberIntVal(p.AuthorizerResultTtlInSeconds)
}

func EncodeApiGatewayAuthorizer_AuthorizerUri(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_uri"] = cty.StringVal(p.AuthorizerUri)
}

func EncodeApiGatewayAuthorizer_IdentityValidationExpression(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["identity_validation_expression"] = cty.StringVal(p.IdentityValidationExpression)
}

func EncodeApiGatewayAuthorizer_ProviderArns(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ProviderArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["provider_arns"] = cty.SetVal(colVals)
}

func EncodeApiGatewayAuthorizer_Type(p ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}