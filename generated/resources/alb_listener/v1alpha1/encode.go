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
	r, ok := mr.(*AlbListener)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AlbListener.")
	}
	return EncodeAlbListener(*r), nil
}

func EncodeAlbListener(r AlbListener) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_CertificateArn(r.Spec.ForProvider, ctyVal)
	EncodeAlbListener_Id(r.Spec.ForProvider, ctyVal)
	EncodeAlbListener_LoadBalancerArn(r.Spec.ForProvider, ctyVal)
	EncodeAlbListener_Port(r.Spec.ForProvider, ctyVal)
	EncodeAlbListener_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeAlbListener_SslPolicy(r.Spec.ForProvider, ctyVal)
	EncodeAlbListener_DefaultAction(r.Spec.ForProvider.DefaultAction, ctyVal)
	EncodeAlbListener_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeAlbListener_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeAlbListener_CertificateArn(p AlbListenerParameters, vals map[string]cty.Value) {
	vals["certificate_arn"] = cty.StringVal(p.CertificateArn)
}

func EncodeAlbListener_Id(p AlbListenerParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAlbListener_LoadBalancerArn(p AlbListenerParameters, vals map[string]cty.Value) {
	vals["load_balancer_arn"] = cty.StringVal(p.LoadBalancerArn)
}

func EncodeAlbListener_Port(p AlbListenerParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeAlbListener_Protocol(p AlbListenerParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAlbListener_SslPolicy(p AlbListenerParameters, vals map[string]cty.Value) {
	vals["ssl_policy"] = cty.StringVal(p.SslPolicy)
}

func EncodeAlbListener_DefaultAction(p []DefaultAction, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAlbListener_DefaultAction_Order(v, ctyVal)
		EncodeAlbListener_DefaultAction_TargetGroupArn(v, ctyVal)
		EncodeAlbListener_DefaultAction_Type(v, ctyVal)
		EncodeAlbListener_DefaultAction_AuthenticateCognito(v.AuthenticateCognito, ctyVal)
		EncodeAlbListener_DefaultAction_AuthenticateOidc(v.AuthenticateOidc, ctyVal)
		EncodeAlbListener_DefaultAction_FixedResponse(v.FixedResponse, ctyVal)
		EncodeAlbListener_DefaultAction_Forward(v.Forward, ctyVal)
		EncodeAlbListener_DefaultAction_Redirect(v.Redirect, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["default_action"] = cty.ListVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_Order(p DefaultAction, vals map[string]cty.Value) {
	vals["order"] = cty.NumberIntVal(p.Order)
}

func EncodeAlbListener_DefaultAction_TargetGroupArn(p DefaultAction, vals map[string]cty.Value) {
	vals["target_group_arn"] = cty.StringVal(p.TargetGroupArn)
}

func EncodeAlbListener_DefaultAction_Type(p DefaultAction, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito(p AuthenticateCognito, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_Scope(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_SessionCookieName(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_SessionTimeout(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_UserPoolArn(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_UserPoolClientId(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_UserPoolDomain(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_AuthenticationRequestExtraParams(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateCognito_OnUnauthenticatedRequest(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authenticate_cognito"] = cty.ListVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_Scope(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_SessionCookieName(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["session_cookie_name"] = cty.StringVal(p.SessionCookieName)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_SessionTimeout(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["session_timeout"] = cty.NumberIntVal(p.SessionTimeout)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_UserPoolArn(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_arn"] = cty.StringVal(p.UserPoolArn)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_UserPoolClientId(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_client_id"] = cty.StringVal(p.UserPoolClientId)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_UserPoolDomain(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["user_pool_domain"] = cty.StringVal(p.UserPoolDomain)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_AuthenticationRequestExtraParams(p AuthenticateCognito, vals map[string]cty.Value) {
	if len(p.AuthenticationRequestExtraParams) == 0 {
		vals["authentication_request_extra_params"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.AuthenticationRequestExtraParams {
		mVals[key] = cty.StringVal(value)
	}
	vals["authentication_request_extra_params"] = cty.MapVal(mVals)
}

func EncodeAlbListener_DefaultAction_AuthenticateCognito_OnUnauthenticatedRequest(p AuthenticateCognito, vals map[string]cty.Value) {
	vals["on_unauthenticated_request"] = cty.StringVal(p.OnUnauthenticatedRequest)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc(p AuthenticateOidc, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_TokenEndpoint(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_UserInfoEndpoint(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_AuthorizationEndpoint(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_ClientSecret(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_Issuer(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_OnUnauthenticatedRequest(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_SessionCookieName(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_AuthenticationRequestExtraParams(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_ClientId(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_Scope(p, ctyVal)
	EncodeAlbListener_DefaultAction_AuthenticateOidc_SessionTimeout(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["authenticate_oidc"] = cty.ListVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_TokenEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["token_endpoint"] = cty.StringVal(p.TokenEndpoint)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_UserInfoEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["user_info_endpoint"] = cty.StringVal(p.UserInfoEndpoint)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_AuthorizationEndpoint(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["authorization_endpoint"] = cty.StringVal(p.AuthorizationEndpoint)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_ClientSecret(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["client_secret"] = cty.StringVal(p.ClientSecret)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_Issuer(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["issuer"] = cty.StringVal(p.Issuer)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_OnUnauthenticatedRequest(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["on_unauthenticated_request"] = cty.StringVal(p.OnUnauthenticatedRequest)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_SessionCookieName(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["session_cookie_name"] = cty.StringVal(p.SessionCookieName)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_AuthenticationRequestExtraParams(p AuthenticateOidc, vals map[string]cty.Value) {
	if len(p.AuthenticationRequestExtraParams) == 0 {
		vals["authentication_request_extra_params"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.AuthenticationRequestExtraParams {
		mVals[key] = cty.StringVal(value)
	}
	vals["authentication_request_extra_params"] = cty.MapVal(mVals)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_ClientId(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["client_id"] = cty.StringVal(p.ClientId)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_Scope(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeAlbListener_DefaultAction_AuthenticateOidc_SessionTimeout(p AuthenticateOidc, vals map[string]cty.Value) {
	vals["session_timeout"] = cty.NumberIntVal(p.SessionTimeout)
}

func EncodeAlbListener_DefaultAction_FixedResponse(p FixedResponse, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_DefaultAction_FixedResponse_MessageBody(p, ctyVal)
	EncodeAlbListener_DefaultAction_FixedResponse_StatusCode(p, ctyVal)
	EncodeAlbListener_DefaultAction_FixedResponse_ContentType(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["fixed_response"] = cty.ListVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_FixedResponse_MessageBody(p FixedResponse, vals map[string]cty.Value) {
	vals["message_body"] = cty.StringVal(p.MessageBody)
}

func EncodeAlbListener_DefaultAction_FixedResponse_StatusCode(p FixedResponse, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeAlbListener_DefaultAction_FixedResponse_ContentType(p FixedResponse, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeAlbListener_DefaultAction_Forward(p Forward, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_DefaultAction_Forward_Stickiness(p.Stickiness, ctyVal)
	EncodeAlbListener_DefaultAction_Forward_TargetGroup(p.TargetGroup, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["forward"] = cty.ListVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_Forward_Stickiness(p Stickiness, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_DefaultAction_Forward_Stickiness_Duration(p, ctyVal)
	EncodeAlbListener_DefaultAction_Forward_Stickiness_Enabled(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["stickiness"] = cty.ListVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_Forward_Stickiness_Duration(p Stickiness, vals map[string]cty.Value) {
	vals["duration"] = cty.NumberIntVal(p.Duration)
}

func EncodeAlbListener_DefaultAction_Forward_Stickiness_Enabled(p Stickiness, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAlbListener_DefaultAction_Forward_TargetGroup(p []TargetGroup, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeAlbListener_DefaultAction_Forward_TargetGroup_Arn(v, ctyVal)
		EncodeAlbListener_DefaultAction_Forward_TargetGroup_Weight(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["target_group"] = cty.SetVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_Forward_TargetGroup_Arn(p TargetGroup, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAlbListener_DefaultAction_Forward_TargetGroup_Weight(p TargetGroup, vals map[string]cty.Value) {
	vals["weight"] = cty.NumberIntVal(p.Weight)
}

func EncodeAlbListener_DefaultAction_Redirect(p Redirect, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_DefaultAction_Redirect_Host(p, ctyVal)
	EncodeAlbListener_DefaultAction_Redirect_Path(p, ctyVal)
	EncodeAlbListener_DefaultAction_Redirect_Port(p, ctyVal)
	EncodeAlbListener_DefaultAction_Redirect_Protocol(p, ctyVal)
	EncodeAlbListener_DefaultAction_Redirect_Query(p, ctyVal)
	EncodeAlbListener_DefaultAction_Redirect_StatusCode(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["redirect"] = cty.ListVal(valsForCollection)
}

func EncodeAlbListener_DefaultAction_Redirect_Host(p Redirect, vals map[string]cty.Value) {
	vals["host"] = cty.StringVal(p.Host)
}

func EncodeAlbListener_DefaultAction_Redirect_Path(p Redirect, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeAlbListener_DefaultAction_Redirect_Port(p Redirect, vals map[string]cty.Value) {
	vals["port"] = cty.StringVal(p.Port)
}

func EncodeAlbListener_DefaultAction_Redirect_Protocol(p Redirect, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAlbListener_DefaultAction_Redirect_Query(p Redirect, vals map[string]cty.Value) {
	vals["query"] = cty.StringVal(p.Query)
}

func EncodeAlbListener_DefaultAction_Redirect_StatusCode(p Redirect, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeAlbListener_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeAlbListener_Timeouts_Read(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeAlbListener_Timeouts_Read(p Timeouts, vals map[string]cty.Value) {
	vals["read"] = cty.StringVal(p.Read)
}

func EncodeAlbListener_Arn(p AlbListenerObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}