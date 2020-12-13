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
	r, ok := mr.(*ApiGatewayIntegrationResponse)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayIntegrationResponse.")
	}
	return EncodeApiGatewayIntegrationResponse(*r), nil
}

func EncodeApiGatewayIntegrationResponse(r ApiGatewayIntegrationResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayIntegrationResponse_SelectionPattern(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_HttpMethod(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_ResponseParameters(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_ContentHandling(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_ResponseTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayIntegrationResponse_StatusCode(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayIntegrationResponse_SelectionPattern(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["selection_pattern"] = cty.StringVal(p.SelectionPattern)
}

func EncodeApiGatewayIntegrationResponse_HttpMethod(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["http_method"] = cty.StringVal(p.HttpMethod)
}

func EncodeApiGatewayIntegrationResponse_ResourceId(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeApiGatewayIntegrationResponse_ResponseParameters(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	if len(p.ResponseParameters) == 0 {
		vals["response_parameters"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseParameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_parameters"] = cty.MapVal(mVals)
}

func EncodeApiGatewayIntegrationResponse_RestApiId(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayIntegrationResponse_ContentHandling(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["content_handling"] = cty.StringVal(p.ContentHandling)
}

func EncodeApiGatewayIntegrationResponse_Id(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayIntegrationResponse_ResponseTemplates(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	if len(p.ResponseTemplates) == 0 {
		vals["response_templates"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_templates"] = cty.MapVal(mVals)
}

func EncodeApiGatewayIntegrationResponse_StatusCode(p ApiGatewayIntegrationResponseParameters, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}