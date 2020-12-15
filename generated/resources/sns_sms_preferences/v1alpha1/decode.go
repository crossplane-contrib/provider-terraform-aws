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

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/hashicorp/terraform/providers"
	"github.com/zclconf/go-cty/cty"
	ctwhy "github.com/crossplane-contrib/terraform-runtime/pkg/plugin/cty"
)

type ctyDecoder struct{}

func (e *ctyDecoder) DecodeCty(mr resource.Managed, ctyValue cty.Value, schema *providers.Schema) (resource.Managed, error) {
	r, ok := mr.(*SnsSmsPreferences)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSnsSmsPreferences(r, ctyValue)
}

func DecodeSnsSmsPreferences(prev *SnsSmsPreferences, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSnsSmsPreferences_DefaultSenderId(&new.Spec.ForProvider, valMap)
	DecodeSnsSmsPreferences_DefaultSmsType(&new.Spec.ForProvider, valMap)
	DecodeSnsSmsPreferences_DeliveryStatusIamRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsSmsPreferences_DeliveryStatusSuccessSamplingRate(&new.Spec.ForProvider, valMap)
	DecodeSnsSmsPreferences_MonthlySpendLimit(&new.Spec.ForProvider, valMap)
	DecodeSnsSmsPreferences_UsageReportS3Bucket(&new.Spec.ForProvider, valMap)

	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSnsSmsPreferences_DefaultSenderId(p *SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	p.DefaultSenderId = ctwhy.ValueAsString(vals["default_sender_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsSmsPreferences_DefaultSmsType(p *SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	p.DefaultSmsType = ctwhy.ValueAsString(vals["default_sms_type"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsSmsPreferences_DeliveryStatusIamRoleArn(p *SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	p.DeliveryStatusIamRoleArn = ctwhy.ValueAsString(vals["delivery_status_iam_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsSmsPreferences_DeliveryStatusSuccessSamplingRate(p *SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	p.DeliveryStatusSuccessSamplingRate = ctwhy.ValueAsString(vals["delivery_status_success_sampling_rate"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsSmsPreferences_MonthlySpendLimit(p *SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	p.MonthlySpendLimit = ctwhy.ValueAsString(vals["monthly_spend_limit"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsSmsPreferences_UsageReportS3Bucket(p *SnsSmsPreferencesParameters, vals map[string]cty.Value) {
	p.UsageReportS3Bucket = ctwhy.ValueAsString(vals["usage_report_s3_bucket"])
}