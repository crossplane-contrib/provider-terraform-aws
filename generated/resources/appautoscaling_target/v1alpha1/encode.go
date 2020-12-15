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
	r, ok := mr.(*AppautoscalingTarget)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a AppautoscalingTarget.")
	}
	return EncodeAppautoscalingTarget(*r), nil
}

func EncodeAppautoscalingTarget(r AppautoscalingTarget) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeAppautoscalingTarget_MaxCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_MinCapacity(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_ResourceId(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_ScalableDimension(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_ServiceNamespace(r.Spec.ForProvider, ctyVal)
	EncodeAppautoscalingTarget_Id(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeAppautoscalingTarget_MaxCapacity(p AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["max_capacity"] = cty.NumberIntVal(p.MaxCapacity)
}

func EncodeAppautoscalingTarget_MinCapacity(p AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["min_capacity"] = cty.NumberIntVal(p.MinCapacity)
}

func EncodeAppautoscalingTarget_ResourceId(p AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["resource_id"] = cty.StringVal(p.ResourceId)
}

func EncodeAppautoscalingTarget_RoleArn(p AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeAppautoscalingTarget_ScalableDimension(p AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["scalable_dimension"] = cty.StringVal(p.ScalableDimension)
}

func EncodeAppautoscalingTarget_ServiceNamespace(p AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["service_namespace"] = cty.StringVal(p.ServiceNamespace)
}

func EncodeAppautoscalingTarget_Id(p AppautoscalingTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}