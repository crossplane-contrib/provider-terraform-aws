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
	r, ok := mr.(*CloudformationStackSet)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudformationStackSet(r, ctyValue)
}

func DecodeCloudformationStackSet(prev *CloudformationStackSet, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudformationStackSet_Tags(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_Capabilities(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_Description(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_Parameters(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_TemplateBody(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_TemplateUrl(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_AdministrationRoleArn(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_ExecutionRoleName(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_Name(&new.Spec.ForProvider, valMap)
	DecodeCloudformationStackSet_Timeouts(&new.Spec.ForProvider.Timeouts, valMap)
	DecodeCloudformationStackSet_StackSetId(&new.Status.AtProvider, valMap)
	DecodeCloudformationStackSet_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveMapTypeDecodeTemplate
func DecodeCloudformationStackSet_Tags(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveCollectionTypeDecodeTemplate
func DecodeCloudformationStackSet_Capabilities(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	goVals := make([]string, 0)
	for _, value := range ctwhy.ValueAsSet(vals["capabilities"]) {
		goVals = append(goVals, ctwhy.ValueAsString(value))
	}
	p.Capabilities = goVals
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_Description(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

//primitiveMapTypeDecodeTemplate
func DecodeCloudformationStackSet_Parameters(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["parameters"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Parameters = vMap
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_TemplateBody(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	p.TemplateBody = ctwhy.ValueAsString(vals["template_body"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_TemplateUrl(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	p.TemplateUrl = ctwhy.ValueAsString(vals["template_url"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_AdministrationRoleArn(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	p.AdministrationRoleArn = ctwhy.ValueAsString(vals["administration_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_ExecutionRoleName(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	p.ExecutionRoleName = ctwhy.ValueAsString(vals["execution_role_name"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_Name(p *CloudformationStackSetParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//containerTypeDecodeTemplate
func DecodeCloudformationStackSet_Timeouts(p *Timeouts, vals map[string]cty.Value) {
	valMap := vals["timeouts"].AsValueMap()
	DecodeCloudformationStackSet_Timeouts_Update(p, valMap)
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_Timeouts_Update(p *Timeouts, vals map[string]cty.Value) {
	p.Update = ctwhy.ValueAsString(vals["update"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_StackSetId(p *CloudformationStackSetObservation, vals map[string]cty.Value) {
	p.StackSetId = ctwhy.ValueAsString(vals["stack_set_id"])
}

//primitiveTypeDecodeTemplate
func DecodeCloudformationStackSet_Arn(p *CloudformationStackSetObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}