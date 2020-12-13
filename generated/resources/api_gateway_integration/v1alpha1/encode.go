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
	r, ok := mr.(*ApiGatewayIntegration)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayIntegration.")
	}
	return EncodeApiGatewayIntegration(*r), nil
}

func EncodeApiGatewayIntegration(r ApiGatewayIntegration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayIntegration_PassthroughBehavior(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_Type(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_Uri(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_IntegrationHttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_TimeoutMilliseconds(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_CacheNamespace(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ConnectionType(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ContentHandling(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_CacheKeyParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_RequestTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_Credentials(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegration_RequestParameters(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayIntegration_PassthroughBehavior(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["passthrough_behavior"] = cty.StringVal(p.PassthroughBehavior)
}

func EncodeApiGatewayIntegration_Type(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeApiGatewayIntegration_Uri(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeApiGatewayIntegration_IntegrationHttpMethod(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["integration_http_method"] = cty.StringVal(p.IntegrationHttpMethod)
}

func EncodeApiGatewayIntegration_TimeoutMilliseconds(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["timeout_milliseconds"] = cty.NumberIntVal(p.TimeoutMilliseconds)
}

func EncodeApiGatewayIntegration_CacheNamespace(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["cache_namespace"] = cty.StringVal(p.CacheNamespace)
}

func EncodeApiGatewayIntegration_ConnectionType(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["connection_type"] = cty.StringVal(p.ConnectionType)
}

func EncodeApiGatewayIntegration_ContentHandling(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["content_handling"] = cty.StringVal(p.ContentHandling)
}

func EncodeApiGatewayIntegration_HttpMethod(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeApiGatewayIntegration_Id(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayIntegration_CacheKeyParameters(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.CacheKeyParameters {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["cache_key_parameters"] = cty.SetVal(colVals)
}

func EncodeApiGatewayIntegration_RequestTemplates(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	if len(p.RequestTemplates) == 0 {
		vals["request_templates"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_templates"] = cty.MapVal(mVals)
}

func EncodeApiGatewayIntegration_ResourceId(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayIntegration_RestApiId(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayIntegration_ConnectionId(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeApiGatewayIntegration_Credentials(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	vals["credentials"] = cty.StringVal(p.Credentials)
}

func EncodeApiGatewayIntegration_RequestParameters(p ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	if len(p.RequestParameters) == 0 {
		vals["request_parameters"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestParameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_parameters"] = cty.MapVal(mVals)
}