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

func EncodeApigatewayv2Integration(r Apigatewayv2Integration) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Integration_PayloadFormatVersion(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_TemplateSelectionExpression(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_TimeoutMilliseconds(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_ConnectionType(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_Description(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_IntegrationType(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_IntegrationSubtype(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_PassthroughBehavior(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_RequestTemplates(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_ConnectionId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_ContentHandlingStrategy(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_CredentialsArn(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_IntegrationMethod(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_IntegrationUri(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_RequestParameters(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Integration_TlsConfig(r.Spec.ForProvider.TlsConfig, ctyVal)
	EncodeApigatewayv2Integration_IntegrationResponseSelectionExpression(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2Integration_PayloadFormatVersion(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["payload_format_version"] = cty.StringVal(p.PayloadFormatVersion)
}

func EncodeApigatewayv2Integration_TemplateSelectionExpression(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["template_selection_expression"] = cty.StringVal(p.TemplateSelectionExpression)
}

func EncodeApigatewayv2Integration_TimeoutMilliseconds(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["timeout_milliseconds"] = cty.NumberIntVal(p.TimeoutMilliseconds)
}

func EncodeApigatewayv2Integration_ApiId(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2Integration_ConnectionType(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["connection_type"] = cty.StringVal(p.ConnectionType)
}

func EncodeApigatewayv2Integration_Description(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApigatewayv2Integration_IntegrationType(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["integration_type"] = cty.StringVal(p.IntegrationType)
}

func EncodeApigatewayv2Integration_IntegrationSubtype(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["integration_subtype"] = cty.StringVal(p.IntegrationSubtype)
}

func EncodeApigatewayv2Integration_PassthroughBehavior(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["passthrough_behavior"] = cty.StringVal(p.PassthroughBehavior)
}

func EncodeApigatewayv2Integration_RequestTemplates(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestTemplates {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_templates"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2Integration_ConnectionId(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["connection_id"] = cty.StringVal(p.ConnectionId)
}

func EncodeApigatewayv2Integration_ContentHandlingStrategy(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["content_handling_strategy"] = cty.StringVal(p.ContentHandlingStrategy)
}

func EncodeApigatewayv2Integration_CredentialsArn(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["credentials_arn"] = cty.StringVal(p.CredentialsArn)
}

func EncodeApigatewayv2Integration_Id(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2Integration_IntegrationMethod(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["integration_method"] = cty.StringVal(p.IntegrationMethod)
}

func EncodeApigatewayv2Integration_IntegrationUri(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	vals["integration_uri"] = cty.StringVal(p.IntegrationUri)
}

func EncodeApigatewayv2Integration_RequestParameters(p Apigatewayv2IntegrationParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.RequestParameters {
		mVals[key] = cty.StringVal(value)
	}
	vals["request_parameters"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2Integration_TlsConfig(p TlsConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Integration_TlsConfig_ServerNameToVerify(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["tls_config"] = cty.ListVal(valsForCollection)
}

func EncodeApigatewayv2Integration_TlsConfig_ServerNameToVerify(p TlsConfig, vals map[string]cty.Value) {
	vals["server_name_to_verify"] = cty.StringVal(p.ServerNameToVerify)
}

func EncodeApigatewayv2Integration_IntegrationResponseSelectionExpression(p Apigatewayv2IntegrationObservation, vals map[string]cty.Value) {
	vals["integration_response_selection_expression"] = cty.StringVal(p.IntegrationResponseSelectionExpression)
}