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

func EncodeLexBot(r LexBot) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLexBot_EnableModelImprovements(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_Id(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_ChildDirected(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_Locale(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_ProcessBehavior(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_Name(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_NluIntentConfidenceThreshold(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_DetectSentiment(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_IdleSessionTtlInSeconds(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_VoiceId(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_CreateVersion(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_Description(r.Spec.ForProvider, ctyVal)
	EncodeLexBot_AbortStatement(r.Spec.ForProvider.AbortStatement, ctyVal)
	EncodeLexBot_ClarificationPrompt(r.Spec.ForProvider.ClarificationPrompt, ctyVal)
	EncodeLexBot_Intent(r.Spec.ForProvider.Intent, ctyVal)
	EncodeLexBot_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeLexBot_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	EncodeLexBot_Status(r.Status.AtProvider, ctyVal)
	EncodeLexBot_Version(r.Status.AtProvider, ctyVal)
	EncodeLexBot_Checksum(r.Status.AtProvider, ctyVal)
	EncodeLexBot_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeLexBot_Arn(r.Status.AtProvider, ctyVal)
	EncodeLexBot_FailureReason(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLexBot_EnableModelImprovements(p LexBotParameters, vals map[string]cty.Value) {
	vals["enable_model_improvements"] = cty.BoolVal(p.EnableModelImprovements)
}

func EncodeLexBot_Id(p LexBotParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLexBot_ChildDirected(p LexBotParameters, vals map[string]cty.Value) {
	vals["child_directed"] = cty.BoolVal(p.ChildDirected)
}

func EncodeLexBot_Locale(p LexBotParameters, vals map[string]cty.Value) {
	vals["locale"] = cty.StringVal(p.Locale)
}

func EncodeLexBot_ProcessBehavior(p LexBotParameters, vals map[string]cty.Value) {
	vals["process_behavior"] = cty.StringVal(p.ProcessBehavior)
}

func EncodeLexBot_Name(p LexBotParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLexBot_NluIntentConfidenceThreshold(p LexBotParameters, vals map[string]cty.Value) {
	vals["nlu_intent_confidence_threshold"] = cty.NumberIntVal(p.NluIntentConfidenceThreshold)
}

func EncodeLexBot_DetectSentiment(p LexBotParameters, vals map[string]cty.Value) {
	vals["detect_sentiment"] = cty.BoolVal(p.DetectSentiment)
}

func EncodeLexBot_IdleSessionTtlInSeconds(p LexBotParameters, vals map[string]cty.Value) {
	vals["idle_session_ttl_in_seconds"] = cty.NumberIntVal(p.IdleSessionTtlInSeconds)
}

func EncodeLexBot_VoiceId(p LexBotParameters, vals map[string]cty.Value) {
	vals["voice_id"] = cty.StringVal(p.VoiceId)
}

func EncodeLexBot_CreateVersion(p LexBotParameters, vals map[string]cty.Value) {
	vals["create_version"] = cty.BoolVal(p.CreateVersion)
}

func EncodeLexBot_Description(p LexBotParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLexBot_AbortStatement(p AbortStatement, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexBot_AbortStatement_ResponseCard(p, ctyVal)
	EncodeLexBot_AbortStatement_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["abort_statement"] = cty.ListVal(valsForCollection)
}

func EncodeLexBot_AbortStatement_ResponseCard(p AbortStatement, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexBot_AbortStatement_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexBot_AbortStatement_Message_Content(v, ctyVal)
		EncodeLexBot_AbortStatement_Message_ContentType(v, ctyVal)
		EncodeLexBot_AbortStatement_Message_GroupNumber(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexBot_AbortStatement_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexBot_AbortStatement_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexBot_AbortStatement_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexBot_ClarificationPrompt(p ClarificationPrompt, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexBot_ClarificationPrompt_MaxAttempts(p, ctyVal)
	EncodeLexBot_ClarificationPrompt_ResponseCard(p, ctyVal)
	EncodeLexBot_ClarificationPrompt_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["clarification_prompt"] = cty.ListVal(valsForCollection)
}

func EncodeLexBot_ClarificationPrompt_MaxAttempts(p ClarificationPrompt, vals map[string]cty.Value) {
	vals["max_attempts"] = cty.NumberIntVal(p.MaxAttempts)
}

func EncodeLexBot_ClarificationPrompt_ResponseCard(p ClarificationPrompt, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexBot_ClarificationPrompt_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexBot_ClarificationPrompt_Message_Content(v, ctyVal)
		EncodeLexBot_ClarificationPrompt_Message_ContentType(v, ctyVal)
		EncodeLexBot_ClarificationPrompt_Message_GroupNumber(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexBot_ClarificationPrompt_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexBot_ClarificationPrompt_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexBot_ClarificationPrompt_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexBot_Intent(p []Intent, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexBot_Intent_IntentVersion(v, ctyVal)
		EncodeLexBot_Intent_IntentName(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["intent"] = cty.SetVal(valsForCollection)
}

func EncodeLexBot_Intent_IntentVersion(p Intent, vals map[string]cty.Value) {
	vals["intent_version"] = cty.StringVal(p.IntentVersion)
}

func EncodeLexBot_Intent_IntentName(p Intent, vals map[string]cty.Value) {
	vals["intent_name"] = cty.StringVal(p.IntentName)
}

func EncodeLexBot_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeLexBot_Timeouts_Create(p, ctyVal)
	EncodeLexBot_Timeouts_Delete(p, ctyVal)
	EncodeLexBot_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeLexBot_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeLexBot_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeLexBot_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeLexBot_LastUpdatedDate(p LexBotObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}

func EncodeLexBot_Status(p LexBotObservation, vals map[string]cty.Value) {
	vals["status"] = cty.StringVal(p.Status)
}

func EncodeLexBot_Version(p LexBotObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}

func EncodeLexBot_Checksum(p LexBotObservation, vals map[string]cty.Value) {
	vals["checksum"] = cty.StringVal(p.Checksum)
}

func EncodeLexBot_CreatedDate(p LexBotObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeLexBot_Arn(p LexBotObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLexBot_FailureReason(p LexBotObservation, vals map[string]cty.Value) {
	vals["failure_reason"] = cty.StringVal(p.FailureReason)
}