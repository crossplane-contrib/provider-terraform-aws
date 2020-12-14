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
	r, ok := mr.(*FlowLog)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a FlowLog.")
	}
	return EncodeFlowLog(*r), nil
}

func EncodeFlowLog(r FlowLog) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeFlowLog_VpcId(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_IamRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_Id(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogFormat(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogGroupName(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_MaxAggregationInterval(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_SubnetId(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_TrafficType(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_EniId(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogDestination(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_LogDestinationType(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_Tags(r.Spec.ForProvider, ctyVal)
	EncodeFlowLog_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeFlowLog_VpcId(p FlowLogParameters, vals map[string]cty.Value) {
	vals["vpc_id"] = cty.StringVal(p.VpcId)
}

func EncodeFlowLog_IamRoleArn(p FlowLogParameters, vals map[string]cty.Value) {
	vals["iam_role_arn"] = cty.StringVal(p.IamRoleArn)
}

func EncodeFlowLog_Id(p FlowLogParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeFlowLog_LogFormat(p FlowLogParameters, vals map[string]cty.Value) {
	vals["log_format"] = cty.StringVal(p.LogFormat)
}

func EncodeFlowLog_LogGroupName(p FlowLogParameters, vals map[string]cty.Value) {
	vals["log_group_name"] = cty.StringVal(p.LogGroupName)
}

func EncodeFlowLog_MaxAggregationInterval(p FlowLogParameters, vals map[string]cty.Value) {
	vals["max_aggregation_interval"] = cty.NumberIntVal(p.MaxAggregationInterval)
}

func EncodeFlowLog_SubnetId(p FlowLogParameters, vals map[string]cty.Value) {
	vals["subnet_id"] = cty.StringVal(p.SubnetId)
}

func EncodeFlowLog_TrafficType(p FlowLogParameters, vals map[string]cty.Value) {
	vals["traffic_type"] = cty.StringVal(p.TrafficType)
}

func EncodeFlowLog_EniId(p FlowLogParameters, vals map[string]cty.Value) {
	vals["eni_id"] = cty.StringVal(p.EniId)
}

func EncodeFlowLog_LogDestination(p FlowLogParameters, vals map[string]cty.Value) {
	vals["log_destination"] = cty.StringVal(p.LogDestination)
}

func EncodeFlowLog_LogDestinationType(p FlowLogParameters, vals map[string]cty.Value) {
	vals["log_destination_type"] = cty.StringVal(p.LogDestinationType)
}

func EncodeFlowLog_Tags(p FlowLogParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeFlowLog_Arn(p FlowLogObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}