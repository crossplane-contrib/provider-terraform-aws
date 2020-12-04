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

func EncodeAlbTargetGroup(r AlbTargetGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAlbTargetGroup_LambdaMultiValueHeadersEnabled(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Protocol(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_DeregistrationDelay(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Port(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_SlowStart(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_ProxyProtocolV2(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_TargetType(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_LoadBalancingAlgorithmType(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeAlbTargetGroup_Stickiness(r.Spec.ForProvider.Stickiness, ctyVal)
	EncodeAlbTargetGroup_HealthCheck(r.Spec.ForProvider.HealthCheck, ctyVal)
	EncodeAlbTargetGroup_Arn(r.Status.AtProvider, ctyVal)
	EncodeAlbTargetGroup_ArnSuffix(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeAlbTargetGroup_LambdaMultiValueHeadersEnabled(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["lambda_multi_value_headers_enabled"] = cty.BoolVal(p.LambdaMultiValueHeadersEnabled)
}

func EncodeAlbTargetGroup_Protocol(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAlbTargetGroup_DeregistrationDelay(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["deregistration_delay"] = cty.NumberIntVal(p.DeregistrationDelay)
}

func EncodeAlbTargetGroup_Port(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["port"] = cty.NumberIntVal(p.Port)
}

func EncodeAlbTargetGroup_SlowStart(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["slow_start"] = cty.NumberIntVal(p.SlowStart)
}

func EncodeAlbTargetGroup_Name(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAlbTargetGroup_ProxyProtocolV2(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["proxy_protocol_v2"] = cty.BoolVal(p.ProxyProtocolV2)
}

func EncodeAlbTargetGroup_Tags(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeAlbTargetGroup_TargetType(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["target_type"] = cty.StringVal(p.TargetType)
}

func EncodeAlbTargetGroup_Id(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAlbTargetGroup_LoadBalancingAlgorithmType(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["load_balancing_algorithm_type"] = cty.StringVal(p.LoadBalancingAlgorithmType)
}

func EncodeAlbTargetGroup_NamePrefix(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeAlbTargetGroup_VpcId(p AlbTargetGroupParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeAlbTargetGroup_Stickiness(p Stickiness, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbTargetGroup_Stickiness_CookieDuration(p, ctyVal)
	EncodeAlbTargetGroup_Stickiness_Enabled(p, ctyVal)
	EncodeAlbTargetGroup_Stickiness_Type(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["stickiness"] = cty.ListVal(valsForCollection)
}

func EncodeAlbTargetGroup_Stickiness_CookieDuration(p Stickiness, vals map[string]cty.Value) {
	vals["cookie_duration"] = cty.NumberIntVal(p.CookieDuration)
}

func EncodeAlbTargetGroup_Stickiness_Enabled(p Stickiness, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAlbTargetGroup_Stickiness_Type(p Stickiness, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeAlbTargetGroup_HealthCheck(p HealthCheck, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeAlbTargetGroup_HealthCheck_Enabled(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_Interval(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_Matcher(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_Protocol(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_Timeout(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_HealthyThreshold(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_Path(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_Port(p, ctyVal)
	EncodeAlbTargetGroup_HealthCheck_UnhealthyThreshold(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["health_check"] = cty.ListVal(valsForCollection)
}

func EncodeAlbTargetGroup_HealthCheck_Enabled(p HealthCheck, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeAlbTargetGroup_HealthCheck_Interval(p HealthCheck, vals map[string]cty.Value) {
	vals["interval"] = cty.NumberIntVal(p.Interval)
}

func EncodeAlbTargetGroup_HealthCheck_Matcher(p HealthCheck, vals map[string]cty.Value) {
	vals["matcher"] = cty.StringVal(p.Matcher)
}

func EncodeAlbTargetGroup_HealthCheck_Protocol(p HealthCheck, vals map[string]cty.Value) {
	vals["protocol"] = cty.StringVal(p.Protocol)
}

func EncodeAlbTargetGroup_HealthCheck_Timeout(p HealthCheck, vals map[string]cty.Value) {
	vals["timeout"] = cty.NumberIntVal(p.Timeout)
}

func EncodeAlbTargetGroup_HealthCheck_HealthyThreshold(p HealthCheck, vals map[string]cty.Value) {
	vals["healthy_threshold"] = cty.NumberIntVal(p.HealthyThreshold)
}

func EncodeAlbTargetGroup_HealthCheck_Path(p HealthCheck, vals map[string]cty.Value) {
	vals["path"] = cty.StringVal(p.Path)
}

func EncodeAlbTargetGroup_HealthCheck_Port(p HealthCheck, vals map[string]cty.Value) {
	vals["port"] = cty.StringVal(p.Port)
}

func EncodeAlbTargetGroup_HealthCheck_UnhealthyThreshold(p HealthCheck, vals map[string]cty.Value) {
	vals["unhealthy_threshold"] = cty.NumberIntVal(p.UnhealthyThreshold)
}

func EncodeAlbTargetGroup_Arn(p AlbTargetGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeAlbTargetGroup_ArnSuffix(p AlbTargetGroupObservation, vals map[string]cty.Value) {
	vals["arn_suffix"] = cty.StringVal(p.ArnSuffix)
}