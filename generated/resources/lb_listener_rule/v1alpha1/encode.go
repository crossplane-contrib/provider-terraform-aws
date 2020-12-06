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

func EncodeLbListenerRule(r LbListenerRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeLbListenerRule_ListenerArn(r.Spec.ForProvider, ctyVal)
	EncodeLbListenerRule_Priority(r.Spec.ForProvider, ctyVal)
	EncodeLbListenerRule_Action(r.Spec.ForProvider.Action, ctyVal)
	EncodeLbListenerRule_Condition(r.Spec.ForProvider.Condition, ctyVal)
	EncodeLbListenerRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLbListenerRule_Id(p LbListenerRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLbListenerRule_ListenerArn(p LbListenerRuleParameters, vals map[string]cty.Value) {
	vals["listener_arn"] = cty.StringVal(p.ListenerArn)
}

func EncodeLbListenerRule_Priority(p LbListenerRuleParameters, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeLbListenerRule_Action(p []Action, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLbListenerRule_Action_Order(v, ctyVal)
		EncodeLbListenerRule_Action_TargetGroupArn(v, ctyVal)
		EncodeLbListenerRule_Action_Type(v, ctyVal)
		EncodeLbListenerRule_Action_AuthenticateCognito(v.AuthenticateCognito, ctyVal)
		EncodeLbListenerRule_Action_AuthenticateOidc(v.AuthenticateOidc, ctyVal)
		EncodeLbListenerRule_Action_FixedResponse(v.FixedResponse, ctyVal)
		EncodeLbListenerRule_Action_Forward(v.Forward, ctyVal)
		EncodeLbListenerRule_Action_Redirect(v.Redirect, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["action"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Action_Order(p Action, vals map[string]cty.Value) {
	vals["order"] = cty.NumberIntVal(p.Order)
}

func EncodeLbListenerRule_Action_TargetGroupArn(p Action, vals map[string]cty.Value) {
	vals["target_group_arn"] = cty.StringVal(p.TargetGroupArn)
}

func EncodeLbListenerRule_Action_Type(p Action, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeLbListenerRule_Action_AuthenticateCognito(p AuthenticateCognito, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Action_AuthenticateCognito_OnUnauthenticatedRequest(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateCognito_Scope(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateCognito_SessionCookieName(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateCognito_SessionTimeout(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateCognito_UserPoolArn(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateCognito_UserPoolClientId(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateCognito_UserPoolDomain(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateCognito_AuthenticationRequestExtraParams(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authenticate_cognito"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_OnUnauthenticatedRequest(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["on_unauthenticated_request"] = cty.StringVal(p.OnUnauthenticatedRequest)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_Scope(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_SessionCookieName(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["session_cookie_name"] = cty.StringVal(p.SessionCookieName)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_SessionTimeout(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["session_timeout"] = cty.NumberIntVal(p.SessionTimeout)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_UserPoolArn(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_arn"] = cty.StringVal(p.UserPoolArn)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_UserPoolClientId(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_client_id"] = cty.StringVal(p.UserPoolClientId)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_UserPoolDomain(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_domain"] = cty.StringVal(p.UserPoolDomain)
}

func EncodeLbListenerRule_Action_AuthenticateCognito_AuthenticationRequestExtraParams(p AuthenticateCognito, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.AuthenticationRequestExtraParams {
		mVals[key] = cty.StringVal(value)
	}
	vals["authentication_request_extra_params"] = cty.MapVal(mVals)
}

func EncodeLbListenerRule_Action_AuthenticateOidc(p AuthenticateOidc, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Action_AuthenticateOidc_AuthorizationEndpoint(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_ClientId(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_Scope(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_SessionTimeout(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_TokenEndpoint(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_AuthenticationRequestExtraParams(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_Issuer(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_OnUnauthenticatedRequest(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_SessionCookieName(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_UserInfoEndpoint(p, ctyVal)
	EncodeLbListenerRule_Action_AuthenticateOidc_ClientSecret(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authenticate_oidc"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_AuthorizationEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["authorization_endpoint"] = cty.StringVal(p.AuthorizationEndpoint)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_ClientId(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_Scope(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_SessionTimeout(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["session_timeout"] = cty.NumberIntVal(p.SessionTimeout)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_TokenEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["token_endpoint"] = cty.StringVal(p.TokenEndpoint)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_AuthenticationRequestExtraParams(p AuthenticateOidc, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.AuthenticationRequestExtraParams {
		mVals[key] = cty.StringVal(value)
	}
	vals["authentication_request_extra_params"] = cty.MapVal(mVals)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_Issuer(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_OnUnauthenticatedRequest(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["on_unauthenticated_request"] = cty.StringVal(p.OnUnauthenticatedRequest)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_SessionCookieName(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["session_cookie_name"] = cty.StringVal(p.SessionCookieName)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_UserInfoEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["user_info_endpoint"] = cty.StringVal(p.UserInfoEndpoint)
}

func EncodeLbListenerRule_Action_AuthenticateOidc_ClientSecret(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["client_secret"] = cty.StringVal(p.ClientSecret)
}

func EncodeLbListenerRule_Action_FixedResponse(p FixedResponse, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Action_FixedResponse_ContentType(p, ctyVal)
	EncodeLbListenerRule_Action_FixedResponse_MessageBody(p, ctyVal)
	EncodeLbListenerRule_Action_FixedResponse_StatusCode(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["fixed_response"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Action_FixedResponse_ContentType(p FixedResponse, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLbListenerRule_Action_FixedResponse_MessageBody(p FixedResponse, vals map[string]cty.Value) {
	vals["message_body"] = cty.StringVal(p.MessageBody)
}

func EncodeLbListenerRule_Action_FixedResponse_StatusCode(p FixedResponse, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeLbListenerRule_Action_Forward(p Forward, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Action_Forward_Stickiness(p.Stickiness, ctyVal)
	EncodeLbListenerRule_Action_Forward_TargetGroup(p.TargetGroup, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["forward"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Action_Forward_Stickiness(p Stickiness, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Action_Forward_Stickiness_Duration(p, ctyVal)
	EncodeLbListenerRule_Action_Forward_Stickiness_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["stickiness"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Action_Forward_Stickiness_Duration(p Stickiness, vals map[string]cty.Value) {
	vals["duration"] = cty.NumberIntVal(p.Duration)
}

func EncodeLbListenerRule_Action_Forward_Stickiness_Enabled(p Stickiness, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeLbListenerRule_Action_Forward_TargetGroup(p []TargetGroup, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLbListenerRule_Action_Forward_TargetGroup_Arn(v, ctyVal)
		EncodeLbListenerRule_Action_Forward_TargetGroup_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["target_group"] = cty.SetVal(valsForCollection)
}

func EncodeLbListenerRule_Action_Forward_TargetGroup_Arn(p TargetGroup, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLbListenerRule_Action_Forward_TargetGroup_Weight(p TargetGroup, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeLbListenerRule_Action_Redirect(p Redirect, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Action_Redirect_Host(p, ctyVal)
	EncodeLbListenerRule_Action_Redirect_Path(p, ctyVal)
	EncodeLbListenerRule_Action_Redirect_Port(p, ctyVal)
	EncodeLbListenerRule_Action_Redirect_Protocol(p, ctyVal)
	EncodeLbListenerRule_Action_Redirect_Query(p, ctyVal)
	EncodeLbListenerRule_Action_Redirect_StatusCode(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["redirect"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Action_Redirect_Host(p Redirect, vals map[string]cty.Value) {
	vals["host"] = cty.StringVal(p.Host)
}

func EncodeLbListenerRule_Action_Redirect_Path(p Redirect, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeLbListenerRule_Action_Redirect_Port(p Redirect, vals map[string]cty.Value) {
	vals["port"] = cty.StringVal(p.Port)
}

func EncodeLbListenerRule_Action_Redirect_Protocol(p Redirect, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeLbListenerRule_Action_Redirect_Query(p Redirect, vals map[string]cty.Value) {
	vals["query"] = cty.StringVal(p.Query)
}

func EncodeLbListenerRule_Action_Redirect_StatusCode(p Redirect, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeLbListenerRule_Condition(p []Condition, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLbListenerRule_Condition_PathPattern(v.PathPattern, ctyVal)
		EncodeLbListenerRule_Condition_QueryString(v.QueryString, ctyVal)
		EncodeLbListenerRule_Condition_SourceIp(v.SourceIp, ctyVal)
		EncodeLbListenerRule_Condition_HostHeader(v.HostHeader, ctyVal)
		EncodeLbListenerRule_Condition_HttpHeader(v.HttpHeader, ctyVal)
		EncodeLbListenerRule_Condition_HttpRequestMethod(v.HttpRequestMethod, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["condition"] = cty.SetVal(valsForCollection)
}

func EncodeLbListenerRule_Condition_PathPattern(p PathPattern, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Condition_PathPattern_Values(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["path_pattern"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Condition_PathPattern_Values(p PathPattern, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.SetVal(colVals)
}

func EncodeLbListenerRule_Condition_QueryString(p QueryString, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Condition_QueryString_Key(p, ctyVal)
	EncodeLbListenerRule_Condition_QueryString_Value(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["query_string"] = cty.SetVal(valsForCollection)
}

func EncodeLbListenerRule_Condition_QueryString_Key(p QueryString, vals map[string]cty.Value) {
	vals["key"] = cty.StringVal(p.Key)
}

func EncodeLbListenerRule_Condition_QueryString_Value(p QueryString, vals map[string]cty.Value) {
	vals["value"] = cty.StringVal(p.Value)
}

func EncodeLbListenerRule_Condition_SourceIp(p SourceIp, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Condition_SourceIp_Values(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["source_ip"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Condition_SourceIp_Values(p SourceIp, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.SetVal(colVals)
}

func EncodeLbListenerRule_Condition_HostHeader(p HostHeader, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Condition_HostHeader_Values(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["host_header"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Condition_HostHeader_Values(p HostHeader, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.SetVal(colVals)
}

func EncodeLbListenerRule_Condition_HttpHeader(p HttpHeader, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Condition_HttpHeader_HttpHeaderName(p, ctyVal)
	EncodeLbListenerRule_Condition_HttpHeader_Values(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["http_header"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Condition_HttpHeader_HttpHeaderName(p HttpHeader, vals map[string]cty.Value) {
	vals["http_header_name"] = cty.StringVal(p.HttpHeaderName)
}

func EncodeLbListenerRule_Condition_HttpHeader_Values(p HttpHeader, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.SetVal(colVals)
}

func EncodeLbListenerRule_Condition_HttpRequestMethod(p HttpRequestMethod, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListenerRule_Condition_HttpRequestMethod_Values(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["http_request_method"] = cty.ListVal(valsForCollection)
}

func EncodeLbListenerRule_Condition_HttpRequestMethod_Values(p HttpRequestMethod, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Values {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["values"] = cty.SetVal(colVals)
}

func EncodeLbListenerRule_Arn(p LbListenerRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}