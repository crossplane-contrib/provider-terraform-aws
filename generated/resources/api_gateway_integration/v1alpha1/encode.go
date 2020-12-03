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

package v1alpha1func EncodeApiGatewayIntegration(r ApiGatewayIntegration) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayIntegration_Credentials(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_Uri(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_IntegrationHttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_PassthroughBehavior(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_RequestTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_CacheKeyParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_RequestParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_TimeoutMilliseconds(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_Type(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_CacheNamespace(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ConnectionType(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ContentHandling(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_Id(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayIntegration_Credentials(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["credentials"] = cty.StringVal(p.Credentials)
}

func EncodeApiGatewayIntegration_RestApiId(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayIntegration_Uri(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeApiGatewayIntegration_ConnectionId(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeApiGatewayIntegration_IntegrationHttpMethod(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["integration_http_method"] = cty.StringVal(p.IntegrationHttpMethod)
}

func EncodeApiGatewayIntegration_PassthroughBehavior(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["passthrough_behavior"] = cty.StringVal(p.PassthroughBehavior)
}

func EncodeApiGatewayIntegration_RequestTemplates(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_templates"] = cty.MapVal(mVals)
}

func EncodeApiGatewayIntegration_CacheKeyParameters(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CacheKeyParameters {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cache_key_parameters"] = cty.SetVal(colVals)
}

func EncodeApiGatewayIntegration_HttpMethod(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeApiGatewayIntegration_RequestParameters(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestParameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_parameters"] = cty.MapVal(mVals)
}

func EncodeApiGatewayIntegration_ResourceId(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayIntegration_TimeoutMilliseconds(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["timeout_milliseconds"] = cty.IntVal(p.TimeoutMilliseconds)
}

func EncodeApiGatewayIntegration_Type(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeApiGatewayIntegration_CacheNamespace(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["cache_namespace"] = cty.StringVal(p.CacheNamespace)
}

func EncodeApiGatewayIntegration_ConnectionType(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["connection_type"] = cty.StringVal(p.ConnectionType)
}

func EncodeApiGatewayIntegration_ContentHandling(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["content_handling"] = cty.StringVal(p.ContentHandling)
}

func EncodeApiGatewayIntegration_Id(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}