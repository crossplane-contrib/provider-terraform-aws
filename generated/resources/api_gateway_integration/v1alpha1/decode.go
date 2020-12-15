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
	r, ok := mr.(*ApiGatewayIntegration)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApiGatewayIntegration(r, ctyValue)
}

func DecodeApiGatewayIntegration(prev *ApiGatewayIntegration, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApiGatewayIntegration_ConnectionId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_Credentials(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_RequestTemplates(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_CacheKeyParameters(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_ContentHandling(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_Id(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_IntegrationHttpMethod(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_PassthroughBehavior(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_RequestParameters(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_TimeoutMilliseconds(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_Type(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_ConnectionType(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_RestApiId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_Uri(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_ResourceId(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_HttpMethod(&new.Spec.ForProvider, valMap)
	DecodeApiGatewayIntegration_CacheNamespace(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_ConnectionId(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.ConnectionId = ctwhy.ValueAsString(vals["connection_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_Credentials(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.Credentials = ctwhy.ValueAsString(vals["credentials"])
}

//primitiveMapTypeDecodeTemplate
func DecodeApiGatewayIntegration_RequestTemplates(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["request_templates"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.RequestTemplates = vMap
}

//primitiveCollectionTypeDecodeTemplate
func DecodeApiGatewayIntegration_CacheKeyParameters(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["cache_key_parameters"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.CacheKeyParameters = goVals
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_ContentHandling(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.ContentHandling = ctwhy.ValueAsString(vals["content_handling"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_Id(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_IntegrationHttpMethod(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.IntegrationHttpMethod = ctwhy.ValueAsString(vals["integration_http_method"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_PassthroughBehavior(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.PassthroughBehavior = ctwhy.ValueAsString(vals["passthrough_behavior"])
}

//primitiveMapTypeDecodeTemplate
func DecodeApiGatewayIntegration_RequestParameters(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["request_parameters"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.RequestParameters = vMap
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_TimeoutMilliseconds(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.TimeoutMilliseconds = ctwhy.ValueAsInt64(vals["timeout_milliseconds"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_Type(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.Type = ctwhy.ValueAsString(vals["type"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_ConnectionType(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.ConnectionType = ctwhy.ValueAsString(vals["connection_type"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_RestApiId(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.RestApiId = ctwhy.ValueAsString(vals["rest_api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_Uri(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.Uri = ctwhy.ValueAsString(vals["uri"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_ResourceId(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.ResourceId = ctwhy.ValueAsString(vals["resource_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_HttpMethod(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.HttpMethod = ctwhy.ValueAsString(vals["http_method"])
}

//primitiveTypeDecodeTemplate
func DecodeApiGatewayIntegration_CacheNamespace(p *ApiGatewayIntegrationParameters, vals map[string]cty.Value) {
	p.CacheNamespace = ctwhy.ValueAsString(vals["cache_namespace"])
}