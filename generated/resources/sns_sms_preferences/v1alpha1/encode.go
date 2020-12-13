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
	r, ok := mr.(*SnsSmsPreferences)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SnsSmsPreferences.")
	}
	return EncodeSnsSmsPreferences(*r), nil
}

func EncodeSnsSmsPreferences(r SnsSmsPreferences) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSnsSmsPreferences_DeliveryStatusSuccessSamplingRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_Id(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_MonthlySpendLimit(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_UsageReportS3Bucket(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_DefaultSenderId(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_DefaultSmsType(r.Spec.ForProvider, ctyVal)
	EncodeSnsSmsPreferences_DeliveryStatusIamRoleArn(r.Spec.ForProvider, ctyVal)

	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
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

func EncodeSnsSmsPreferences_DefaultSenderId(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["default_sender_id"] = cty.StringVal(p.DefaultSenderId)
}

func EncodeSnsSmsPreferences_DefaultSmsType(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["default_sms_type"] = cty.StringVal(p.DefaultSmsType)
}

func EncodeSnsSmsPreferences_DeliveryStatusIamRoleArn(p SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	vals["delivery_status_iam_role_arn"] = cty.StringVal(p.DeliveryStatusIamRoleArn)
}