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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*ApiGatewayAuthorizer)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayAuthorizer(r, ctyValue)
}

func DecodeApiGatewayAuthorizer(prev *ApiGatewayAuthorizer, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayAuthorizer_Name(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_Type(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_IdentityValidationExpression(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_ProviderArns(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_RestApiId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_AuthorizerCredentials(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_AuthorizerUri(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayAuthorizer_IdentitySource(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_Name(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_Type(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_IdentityValidationExpression(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.IdentityValidationExpression = ctwhy.ValueAsString(vals["identity_validation_expression"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_ProviderArns(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["provider_arns"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.ProviderArns = goVals
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_RestApiId(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.RestApiId = ctwhy.ValueAsString(vals["rest_api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_AuthorizerCredentials(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.AuthorizerCredentials = ctwhy.ValueAsString(vals["authorizer_credentials"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_AuthorizerResultTtlInSeconds(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.AuthorizerResultTtlInSeconds = ctwhy.ValueAsInt64(vals["authorizer_result_ttl_in_seconds"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_AuthorizerUri(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.AuthorizerUri = ctwhy.ValueAsString(vals["authorizer_uri"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayAuthorizer_IdentitySource(p *ApiGatewayAuthorizerParameters, vals map[string]cty.Value) {
	p.IdentitySource = ctwhy.ValueAsString(vals["identity_source"])
}