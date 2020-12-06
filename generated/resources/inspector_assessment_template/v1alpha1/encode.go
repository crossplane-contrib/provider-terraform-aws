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

func EncodeInspectorAssessmentTemplate(r InspectorAssessmentTemplate) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeInspectorAssessmentTemplate_TargetArn(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTemplate_Duration(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTemplate_Id(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTemplate_Name(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTemplate_RulesPackageArns(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTemplate_Tags(r.Spec.ForProvider, ctyVal)
	EncodeInspectorAssessmentTemplate_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeInspectorAssessmentTemplate_TargetArn(p InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	vals["target_arn"] = cty.StringVal(p.TargetArn)
}

func EncodeInspectorAssessmentTemplate_Duration(p InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	vals["duration"] = cty.NumberIntVal(p.Duration)
}

func EncodeInspectorAssessmentTemplate_Id(p InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeInspectorAssessmentTemplate_Name(p InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeInspectorAssessmentTemplate_RulesPackageArns(p InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.RulesPackageArns {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["rules_package_arns"] = cty.SetVal(colVals)
}

func EncodeInspectorAssessmentTemplate_Tags(p InspectorAssessmentTemplateParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeInspectorAssessmentTemplate_Arn(p InspectorAssessmentTemplateObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}