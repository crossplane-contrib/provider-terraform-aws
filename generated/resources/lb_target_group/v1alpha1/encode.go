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

package v1alpha1func EncodeLbTargetGroup(r LbTargetGroup) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeLbTargetGroup_LoadBalancingAlgorithmType(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_LambdaMultiValueHeadersEnabled(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_DeregistrationDelay(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_Port(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_SlowStart(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_ProxyProtocolV2(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_TargetType(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeLbTargetGroup_HealthCheck(r.Spec.ForProvider.HealthCheck, ctyVal)
	EncodeLbTargetGroup_Stickiness(r.Spec.ForProvider.Stickiness, ctyVal)
	EncodeLbTargetGroup_ArnSuffix(r.Status.AtProvider, ctyVal)
	EncodeLbTargetGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeLbTargetGroup_LoadBalancingAlgorithmType(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["load_balancing_algorithm_type"] = cty.StringVal(p.LoadBalancingAlgorithmType)
}

func EncodeLbTargetGroup_NamePrefix(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeLbTargetGroup_Protocol(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeLbTargetGroup_LambdaMultiValueHeadersEnabled(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["lambda_multi_value_headers_enabled"] = cty.BoolVal(p.LambdaMultiValueHeadersEnabled)
}

func EncodeLbTargetGroup_DeregistrationDelay(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["deregistration_delay"] = cty.IntVal(p.DeregistrationDelay)
}

func EncodeLbTargetGroup_Id(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLbTargetGroup_Port(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["port"] = cty.IntVal(p.Port)
}

func EncodeLbTargetGroup_SlowStart(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["slow_start"] = cty.IntVal(p.SlowStart)
}

func EncodeLbTargetGroup_Tags(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeLbTargetGroup_Name(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLbTargetGroup_ProxyProtocolV2(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["proxy_protocol_v2"] = cty.BoolVal(p.ProxyProtocolV2)
}

func EncodeLbTargetGroup_TargetType(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["target_type"] = cty.StringVal(p.TargetType)
}

func EncodeLbTargetGroup_VpcId(p *LbTargetGroupParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeLbTargetGroup_HealthCheck(p *HealthCheck, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.HealthCheck {
		ctyVal = make(map[string]cty.Value)
		EncodeLbTargetGroup_HealthCheck_UnhealthyThreshold(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_Enabled(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_Interval(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_Matcher(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_Port(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_HealthyThreshold(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_Path(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_Protocol(v, ctyVal)
		EncodeLbTargetGroup_HealthCheck_Timeout(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["health_check"] = cty.ListVal(valsForCollection)
}

func EncodeLbTargetGroup_HealthCheck_UnhealthyThreshold(p *HealthCheck, vals map[string]cty.Value) {
	vals["unhealthy_threshold"] = cty.IntVal(p.UnhealthyThreshold)
}

func EncodeLbTargetGroup_HealthCheck_Enabled(p *HealthCheck, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeLbTargetGroup_HealthCheck_Interval(p *HealthCheck, vals map[string]cty.Value) {
	vals["interval"] = cty.IntVal(p.Interval)
}

func EncodeLbTargetGroup_HealthCheck_Matcher(p *HealthCheck, vals map[string]cty.Value) {
	vals["matcher"] = cty.StringVal(p.Matcher)
}

func EncodeLbTargetGroup_HealthCheck_Port(p *HealthCheck, vals map[string]cty.Value) {
	vals["port"] = cty.StringVal(p.Port)
}

func EncodeLbTargetGroup_HealthCheck_HealthyThreshold(p *HealthCheck, vals map[string]cty.Value) {
	vals["healthy_threshold"] = cty.IntVal(p.HealthyThreshold)
}

func EncodeLbTargetGroup_HealthCheck_Path(p *HealthCheck, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeLbTargetGroup_HealthCheck_Protocol(p *HealthCheck, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeLbTargetGroup_HealthCheck_Timeout(p *HealthCheck, vals map[string]cty.Value) {
	vals["timeout"] = cty.IntVal(p.Timeout)
}

func EncodeLbTargetGroup_Stickiness(p *Stickiness, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Stickiness {
		ctyVal = make(map[string]cty.Value)
		EncodeLbTargetGroup_Stickiness_CookieDuration(v, ctyVal)
		EncodeLbTargetGroup_Stickiness_Enabled(v, ctyVal)
		EncodeLbTargetGroup_Stickiness_Type(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["stickiness"] = cty.ListVal(valsForCollection)
}

func EncodeLbTargetGroup_Stickiness_CookieDuration(p *Stickiness, vals map[string]cty.Value) {
	vals["cookie_duration"] = cty.IntVal(p.CookieDuration)
}

func EncodeLbTargetGroup_Stickiness_Enabled(p *Stickiness, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeLbTargetGroup_Stickiness_Type(p *Stickiness, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeLbTargetGroup_ArnSuffix(p *LbTargetGroupObservation, vals map[string]cty.Value) {
	vals["arn_suffix"] = cty.StringVal(p.ArnSuffix)
}

func EncodeLbTargetGroup_Arn(p *LbTargetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}