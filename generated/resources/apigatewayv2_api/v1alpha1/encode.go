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

func EncodeApigatewayv2Api(r Apigatewayv2Api) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Api_Description(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_RouteSelectionExpression(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_Body(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_CredentialsArn(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_Version(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_ApiKeySelectionExpression(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_Name(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_Target(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_DisableExecuteApiEndpoint(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_ProtocolType(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_RouteKey(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Api_CorsConfiguration(r.Spec.ForProvider.CorsConfiguration, ctyVal)
	EncodeApigatewayv2Api_ExecutionArn(r.Status.AtProvider, ctyVal)
	EncodeApigatewayv2Api_ApiEndpoint(r.Status.AtProvider, ctyVal)
	EncodeApigatewayv2Api_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2Api_Description(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApigatewayv2Api_RouteSelectionExpression(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["route_selection_expression"] = cty.StringVal(p.RouteSelectionExpression)
}

func EncodeApigatewayv2Api_Body(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["body"] = cty.StringVal(p.Body)
}

func EncodeApigatewayv2Api_CredentialsArn(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["credentials_arn"] = cty.StringVal(p.CredentialsArn)
}

func EncodeApigatewayv2Api_Version(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeApigatewayv2Api_ApiKeySelectionExpression(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["api_key_selection_expression"] = cty.StringVal(p.ApiKeySelectionExpression)
}

func EncodeApigatewayv2Api_Name(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApigatewayv2Api_Target(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["target"] = cty.StringVal(p.Target)
}

func EncodeApigatewayv2Api_Tags(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2Api_DisableExecuteApiEndpoint(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["disable_execute_api_endpoint"] = cty.BoolVal(p.DisableExecuteApiEndpoint)
}

func EncodeApigatewayv2Api_Id(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2Api_ProtocolType(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["protocol_type"] = cty.StringVal(p.ProtocolType)
}

func EncodeApigatewayv2Api_RouteKey(p Apigatewayv2ApiParameters, vals map[string]cty.Value) {
	vals["route_key"] = cty.StringVal(p.RouteKey)
}

func EncodeApigatewayv2Api_CorsConfiguration(p CorsConfiguration, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Api_CorsConfiguration_AllowCredentials(p, ctyVal)
	EncodeApigatewayv2Api_CorsConfiguration_AllowHeaders(p, ctyVal)
	EncodeApigatewayv2Api_CorsConfiguration_AllowMethods(p, ctyVal)
	EncodeApigatewayv2Api_CorsConfiguration_AllowOrigins(p, ctyVal)
	EncodeApigatewayv2Api_CorsConfiguration_ExposeHeaders(p, ctyVal)
	EncodeApigatewayv2Api_CorsConfiguration_MaxAge(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["cors_configuration"] = cty.ListVal(valsForCollection)
}

func EncodeApigatewayv2Api_CorsConfiguration_AllowCredentials(p CorsConfiguration, vals map[string]cty.Value) {
	vals["allow_credentials"] = cty.BoolVal(p.AllowCredentials)
}

func EncodeApigatewayv2Api_CorsConfiguration_AllowHeaders(p CorsConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowHeaders {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allow_headers"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2Api_CorsConfiguration_AllowMethods(p CorsConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowMethods {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allow_methods"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2Api_CorsConfiguration_AllowOrigins(p CorsConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.AllowOrigins {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["allow_origins"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2Api_CorsConfiguration_ExposeHeaders(p CorsConfiguration, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ExposeHeaders {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["expose_headers"] = cty.SetVal(colVals)
}

func EncodeApigatewayv2Api_CorsConfiguration_MaxAge(p CorsConfiguration, vals map[string]cty.Value) {
	vals["max_age"] = cty.NumberIntVal(p.MaxAge)
}

func EncodeApigatewayv2Api_ExecutionArn(p Apigatewayv2ApiObservation, vals map[string]cty.Value) {
	vals["execution_arn"] = cty.StringVal(p.ExecutionArn)
}

func EncodeApigatewayv2Api_ApiEndpoint(p Apigatewayv2ApiObservation, vals map[string]cty.Value) {
	vals["api_endpoint"] = cty.StringVal(p.ApiEndpoint)
}

func EncodeApigatewayv2Api_Arn(p Apigatewayv2ApiObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}