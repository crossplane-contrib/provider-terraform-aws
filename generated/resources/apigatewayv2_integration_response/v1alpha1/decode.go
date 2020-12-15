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
	r, ok := mr.(*Apigatewayv2IntegrationResponse)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApigatewayv2IntegrationResponse(r, ctyValue)
}

func DecodeApigatewayv2IntegrationResponse(prev *Apigatewayv2IntegrationResponse, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApigatewayv2IntegrationResponse_IntegrationId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2IntegrationResponse_IntegrationResponseKey(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2IntegrationResponse_ResponseTemplates(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2IntegrationResponse_TemplateSelectionExpression(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2IntegrationResponse_ApiId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2IntegrationResponse_ContentHandlingStrategy(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2IntegrationResponse_IntegrationId(p *Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	p.IntegrationId = ctwhy.ValueAsString(vals["integration_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2IntegrationResponse_IntegrationResponseKey(p *Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	p.IntegrationResponseKey = ctwhy.ValueAsString(vals["integration_response_key"])
}

//primitiveMapTypeDecodeTemplate
func DecodeApigatewayv2IntegrationResponse_ResponseTemplates(p *Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["response_templates"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ResponseTemplates = vMap
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2IntegrationResponse_TemplateSelectionExpression(p *Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	p.TemplateSelectionExpression = ctwhy.ValueAsString(vals["template_selection_expression"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2IntegrationResponse_ApiId(p *Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	p.ApiId = ctwhy.ValueAsString(vals["api_id"])
}

//primitiveTypeDecodeTemplate
func DecodeApigatewayv2IntegrationResponse_ContentHandlingStrategy(p *Apigatewayv2IntegrationResponseParameters, vals map[string]cty.Value) {
	p.ContentHandlingStrategy = ctwhy.ValueAsString(vals["content_handling_strategy"])
}