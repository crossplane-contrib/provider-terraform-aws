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
	r, ok := mr.(*SnsTopic)
	if !ok {
		return nil, fmt.Errorf("DecodeCty received a resource.Managed value that does not assert to the expected type")
	}
	return DecodeSnsTopic(r, ctyValue)
}

func DecodeSnsTopic(prev *SnsTopic, ctyValue cty.Value) (resource.Managed, error) {
	valMap := ctyValue.AsValueMap()
	new := prev.DeepCopy()
	DecodeSnsTopic_ApplicationSuccessFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_SqsSuccessFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_ApplicationSuccessFeedbackSampleRate(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_DisplayName(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_HttpSuccessFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_HttpSuccessFeedbackSampleRate(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_LambdaFailureFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_Name(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_SqsSuccessFeedbackSampleRate(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_KmsMasterKeyId(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_LambdaSuccessFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_Policy(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_SqsFailureFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_ApplicationFailureFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_DeliveryPolicy(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_HttpFailureFeedbackRoleArn(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_Id(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_LambdaSuccessFeedbackSampleRate(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_NamePrefix(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_Tags(&new.Spec.ForProvider, valMap)
	DecodeSnsTopic_Arn(&new.Status.AtProvider, valMap)
	eid := valMap["id"].AsString()
	if len(eid) > 0 {
		meta.SetExternalName(new, eid)
	}
	return new, nil
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_ApplicationSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.ApplicationSuccessFeedbackRoleArn = ctwhy.ValueAsString(vals["application_success_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_SqsSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.SqsSuccessFeedbackRoleArn = ctwhy.ValueAsString(vals["sqs_success_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_ApplicationSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.ApplicationSuccessFeedbackSampleRate = ctwhy.ValueAsInt64(vals["application_success_feedback_sample_rate"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_DisplayName(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.DisplayName = ctwhy.ValueAsString(vals["display_name"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_HttpSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.HttpSuccessFeedbackRoleArn = ctwhy.ValueAsString(vals["http_success_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_HttpSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.HttpSuccessFeedbackSampleRate = ctwhy.ValueAsInt64(vals["http_success_feedback_sample_rate"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_LambdaFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.LambdaFailureFeedbackRoleArn = ctwhy.ValueAsString(vals["lambda_failure_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_Name(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.Name = ctwhy.ValueAsString(vals["name"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_SqsSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.SqsSuccessFeedbackSampleRate = ctwhy.ValueAsInt64(vals["sqs_success_feedback_sample_rate"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_KmsMasterKeyId(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.KmsMasterKeyId = ctwhy.ValueAsString(vals["kms_master_key_id"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_LambdaSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.LambdaSuccessFeedbackRoleArn = ctwhy.ValueAsString(vals["lambda_success_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_Policy(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.Policy = ctwhy.ValueAsString(vals["policy"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_SqsFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.SqsFailureFeedbackRoleArn = ctwhy.ValueAsString(vals["sqs_failure_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_ApplicationFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.ApplicationFailureFeedbackRoleArn = ctwhy.ValueAsString(vals["application_failure_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_DeliveryPolicy(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.DeliveryPolicy = ctwhy.ValueAsString(vals["delivery_policy"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_HttpFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.HttpFailureFeedbackRoleArn = ctwhy.ValueAsString(vals["http_failure_feedback_role_arn"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_Id(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.Id = ctwhy.ValueAsString(vals["id"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_LambdaSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.LambdaSuccessFeedbackSampleRate = ctwhy.ValueAsInt64(vals["lambda_success_feedback_sample_rate"])
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_NamePrefix(p *SnsTopicParameters, vals map[string]cty.Value) {
	p.NamePrefix = ctwhy.ValueAsString(vals["name_prefix"])
}

//primitiveMapTypeDecodeTemplate
func DecodeSnsTopic_Tags(p *SnsTopicParameters, vals map[string]cty.Value) {
	// TODO: generalize generation of the element type, string elements are hard-coded atm
	vMap := make(map[string]string)
	v := vals["tags"].AsValueMap()
	for key, value := range v {
		vMap[key] = ctwhy.ValueAsString(value)
	}
	p.Tags = vMap
}

//primitiveTypeDecodeTemplate
func DecodeSnsTopic_Arn(p *SnsTopicObservation, vals map[string]cty.Value) {
	p.Arn = ctwhy.ValueAsString(vals["arn"])
}