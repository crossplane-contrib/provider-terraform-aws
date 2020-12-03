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

package v1alpha1func EncodeAlbTargetGroup(r AlbTargetGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAlbTargetGroup_ProxyProtocolV2(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_TargetType(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_LoadBalancingAlgorithmType(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_LambdaMultiValueHeadersEnabled(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Port(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_DeregistrationDelay(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_SlowStart(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_HealthCheck(r.Spec.ForProvider.HealthCheck, ctyVal)
	EncodeAlbTargetGroup_Stickiness(r.Spec.ForProvider.Stickiness, ctyVal)
	EncodeAlbTargetGroup_ArnSuffix(r.Status.AtProvider, ctyVal)
	EncodeAlbTargetGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAlbTargetGroup_ProxyProtocolV2(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["proxy_protocol_v2"] = cty.BoolVal(p.ProxyProtocolV2)
}

func EncodeAlbTargetGroup_Tags(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAlbTargetGroup_TargetType(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["target_type"] = cty.StringVal(p.TargetType)
}

func EncodeAlbTargetGroup_Id(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAlbTargetGroup_LoadBalancingAlgorithmType(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["load_balancing_algorithm_type"] = cty.StringVal(p.LoadBalancingAlgorithmType)
}

func EncodeAlbTargetGroup_Protocol(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAlbTargetGroup_LambdaMultiValueHeadersEnabled(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["lambda_multi_value_headers_enabled"] = cty.BoolVal(p.LambdaMultiValueHeadersEnabled)
}

func EncodeAlbTargetGroup_Name(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAlbTargetGroup_NamePrefix(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeAlbTargetGroup_Port(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["port"] = cty.IntVal(p.Port)
}

func EncodeAlbTargetGroup_VpcId(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeAlbTargetGroup_DeregistrationDelay(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["deregistration_delay"] = cty.IntVal(p.DeregistrationDelay)
}

func EncodeAlbTargetGroup_SlowStart(p *AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["slow_start"] = cty.IntVal(p.SlowStart)
}

func EncodeAlbTargetGroup_HealthCheck(p *HealthCheck, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.HealthCheck {
		ctyVal = make(map[string]cty.Value)
		EncodeAlbTargetGroup_HealthCheck_Path(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_Enabled(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_Interval(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_Matcher(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_Port(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_Protocol(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_Timeout(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_UnhealthyThreshold(v, ctyVal)
		EncodeAlbTargetGroup_HealthCheck_HealthyThreshold(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["health_check"] = cty.ListVal(valsForCollection)
}

func EncodeAlbTargetGroup_HealthCheck_Path(p *HealthCheck, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeAlbTargetGroup_HealthCheck_Enabled(p *HealthCheck, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAlbTargetGroup_HealthCheck_Interval(p *HealthCheck, vals map[string]cty.Value) {
	vals["interval"] = cty.IntVal(p.Interval)
}

func EncodeAlbTargetGroup_HealthCheck_Matcher(p *HealthCheck, vals map[string]cty.Value) {
	vals["matcher"] = cty.StringVal(p.Matcher)
}

func EncodeAlbTargetGroup_HealthCheck_Port(p *HealthCheck, vals map[string]cty.Value) {
	vals["port"] = cty.StringVal(p.Port)
}

func EncodeAlbTargetGroup_HealthCheck_Protocol(p *HealthCheck, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAlbTargetGroup_HealthCheck_Timeout(p *HealthCheck, vals map[string]cty.Value) {
	vals["timeout"] = cty.IntVal(p.Timeout)
}

func EncodeAlbTargetGroup_HealthCheck_UnhealthyThreshold(p *HealthCheck, vals map[string]cty.Value) {
	vals["unhealthy_threshold"] = cty.IntVal(p.UnhealthyThreshold)
}

func EncodeAlbTargetGroup_HealthCheck_HealthyThreshold(p *HealthCheck, vals map[string]cty.Value) {
	vals["healthy_threshold"] = cty.IntVal(p.HealthyThreshold)
}

func EncodeAlbTargetGroup_Stickiness(p *Stickiness, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Stickiness {
		ctyVal = make(map[string]cty.Value)
		EncodeAlbTargetGroup_Stickiness_CookieDuration(v, ctyVal)
		EncodeAlbTargetGroup_Stickiness_Enabled(v, ctyVal)
		EncodeAlbTargetGroup_Stickiness_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["stickiness"] = cty.ListVal(valsForCollection)
}

func EncodeAlbTargetGroup_Stickiness_CookieDuration(p *Stickiness, vals map[string]cty.Value) {
	vals["cookie_duration"] = cty.IntVal(p.CookieDuration)
}

func EncodeAlbTargetGroup_Stickiness_Enabled(p *Stickiness, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAlbTargetGroup_Stickiness_Type(p *Stickiness, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAlbTargetGroup_ArnSuffix(p *AlbTargetGroupObservation, vals map[string]cty.Value) {
	vals["arn_suffix"] = cty.StringVal(p.ArnSuffix)
}

func EncodeAlbTargetGroup_Arn(p *AlbTargetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}