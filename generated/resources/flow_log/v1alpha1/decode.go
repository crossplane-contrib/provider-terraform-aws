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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*FlowLog)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeFlowLog(r, ctyValue)
}

func DecodeFlowLog(prev *FlowLog, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeFlowLog_LogDestination(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_IamRoleArn(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_LogGroupName(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_MaxAggregationInterval(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_SubnetId(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_Tags(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_TrafficType(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_EniId(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_LogDestinationType(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_LogFormat(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_VpcId(&new.Spec.ForProvider, valMap)
	DecodeFlowLog_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_LogDestination(p *FlowLogParameters, vals map[string]cty.Value) {
	p.LogDestination = ctwhy.ValueAsString(vals["log_destination"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_IamRoleArn(p *FlowLogParameters, vals map[string]cty.Value) {
	p.IamRoleArn = ctwhy.ValueAsString(vals["iam_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_LogGroupName(p *FlowLogParameters, vals map[string]cty.Value) {
	p.LogGroupName = ctwhy.ValueAsString(vals["log_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_MaxAggregationInterval(p *FlowLogParameters, vals map[string]cty.Value) {
	p.MaxAggregationInterval = ctwhy.ValueAsInt64(vals["max_aggregation_interval"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_SubnetId(p *FlowLogParameters, vals map[string]cty.Value) {
	p.SubnetId = ctwhy.ValueAsString(vals["subnet_id"])
}

//primitiveMapTypeDecodeTemplate
func DecodeFlowLog_Tags(p *FlowLogParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_TrafficType(p *FlowLogParameters, vals map[string]cty.Value) {
	p.TrafficType = ctwhy.ValueAsString(vals["traffic_type"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_EniId(p *FlowLogParameters, vals map[string]cty.Value) {
	p.EniId = ctwhy.ValueAsString(vals["eni_id"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_LogDestinationType(p *FlowLogParameters, vals map[string]cty.Value) {
	p.LogDestinationType = ctwhy.ValueAsString(vals["log_destination_type"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_LogFormat(p *FlowLogParameters, vals map[string]cty.Value) {
	p.LogFormat = ctwhy.ValueAsString(vals["log_format"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_VpcId(p *FlowLogParameters, vals map[string]cty.Value) {
	p.VpcId = ctwhy.ValueAsString(vals["vpc_id"])
}

//primitiveTypeDecodeTemplate
func DecodeFlowLog_Arn(p *FlowLogObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}