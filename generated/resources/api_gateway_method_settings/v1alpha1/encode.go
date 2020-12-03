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

package v1alpha1func EncodeApiGatewayMethodSettings(r ApiGatewayMethodSettings) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayMethodSettings_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_MethodPath(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_StageName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_Settings(r.Spec.ForProvider.Settings, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayMethodSettings_Id(p *ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayMethodSettings_MethodPath(p *ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["method_path"] = cty.StringVal(p.MethodPath)
}

func EncodeApiGatewayMethodSettings_RestApiId(p *ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayMethodSettings_StageName(p *ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["stage_name"] = cty.StringVal(p.StageName)
}

func EncodeApiGatewayMethodSettings_Settings(p *Settings, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Settings {
		ctyVal = make(map[string]cty.Value)
		EncodeApiGatewayMethodSettings_Settings_CacheTtlInSeconds(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_DataTraceEnabled(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_MetricsEnabled(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_ThrottlingBurstLimit(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_ThrottlingRateLimit(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_UnauthorizedCacheControlHeaderStrategy(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_CacheDataEncrypted(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_CachingEnabled(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_LoggingLevel(v, ctyVal)
		EncodeApiGatewayMethodSettings_Settings_RequireAuthorizationForCacheControl(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["settings"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayMethodSettings_Settings_CacheTtlInSeconds(p *Settings, vals map[string]cty.Value) {
	vals["cache_ttl_in_seconds"] = cty.IntVal(p.CacheTtlInSeconds)
}

func EncodeApiGatewayMethodSettings_Settings_DataTraceEnabled(p *Settings, vals map[string]cty.Value) {
	vals["data_trace_enabled"] = cty.BoolVal(p.DataTraceEnabled)
}

func EncodeApiGatewayMethodSettings_Settings_MetricsEnabled(p *Settings, vals map[string]cty.Value) {
	vals["metrics_enabled"] = cty.BoolVal(p.MetricsEnabled)
}

func EncodeApiGatewayMethodSettings_Settings_ThrottlingBurstLimit(p *Settings, vals map[string]cty.Value) {
	vals["throttling_burst_limit"] = cty.IntVal(p.ThrottlingBurstLimit)
}

func EncodeApiGatewayMethodSettings_Settings_ThrottlingRateLimit(p *Settings, vals map[string]cty.Value) {
	vals["throttling_rate_limit"] = cty.IntVal(p.ThrottlingRateLimit)
}

func EncodeApiGatewayMethodSettings_Settings_UnauthorizedCacheControlHeaderStrategy(p *Settings, vals map[string]cty.Value) {
	vals["unauthorized_cache_control_header_strategy"] = cty.StringVal(p.UnauthorizedCacheControlHeaderStrategy)
}

func EncodeApiGatewayMethodSettings_Settings_CacheDataEncrypted(p *Settings, vals map[string]cty.Value) {
	vals["cache_data_encrypted"] = cty.BoolVal(p.CacheDataEncrypted)
}

func EncodeApiGatewayMethodSettings_Settings_CachingEnabled(p *Settings, vals map[string]cty.Value) {
	vals["caching_enabled"] = cty.BoolVal(p.CachingEnabled)
}

func EncodeApiGatewayMethodSettings_Settings_LoggingLevel(p *Settings, vals map[string]cty.Value) {
	vals["logging_level"] = cty.StringVal(p.LoggingLevel)
}

func EncodeApiGatewayMethodSettings_Settings_RequireAuthorizationForCacheControl(p *Settings, vals map[string]cty.Value) {
	vals["require_authorization_for_cache_control"] = cty.BoolVal(p.RequireAuthorizationForCacheControl)
}