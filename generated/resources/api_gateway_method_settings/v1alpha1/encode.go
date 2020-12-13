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
	r, ok := mr.(*ApiGatewayMethodSettings)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a ApiGatewayMethodSettings.")
	}
	return EncodeApiGatewayMethodSettings(*r), nil
}

func EncodeApiGatewayMethodSettings(r ApiGatewayMethodSettings) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayMethodSettings_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_MethodPath(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_RestApiId(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_StageName(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayMethodSettings_Settings(r.Spec.ForProvider.Settings, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeApiGatewayMethodSettings_Id(p ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayMethodSettings_MethodPath(p ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["method_path"] = cty.StringVal(p.MethodPath)
}

func EncodeApiGatewayMethodSettings_RestApiId(p ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["rest_api_id"] = cty.StringVal(p.RestApiId)
}

func EncodeApiGatewayMethodSettings_StageName(p ApiGatewayMethodSettingsParameters, vals map[string]cty.Value) {
	vals["stage_name"] = cty.StringVal(p.StageName)
}

func EncodeApiGatewayMethodSettings_Settings(p Settings, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeApiGatewayMethodSettings_Settings_CacheTtlInSeconds(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_LoggingLevel(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_MetricsEnabled(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_UnauthorizedCacheControlHeaderStrategy(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_CacheDataEncrypted(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_CachingEnabled(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_DataTraceEnabled(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_RequireAuthorizationForCacheControl(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_ThrottlingBurstLimit(p, ctyVal)
	EncodeApiGatewayMethodSettings_Settings_ThrottlingRateLimit(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["settings"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayMethodSettings_Settings_CacheTtlInSeconds(p Settings, vals map[string]cty.Value) {
	vals["cache_ttl_in_seconds"] = cty.NumberIntVal(p.CacheTtlInSeconds)
}

func EncodeApiGatewayMethodSettings_Settings_LoggingLevel(p Settings, vals map[string]cty.Value) {
	vals["logging_level"] = cty.StringVal(p.LoggingLevel)
}

func EncodeApiGatewayMethodSettings_Settings_MetricsEnabled(p Settings, vals map[string]cty.Value) {
	vals["metrics_enabled"] = cty.BoolVal(p.MetricsEnabled)
}

func EncodeApiGatewayMethodSettings_Settings_UnauthorizedCacheControlHeaderStrategy(p Settings, vals map[string]cty.Value) {
	vals["unauthorized_cache_control_header_strategy"] = cty.StringVal(p.UnauthorizedCacheControlHeaderStrategy)
}

func EncodeApiGatewayMethodSettings_Settings_CacheDataEncrypted(p Settings, vals map[string]cty.Value) {
	vals["cache_data_encrypted"] = cty.BoolVal(p.CacheDataEncrypted)
}

func EncodeApiGatewayMethodSettings_Settings_CachingEnabled(p Settings, vals map[string]cty.Value) {
	vals["caching_enabled"] = cty.BoolVal(p.CachingEnabled)
}

func EncodeApiGatewayMethodSettings_Settings_DataTraceEnabled(p Settings, vals map[string]cty.Value) {
	vals["data_trace_enabled"] = cty.BoolVal(p.DataTraceEnabled)
}

func EncodeApiGatewayMethodSettings_Settings_RequireAuthorizationForCacheControl(p Settings, vals map[string]cty.Value) {
	vals["require_authorization_for_cache_control"] = cty.BoolVal(p.RequireAuthorizationForCacheControl)
}

func EncodeApiGatewayMethodSettings_Settings_ThrottlingBurstLimit(p Settings, vals map[string]cty.Value) {
	vals["throttling_burst_limit"] = cty.NumberIntVal(p.ThrottlingBurstLimit)
}

func EncodeApiGatewayMethodSettings_Settings_ThrottlingRateLimit(p Settings, vals map[string]cty.Value) {
	vals["throttling_rate_limit"] = cty.NumberIntVal(p.ThrottlingRateLimit)
}