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
	r, ok := mr.(*SnsTopic)
	if !ok {
		return cty.NilVal, fmt.Errorf("EncodeType received a resource.Managed value which is not a SnsTopic.")
	}
	return EncodeSnsTopic(*r), nil
}

func EncodeSnsTopic(r SnsTopic) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeSnsTopic_ApplicationFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_DeliveryPolicy(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Id(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_SqsFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_SqsSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Name(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_SqsSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_ApplicationSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_LambdaSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_ApplicationSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_DisplayName(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_HttpFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_HttpSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_HttpSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_KmsMasterKeyId(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_LambdaFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_LambdaSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Policy(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Arn(r.Status.AtProvider, ctyVal)
	// always set id = external-name if it exists
	// TODO: we should trim Id off schemas in an "optimize" pass
	// before code generation
	en := meta.GetExternalName(&r)
	if len(en) > 0 {
		ctyVal["id"] = cty.StringVal(en)
	}
	return cty.ObjectVal(ctyVal)
}

func EncodeSnsTopic_ApplicationFailureFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["application_failure_feedback_role_arn"] = cty.StringVal(p.ApplicationFailureFeedbackRoleArn)
}

func EncodeSnsTopic_DeliveryPolicy(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["delivery_policy"] = cty.StringVal(p.DeliveryPolicy)
}

func EncodeSnsTopic_Id(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSnsTopic_SqsFailureFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["sqs_failure_feedback_role_arn"] = cty.StringVal(p.SqsFailureFeedbackRoleArn)
}

func EncodeSnsTopic_SqsSuccessFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["sqs_success_feedback_role_arn"] = cty.StringVal(p.SqsSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_Name(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSnsTopic_NamePrefix(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeSnsTopic_SqsSuccessFeedbackSampleRate(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["sqs_success_feedback_sample_rate"] = cty.NumberIntVal(p.SqsSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_ApplicationSuccessFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["application_success_feedback_role_arn"] = cty.StringVal(p.ApplicationSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_Tags(p SnsTopicParameters, vals map[string]cty.Value) {
	if len(p.Tags) == 0 {
		vals["tags"] = cty.NullVal(cty.Map(cty.String))
		return
	}
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSnsTopic_LambdaSuccessFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["lambda_success_feedback_role_arn"] = cty.StringVal(p.LambdaSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_ApplicationSuccessFeedbackSampleRate(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["application_success_feedback_sample_rate"] = cty.NumberIntVal(p.ApplicationSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_DisplayName(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["display_name"] = cty.StringVal(p.DisplayName)
}

func EncodeSnsTopic_HttpFailureFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["http_failure_feedback_role_arn"] = cty.StringVal(p.HttpFailureFeedbackRoleArn)
}

func EncodeSnsTopic_HttpSuccessFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["http_success_feedback_role_arn"] = cty.StringVal(p.HttpSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_HttpSuccessFeedbackSampleRate(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["http_success_feedback_sample_rate"] = cty.NumberIntVal(p.HttpSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_KmsMasterKeyId(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["kms_master_key_id"] = cty.StringVal(p.KmsMasterKeyId)
}

func EncodeSnsTopic_LambdaFailureFeedbackRoleArn(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["lambda_failure_feedback_role_arn"] = cty.StringVal(p.LambdaFailureFeedbackRoleArn)
}

func EncodeSnsTopic_LambdaSuccessFeedbackSampleRate(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["lambda_success_feedback_sample_rate"] = cty.NumberIntVal(p.LambdaSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_Policy(p SnsTopicParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeSnsTopic_Arn(p SnsTopicObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}