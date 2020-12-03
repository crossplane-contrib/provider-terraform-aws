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

package v1alpha1func EncodeRoute53QueryLog(r Route53QueryLog) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeRoute53QueryLog_CloudwatchLogGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeRoute53QueryLog_Id(r.Spec.ForProvider, ctyVal)
	EncodeRoute53QueryLog_ZoneId(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeRoute53QueryLog_CloudwatchLogGroupArn(p *Route53QueryLogParameters, vals map[string]cty.Value) {
	vals["cloudwatch_log_group_arn"] = cty.StringVal(p.CloudwatchLogGroupArn)
}

func EncodeRoute53QueryLog_Id(p *Route53QueryLogParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeRoute53QueryLog_ZoneId(p *Route53QueryLogParameters, vals map[string]cty.Value) {
	vals["zone_id"] = cty.StringVal(p.ZoneId)
}