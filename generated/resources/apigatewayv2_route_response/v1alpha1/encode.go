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

func EncodeApigatewayv2RouteResponse(r Apigatewayv2RouteResponse) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2RouteResponse_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_ModelSelectionExpression(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_ResponseModels(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_RouteId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_RouteResponseKey(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2RouteResponse_ApiId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2RouteResponse_Id(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2RouteResponse_ModelSelectionExpression(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	vals["model_selection_expression"] = cty.StringVal(p.ModelSelectionExpression)
}

func EncodeApigatewayv2RouteResponse_ResponseModels(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
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

func EncodeApigatewayv2RouteResponse_ApiId(p Apigatewayv2RouteResponseParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}