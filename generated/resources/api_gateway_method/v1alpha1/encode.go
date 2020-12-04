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

func EncodeApiGatewayMethod(r ApiGatewayMethod) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayMethod_RequestModels(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_RequestParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_Authorization(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_AuthorizerId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_RequestValidatorId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_ApiKeyRequired(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethod_AuthorizationScopes(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayMethod_RequestModels(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestModels {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_models"] = cty.MapVal(mVals)
}

func EncodeApiGatewayMethod_RequestParameters(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestParameters {
		mVals[key] = cty.BoolVal(value)
	}
	vals["request_parameters"] = cty.MapVal(mVals)
}

func EncodeApiGatewayMethod_ResourceId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayMethod_RestApiId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayMethod_Authorization(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["authorization"] = cty.StringVal(p.Authorization)
}

func EncodeApiGatewayMethod_AuthorizerId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["authorizer_id"] = cty.StringVal(p.AuthorizerId)
}

func EncodeApiGatewayMethod_HttpMethod(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeApiGatewayMethod_Id(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayMethod_RequestValidatorId(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["request_validator_id"] = cty.StringVal(p.RequestValidatorId)
}

func EncodeApiGatewayMethod_ApiKeyRequired(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	vals["api_key_required"] = cty.BoolVal(p.ApiKeyRequired)
}

func EncodeApiGatewayMethod_AuthorizationScopes(p ApiGatewayMethodParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AuthorizationScopes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["authorization_scopes"] = cty.SetVal(colVals)
}