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

func EncodeApigatewayv2Stage(r Apigatewayv2Stage) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Stage_Id(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_AutoDeploy(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_DeploymentId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_Description(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_Name(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_StageVariables(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_ApiId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_ClientCertificateId(r.Spec.ForProvider, ctyVal)
	EncodeApigatewayv2Stage_AccessLogSettings(r.Spec.ForProvider.AccessLogSettings, ctyVal)
	EncodeApigatewayv2Stage_DefaultRouteSettings(r.Spec.ForProvider.DefaultRouteSettings, ctyVal)
	EncodeApigatewayv2Stage_RouteSettings(r.Spec.ForProvider.RouteSettings, ctyVal)
	EncodeApigatewayv2Stage_ExecutionArn(r.Status.AtProvider, ctyVal)
	EncodeApigatewayv2Stage_InvokeUrl(r.Status.AtProvider, ctyVal)
	EncodeApigatewayv2Stage_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeApigatewayv2Stage_Id(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApigatewayv2Stage_Tags(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2Stage_AutoDeploy(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	vals["auto_deploy"] = cty.BoolVal(p.AutoDeploy)
}

func EncodeApigatewayv2Stage_DeploymentId(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	vals["deployment_id"] = cty.StringVal(p.DeploymentId)
}

func EncodeApigatewayv2Stage_Description(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApigatewayv2Stage_Name(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApigatewayv2Stage_StageVariables(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.StageVariables {
		mVals[key] = cty.StringVal(value)
	}
	vals["stage_variables"] = cty.MapVal(mVals)
}

func EncodeApigatewayv2Stage_ApiId(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApigatewayv2Stage_ClientCertificateId(p Apigatewayv2StageParameters, vals map[string]cty.Value) {
	vals["client_certificate_id"] = cty.StringVal(p.ClientCertificateId)
}

func EncodeApigatewayv2Stage_AccessLogSettings(p AccessLogSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Stage_AccessLogSettings_DestinationArn(p, ctyVal)
	EncodeApigatewayv2Stage_AccessLogSettings_Format(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["access_log_settings"] = cty.ListVal(valsForCollection)
}

func EncodeApigatewayv2Stage_AccessLogSettings_DestinationArn(p AccessLogSettings, vals map[string]cty.Value) {
	vals["destination_arn"] = cty.StringVal(p.DestinationArn)
}

func EncodeApigatewayv2Stage_AccessLogSettings_Format(p AccessLogSettings, vals map[string]cty.Value) {
	vals["format"] = cty.StringVal(p.Format)
}

func EncodeApigatewayv2Stage_DefaultRouteSettings(p DefaultRouteSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Stage_DefaultRouteSettings_ThrottlingBurstLimit(p, ctyVal)
	EncodeApigatewayv2Stage_DefaultRouteSettings_ThrottlingRateLimit(p, ctyVal)
	EncodeApigatewayv2Stage_DefaultRouteSettings_DataTraceEnabled(p, ctyVal)
	EncodeApigatewayv2Stage_DefaultRouteSettings_DetailedMetricsEnabled(p, ctyVal)
	EncodeApigatewayv2Stage_DefaultRouteSettings_LoggingLevel(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["default_route_settings"] = cty.ListVal(valsForCollection)
}

func EncodeApigatewayv2Stage_DefaultRouteSettings_ThrottlingBurstLimit(p DefaultRouteSettings, vals map[string]cty.Value) {
	vals["throttling_burst_limit"] = cty.NumberIntVal(p.ThrottlingBurstLimit)
}

func EncodeApigatewayv2Stage_DefaultRouteSettings_ThrottlingRateLimit(p DefaultRouteSettings, vals map[string]cty.Value) {
	vals["throttling_rate_limit"] = cty.NumberIntVal(p.ThrottlingRateLimit)
}

func EncodeApigatewayv2Stage_DefaultRouteSettings_DataTraceEnabled(p DefaultRouteSettings, vals map[string]cty.Value) {
	vals["data_trace_enabled"] = cty.BoolVal(p.DataTraceEnabled)
}

func EncodeApigatewayv2Stage_DefaultRouteSettings_DetailedMetricsEnabled(p DefaultRouteSettings, vals map[string]cty.Value) {
	vals["detailed_metrics_enabled"] = cty.BoolVal(p.DetailedMetricsEnabled)
}

func EncodeApigatewayv2Stage_DefaultRouteSettings_LoggingLevel(p DefaultRouteSettings, vals map[string]cty.Value) {
	vals["logging_level"] = cty.StringVal(p.LoggingLevel)
}

func EncodeApigatewayv2Stage_RouteSettings(p RouteSettings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApigatewayv2Stage_RouteSettings_LoggingLevel(p, ctyVal)
	EncodeApigatewayv2Stage_RouteSettings_RouteKey(p, ctyVal)
	EncodeApigatewayv2Stage_RouteSettings_ThrottlingBurstLimit(p, ctyVal)
	EncodeApigatewayv2Stage_RouteSettings_ThrottlingRateLimit(p, ctyVal)
	EncodeApigatewayv2Stage_RouteSettings_DataTraceEnabled(p, ctyVal)
	EncodeApigatewayv2Stage_RouteSettings_DetailedMetricsEnabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["route_settings"] = cty.SetVal(valsForCollection)
}

func EncodeApigatewayv2Stage_RouteSettings_LoggingLevel(p RouteSettings, vals map[string]cty.Value) {
	vals["logging_level"] = cty.StringVal(p.LoggingLevel)
}

func EncodeApigatewayv2Stage_RouteSettings_RouteKey(p RouteSettings, vals map[string]cty.Value) {
	vals["route_key"] = cty.StringVal(p.RouteKey)
}

func EncodeApigatewayv2Stage_RouteSettings_ThrottlingBurstLimit(p RouteSettings, vals map[string]cty.Value) {
	vals["throttling_burst_limit"] = cty.NumberIntVal(p.ThrottlingBurstLimit)
}

func EncodeApigatewayv2Stage_RouteSettings_ThrottlingRateLimit(p RouteSettings, vals map[string]cty.Value) {
	vals["throttling_rate_limit"] = cty.NumberIntVal(p.ThrottlingRateLimit)
}

func EncodeApigatewayv2Stage_RouteSettings_DataTraceEnabled(p RouteSettings, vals map[string]cty.Value) {
	vals["data_trace_enabled"] = cty.BoolVal(p.DataTraceEnabled)
}

func EncodeApigatewayv2Stage_RouteSettings_DetailedMetricsEnabled(p RouteSettings, vals map[string]cty.Value) {
	vals["detailed_metrics_enabled"] = cty.BoolVal(p.DetailedMetricsEnabled)
}

func EncodeApigatewayv2Stage_ExecutionArn(p Apigatewayv2StageObservation, vals map[string]cty.Value) {
	vals["execution_arn"] = cty.StringVal(p.ExecutionArn)
}

func EncodeApigatewayv2Stage_InvokeUrl(p Apigatewayv2StageObservation, vals map[string]cty.Value) {
	vals["invoke_url"] = cty.StringVal(p.InvokeUrl)
}

func EncodeApigatewayv2Stage_Arn(p Apigatewayv2StageObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}