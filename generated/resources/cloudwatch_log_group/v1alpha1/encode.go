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

func EncodeCloudwatchLogGroup(r CloudwatchLogGroup) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeCloudwatchLogGroup_KmsKeyId(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_Name(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_RetentionInDays(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_Tags(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_Id(r.Spec.ForProvider, ctyVal)
	EncodeCloudwatchLogGroup_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeCloudwatchLogGroup_KmsKeyId(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["kms_key_id"] = cty.StringVal(p.KmsKeyId)
}

func EncodeCloudwatchLogGroup_Name(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeCloudwatchLogGroup_NamePrefix(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeCloudwatchLogGroup_RetentionInDays(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["retention_in_days"] = cty.NumberIntVal(p.RetentionInDays)
}

func EncodeCloudwatchLogGroup_Tags(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeCloudwatchLogGroup_Id(p CloudwatchLogGroupParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeCloudwatchLogGroup_Arn(p CloudwatchLogGroupObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}