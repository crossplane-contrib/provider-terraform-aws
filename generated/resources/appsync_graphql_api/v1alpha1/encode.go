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

package v1alpha1func EncodeAppsyncGraphqlApi(r AppsyncGraphqlApi) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_XrayEnabled(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_AuthenticationType(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Schema(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider(r.Spec.ForProvider.AdditionalAuthenticationProvider, ctyVal)
	EncodeAppsyncGraphqlApi_LogConfig(r.Spec.ForProvider.LogConfig, ctyVal)
	EncodeAppsyncGraphqlApi_OpenidConnectConfig(r.Spec.ForProvider.OpenidConnectConfig, ctyVal)
	EncodeAppsyncGraphqlApi_UserPoolConfig(r.Spec.ForProvider.UserPoolConfig, ctyVal)
	EncodeAppsyncGraphqlApi_Uris(r.Status.AtProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAppsyncGraphqlApi_Tags(p *AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAppsyncGraphqlApi_XrayEnabled(p *AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["xray_enabled"] = cty.BoolVal(p.XrayEnabled)
}

func EncodeAppsyncGraphqlApi_AuthenticationType(p *AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["authentication_type"] = cty.StringVal(p.AuthenticationType)
}

func EncodeAppsyncGraphqlApi_Id(p *AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppsyncGraphqlApi_Name(p *AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppsyncGraphqlApi_Schema(p *AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["schema"] = cty.StringVal(p.Schema)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider(p *AdditionalAuthenticationProvider, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AdditionalAuthenticationProvider {
		ctyVal = make(map[string]cty.Value)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_AuthenticationType(v, ctyVal)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig(v.OpenidConnectConfig, ctyVal)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig(v.UserPoolConfig, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["additional_authentication_provider"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_AuthenticationType(p *AdditionalAuthenticationProvider, vals map[string]cty.Value) {
	vals["authentication_type"] = cty.StringVal(p.AuthenticationType)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig(p *OpenidConnectConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OpenidConnectConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_IatTtl(v, ctyVal)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_Issuer(v, ctyVal)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_AuthTtl(v, ctyVal)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_ClientId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["openid_connect_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_IatTtl(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["iat_ttl"] = cty.IntVal(p.IatTtl)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_Issuer(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_AuthTtl(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["auth_ttl"] = cty.IntVal(p.AuthTtl)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_ClientId(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig(p *UserPoolConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UserPoolConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AwsRegion(v, ctyVal)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_UserPoolId(v, ctyVal)
		EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AppIdClientRegex(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["user_pool_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AwsRegion(p *UserPoolConfig, vals map[string]cty.Value) {
	vals["aws_region"] = cty.StringVal(p.AwsRegion)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_UserPoolId(p *UserPoolConfig, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AppIdClientRegex(p *UserPoolConfig, vals map[string]cty.Value) {
	vals["app_id_client_regex"] = cty.StringVal(p.AppIdClientRegex)
}

func EncodeAppsyncGraphqlApi_LogConfig(p *LogConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LogConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeAppsyncGraphqlApi_LogConfig_CloudwatchLogsRoleArn(v, ctyVal)
		EncodeAppsyncGraphqlApi_LogConfig_ExcludeVerboseContent(v, ctyVal)
		EncodeAppsyncGraphqlApi_LogConfig_FieldLogLevel(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["log_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_LogConfig_CloudwatchLogsRoleArn(p *LogConfig, vals map[string]cty.Value) {
	vals["cloudwatch_logs_role_arn"] = cty.StringVal(p.CloudwatchLogsRoleArn)
}

func EncodeAppsyncGraphqlApi_LogConfig_ExcludeVerboseContent(p *LogConfig, vals map[string]cty.Value) {
	vals["exclude_verbose_content"] = cty.BoolVal(p.ExcludeVerboseContent)
}

func EncodeAppsyncGraphqlApi_LogConfig_FieldLogLevel(p *LogConfig, vals map[string]cty.Value) {
	vals["field_log_level"] = cty.StringVal(p.FieldLogLevel)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig(p *OpenidConnectConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.OpenidConnectConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeAppsyncGraphqlApi_OpenidConnectConfig_IatTtl(v, ctyVal)
		EncodeAppsyncGraphqlApi_OpenidConnectConfig_Issuer(v, ctyVal)
		EncodeAppsyncGraphqlApi_OpenidConnectConfig_AuthTtl(v, ctyVal)
		EncodeAppsyncGraphqlApi_OpenidConnectConfig_ClientId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["openid_connect_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_IatTtl(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["iat_ttl"] = cty.IntVal(p.IatTtl)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_Issuer(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_AuthTtl(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["auth_ttl"] = cty.IntVal(p.AuthTtl)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_ClientId(p *OpenidConnectConfig, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig(p *UserPoolConfig, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.UserPoolConfig {
		ctyVal = make(map[string]cty.Value)
		EncodeAppsyncGraphqlApi_UserPoolConfig_AppIdClientRegex(v, ctyVal)
		EncodeAppsyncGraphqlApi_UserPoolConfig_AwsRegion(v, ctyVal)
		EncodeAppsyncGraphqlApi_UserPoolConfig_DefaultAction(v, ctyVal)
		EncodeAppsyncGraphqlApi_UserPoolConfig_UserPoolId(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["user_pool_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_AppIdClientRegex(p *UserPoolConfig, vals map[string]cty.Value) {
	vals["app_id_client_regex"] = cty.StringVal(p.AppIdClientRegex)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_AwsRegion(p *UserPoolConfig, vals map[string]cty.Value) {
	vals["aws_region"] = cty.StringVal(p.AwsRegion)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_DefaultAction(p *UserPoolConfig, vals map[string]cty.Value) {
	vals["default_action"] = cty.StringVal(p.DefaultAction)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_UserPoolId(p *UserPoolConfig, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeAppsyncGraphqlApi_Uris(p *AppsyncGraphqlApiObservation, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Uris {
		mVals[key] = cty.StringVal(value)
	}
	vals["uris"] = cty.MapVal(mVals)
}

func EncodeAppsyncGraphqlApi_Arn(p *AppsyncGraphqlApiObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}