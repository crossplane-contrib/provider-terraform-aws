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

package v1alpha1func EncodeApigatewayv2Route(r Apigatewayv2Route) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApigatewayv2Route_OperationName(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_RequestModels(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_Target(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_ApiKeyRequired(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_AuthorizationScopes(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_AuthorizerId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_ModelSelectionExpression(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_AuthorizationType(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_RouteKey(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Route_RouteResponseSelectionExpression(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeApigatewayv2Route_OperationName(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["operation_name"] = cty.StringVal(p.OperationName)
}

func EncodeApigatewayv2Route_RequestModels(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestModels {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_models"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2Route_Target(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["target"] = cty.StringVal(p.Target)
}

func EncodeApigatewayv2Route_ApiId(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2Route_ApiKeyRequired(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["api_key_required"] = cty.BoolVal(p.ApiKeyRequired)
}

func EncodeApigatewayv2Route_AuthorizationScopes(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AuthorizationScopes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["authorization_scopes"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2Route_AuthorizerId(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["authorizer_id"] = cty.StringVal(p.AuthorizerId)
}

func EncodeApigatewayv2Route_ModelSelectionExpression(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["model_selection_expression"] = cty.StringVal(p.ModelSelectionExpression)
}

func EncodeApigatewayv2Route_AuthorizationType(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["authorization_type"] = cty.StringVal(p.AuthorizationType)
}

func EncodeApigatewayv2Route_Id(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2Route_RouteKey(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["route_key"] = cty.StringVal(p.RouteKey)
}

func EncodeApigatewayv2Route_RouteResponseSelectionExpression(p *Apigatewayv2RouteParameters, vals map[string]cty.Value) {
	vals["route_response_selection_expression"] = cty.StringVal(p.RouteResponseSelectionExpression)
}