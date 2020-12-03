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

package v1alpha1func EncodeFlowLog(r FlowLog) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeFlowLog_MaxAggregationInterval(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_EniId(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_Id(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogFormat(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogGroupName(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_Tags(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_TrafficType(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_IamRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogDestination(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogDestinationType(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeFlowLog_MaxAggregationInterval(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["max_aggregation_interval"] = cty.IntVal(p.MaxAggregationInterval)
}

func EncodeFlowLog_SubnetId(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeFlowLog_EniId(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["eni_id"] = cty.StringVal(p.EniId)
}

func EncodeFlowLog_Id(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeFlowLog_LogFormat(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["log_format"] = cty.StringVal(p.LogFormat)
}

func EncodeFlowLog_LogGroupName(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeFlowLog_Tags(p *FlowLogParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeFlowLog_TrafficType(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["traffic_type"] = cty.StringVal(p.TrafficType)
}

func EncodeFlowLog_VpcId(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeFlowLog_IamRoleArn(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["iam_role_arn"] = cty.StringVal(p.IamRoleArn)
}

func EncodeFlowLog_LogDestination(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["log_destination"] = cty.StringVal(p.LogDestination)
}

func EncodeFlowLog_LogDestinationType(p *FlowLogParameters, vals map[string]cty.Value) {
	vals["log_destination_type"] = cty.StringVal(p.LogDestinationType)
}

func EncodeFlowLog_Arn(p *FlowLogObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}