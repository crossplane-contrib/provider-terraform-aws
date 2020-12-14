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
	r, ok := mr.(*CloudwatchEventRule)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a CloudwatchEventRule.")
	}
	return EncodeCloudwatchEventRule(*r), nil
}

func EncodeCloudwatchEventRule(r CloudwatchEventRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchEventRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_EventPattern(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_ScheduleExpression(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_IsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchEventRule_Description(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeCloudwatchEventRule_EventPattern(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["event_pattern"] = cty.StringVal(p.EventPattern)
}

func EncodeCloudwatchEventRule_NamePrefix(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeCloudwatchEventRule_ScheduleExpression(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["schedule_expression"] = cty.StringVal(p.ScheduleExpression)
}

func EncodeCloudwatchEventRule_Tags(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudwatchEventRule_Id(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchEventRule_IsEnabled(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["is_enabled"] = cty.BoolVal(p.IsEnabled)
}

func EncodeCloudwatchEventRule_Name(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchEventRule_RoleArn(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["role_arn"] = cty.StringVal(p.RoleArn)
}

func EncodeCloudwatchEventRule_Arn(p CloudwatchEventRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}