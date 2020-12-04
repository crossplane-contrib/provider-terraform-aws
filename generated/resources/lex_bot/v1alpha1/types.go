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
	EnableModelImprovements      bool                `json:"enable_model_improvements"`
	Id                           string              `json:"id"`
	ChildDirected                bool                `json:"child_directed"`
	Locale                       string              `json:"locale"`
	ProcessBehavior              string              `json:"process_behavior"`
	Name                         string              `json:"name"`
	NluIntentConfidenceThreshold int64               `json:"nlu_intent_confidence_threshold"`
	DetectSentiment              bool                `json:"detect_sentiment"`
	IdleSessionTtlInSeconds      int64               `json:"idle_session_ttl_in_seconds"`
	VoiceId                      string              `json:"voice_id"`
	CreateVersion                bool                `json:"create_version"`
	Description                  string              `json:"description"`
	AbortStatement               AbortStatement      `json:"abort_statement"`
	ClarificationPrompt          ClarificationPrompt `json:"clarification_prompt"`
	Intent                       []Intent            `json:"intent"`
	Timeouts                     Timeouts            `json:"timeouts"`
}

type AbortStatement struct {
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
}

type ClarificationPrompt struct {
	MaxAttempts  int64     `json:"max_attempts"`
	ResponseCard string    `json:"response_card"`
	Message      []Message `json:"message"`
}

type Message struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	GroupNumber int64  `json:"group_number"`
}

type Intent struct {
	IntentVersion string `json:"intent_version"`
	IntentName    string `json:"intent_name"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A LexBotStatus defines the observed state of a LexBot
type LexBotStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LexBotObservation `json:",inline"`
}

// A LexBotObservation records the observed state of a LexBot
type LexBotObservation struct {
	LastUpdatedDate string `json:"last_updated_date"`
	Status          string `json:"status"`
	Version         string `json:"version"`
	Checksum        string `json:"checksum"`
	CreatedDate     string `json:"created_date"`
	Arn             string `json:"arn"`
	FailureReason   string `json:"failure_reason"`
}