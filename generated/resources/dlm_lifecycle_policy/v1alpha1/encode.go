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

package v1alpha1func EncodeDlmLifecyclePolicy(r DlmLifecyclePolicy) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeDlmLifecyclePolicy_Description(r.Spec.ForProvider, ctyVal)
	EncodeDlmLifecyclePolicy_ExecutionRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeDlmLifecyclePolicy_Id(r.Spec.ForProvider, ctyVal)
	EncodeDlmLifecyclePolicy_State(r.Spec.ForProvider, ctyVal)
	EncodeDlmLifecyclePolicy_Tags(r.Spec.ForProvider, ctyVal)
	EncodeDlmLifecyclePolicy_PolicyDetails(r.Spec.ForProvider.PolicyDetails, ctyVal)
	EncodeDlmLifecyclePolicy_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeDlmLifecyclePolicy_Description(p *DlmLifecyclePolicyParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeDlmLifecyclePolicy_ExecutionRoleArn(p *DlmLifecyclePolicyParameters, vals map[string]cty.Value) {
	vals["execution_role_arn"] = cty.StringVal(p.ExecutionRoleArn)
}

func EncodeDlmLifecyclePolicy_Id(p *DlmLifecyclePolicyParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeDlmLifecyclePolicy_State(p *DlmLifecyclePolicyParameters, vals map[string]cty.Value) {
	vals["state"] = cty.StringVal(p.State)
}

func EncodeDlmLifecyclePolicy_Tags(p *DlmLifecyclePolicyParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeDlmLifecyclePolicy_PolicyDetails(p *PolicyDetails, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.PolicyDetails {
		ctyVal = make(map[string]cty.Value)
		EncodeDlmLifecyclePolicy_PolicyDetails_TargetTags(v, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_ResourceTypes(v, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule(v.Schedule, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["policy_details"] = cty.ListVal(valsForCollection)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_TargetTags(p *PolicyDetails, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.TargetTags {
		mVals[key] = cty.StringVal(value)
	}
	vals["target_tags"] = cty.MapVal(mVals)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_ResourceTypes(p *PolicyDetails, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.ResourceTypes {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["resource_types"] = cty.ListVal(colVals)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule(p *Schedule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.Schedule {
		ctyVal = make(map[string]cty.Value)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_TagsToAdd(v, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CopyTags(v, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_Name(v, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule(v.CreateRule, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_RetainRule(v.RetainRule, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["schedule"] = cty.ListVal(valsForCollection)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_TagsToAdd(p *Schedule, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.TagsToAdd {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags_to_add"] = cty.MapVal(mVals)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CopyTags(p *Schedule, vals map[string]cty.Value) {
	vals["copy_tags"] = cty.BoolVal(p.CopyTags)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_Name(p *Schedule, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule(p *CreateRule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.CreateRule {
		ctyVal = make(map[string]cty.Value)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule_Interval(v, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule_IntervalUnit(v, ctyVal)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule_Times(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["create_rule"] = cty.ListVal(valsForCollection)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule_Interval(p *CreateRule, vals map[string]cty.Value) {
	vals["interval"] = cty.IntVal(p.Interval)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule_IntervalUnit(p *CreateRule, vals map[string]cty.Value) {
	vals["interval_unit"] = cty.StringVal(p.IntervalUnit)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_CreateRule_Times(p *CreateRule, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Times {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["times"] = cty.ListVal(colVals)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_RetainRule(p *RetainRule, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.RetainRule {
		ctyVal = make(map[string]cty.Value)
		EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_RetainRule_Count(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["retain_rule"] = cty.ListVal(valsForCollection)
}

func EncodeDlmLifecyclePolicy_PolicyDetails_Schedule_RetainRule_Count(p *RetainRule, vals map[string]cty.Value) {
	vals["count"] = cty.IntVal(p.Count)
}

func EncodeDlmLifecyclePolicy_Arn(p *DlmLifecyclePolicyObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}