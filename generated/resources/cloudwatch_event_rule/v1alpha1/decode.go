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
	r, ok := mr.(*CloudwatchEventRule)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeCloudwatchEventRule(r, ctyValue)
}

func DecodeCloudwatchEventRule(prev *CloudwatchEventRule, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeCloudwatchEventRule_Description(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_EventPattern(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_ScheduleExpression(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_Tags(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_Id(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_IsEnabled(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_Name(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_RoleArn(&new.Spec.ForProvider, valMap)
	DecodeCloudwatchEventRule_Arn(&new.Status.AtProvider, valMap)
	meta.SetExternalName(new, valMap["id"].AsString())
	return new, nil
}

func DecodeCloudwatchEventRule_Description(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.Description = ctwhy.ValueAsString(vals["description"])
}

func DecodeCloudwatchEventRule_EventPattern(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.EventPattern = ctwhy.ValueAsString(vals["event_pattern"])
}

func DecodeCloudwatchEventRule_NamePrefix(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

func DecodeCloudwatchEventRule_ScheduleExpression(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.ScheduleExpression = ctwhy.ValueAsString(vals["schedule_expression"])
}

func DecodeCloudwatchEventRule_Tags(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

func DecodeCloudwatchEventRule_Id(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

func DecodeCloudwatchEventRule_IsEnabled(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.IsEnabled = ctwhy.ValueAsBool(vals["is_enabled"])
}

func DecodeCloudwatchEventRule_Name(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

func DecodeCloudwatchEventRule_RoleArn(p *CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	p.RoleArn = ctwhy.ValueAsString(vals["role_arn"])
}

func DecodeCloudwatchEventRule_Arn(p *CloudwatchEventRuleObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}