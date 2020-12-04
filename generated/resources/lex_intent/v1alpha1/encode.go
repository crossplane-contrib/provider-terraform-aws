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

func EncodeLexIntent(r LexIntent) cty.Value {
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_Id(r.Spec.ForProvider, ctyVal)
	EncodeLexIntent_Name(r.Spec.ForProvider, ctyVal)
	EncodeLexIntent_ParentIntentSignature(r.Spec.ForProvider, ctyVal)
	EncodeLexIntent_CreateVersion(r.Spec.ForProvider, ctyVal)
	EncodeLexIntent_Description(r.Spec.ForProvider, ctyVal)
	EncodeLexIntent_SampleUtterances(r.Spec.ForProvider, ctyVal)
	EncodeLexIntent_ConclusionStatement(r.Spec.ForProvider.ConclusionStatement, ctyVal)
	EncodeLexIntent_ConfirmationPrompt(r.Spec.ForProvider.ConfirmationPrompt, ctyVal)
	EncodeLexIntent_DialogCodeHook(r.Spec.ForProvider.DialogCodeHook, ctyVal)
	EncodeLexIntent_FollowUpPrompt(r.Spec.ForProvider.FollowUpPrompt, ctyVal)
	EncodeLexIntent_FulfillmentActivity(r.Spec.ForProvider.FulfillmentActivity, ctyVal)
	EncodeLexIntent_RejectionStatement(r.Spec.ForProvider.RejectionStatement, ctyVal)
	EncodeLexIntent_Slot(r.Spec.ForProvider.Slot, ctyVal)
	EncodeLexIntent_Timeouts(r.Spec.ForProvider.Timeouts, ctyVal)
	EncodeLexIntent_Checksum(r.Status.AtProvider, ctyVal)
	EncodeLexIntent_CreatedDate(r.Status.AtProvider, ctyVal)
	EncodeLexIntent_LastUpdatedDate(r.Status.AtProvider, ctyVal)
	EncodeLexIntent_Arn(r.Status.AtProvider, ctyVal)
	EncodeLexIntent_Version(r.Status.AtProvider, ctyVal)
	return cty.ObjectVal(ctyVal)
}

func EncodeLexIntent_Id(p LexIntentParameters, vals map[string]cty.Value) {
	vals["id"] = cty.StringVal(p.Id)
}

func EncodeLexIntent_Name(p LexIntentParameters, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLexIntent_ParentIntentSignature(p LexIntentParameters, vals map[string]cty.Value) {
	vals["parent_intent_signature"] = cty.StringVal(p.ParentIntentSignature)
}

func EncodeLexIntent_CreateVersion(p LexIntentParameters, vals map[string]cty.Value) {
	vals["create_version"] = cty.BoolVal(p.CreateVersion)
}

func EncodeLexIntent_Description(p LexIntentParameters, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLexIntent_SampleUtterances(p LexIntentParameters, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SampleUtterances {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["sample_utterances"] = cty.SetVal(colVals)
}

func EncodeLexIntent_ConclusionStatement(p ConclusionStatement, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_ConclusionStatement_ResponseCard(p, ctyVal)
	EncodeLexIntent_ConclusionStatement_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["conclusion_statement"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_ConclusionStatement_ResponseCard(p ConclusionStatement, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexIntent_ConclusionStatement_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexIntent_ConclusionStatement_Message_Content(v, ctyVal)
		EncodeLexIntent_ConclusionStatement_Message_ContentType(v, ctyVal)
		EncodeLexIntent_ConclusionStatement_Message_GroupNumber(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexIntent_ConclusionStatement_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexIntent_ConclusionStatement_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexIntent_ConclusionStatement_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexIntent_ConfirmationPrompt(p ConfirmationPrompt, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_ConfirmationPrompt_MaxAttempts(p, ctyVal)
	EncodeLexIntent_ConfirmationPrompt_ResponseCard(p, ctyVal)
	EncodeLexIntent_ConfirmationPrompt_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["confirmation_prompt"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_ConfirmationPrompt_MaxAttempts(p ConfirmationPrompt, vals map[string]cty.Value) {
	vals["max_attempts"] = cty.NumberIntVal(p.MaxAttempts)
}

func EncodeLexIntent_ConfirmationPrompt_ResponseCard(p ConfirmationPrompt, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexIntent_ConfirmationPrompt_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexIntent_ConfirmationPrompt_Message_Content(v, ctyVal)
		EncodeLexIntent_ConfirmationPrompt_Message_ContentType(v, ctyVal)
		EncodeLexIntent_ConfirmationPrompt_Message_GroupNumber(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexIntent_ConfirmationPrompt_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexIntent_ConfirmationPrompt_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexIntent_ConfirmationPrompt_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexIntent_DialogCodeHook(p DialogCodeHook, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_DialogCodeHook_MessageVersion(p, ctyVal)
	EncodeLexIntent_DialogCodeHook_Uri(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["dialog_code_hook"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_DialogCodeHook_MessageVersion(p DialogCodeHook, vals map[string]cty.Value) {
	vals["message_version"] = cty.StringVal(p.MessageVersion)
}

func EncodeLexIntent_DialogCodeHook_Uri(p DialogCodeHook, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeLexIntent_FollowUpPrompt(p FollowUpPrompt, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_FollowUpPrompt_Prompt(p.Prompt, ctyVal)
	EncodeLexIntent_FollowUpPrompt_RejectionStatement(p.RejectionStatement, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["follow_up_prompt"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_FollowUpPrompt_Prompt(p Prompt, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_FollowUpPrompt_Prompt_MaxAttempts(p, ctyVal)
	EncodeLexIntent_FollowUpPrompt_Prompt_ResponseCard(p, ctyVal)
	EncodeLexIntent_FollowUpPrompt_Prompt_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["prompt"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_FollowUpPrompt_Prompt_MaxAttempts(p Prompt, vals map[string]cty.Value) {
	vals["max_attempts"] = cty.NumberIntVal(p.MaxAttempts)
}

func EncodeLexIntent_FollowUpPrompt_Prompt_ResponseCard(p Prompt, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexIntent_FollowUpPrompt_Prompt_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexIntent_FollowUpPrompt_Prompt_Message_ContentType(v, ctyVal)
		EncodeLexIntent_FollowUpPrompt_Prompt_Message_GroupNumber(v, ctyVal)
		EncodeLexIntent_FollowUpPrompt_Prompt_Message_Content(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexIntent_FollowUpPrompt_Prompt_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexIntent_FollowUpPrompt_Prompt_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexIntent_FollowUpPrompt_Prompt_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexIntent_FollowUpPrompt_RejectionStatement(p RejectionStatement, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_FollowUpPrompt_RejectionStatement_ResponseCard(p, ctyVal)
	EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["rejection_statement"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_FollowUpPrompt_RejectionStatement_ResponseCard(p RejectionStatement, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message_ContentType(v, ctyVal)
		EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message_GroupNumber(v, ctyVal)
		EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message_Content(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexIntent_FollowUpPrompt_RejectionStatement_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexIntent_FulfillmentActivity(p FulfillmentActivity, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_FulfillmentActivity_Type(p, ctyVal)
	EncodeLexIntent_FulfillmentActivity_CodeHook(p.CodeHook, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["fulfillment_activity"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_FulfillmentActivity_Type(p FulfillmentActivity, vals map[string]cty.Value) {
	vals["type"] = cty.StringVal(p.Type)
}

func EncodeLexIntent_FulfillmentActivity_CodeHook(p CodeHook, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_FulfillmentActivity_CodeHook_MessageVersion(p, ctyVal)
	EncodeLexIntent_FulfillmentActivity_CodeHook_Uri(p, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["code_hook"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_FulfillmentActivity_CodeHook_MessageVersion(p CodeHook, vals map[string]cty.Value) {
	vals["message_version"] = cty.StringVal(p.MessageVersion)
}

func EncodeLexIntent_FulfillmentActivity_CodeHook_Uri(p CodeHook, vals map[string]cty.Value) {
	vals["uri"] = cty.StringVal(p.Uri)
}

func EncodeLexIntent_RejectionStatement(p RejectionStatement, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_RejectionStatement_ResponseCard(p, ctyVal)
	EncodeLexIntent_RejectionStatement_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["rejection_statement"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_RejectionStatement_ResponseCard(p RejectionStatement, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexIntent_RejectionStatement_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexIntent_RejectionStatement_Message_ContentType(v, ctyVal)
		EncodeLexIntent_RejectionStatement_Message_GroupNumber(v, ctyVal)
		EncodeLexIntent_RejectionStatement_Message_Content(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexIntent_RejectionStatement_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexIntent_RejectionStatement_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexIntent_RejectionStatement_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexIntent_Slot(p []Slot, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexIntent_Slot_SampleUtterances(v, ctyVal)
		EncodeLexIntent_Slot_SlotConstraint(v, ctyVal)
		EncodeLexIntent_Slot_SlotType(v, ctyVal)
		EncodeLexIntent_Slot_SlotTypeVersion(v, ctyVal)
		EncodeLexIntent_Slot_Description(v, ctyVal)
		EncodeLexIntent_Slot_Name(v, ctyVal)
		EncodeLexIntent_Slot_Priority(v, ctyVal)
		EncodeLexIntent_Slot_ResponseCard(v, ctyVal)
		EncodeLexIntent_Slot_ValueElicitationPrompt(v.ValueElicitationPrompt, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["slot"] = cty.SetVal(valsForCollection)
}

func EncodeLexIntent_Slot_SampleUtterances(p Slot, vals map[string]cty.Value) {
	colVals := make([]cty.Value, 0)
	for _, value := range p.SampleUtterances {
		colVals = append(colVals, cty.StringVal(value))
	}
	vals["sample_utterances"] = cty.ListVal(colVals)
}

func EncodeLexIntent_Slot_SlotConstraint(p Slot, vals map[string]cty.Value) {
	vals["slot_constraint"] = cty.StringVal(p.SlotConstraint)
}

func EncodeLexIntent_Slot_SlotType(p Slot, vals map[string]cty.Value) {
	vals["slot_type"] = cty.StringVal(p.SlotType)
}

func EncodeLexIntent_Slot_SlotTypeVersion(p Slot, vals map[string]cty.Value) {
	vals["slot_type_version"] = cty.StringVal(p.SlotTypeVersion)
}

func EncodeLexIntent_Slot_Description(p Slot, vals map[string]cty.Value) {
	vals["description"] = cty.StringVal(p.Description)
}

func EncodeLexIntent_Slot_Name(p Slot, vals map[string]cty.Value) {
	vals["name"] = cty.StringVal(p.Name)
}

func EncodeLexIntent_Slot_Priority(p Slot, vals map[string]cty.Value) {
	vals["priority"] = cty.NumberIntVal(p.Priority)
}

func EncodeLexIntent_Slot_ResponseCard(p Slot, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexIntent_Slot_ValueElicitationPrompt(p ValueElicitationPrompt, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 1)
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_Slot_ValueElicitationPrompt_MaxAttempts(p, ctyVal)
	EncodeLexIntent_Slot_ValueElicitationPrompt_ResponseCard(p, ctyVal)
	EncodeLexIntent_Slot_ValueElicitationPrompt_Message(p.Message, ctyVal)
	valsForCollection[0] = cty.ObjectVal(ctyVal)
	vals["value_elicitation_prompt"] = cty.ListVal(valsForCollection)
}

func EncodeLexIntent_Slot_ValueElicitationPrompt_MaxAttempts(p ValueElicitationPrompt, vals map[string]cty.Value) {
	vals["max_attempts"] = cty.NumberIntVal(p.MaxAttempts)
}

func EncodeLexIntent_Slot_ValueElicitationPrompt_ResponseCard(p ValueElicitationPrompt, vals map[string]cty.Value) {
	vals["response_card"] = cty.StringVal(p.ResponseCard)
}

func EncodeLexIntent_Slot_ValueElicitationPrompt_Message(p []Message, vals map[string]cty.Value) {
	valsForCollection := make([]cty.Value, 0)
	for _, v := range p {
		ctyVal := make(map[string]cty.Value)
		EncodeLexIntent_Slot_ValueElicitationPrompt_Message_Content(v, ctyVal)
		EncodeLexIntent_Slot_ValueElicitationPrompt_Message_ContentType(v, ctyVal)
		EncodeLexIntent_Slot_ValueElicitationPrompt_Message_GroupNumber(v, ctyVal)
		valsForCollection = append(valsForCollection, cty.ObjectVal(ctyVal))
	}
	vals["message"] = cty.SetVal(valsForCollection)
}

func EncodeLexIntent_Slot_ValueElicitationPrompt_Message_Content(p Message, vals map[string]cty.Value) {
	vals["content"] = cty.StringVal(p.Content)
}

func EncodeLexIntent_Slot_ValueElicitationPrompt_Message_ContentType(p Message, vals map[string]cty.Value) {
	vals["content_type"] = cty.StringVal(p.ContentType)
}

func EncodeLexIntent_Slot_ValueElicitationPrompt_Message_GroupNumber(p Message, vals map[string]cty.Value) {
	vals["group_number"] = cty.NumberIntVal(p.GroupNumber)
}

func EncodeLexIntent_Timeouts(p Timeouts, vals map[string]cty.Value) {
	ctyVal := make(map[string]cty.Value)
	EncodeLexIntent_Timeouts_Create(p, ctyVal)
	EncodeLexIntent_Timeouts_Delete(p, ctyVal)
	EncodeLexIntent_Timeouts_Update(p, ctyVal)
	vals["timeouts"] = cty.ObjectVal(ctyVal)
}

func EncodeLexIntent_Timeouts_Create(p Timeouts, vals map[string]cty.Value) {
	vals["create"] = cty.StringVal(p.Create)
}

func EncodeLexIntent_Timeouts_Delete(p Timeouts, vals map[string]cty.Value) {
	vals["delete"] = cty.StringVal(p.Delete)
}

func EncodeLexIntent_Timeouts_Update(p Timeouts, vals map[string]cty.Value) {
	vals["update"] = cty.StringVal(p.Update)
}

func EncodeLexIntent_Checksum(p LexIntentObservation, vals map[string]cty.Value) {
	vals["checksum"] = cty.StringVal(p.Checksum)
}

func EncodeLexIntent_CreatedDate(p LexIntentObservation, vals map[string]cty.Value) {
	vals["created_date"] = cty.StringVal(p.CreatedDate)
}

func EncodeLexIntent_LastUpdatedDate(p LexIntentObservation, vals map[string]cty.Value) {
	vals["last_updated_date"] = cty.StringVal(p.LastUpdatedDate)
}

func EncodeLexIntent_Arn(p LexIntentObservation, vals map[string]cty.Value) {
	vals["arn"] = cty.StringVal(p.Arn)
}

func EncodeLexIntent_Version(p LexIntentObservation, vals map[string]cty.Value) {
	vals["version"] = cty.StringVal(p.Version)
}