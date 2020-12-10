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
	r, ok := mr.(*Apigatewayv2IntegrationResponse)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Apigatewayv2IntegrationResponse.")
	}
	return EncodeApigatewayv2IntegrationResponse(*r), nil
}

func EncodeApigatewayv2IntegrationResponse(r Apigatewayv2IntegrationResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2IntegrationResponse_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2IntegrationResponse_ContentHandlingStrategy(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2IntegrationResponse_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2IntegrationResponse_IntegrationId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2IntegrationResponse_IntegrationResponseKey(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2IntegrationResponse_ResponseTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2IntegrationResponse_TemplateSelectionExpression(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2IntegrationResponse_ApiId(p Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2IntegrationResponse_ContentHandlingStrategy(p Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	vals["content_handling_strategy"] = cty.StringVal(p.ContentHandlingStrategy)
}

func EncodeApigatewayv2IntegrationResponse_Id(p Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2IntegrationResponse_IntegrationId(p Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	vals["integration_id"] = cty.StringVal(p.IntegrationId)
}

func EncodeApigatewayv2IntegrationResponse_IntegrationResponseKey(p Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	vals["integration_response_key"] = cty.StringVal(p.IntegrationResponseKey)
}

func EncodeApigatewayv2IntegrationResponse_ResponseTemplates(p Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_templates"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2IntegrationResponse_TemplateSelectionExpression(p Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	vals["template_selection_expression"] = cty.StringVal(p.TemplateSelectionExpression)
}