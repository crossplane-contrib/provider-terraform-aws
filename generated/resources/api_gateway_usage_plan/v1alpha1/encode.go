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

package v1alpha1func EncodeApiGatewayUsagePlan(r ApiGatewayUsagePlan) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeApiGatewayUsagePlan_Description(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlan_Id(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlan_Name(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlan_ProductCode(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlan_Tags(r.Spec.ForProvider, ctyVal)
	EncodeApiGatewayUsagePlan_ApiStages(r.Spec.ForProvider.ApiStages, ctyVal)
	EncodeApiGatewayUsagePlan_QuotaSettings(r.Spec.ForProvider.QuotaSettings, ctyVal)
	EncodeApiGatewayUsagePlan_ThrottleSettings(r.Spec.ForProvider.ThrottleSettings, ctyVal)
	EncodeApiGatewayUsagePlan_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeApiGatewayUsagePlan_Description(p *ApiGatewayUsagePlanParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeApiGatewayUsagePlan_Id(p *ApiGatewayUsagePlanParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeApiGatewayUsagePlan_Name(p *ApiGatewayUsagePlanParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeApiGatewayUsagePlan_ProductCode(p *ApiGatewayUsagePlanParameters, vals map[string]cty.Value) {
	vals["product_code"] = cty.StringVal(p.ProductCode)
}

func EncodeApiGatewayUsagePlan_Tags(p *ApiGatewayUsagePlanParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeApiGatewayUsagePlan_ApiStages(p *ApiStages, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ApiStages {
		ctyVal = make(map[string]cty.Value)
		EncodeApiGatewayUsagePlan_ApiStages_ApiId(v, ctyVal)
		EncodeApiGatewayUsagePlan_ApiStages_Stage(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["api_stages"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayUsagePlan_ApiStages_ApiId(p *ApiStages, vals map[string]cty.Value) {
	vals["api_id"] = cty.StringVal(p.ApiId)
}

func EncodeApiGatewayUsagePlan_ApiStages_Stage(p *ApiStages, vals map[string]cty.Value) {
	vals["stage"] = cty.StringVal(p.Stage)
}

func EncodeApiGatewayUsagePlan_QuotaSettings(p *QuotaSettings, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.QuotaSettings {
		ctyVal = make(map[string]cty.Value)
		EncodeApiGatewayUsagePlan_QuotaSettings_Limit(v, ctyVal)
		EncodeApiGatewayUsagePlan_QuotaSettings_Offset(v, ctyVal)
		EncodeApiGatewayUsagePlan_QuotaSettings_Period(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["quota_settings"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayUsagePlan_QuotaSettings_Limit(p *QuotaSettings, vals map[string]cty.Value) {
	vals["limit"] = cty.IntVal(p.Limit)
}

func EncodeApiGatewayUsagePlan_QuotaSettings_Offset(p *QuotaSettings, vals map[string]cty.Value) {
	vals["offset"] = cty.IntVal(p.Offset)
}

func EncodeApiGatewayUsagePlan_QuotaSettings_Period(p *QuotaSettings, vals map[string]cty.Value) {
	vals["period"] = cty.StringVal(p.Period)
}

func EncodeApiGatewayUsagePlan_ThrottleSettings(p *ThrottleSettings, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ThrottleSettings {
		ctyVal = make(map[string]cty.Value)
		EncodeApiGatewayUsagePlan_ThrottleSettings_BurstLimit(v, ctyVal)
		EncodeApiGatewayUsagePlan_ThrottleSettings_RateLimit(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["throttle_settings"] = cty.ListVal(valsForCollection)
}

func EncodeApiGatewayUsagePlan_ThrottleSettings_BurstLimit(p *ThrottleSettings, vals map[string]cty.Value) {
	vals["burst_limit"] = cty.IntVal(p.BurstLimit)
}

func EncodeApiGatewayUsagePlan_ThrottleSettings_RateLimit(p *ThrottleSettings, vals map[string]cty.Value) {
	vals["rate_limit"] = cty.IntVal(p.RateLimit)
}

func EncodeApiGatewayUsagePlan_Arn(p *ApiGatewayUsagePlanObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}