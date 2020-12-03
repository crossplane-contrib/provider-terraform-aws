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

package v1alpha1func EncodeApiGatewayAccount(r ApiGatewayAccount) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayAccount_CloudwatchRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAccount_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayAccount_ThrottleSettings(r.Status.AtProvider.ThrottleSettings, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayAccount_CloudwatchRoleArn(p *ApiGatewayAccountParameters, vals map[string]cty.Value) {
	vals["cloudwatch_role_arn"] = cty.StringVal(p.CloudwatchRoleArn)
}

func EncodeApiGatewayAccount_Id(p *ApiGatewayAccountParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayAccount_ThrottleSettings(p *ThrottleSettings, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ThrottleSettings {
		ctyVal = make(map[string]cty.Value)
		EncodeApiGatewayAccount_ThrottleSettings_BurstLimit(v, ctyVal)
		EncodeApiGatewayAccount_ThrottleSettings_RateLimit(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["throttle_settings"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayAccount_ThrottleSettings_BurstLimit(p *ThrottleSettings, vals map[string]cty.Value) {
	vals["burst_limit"] = cty.IntVal(p.BurstLimit)
}

func EncodeApiGatewayAccount_ThrottleSettings_RateLimit(p *ThrottleSettings, vals map[string]cty.Value) {
	vals["rate_limit"] = cty.IntVal(p.RateLimit)
}