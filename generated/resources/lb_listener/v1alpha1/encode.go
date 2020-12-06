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

func EncodeLbListener(r LbListener) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_CertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeLbListener_Id(r.Spec.ForProvider, ctyVal)
	EncodeLbListener_LoadBalancerArn(r.Spec.ForProvider, ctyVal)
	EncodeLbListener_Port(r.Spec.ForProvider, ctyVal)
	EncodeLbListener_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeLbListener_SslPolicy(r.Spec.ForProvider, ctyVal)
	EncodeLbListener_DefaultAction(r.Spec.ForProvider.DefaultAction, ctyVal)
	EncodeLbListener_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeLbListener_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLbListener_CertificateArn(p LbListenerParameters, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeLbListener_Id(p LbListenerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLbListener_LoadBalancerArn(p LbListenerParameters, vals map[string]cty.Value) {
	vals["load_balancer_arn"] = cty.StringVal(p.LoadBalancerArn)
}

func EncodeLbListener_Port(p LbListenerParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeLbListener_Protocol(p LbListenerParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeLbListener_SslPolicy(p LbListenerParameters, vals map[string]cty.Value) {
	vals["ssl_policy"] = cty.StringVal(p.SslPolicy)
}

func EncodeLbListener_DefaultAction(p []DefaultAction, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLbListener_DefaultAction_Order(v, ctyVal)
		EncodeLbListener_DefaultAction_TargetGroupArn(v, ctyVal)
		EncodeLbListener_DefaultAction_Type(v, ctyVal)
		EncodeLbListener_DefaultAction_Redirect(v.Redirect, ctyVal)
		EncodeLbListener_DefaultAction_AuthenticateCognito(v.AuthenticateCognito, ctyVal)
		EncodeLbListener_DefaultAction_AuthenticateOidc(v.AuthenticateOidc, ctyVal)
		EncodeLbListener_DefaultAction_FixedResponse(v.FixedResponse, ctyVal)
		EncodeLbListener_DefaultAction_Forward(v.Forward, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["default_action"] = cty.ListVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_Order(p DefaultAction, vals map[string]cty.Value) {
	vals["order"] = cty.NumberIntVal(p.Order)
}

func EncodeLbListener_DefaultAction_TargetGroupArn(p DefaultAction, vals map[string]cty.Value) {
	vals["target_group_arn"] = cty.StringVal(p.TargetGroupArn)
}

func EncodeLbListener_DefaultAction_Type(p DefaultAction, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeLbListener_DefaultAction_Redirect(p Redirect, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_DefaultAction_Redirect_Path(p, ctyVal)
	EncodeLbListener_DefaultAction_Redirect_Port(p, ctyVal)
	EncodeLbListener_DefaultAction_Redirect_Protocol(p, ctyVal)
	EncodeLbListener_DefaultAction_Redirect_Query(p, ctyVal)
	EncodeLbListener_DefaultAction_Redirect_StatusCode(p, ctyVal)
	EncodeLbListener_DefaultAction_Redirect_Host(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["redirect"] = cty.ListVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_Redirect_Path(p Redirect, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeLbListener_DefaultAction_Redirect_Port(p Redirect, vals map[string]cty.Value) {
	vals["port"] = cty.StringVal(p.Port)
}

func EncodeLbListener_DefaultAction_Redirect_Protocol(p Redirect, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeLbListener_DefaultAction_Redirect_Query(p Redirect, vals map[string]cty.Value) {
	vals["query"] = cty.StringVal(p.Query)
}

func EncodeLbListener_DefaultAction_Redirect_StatusCode(p Redirect, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeLbListener_DefaultAction_Redirect_Host(p Redirect, vals map[string]cty.Value) {
	vals["host"] = cty.StringVal(p.Host)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito(p AuthenticateCognito, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_DefaultAction_AuthenticateCognito_AuthenticationRequestExtraParams(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateCognito_OnUnauthenticatedRequest(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateCognito_Scope(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateCognito_SessionCookieName(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateCognito_SessionTimeout(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateCognito_UserPoolArn(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateCognito_UserPoolClientId(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateCognito_UserPoolDomain(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authenticate_cognito"] = cty.ListVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_AuthenticationRequestExtraParams(p AuthenticateCognito, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.AuthenticationRequestExtraParams {
		mVals[key] = cty.StringVal(value)
	}
	vals["authentication_request_extra_params"] = cty.MapVal(mVals)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_OnUnauthenticatedRequest(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["on_unauthenticated_request"] = cty.StringVal(p.OnUnauthenticatedRequest)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_Scope(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_SessionCookieName(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["session_cookie_name"] = cty.StringVal(p.SessionCookieName)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_SessionTimeout(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["session_timeout"] = cty.NumberIntVal(p.SessionTimeout)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_UserPoolArn(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_arn"] = cty.StringVal(p.UserPoolArn)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_UserPoolClientId(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_client_id"] = cty.StringVal(p.UserPoolClientId)
}

func EncodeLbListener_DefaultAction_AuthenticateCognito_UserPoolDomain(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_domain"] = cty.StringVal(p.UserPoolDomain)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc(p AuthenticateOidc, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_DefaultAction_AuthenticateOidc_SessionTimeout(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_AuthorizationEndpoint(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_OnUnauthenticatedRequest(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_Scope(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_Issuer(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_SessionCookieName(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_TokenEndpoint(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_UserInfoEndpoint(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_AuthenticationRequestExtraParams(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_ClientId(p, ctyVal)
	EncodeLbListener_DefaultAction_AuthenticateOidc_ClientSecret(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authenticate_oidc"] = cty.ListVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_SessionTimeout(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["session_timeout"] = cty.NumberIntVal(p.SessionTimeout)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_AuthorizationEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["authorization_endpoint"] = cty.StringVal(p.AuthorizationEndpoint)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_OnUnauthenticatedRequest(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["on_unauthenticated_request"] = cty.StringVal(p.OnUnauthenticatedRequest)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_Scope(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_Issuer(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_SessionCookieName(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["session_cookie_name"] = cty.StringVal(p.SessionCookieName)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_TokenEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["token_endpoint"] = cty.StringVal(p.TokenEndpoint)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_UserInfoEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["user_info_endpoint"] = cty.StringVal(p.UserInfoEndpoint)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_AuthenticationRequestExtraParams(p AuthenticateOidc, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.AuthenticationRequestExtraParams {
		mVals[key] = cty.StringVal(value)
	}
	vals["authentication_request_extra_params"] = cty.MapVal(mVals)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_ClientId(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeLbListener_DefaultAction_AuthenticateOidc_ClientSecret(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["client_secret"] = cty.StringVal(p.ClientSecret)
}

func EncodeLbListener_DefaultAction_FixedResponse(p FixedResponse, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_DefaultAction_FixedResponse_ContentType(p, ctyVal)
	EncodeLbListener_DefaultAction_FixedResponse_MessageBody(p, ctyVal)
	EncodeLbListener_DefaultAction_FixedResponse_StatusCode(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["fixed_response"] = cty.ListVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_FixedResponse_ContentType(p FixedResponse, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLbListener_DefaultAction_FixedResponse_MessageBody(p FixedResponse, vals map[string]cty.Value) {
	vals["message_body"] = cty.StringVal(p.MessageBody)
}

func EncodeLbListener_DefaultAction_FixedResponse_StatusCode(p FixedResponse, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeLbListener_DefaultAction_Forward(p Forward, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_DefaultAction_Forward_Stickiness(p.Stickiness, ctyVal)
	EncodeLbListener_DefaultAction_Forward_TargetGroup(p.TargetGroup, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["forward"] = cty.ListVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_Forward_Stickiness(p Stickiness, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_DefaultAction_Forward_Stickiness_Duration(p, ctyVal)
	EncodeLbListener_DefaultAction_Forward_Stickiness_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["stickiness"] = cty.ListVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_Forward_Stickiness_Duration(p Stickiness, vals map[string]cty.Value) {
	vals["duration"] = cty.NumberIntVal(p.Duration)
}

func EncodeLbListener_DefaultAction_Forward_Stickiness_Enabled(p Stickiness, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeLbListener_DefaultAction_Forward_TargetGroup(p []TargetGroup, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLbListener_DefaultAction_Forward_TargetGroup_Arn(v, ctyVal)
		EncodeLbListener_DefaultAction_Forward_TargetGroup_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["target_group"] = cty.SetVal(valsForCollection)
}

func EncodeLbListener_DefaultAction_Forward_TargetGroup_Arn(p TargetGroup, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLbListener_DefaultAction_Forward_TargetGroup_Weight(p TargetGroup, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeLbListener_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeLbListener_Timeouts_Read(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeLbListener_Timeouts_Read(p Timeouts, vals map[string]cty.Value) {
	vals["read"] = cty.StringVal(p.Read)
}

func EncodeLbListener_Arn(p LbListenerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}