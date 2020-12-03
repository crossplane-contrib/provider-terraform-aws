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

package v1alpha1func EncodeSnsTopic(r SnsTopic) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSnsTopic_HttpFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Id(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_LambdaFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_LambdaSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_SqsSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_SqsSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Tags(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_ApplicationSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_DeliveryPolicy(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_HttpSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_SqsFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_NamePrefix(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_DisplayName(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_KmsMasterKeyId(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_LambdaSuccessFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Policy(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_ApplicationFailureFeedbackRoleArn(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_ApplicationSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_HttpSuccessFeedbackSampleRate(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Name(r.Spec.ForProvider, ctyVal)
	EncodeSnsTopic_Arn(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVals)
}

func EncodeSnsTopic_HttpFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["http_failure_feedback_role_arn"] = cty.StringVal(p.HttpFailureFeedbackRoleArn)
}

func EncodeSnsTopic_Id(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSnsTopic_LambdaFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["lambda_failure_feedback_role_arn"] = cty.StringVal(p.LambdaFailureFeedbackRoleArn)
}

func EncodeSnsTopic_LambdaSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["lambda_success_feedback_sample_rate"] = cty.IntVal(p.LambdaSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_SqsSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["sqs_success_feedback_role_arn"] = cty.StringVal(p.SqsSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_SqsSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["sqs_success_feedback_sample_rate"] = cty.IntVal(p.SqsSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_Tags(p *SnsTopicParameters, vals map[string]cty.Value) {
	mVals := make(map[string]cty.Value)
	for key, value := range p.Tags {
		mVals[key] = cty.StringVal(value)
	}
	vals["tags"] = cty.MapVal(mVals)
}

func EncodeSnsTopic_ApplicationSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["application_success_feedback_role_arn"] = cty.StringVal(p.ApplicationSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_DeliveryPolicy(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["delivery_policy"] = cty.StringVal(p.DeliveryPolicy)
}

func EncodeSnsTopic_HttpSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["http_success_feedback_role_arn"] = cty.StringVal(p.HttpSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_SqsFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["sqs_failure_feedback_role_arn"] = cty.StringVal(p.SqsFailureFeedbackRoleArn)
}

func EncodeSnsTopic_NamePrefix(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["name_prefix"] = cty.StringVal(p.NamePrefix)
}

func EncodeSnsTopic_DisplayName(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["display_name"] = cty.StringVal(p.DisplayName)
}

func EncodeSnsTopic_KmsMasterKeyId(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["kms_master_key_id"] = cty.StringVal(p.KmsMasterKeyId)
}

func EncodeSnsTopic_LambdaSuccessFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["lambda_success_feedback_role_arn"] = cty.StringVal(p.LambdaSuccessFeedbackRoleArn)
}

func EncodeSnsTopic_Policy(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["policy"] = cty.StringVal(p.Policy)
}

func EncodeSnsTopic_ApplicationFailureFeedbackRoleArn(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["application_failure_feedback_role_arn"] = cty.StringVal(p.ApplicationFailureFeedbackRoleArn)
}

func EncodeSnsTopic_ApplicationSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["application_success_feedback_sample_rate"] = cty.IntVal(p.ApplicationSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_HttpSuccessFeedbackSampleRate(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["http_success_feedback_sample_rate"] = cty.IntVal(p.HttpSuccessFeedbackSampleRate)
}

func EncodeSnsTopic_Name(p *SnsTopicParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSnsTopic_Arn(p *SnsTopicObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}