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

func EncodeGuarddutyFilter(r GuarddutyFilter) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyFilter_DetectorId(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyFilter_Id(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyFilter_Name(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyFilter_Rank(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyFilter_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyFilter_Action(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyFilter_Description(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyFilter_FindingCriteria(r.Spec.ForProvider.FindingCriteria, ctyVal)
	EncodeGuarddutyFilter_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyFilter_DetectorId(p GuarddutyFilterParameters, vals map[string]cty.Value) {
	vals["detector_id"] = cty.StringVal(p.DetectorId)
}

func EncodeGuarddutyFilter_Id(p GuarddutyFilterParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGuarddutyFilter_Name(p GuarddutyFilterParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeGuarddutyFilter_Rank(p GuarddutyFilterParameters, vals map[string]cty.Value) {
	vals["rank"] = cty.NumberIntVal(p.Rank)
}

func EncodeGuarddutyFilter_Tags(p GuarddutyFilterParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGuarddutyFilter_Action(p GuarddutyFilterParameters, vals map[string]cty.Value) {
	vals["action"] = cty.StringVal(p.Action)
}

func EncodeGuarddutyFilter_Description(p GuarddutyFilterParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeGuarddutyFilter_FindingCriteria(p FindingCriteria, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyFilter_FindingCriteria_Criterion(p.Criterion, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["finding_criteria"] = cty.ListVal(valsForCollection)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion(p []Criterion, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeGuarddutyFilter_FindingCriteria_Criterion_Equals(v, ctyVal)
		EncodeGuarddutyFilter_FindingCriteria_Criterion_Field(v, ctyVal)
		EncodeGuarddutyFilter_FindingCriteria_Criterion_GreaterThan(v, ctyVal)
		EncodeGuarddutyFilter_FindingCriteria_Criterion_GreaterThanOrEqual(v, ctyVal)
		EncodeGuarddutyFilter_FindingCriteria_Criterion_LessThan(v, ctyVal)
		EncodeGuarddutyFilter_FindingCriteria_Criterion_LessThanOrEqual(v, ctyVal)
		EncodeGuarddutyFilter_FindingCriteria_Criterion_NotEquals(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["criterion"] = cty.SetVal(valsForCollection)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion_Equals(p Criterion, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Equals {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["equals"] = cty.ListVal(colVals)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion_Field(p Criterion, vals map[string]cty.Value) {
	vals["field"] = cty.StringVal(p.Field)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion_GreaterThan(p Criterion, vals map[string]cty.Value) {
	vals["greater_than"] = cty.StringVal(p.GreaterThan)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion_GreaterThanOrEqual(p Criterion, vals map[string]cty.Value) {
	vals["greater_than_or_equal"] = cty.StringVal(p.GreaterThanOrEqual)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion_LessThan(p Criterion, vals map[string]cty.Value) {
	vals["less_than"] = cty.StringVal(p.LessThan)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion_LessThanOrEqual(p Criterion, vals map[string]cty.Value) {
	vals["less_than_or_equal"] = cty.StringVal(p.LessThanOrEqual)
}

func EncodeGuarddutyFilter_FindingCriteria_Criterion_NotEquals(p Criterion, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.NotEquals {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["not_equals"] = cty.ListVal(colVals)
}

func EncodeGuarddutyFilter_Arn(p GuarddutyFilterObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}