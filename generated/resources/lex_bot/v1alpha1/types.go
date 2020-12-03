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

// LexBot is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LexBot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LexBotSpec   `json:"spec"`
	Status LexBotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LexBot contains a list of LexBotList
type LexBotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LexBot `json:"items"`
}

// A LexBotSpec defines the desired state of a LexBot
type LexBotSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LexBotParameters `json:",inline"`
}

// A LexBotParameters defines the desired state of a LexBot
type LexBotParameters struct {
	IdleSessionTtlInSeconds      int                 `json:"idle_session_ttl_in_seconds"`
	Locale                       string              `json:"locale"`
	NluIntentConfidenceThreshold int                 `json:"nlu_intent_confidence_threshold"`
	VoiceId                      string              `json:"voice_id"`
	CreateVersion                bool                `json:"create_version"`
	Id                           string              `json:"id"`
	EnableModelImprovements      bool                `json:"enable_model_improvements"`
	ProcessBehavior              string              `json:"process_behavior"`
	ChildDirected                bool                `json:"child_directed"`
	Description                  string              `json:"description"`
	DetectSentiment              bool                `json:"detect_sentiment"`
	Name                         string              `json:"name"`
	Intent                       []Intent            `json:"intent"`
	Timeouts                     []Timeouts          `json:"timeouts"`
	AbortStatement               AbortStatement      `json:"abort_statement"`
	ClarificationPrompt          ClarificationPrompt `json:"clarification_prompt"`
}

type Intent struct {
	IntentName    string `json:"intent_name"`
	IntentVersion string `json:"intent_version"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type AbortStatement struct {
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	GroupNumber int    `json:"group_number"`
}

type ClarificationPrompt struct {
	MaxAttempts  int       `json:"max_attempts"`
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	GroupNumber int    `json:"group_number"`
}

// A LexBotStatus defines the observed state of a LexBot
type LexBotStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LexBotObservation `json:",inline"`
}

// A LexBotObservation records the observed state of a LexBot
type LexBotObservation struct {
	Arn             string `json:"arn"`
	CreatedDate     string `json:"created_date"`
	Status          string `json:"status"`
	LastUpdatedDate string `json:"last_updated_date"`
	Checksum        string `json:"checksum"`
	FailureReason   string `json:"failure_reason"`
	Version         string `json:"version"`
}