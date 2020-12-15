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
	r, ok := mr.(*AutoscalingSchedule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeAutoscalingSchedule(r, ctyValue)
}

func DecodeAutoscalingSchedule(prev *AutoscalingSchedule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeAutoscalingSchedule_Id(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_Recurrence(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_StartTime(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_AutoscalingGroupName(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_EndTime(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_ScheduledActionName(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_DesiredCapacity(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_MaxSize(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_MinSize(&new.Spec.ForProvider, valMap)
	DecodeAutoscalingSchedule_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_Id(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_Recurrence(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.Recurrence = ctwhy.ValueAsString(vals["recurrence"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_StartTime(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.StartTime = ctwhy.ValueAsString(vals["start_time"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_AutoscalingGroupName(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.AutoscalingGroupName = ctwhy.ValueAsString(vals["autoscaling_group_name"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_EndTime(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.EndTime = ctwhy.ValueAsString(vals["end_time"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_ScheduledActionName(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.ScheduledActionName = ctwhy.ValueAsString(vals["scheduled_action_name"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_DesiredCapacity(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.DesiredCapacity = ctwhy.ValueAsInt64(vals["desired_capacity"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_MaxSize(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.MaxSize = ctwhy.ValueAsInt64(vals["max_size"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_MinSize(p *AutoscalingScheduleParameters, vals map[string]cty.Value) {
	p.MinSize = ctwhy.ValueAsInt64(vals["min_size"])
}

//primitiveTypeDecodeTemplate
func DecodeAutoscalingSchedule_Arn(p *AutoscalingScheduleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}