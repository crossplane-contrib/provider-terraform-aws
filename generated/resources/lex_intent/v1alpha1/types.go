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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
)

// +kubebuilder:object:root=true

// LexIntent is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LexIntent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LexIntentSpec   `json:"spec"`
	Status LexIntentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LexIntent contains a list of LexIntentList
type LexIntentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LexIntent `json:"items"`
}

// A LexIntentSpec defines the desired state of a LexIntent
type LexIntentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LexIntentParameters `json:",inline"`
}

// A LexIntentParameters defines the desired state of a LexIntent
type LexIntentParameters struct {
	Id                    string              `json:"id"`
	Name                  string              `json:"name"`
	ParentIntentSignature string              `json:"parent_intent_signature"`
	CreateVersion         bool                `json:"create_version"`
	Description           string              `json:"description"`
	SampleUtterances      []string            `json:"sample_utterances"`
	ConclusionStatement   ConclusionStatement `json:"conclusion_statement"`
	ConfirmationPrompt    ConfirmationPrompt  `json:"confirmation_prompt"`
	DialogCodeHook        DialogCodeHook      `json:"dialog_code_hook"`
	FollowUpPrompt        FollowUpPrompt      `json:"follow_up_prompt"`
	FulfillmentActivity   FulfillmentActivity `json:"fulfillment_activity"`
	RejectionStatement    RejectionStatement  `json:"rejection_statement"`
	Slot                  []Slot              `json:"slot"`
	Timeouts              Timeouts            `json:"timeouts"`
}

type ConclusionStatement struct {
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
}

type ConfirmationPrompt struct {
	MaxAttempts  int64     `json:"max_attempts"`
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
}

type DialogCodeHook struct {
	MessageVersion string `json:"message_version"`
	Uri            string `json:"uri"`
}

type FollowUpPrompt struct {
	Prompt             Prompt             `json:"prompt"`
	RejectionStatement RejectionStatement `json:"rejection_statement"`
}

type Prompt struct {
	MaxAttempts  int64     `json:"max_attempts"`
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
	Content     string `json:"content"`
}

type RejectionStatement struct {
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
	Content     string `json:"content"`
}

type FulfillmentActivity struct {
	Type     string   `json:"type"`
	CodeHook CodeHook `json:"code_hook"`
}

type CodeHook struct {
	MessageVersion string `json:"message_version"`
	Uri            string `json:"uri"`
}

type RejectionStatement struct {
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
	Content     string `json:"content"`
}

type Slot struct {
	SampleUtterances       []string               `json:"sample_utterances"`
	SlotConstraint         string                 `json:"slot_constraint"`
	SlotType               string                 `json:"slot_type"`
	SlotTypeVersion        string                 `json:"slot_type_version"`
	Description            string                 `json:"description"`
	Name                   string                 `json:"name"`
	Priority               int64                  `json:"priority"`
	ResponseCard           string                 `json:"response_card"`
	ValueElicitationPrompt ValueElicitationPrompt `json:"value_elicitation_prompt"`
}

type ValueElicitationPrompt struct {
	MaxAttempts  int64     `json:"max_attempts"`
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A LexIntentStatus defines the observed state of a LexIntent
type LexIntentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LexIntentObservation `json:",inline"`
}

// A LexIntentObservation records the observed state of a LexIntent
type LexIntentObservation struct {
	Checksum        string `json:"checksum"`
	CreatedDate     string `json:"created_date"`
	LastUpdatedDate string `json:"last_updated_date"`
	Arn             string `json:"arn"`
	Version         string `json:"version"`
}