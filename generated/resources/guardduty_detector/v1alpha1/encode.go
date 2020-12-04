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

func EncodeGuarddutyDetector(r GuarddutyDetector) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeGuarddutyDetector_FindingPublishingFrequency(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyDetector_Id(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyDetector_Tags(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyDetector_Enable(r.Spec.ForProvider, ctyVal)
	EncodeGuarddutyDetector_AccountId(r.Status.AtProvider, ctyVal)
	EncodeGuarddutyDetector_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeGuarddutyDetector_FindingPublishingFrequency(p GuarddutyDetectorParameters, vals map[string]cty.Value) {
	vals["finding_publishing_frequency"] = cty.StringVal(p.FindingPublishingFrequency)
}

func EncodeGuarddutyDetector_Id(p GuarddutyDetectorParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeGuarddutyDetector_Tags(p GuarddutyDetectorParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeGuarddutyDetector_Enable(p GuarddutyDetectorParameters, vals map[string]cty.Value) {
	vals["enable"] = cty.BoolVal(p.Enable)
}

func EncodeGuarddutyDetector_AccountId(p GuarddutyDetectorObservation, vals map[string]cty.Value) {
	vals["account_id"] = cty.StringVal(p.AccountId)
}

func EncodeGuarddutyDetector_Arn(p GuarddutyDetectorObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}