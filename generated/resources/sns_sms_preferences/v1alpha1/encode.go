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

func EncodeSnsSmsPreferences(r SnsSmsPreferences) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSnsSmsPreferences_DefaultSenderId(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_DefaultSmsType(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_DeliveryStatusIamRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_DeliveryStatusSuccessSamplingRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_Id(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_MonthlySpendLimit(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_UsageReportS3Bucket(r.Spec.ForProvider, ctyVal)

	return cty.ObjectVal(ctyVal)
}

func EncodeSnsSmsPreferences_DefaultSenderId(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["default_sender_id"] = cty.StringVal(p.DefaultSenderId)
}

func EncodeSnsSmsPreferences_DefaultSmsType(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["default_sms_type"] = cty.StringVal(p.DefaultSmsType)
}

func EncodeSnsSmsPreferences_DeliveryStatusIamRoleArn(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["delivery_status_iam_role_arn"] = cty.StringVal(p.DeliveryStatusIamRoleArn)
}

func EncodeSnsSmsPreferences_DeliveryStatusSuccessSamplingRate(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["delivery_status_success_sampling_rate"] = cty.StringVal(p.DeliveryStatusSuccessSamplingRate)
}

func EncodeSnsSmsPreferences_Id(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSnsSmsPreferences_MonthlySpendLimit(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["monthly_spend_limit"] = cty.StringVal(p.MonthlySpendLimit)
}

func EncodeSnsSmsPreferences_UsageReportS3Bucket(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["usage_report_s3_bucket"] = cty.StringVal(p.UsageReportS3Bucket)
}