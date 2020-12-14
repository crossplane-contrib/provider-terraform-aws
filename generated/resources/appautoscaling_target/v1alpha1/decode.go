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
	r, ok := mr.(*AppautoscalingTarget)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeAppautoscalingTarget(r, ctyValue)
}

func DecodeAppautoscalingTarget(prev *AppautoscalingTarget, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeAppautoscalingTarget_ServiceNamespace(&new.Spec.ForProvider, valMap)
	DecodeAppautoscalingTarget_Id(&new.Spec.ForProvider, valMap)
	DecodeAppautoscalingTarget_MaxCapacity(&new.Spec.ForProvider, valMap)
	DecodeAppautoscalingTarget_MinCapacity(&new.Spec.ForProvider, valMap)
	DecodeAppautoscalingTarget_ResourceId(&new.Spec.ForProvider, valMap)
	DecodeAppautoscalingTarget_RoleArn(&new.Spec.ForProvider, valMap)
	DecodeAppautoscalingTarget_ScalableDimension(&new.Spec.ForProvider, valMap)

	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeAppautoscalingTarget_ServiceNamespace(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	p.ServiceNamespace = ctwhy.ValueAsString(vals["service_namespace"])
}

func DecodeAppautoscalingTarget_Id(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeAppautoscalingTarget_MaxCapacity(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	p.MaxCapacity = ctwhy.ValueAsInt64(vals["max_capacity"])
}

func DecodeAppautoscalingTarget_MinCapacity(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	p.MinCapacity = ctwhy.ValueAsInt64(vals["min_capacity"])
}

func DecodeAppautoscalingTarget_ResourceId(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	p.ResourceId = ctwhy.ValueAsString(vals["resource_id"])
}

func DecodeAppautoscalingTarget_RoleArn(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

func DecodeAppautoscalingTarget_ScalableDimension(p *AppautoscalingTargetParameters, vals map[string]cty.Value) {
	p.ScalableDimension = ctwhy.ValueAsString(vals["scalable_dimension"])
}