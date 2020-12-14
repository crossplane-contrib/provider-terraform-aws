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
	r, ok := mr.(*Apigatewayv2RouteResponse)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeApigatewayv2RouteResponse(r, ctyValue)
}

func DecodeApigatewayv2RouteResponse(prev *Apigatewayv2RouteResponse, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeApigatewayv2RouteResponse_ApiId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2RouteResponse_Id(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2RouteResponse_ModelSelectionExpression(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2RouteResponse_ResponseModels(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2RouteResponse_RouteId(&new.Spec.ForProvider, valMap)
	DecodeApigatewayv2RouteResponse_RouteResponseKey(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeApigatewayv2RouteResponse_ApiId(p *Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	p.ApiId = ctwhy.ValueAsString(vals["api_id"])
}

func DecodeApigatewayv2RouteResponse_Id(p *Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeApigatewayv2RouteResponse_ModelSelectionExpression(p *Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	p.ModelSelectionExpression = ctwhy.ValueAsString(vals["model_selection_expression"])
}

func DecodeApigatewayv2RouteResponse_ResponseModels(p *Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["response_models"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.ResponseModels = vMap
}

func DecodeApigatewayv2RouteResponse_RouteId(p *Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	p.RouteId = ctwhy.ValueAsString(vals["route_id"])
}

func DecodeApigatewayv2RouteResponse_RouteResponseKey(p *Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	p.RouteResponseKey = ctwhy.ValueAsString(vals["route_response_key"])
}