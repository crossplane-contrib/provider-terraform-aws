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

package v1alpha1func EncodeDbProxy(r DbProxy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDbProxy_Id(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_DebugLogging(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_EngineFamily(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_IdleClientTimeout(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_Name(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_RequireTls(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_VpcSecurityGroupIds(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_VpcSubnetIds(r.Spec.ForProvider, ctyVal)
	EncodeDbProxy_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeDbProxy_Auth(r.Spec.ForProvider.Auth, ctyVal)
	EncodeDbProxy_Arn(r.Status.AtProvider, ctyVal)
	EncodeDbProxy_Endpoint(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDbProxy_Id(p *DbProxyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDbProxy_Tags(p *DbProxyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDbProxy_RoleArn(p *DbProxyParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeDbProxy_DebugLogging(p *DbProxyParameters, vals map[string]cty.Value) {
	vals["debug_logging"] = cty.BoolVal(p.DebugLogging)
}

func EncodeDbProxy_EngineFamily(p *DbProxyParameters, vals map[string]cty.Value) {
	vals["engine_family"] = cty.StringVal(p.EngineFamily)
}

func EncodeDbProxy_IdleClientTimeout(p *DbProxyParameters, vals map[string]cty.Value) {
	vals["idle_client_timeout"] = cty.IntVal(p.IdleClientTimeout)
}

func EncodeDbProxy_Name(p *DbProxyParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDbProxy_RequireTls(p *DbProxyParameters, vals map[string]cty.Value) {
	vals["require_tls"] = cty.BoolVal(p.RequireTls)
}

func EncodeDbProxy_VpcSecurityGroupIds(p *DbProxyParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSecurityGroupIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_security_group_ids"] = cty.SetVal(colVals)
}

func EncodeDbProxy_VpcSubnetIds(p *DbProxyParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.VpcSubnetIds {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["vpc_subnet_ids"] = cty.SetVal(colVals)
}

func EncodeDbProxy_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	ctyVal = make(map[string]cty.Value)
	EncodeDbProxy_Timeouts_Create(p, ctyVal)
	EncodeDbProxy_Timeouts_Delete(p, ctyVal)
	EncodeDbProxy_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeDbProxy_Timeouts_Create(p *Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeDbProxy_Timeouts_Delete(p *Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeDbProxy_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeDbProxy_Auth(p *Auth, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Auth {
		ctyVal = make(map[string]cty.Value)
		EncodeDbProxy_Auth_Description(v, ctyVal)
		EncodeDbProxy_Auth_IamAuth(v, ctyVal)
		EncodeDbProxy_Auth_SecretArn(v, ctyVal)
		EncodeDbProxy_Auth_AuthScheme(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["auth"] = cty.SetVal(valsForCollection)
}

func EncodeDbProxy_Auth_Description(p *Auth, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDbProxy_Auth_IamAuth(p *Auth, vals map[string]cty.Value) {
	vals["iam_auth"] = cty.StringVal(p.IamAuth)
}

func EncodeDbProxy_Auth_SecretArn(p *Auth, vals map[string]cty.Value) {
	vals["secret_arn"] = cty.StringVal(p.SecretArn)
}

func EncodeDbProxy_Auth_AuthScheme(p *Auth, vals map[string]cty.Value) {
	vals["auth_scheme"] = cty.StringVal(p.AuthScheme)
}

func EncodeDbProxy_Arn(p *DbProxyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeDbProxy_Endpoint(p *DbProxyObservation, vals map[string]cty.Value) {
	vals["endpoint"] = cty.StringVal(p.Endpoint)
}