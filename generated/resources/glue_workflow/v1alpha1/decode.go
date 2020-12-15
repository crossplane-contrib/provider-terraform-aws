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
	r, ok := mr.(*GlueWorkflow)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeGlueWorkflow(r, ctyValue)
}

func DecodeGlueWorkflow(prev *GlueWorkflow, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeGlueWorkflow_MaxConcurrentRuns(&new.Spec.ForProvider, valMap)
	DecodeGlueWorkflow_Name(&new.Spec.ForProvider, valMap)
	DecodeGlueWorkflow_Tags(&new.Spec.ForProvider, valMap)
	DecodeGlueWorkflow_DefaultRunProperties(&new.Spec.ForProvider, valMap)
	DecodeGlueWorkflow_Description(&new.Spec.ForProvider, valMap)
	DecodeGlueWorkflow_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeGlueWorkflow_MaxConcurrentRuns(p *GlueWorkflowParameters, vals map[string]cty.Value) {
	p.MaxConcurrentRuns = ctwhy.ValueAsInt64(vals["max_concurrent_runs"])
}

//primitiveTypeDecodeTemplate
func DecodeGlueWorkflow_Name(p *GlueWorkflowParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveMapTypeDecodeTemplate
func DecodeGlueWorkflow_Tags(p *GlueWorkflowParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveMapTypeDecodeTemplate
func DecodeGlueWorkflow_DefaultRunProperties(p *GlueWorkflowParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["default_run_properties"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.DefaultRunProperties = vMap
}

//primitiveTypeDecodeTemplate
func DecodeGlueWorkflow_Description(p *GlueWorkflowParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveTypeDecodeTemplate
func DecodeGlueWorkflow_Arn(p *GlueWorkflowObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}