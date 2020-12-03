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

package v1alpha1func EncodeAppautoscalingScheduledAction(r AppautoscalingScheduledAction) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAppautoscalingScheduledAction_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_Schedule(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_EndTime(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_Name(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_ScalableDimension(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_ServiceNamespace(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_StartTime(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingScheduledAction_ScalableTargetAction(r.Spec.ForProvider.ScalableTargetAction, ctyVal)
	EncodeAppautoscalingScheduledAction_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeAppautoscalingScheduledAction_Id(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppautoscalingScheduledAction_Schedule(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["schedule"] = cty.StringVal(p.Schedule)
}

func EncodeAppautoscalingScheduledAction_EndTime(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["end_time"] = cty.StringVal(p.EndTime)
}

func EncodeAppautoscalingScheduledAction_Name(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeAppautoscalingScheduledAction_ResourceId(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeAppautoscalingScheduledAction_ScalableDimension(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["scalable_dimension"] = cty.StringVal(p.ScalableDimension)
}

func EncodeAppautoscalingScheduledAction_ServiceNamespace(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["service_namespace"] = cty.StringVal(p.ServiceNamespace)
}

func EncodeAppautoscalingScheduledAction_StartTime(p *AppautoscalingScheduledActionParameters, vals map[string]cty.Value) {
	vals["start_time"] = cty.StringVal(p.StartTime)
}

func EncodeAppautoscalingScheduledAction_ScalableTargetAction(p *ScalableTargetAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.ScalableTargetAction {
		ctyVal = make(map[string]cty.Value)
		EncodeAppautoscalingScheduledAction_ScalableTargetAction_MinCapacity(v, ctyVal)
		EncodeAppautoscalingScheduledAction_ScalableTargetAction_MaxCapacity(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["scalable_target_action"] = cty.ListVal(valsForCollection)
}

func EncodeAppautoscalingScheduledAction_ScalableTargetAction_MinCapacity(p *ScalableTargetAction, vals map[string]cty.Value) {
	vals["min_capacity"] = cty.IntVal(p.MinCapacity)
}

func EncodeAppautoscalingScheduledAction_ScalableTargetAction_MaxCapacity(p *ScalableTargetAction, vals map[string]cty.Value) {
	vals["max_capacity"] = cty.IntVal(p.MaxCapacity)
}

func EncodeAppautoscalingScheduledAction_Arn(p *AppautoscalingScheduledActionObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}