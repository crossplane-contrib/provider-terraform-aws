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
	r, ok := mr.(*Apigatewayv2RouteResponse)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a Apigatewayv2RouteResponse.")
	}
	return EncodeApigatewayv2RouteResponse(*r), nil
}

func EncodeApigatewayv2RouteResponse(r Apigatewayv2RouteResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2RouteResponse_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_ModelSelectionExpression(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_ResponseModels(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_RouteId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_RouteResponseKey(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	ctyVal["id"] = cty.StringVal(en)
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2RouteResponse_ApiId(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2RouteResponse_ModelSelectionExpression(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	vals["model_selection_expression"] = cty.StringVal(p.ModelSelectionExpression)
}

func EncodeApigatewayv2RouteResponse_ResponseModels(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	if len(p.ResponseModels) == 0 {
		vals["response_models"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.ResponseModels {
		mVals[key] = cty.StringVal(value)
	}
	vals["response_models"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2RouteResponse_RouteId(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	vals["route_id"] = cty.StringVal(p.RouteId)
}

func EncodeApigatewayv2RouteResponse_RouteResponseKey(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	vals["route_response_key"] = cty.StringVal(p.RouteResponseKey)
}