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

func EncodeInspectorAssessmentTarget(r InspectorAssessmentTarget) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeInspectorAssessmentTarget_Name(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTarget_ResourceGroupArn(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTarget_Id(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTarget_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeInspectorAssessmentTarget_Name(p InspectorAssessmentTargetParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeInspectorAssessmentTarget_ResourceGroupArn(p InspectorAssessmentTargetParameters, vals map[string]cty.Value) {
	vals["resource_group_arn"] = cty.StringVal(p.ResourceGroupArn)
}

func EncodeInspectorAssessmentTarget_Id(p InspectorAssessmentTargetParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeInspectorAssessmentTarget_Arn(p InspectorAssessmentTargetObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}