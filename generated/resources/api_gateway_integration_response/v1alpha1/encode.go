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

func EncodeApiGatewayIntegrationResponse(r ApiGatewayIntegrationResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayIntegrationResponse_ContentHandling(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_StatusCode(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_ResponseParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_ResponseTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_SelectionPattern(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayIntegrationResponse_ContentHandling(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["content_handling"] = cty.StringVal(p.ContentHandling)
}

func EncodeApiGatewayIntegrationResponse_ResourceId(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayIntegrationResponse_RestApiId(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayIntegrationResponse_StatusCode(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeApiGatewayIntegrationResponse_HttpMethod(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeApiGatewayIntegrationResponse_Id(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayIntegrationResponse_ResponseParameters(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseParameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_parameters"] = cty.MapVal(mVals)
}

func EncodeApiGatewayIntegrationResponse_ResponseTemplates(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_templates"] = cty.MapVal(mVals)
}

func EncodeApiGatewayIntegrationResponse_SelectionPattern(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["selection_pattern"] = cty.StringVal(p.SelectionPattern)
}