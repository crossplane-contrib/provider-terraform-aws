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

package v1alpha1func EncodeAppautoscalingTarget(r AppautoscalingTarget) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeAppautoscalingTarget_ScalableDimension(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_ServiceNamespace(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_MaxCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_MinCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_RoleArn(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeAppautoscalingTarget_ScalableDimension(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["scalable_dimension"] = cty.StringVal(p.ScalableDimension)
}

func EncodeAppautoscalingTarget_ServiceNamespace(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["service_namespace"] = cty.StringVal(p.ServiceNamespace)
}

func EncodeAppautoscalingTarget_Id(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeAppautoscalingTarget_MaxCapacity(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["max_capacity"] = cty.IntVal(p.MaxCapacity)
}

func EncodeAppautoscalingTarget_MinCapacity(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["min_capacity"] = cty.IntVal(p.MinCapacity)
}

func EncodeAppautoscalingTarget_ResourceId(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeAppautoscalingTarget_RoleArn(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}