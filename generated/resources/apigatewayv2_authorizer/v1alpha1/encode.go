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
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
)

type ctyEncoder struct{}

func (e *ctyEncoder) EncodeCty(mr resource.Managed, schema *providers.Schema) (cty.Value, error) {
	r, ok := mr.(*Apigatewayv2Authorizer)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Apigatewayv2Authorizer.")
	}
	return EncodeApigatewayv2Authorizer(*r), nil
}

func EncodeApigatewayv2Authorizer(r Apigatewayv2Authorizer) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Authorizer_AuthorizerType(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_AuthorizerUri(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_Name(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_AuthorizerCredentialsArn(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_AuthorizerPayloadFormatVersion(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_AuthorizerResultTtlInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_EnableSimpleResponses(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_IdentitySources(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Authorizer_JwtConfiguration(r.Spec.ForProvider.JwtConfiguration, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2Authorizer_AuthorizerType(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_type"] = cty.StringVal(p.AuthorizerType)
}

func EncodeApigatewayv2Authorizer_AuthorizerUri(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_uri"] = cty.StringVal(p.AuthorizerUri)
}

func EncodeApigatewayv2Authorizer_Id(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2Authorizer_Name(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApigatewayv2Authorizer_ApiId(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2Authorizer_AuthorizerCredentialsArn(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_credentials_arn"] = cty.StringVal(p.AuthorizerCredentialsArn)
}

func EncodeApigatewayv2Authorizer_AuthorizerPayloadFormatVersion(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_payload_format_version"] = cty.StringVal(p.AuthorizerPayloadFormatVersion)
}

func EncodeApigatewayv2Authorizer_AuthorizerResultTtlInSeconds(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["authorizer_result_ttl_in_seconds"] = cty.NumberIntVal(p.AuthorizerResultTtlInSeconds)
}

func EncodeApigatewayv2Authorizer_EnableSimpleResponses(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	vals["enable_simple_responses"] = cty.BoolVal(p.EnableSimpleResponses)
}

func EncodeApigatewayv2Authorizer_IdentitySources(p Apigatewayv2AuthorizerParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.IdentitySources {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["identity_sources"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2Authorizer_JwtConfiguration(p JwtConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Authorizer_JwtConfiguration_Audience(p, ctyVal)
	EncodeApigatewayv2Authorizer_JwtConfiguration_Issuer(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["jwt_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeApigatewayv2Authorizer_JwtConfiguration_Audience(p JwtConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Audience {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["audience"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2Authorizer_JwtConfiguration_Issuer(p JwtConfiguration, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}