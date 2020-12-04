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

func EncodeAppsyncGraphqlApi(r AppsyncGraphqlApi) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Schema(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_XrayEnabled(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_AuthenticationType(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider(r.Spec.ForProvider.AdditionalAuthenticationProvider, ctyVal)
	EncodeAppsyncGraphqlApi_LogConfig(r.Spec.ForProvider.LogConfig, ctyVal)
	EncodeAppsyncGraphqlApi_OpenidConnectConfig(r.Spec.ForProvider.OpenidConnectConfig, ctyVal)
	EncodeAppsyncGraphqlApi_UserPoolConfig(r.Spec.ForProvider.UserPoolConfig, ctyVal)
	EncodeAppsyncGraphqlApi_Uris(r.Status.AtProvider, ctyVal)
	EncodeAppsyncGraphqlApi_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAppsyncGraphqlApi_Name(p AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppsyncGraphqlApi_Schema(p AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["schema"] = cty.StringVal(p.Schema)
}

func EncodeAppsyncGraphqlApi_Tags(p AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAppsyncGraphqlApi_XrayEnabled(p AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["xray_enabled"] = cty.BoolVal(p.XrayEnabled)
}

func EncodeAppsyncGraphqlApi_AuthenticationType(p AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["authentication_type"] = cty.StringVal(p.AuthenticationType)
}

func EncodeAppsyncGraphqlApi_Id(p AppsyncGraphqlApiParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider(p AdditionalAuthenticationProvider, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_AuthenticationType(p, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig(p.OpenidConnectConfig, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig(p.UserPoolConfig, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["additional_authentication_provider"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_AuthenticationType(p AdditionalAuthenticationProvider, vals map[string]cty.Value) {
	vals["authentication_type"] = cty.StringVal(p.AuthenticationType)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig(p OpenidConnectConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_AuthTtl(p, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_ClientId(p, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_IatTtl(p, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_Issuer(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["openid_connect_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_AuthTtl(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["auth_ttl"] = cty.NumberIntVal(p.AuthTtl)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_ClientId(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_IatTtl(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["iat_ttl"] = cty.NumberIntVal(p.IatTtl)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_OpenidConnectConfig_Issuer(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig(p UserPoolConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AppIdClientRegex(p, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AwsRegion(p, ctyVal)
	EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_UserPoolId(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["user_pool_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AppIdClientRegex(p UserPoolConfig, vals map[string]cty.Value) {
	vals["app_id_client_regex"] = cty.StringVal(p.AppIdClientRegex)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_AwsRegion(p UserPoolConfig, vals map[string]cty.Value) {
	vals["aws_region"] = cty.StringVal(p.AwsRegion)
}

func EncodeAppsyncGraphqlApi_AdditionalAuthenticationProvider_UserPoolConfig_UserPoolId(p UserPoolConfig, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeAppsyncGraphqlApi_LogConfig(p LogConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_LogConfig_CloudwatchLogsRoleArn(p, ctyVal)
	EncodeAppsyncGraphqlApi_LogConfig_ExcludeVerboseContent(p, ctyVal)
	EncodeAppsyncGraphqlApi_LogConfig_FieldLogLevel(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["log_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_LogConfig_CloudwatchLogsRoleArn(p LogConfig, vals map[string]cty.Value) {
	vals["cloudwatch_logs_role_arn"] = cty.StringVal(p.CloudwatchLogsRoleArn)
}

func EncodeAppsyncGraphqlApi_LogConfig_ExcludeVerboseContent(p LogConfig, vals map[string]cty.Value) {
	vals["exclude_verbose_content"] = cty.BoolVal(p.ExcludeVerboseContent)
}

func EncodeAppsyncGraphqlApi_LogConfig_FieldLogLevel(p LogConfig, vals map[string]cty.Value) {
	vals["field_log_level"] = cty.StringVal(p.FieldLogLevel)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig(p OpenidConnectConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_OpenidConnectConfig_AuthTtl(p, ctyVal)
	EncodeAppsyncGraphqlApi_OpenidConnectConfig_ClientId(p, ctyVal)
	EncodeAppsyncGraphqlApi_OpenidConnectConfig_IatTtl(p, ctyVal)
	EncodeAppsyncGraphqlApi_OpenidConnectConfig_Issuer(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["openid_connect_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_AuthTtl(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["auth_ttl"] = cty.NumberIntVal(p.AuthTtl)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_ClientId(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_IatTtl(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["iat_ttl"] = cty.NumberIntVal(p.IatTtl)
}

func EncodeAppsyncGraphqlApi_OpenidConnectConfig_Issuer(p OpenidConnectConfig, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig(p UserPoolConfig, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAppsyncGraphqlApi_UserPoolConfig_DefaultAction(p, ctyVal)
	EncodeAppsyncGraphqlApi_UserPoolConfig_UserPoolId(p, ctyVal)
	EncodeAppsyncGraphqlApi_UserPoolConfig_AppIdClientRegex(p, ctyVal)
	EncodeAppsyncGraphqlApi_UserPoolConfig_AwsRegion(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["user_pool_config"] = cty.ListVal(valsForCollection)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_DefaultAction(p UserPoolConfig, vals map[string]cty.Value) {
	vals["default_action"] = cty.StringVal(p.DefaultAction)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_UserPoolId(p UserPoolConfig, vals map[string]cty.Value) {
	vals["user_pool_id"] = cty.StringVal(p.UserPoolId)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_AppIdClientRegex(p UserPoolConfig, vals map[string]cty.Value) {
	vals["app_id_client_regex"] = cty.StringVal(p.AppIdClientRegex)
}

func EncodeAppsyncGraphqlApi_UserPoolConfig_AwsRegion(p UserPoolConfig, vals map[string]cty.Value) {
	vals["aws_region"] = cty.StringVal(p.AwsRegion)
}

func EncodeAppsyncGraphqlApi_Uris(p AppsyncGraphqlApiObservation, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Uris {
		mVals[key] = cty.StringVal(value)
	}
	vals["uris"] = cty.MapVal(mVals)
}

func EncodeAppsyncGraphqlApi_Arn(p AppsyncGraphqlApiObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}