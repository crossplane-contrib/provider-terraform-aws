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

package v1alpha1func EncodeSesReceiptRule(r SesReceiptRule) cty.Value {
	ctyVals := make(map[string]cty.Value)
	EncodeSesReceiptRule_TlsPolicy(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_After(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_Enabled(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_Id(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_Name(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_Recipients(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_RuleSetName(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_ScanEnabled(r.Spec.ForProvider, ctyVal)
	EncodeSesReceiptRule_SnsAction(r.Spec.ForProvider.SnsAction, ctyVal)
	EncodeSesReceiptRule_StopAction(r.Spec.ForProvider.StopAction, ctyVal)
	EncodeSesReceiptRule_WorkmailAction(r.Spec.ForProvider.WorkmailAction, ctyVal)
	EncodeSesReceiptRule_AddHeaderAction(r.Spec.ForProvider.AddHeaderAction, ctyVal)
	EncodeSesReceiptRule_BounceAction(r.Spec.ForProvider.BounceAction, ctyVal)
	EncodeSesReceiptRule_LambdaAction(r.Spec.ForProvider.LambdaAction, ctyVal)
	EncodeSesReceiptRule_S3Action(r.Spec.ForProvider.S3Action, ctyVal)

	return cty.ObjectVal(ctyVals)
}

func EncodeSesReceiptRule_TlsPolicy(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	vals["tls_policy"] = cty.StringVal(p.TlsPolicy)
}

func EncodeSesReceiptRule_After(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	vals["after"] = cty.StringVal(p.After)
}

func EncodeSesReceiptRule_Enabled(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	vals["enabled"] = cty.BoolVal(p.Enabled)
}

func EncodeSesReceiptRule_Id(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeSesReceiptRule_Name(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeSesReceiptRule_Recipients(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.Recipients {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["recipients"] = cty.SetVal(colVals)
}

func EncodeSesReceiptRule_RuleSetName(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	vals["rule_set_name"] = cty.StringVal(p.RuleSetName)
}

func EncodeSesReceiptRule_ScanEnabled(p *SesReceiptRuleParameters, vals map[string]cty.Value) {
	vals["scan_enabled"] = cty.BoolVal(p.ScanEnabled)
}

func EncodeSesReceiptRule_SnsAction(p *SnsAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.SnsAction {
		ctyVal = make(map[string]cty.Value)
		EncodeSesReceiptRule_SnsAction_Position(v, ctyVal)
		EncodeSesReceiptRule_SnsAction_TopicArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["sns_action"] = cty.SetVal(valsForCollection)
}

func EncodeSesReceiptRule_SnsAction_Position(p *SnsAction, vals map[string]cty.Value) {
	vals["position"] = cty.IntVal(p.Position)
}

func EncodeSesReceiptRule_SnsAction_TopicArn(p *SnsAction, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}

func EncodeSesReceiptRule_StopAction(p *StopAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.StopAction {
		ctyVal = make(map[string]cty.Value)
		EncodeSesReceiptRule_StopAction_Position(v, ctyVal)
		EncodeSesReceiptRule_StopAction_Scope(v, ctyVal)
		EncodeSesReceiptRule_StopAction_TopicArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["stop_action"] = cty.SetVal(valsForCollection)
}

func EncodeSesReceiptRule_StopAction_Position(p *StopAction, vals map[string]cty.Value) {
	vals["position"] = cty.IntVal(p.Position)
}

func EncodeSesReceiptRule_StopAction_Scope(p *StopAction, vals map[string]cty.Value) {
	vals["scope"] = cty.StringVal(p.Scope)
}

func EncodeSesReceiptRule_StopAction_TopicArn(p *StopAction, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}

func EncodeSesReceiptRule_WorkmailAction(p *WorkmailAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.WorkmailAction {
		ctyVal = make(map[string]cty.Value)
		EncodeSesReceiptRule_WorkmailAction_Position(v, ctyVal)
		EncodeSesReceiptRule_WorkmailAction_TopicArn(v, ctyVal)
		EncodeSesReceiptRule_WorkmailAction_OrganizationArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["workmail_action"] = cty.SetVal(valsForCollection)
}

func EncodeSesReceiptRule_WorkmailAction_Position(p *WorkmailAction, vals map[string]cty.Value) {
	vals["position"] = cty.IntVal(p.Position)
}

func EncodeSesReceiptRule_WorkmailAction_TopicArn(p *WorkmailAction, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}

func EncodeSesReceiptRule_WorkmailAction_OrganizationArn(p *WorkmailAction, vals map[string]cty.Value) {
	vals["organization_arn"] = cty.StringVal(p.OrganizationArn)
}

func EncodeSesReceiptRule_AddHeaderAction(p *AddHeaderAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.AddHeaderAction {
		ctyVal = make(map[string]cty.Value)
		EncodeSesReceiptRule_AddHeaderAction_Position(v, ctyVal)
		EncodeSesReceiptRule_AddHeaderAction_HeaderName(v, ctyVal)
		EncodeSesReceiptRule_AddHeaderAction_HeaderValue(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["add_header_action"] = cty.SetVal(valsForCollection)
}

func EncodeSesReceiptRule_AddHeaderAction_Position(p *AddHeaderAction, vals map[string]cty.Value) {
	vals["position"] = cty.IntVal(p.Position)
}

func EncodeSesReceiptRule_AddHeaderAction_HeaderName(p *AddHeaderAction, vals map[string]cty.Value) {
	vals["header_name"] = cty.StringVal(p.HeaderName)
}

func EncodeSesReceiptRule_AddHeaderAction_HeaderValue(p *AddHeaderAction, vals map[string]cty.Value) {
	vals["header_value"] = cty.StringVal(p.HeaderValue)
}

func EncodeSesReceiptRule_BounceAction(p *BounceAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.BounceAction {
		ctyVal = make(map[string]cty.Value)
		EncodeSesReceiptRule_BounceAction_Position(v, ctyVal)
		EncodeSesReceiptRule_BounceAction_Sender(v, ctyVal)
		EncodeSesReceiptRule_BounceAction_SmtpReplyCode(v, ctyVal)
		EncodeSesReceiptRule_BounceAction_StatusCode(v, ctyVal)
		EncodeSesReceiptRule_BounceAction_TopicArn(v, ctyVal)
		EncodeSesReceiptRule_BounceAction_Message(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["bounce_action"] = cty.SetVal(valsForCollection)
}

func EncodeSesReceiptRule_BounceAction_Position(p *BounceAction, vals map[string]cty.Value) {
	vals["position"] = cty.IntVal(p.Position)
}

func EncodeSesReceiptRule_BounceAction_Sender(p *BounceAction, vals map[string]cty.Value) {
	vals["sender"] = cty.StringVal(p.Sender)
}

func EncodeSesReceiptRule_BounceAction_SmtpReplyCode(p *BounceAction, vals map[string]cty.Value) {
	vals["smtp_reply_code"] = cty.StringVal(p.SmtpReplyCode)
}

func EncodeSesReceiptRule_BounceAction_StatusCode(p *BounceAction, vals map[string]cty.Value) {
	vals["status_code"] = cty.StringVal(p.StatusCode)
}

func EncodeSesReceiptRule_BounceAction_TopicArn(p *BounceAction, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}

func EncodeSesReceiptRule_BounceAction_Message(p *BounceAction, vals map[string]cty.Value) {
	vals["message"] = cty.StringVal(p.Message)
}

func EncodeSesReceiptRule_LambdaAction(p *LambdaAction, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.LambdaAction {
		ctyVal = make(map[string]cty.Value)
		EncodeSesReceiptRule_LambdaAction_FunctionArn(v, ctyVal)
		EncodeSesReceiptRule_LambdaAction_InvocationType(v, ctyVal)
		EncodeSesReceiptRule_LambdaAction_Position(v, ctyVal)
		EncodeSesReceiptRule_LambdaAction_TopicArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["lambda_action"] = cty.SetVal(valsForCollection)
}

func EncodeSesReceiptRule_LambdaAction_FunctionArn(p *LambdaAction, vals map[string]cty.Value) {
	vals["function_arn"] = cty.StringVal(p.FunctionArn)
}

func EncodeSesReceiptRule_LambdaAction_InvocationType(p *LambdaAction, vals map[string]cty.Value) {
	vals["invocation_type"] = cty.StringVal(p.InvocationType)
}

func EncodeSesReceiptRule_LambdaAction_Position(p *LambdaAction, vals map[string]cty.Value) {
	vals["position"] = cty.IntVal(p.Position)
}

func EncodeSesReceiptRule_LambdaAction_TopicArn(p *LambdaAction, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}

func EncodeSesReceiptRule_S3Action(p *S3Action, vals map[string]cty.Value) {
	valsForCollection = make([]cty.Value, 0)
	for _, v := range p.S3Action {
		ctyVal = make(map[string]cty.Value)
		EncodeSesReceiptRule_S3Action_BucketName(v, ctyVal)
		EncodeSesReceiptRule_S3Action_KmsKeyArn(v, ctyVal)
		EncodeSesReceiptRule_S3Action_ObjectKeyPrefix(v, ctyVal)
		EncodeSesReceiptRule_S3Action_Position(v, ctyVal)
		EncodeSesReceiptRule_S3Action_TopicArn(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["s3_action"] = cty.SetVal(valsForCollection)
}

func EncodeSesReceiptRule_S3Action_BucketName(p *S3Action, vals map[string]cty.Value) {
	vals["bucket_name"] = cty.StringVal(p.BucketName)
}

func EncodeSesReceiptRule_S3Action_KmsKeyArn(p *S3Action, vals map[string]cty.Value) {
	vals["kms_key_arn"] = cty.StringVal(p.KmsKeyArn)
}

func EncodeSesReceiptRule_S3Action_ObjectKeyPrefix(p *S3Action, vals map[string]cty.Value) {
	vals["object_key_prefix"] = cty.StringVal(p.ObjectKeyPrefix)
}

func EncodeSesReceiptRule_S3Action_Position(p *S3Action, vals map[string]cty.Value) {
	vals["position"] = cty.IntVal(p.Position)
}

func EncodeSesReceiptRule_S3Action_TopicArn(p *S3Action, vals map[string]cty.Value) {
	vals["topic_arn"] = cty.StringVal(p.TopicArn)
}