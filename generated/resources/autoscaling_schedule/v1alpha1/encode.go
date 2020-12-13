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
	r, ok := mr.(*AutoscalingSchedule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AutoscalingSchedule.")
	}
	return EncodeAutoscalingSchedule(*r), nil
}

func EncodeAutoscalingSchedule(r AutoscalingSchedule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAutoscalingSchedule_AutoscalingGroupName(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_EndTime(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_MaxSize(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_Recurrence(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_ScheduledActionName(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_DesiredCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_Id(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_MinSize(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_StartTime(r.Spec.ForProvider, ctyVal)
	EncodeAutoscalingSchedule_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeAutoscalingSchedule_AutoscalingGroupName(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["autoscaling_group_name"] = cty.StringVal(p.AutoscalingGroupName)
}

func EncodeAutoscalingSchedule_EndTime(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["end_time"] = cty.StringVal(p.EndTime)
}

func EncodeAutoscalingSchedule_MaxSize(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["max_size"] = cty.NumberIntVal(p.MaxSize)
}

func EncodeAutoscalingSchedule_Recurrence(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["recurrence"] = cty.StringVal(p.Recurrence)
}

func EncodeAutoscalingSchedule_ScheduledActionName(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["scheduled_action_name"] = cty.StringVal(p.ScheduledActionName)
}

func EncodeAutoscalingSchedule_DesiredCapacity(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["desired_capacity"] = cty.NumberIntVal(p.DesiredCapacity)
}

func EncodeAutoscalingSchedule_Id(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAutoscalingSchedule_MinSize(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["min_size"] = cty.NumberIntVal(p.MinSize)
}

func EncodeAutoscalingSchedule_StartTime(p AutoscalingScheduleParameters, vals map[string]cty.Value) {
	vals["start_time"] = cty.StringVal(p.StartTime)
}

func EncodeAutoscalingSchedule_Arn(p AutoscalingScheduleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}