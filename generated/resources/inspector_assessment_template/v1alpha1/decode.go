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
	r, ok := mr.(*InspectorAssessmentTemplate)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeInspectorAssessmentTemplate(r, ctyValue)
}

func DecodeInspectorAssessmentTemplate(prev *InspectorAssessmentTemplate, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeInspectorAssessmentTemplate_Duration(&new.Spec.ForProvider, valMap)
	DecodeInspectorAssessmentTemplate_Id(&new.Spec.ForProvider, valMap)
	DecodeInspectorAssessmentTemplate_Name(&new.Spec.ForProvider, valMap)
	DecodeInspectorAssessmentTemplate_RulesPackageArns(&new.Spec.ForProvider, valMap)
	DecodeInspectorAssessmentTemplate_Tags(&new.Spec.ForProvider, valMap)
	DecodeInspectorAssessmentTemplate_TargetArn(&new.Spec.ForProvider, valMap)
	DecodeInspectorAssessmentTemplate_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeInspectorAssessmentTemplate_Duration(p *InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	p.Duration = ctwhy.ValueAsInt64(vals["duration"])
}

//primitiveTypeDecodeTemplate
func DecodeInspectorAssessmentTemplate_Id(p *InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeInspectorAssessmentTemplate_Name(p *InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveCollectionTypeDecodeTemplate
func DecodeInspectorAssessmentTemplate_RulesPackageArns(p *InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["rules_package_arns"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.RulesPackageArns = goVals
}

//primitiveMapTypeDecodeTemplate
func DecodeInspectorAssessmentTemplate_Tags(p *InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeInspectorAssessmentTemplate_TargetArn(p *InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	p.TargetArn = ctwhy.ValueAsString(vals["target_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeInspectorAssessmentTemplate_Arn(p *InspectorAssessmentTemplateObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}