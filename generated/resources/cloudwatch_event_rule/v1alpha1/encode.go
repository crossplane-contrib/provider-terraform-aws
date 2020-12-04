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
	"github.com/zclconf/go-cty/cty"
)

func EncodeCloudwatchEventRule(r CloudwatchEventRule) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchEventRule_EventPattern(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_ScheduleExpression(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Description(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_IsEnabled(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_RoleArn(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchEventRule_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchEventRule_EventPattern(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["event_pattern"] = cty.StringVal(p.EventPattern)
}

func EncodeCloudwatchEventRule_Id(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchEventRule_NamePrefix(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeCloudwatchEventRule_ScheduleExpression(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["schedule_expression"] = cty.StringVal(p.ScheduleExpression)
}

func EncodeCloudwatchEventRule_Description(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
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

func EncodeCloudwatchEventRule_Tags(p CloudwatchEventRuleParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudwatchEventRule_Arn(p CloudwatchEventRuleObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}