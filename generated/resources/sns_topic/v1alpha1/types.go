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

// SnsTopic is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SnsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnsTopicSpec   `json:"spec"`
	Status SnsTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnsTopic contains a list of SnsTopicList
type SnsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnsTopic `json:"items"`
}

// A SnsTopicSpec defines the desired state of a SnsTopic
type SnsTopicSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SnsTopicParameters `json:"forProvider"`
}

// A SnsTopicParameters defines the desired state of a SnsTopic
type SnsTopicParameters struct {
	ApplicationSuccessFeedbackRoleArn    string            `json:"application_success_feedback_role_arn"`
	LambdaSuccessFeedbackRoleArn         string            `json:"lambda_success_feedback_role_arn"`
	Name                                 string            `json:"name"`
	SqsFailureFeedbackRoleArn            string            `json:"sqs_failure_feedback_role_arn"`
	HttpFailureFeedbackRoleArn           string            `json:"http_failure_feedback_role_arn"`
	LambdaFailureFeedbackRoleArn         string            `json:"lambda_failure_feedback_role_arn"`
	LambdaSuccessFeedbackSampleRate      int64             `json:"lambda_success_feedback_sample_rate"`
	Tags                                 map[string]string `json:"tags"`
	ApplicationSuccessFeedbackSampleRate int64             `json:"application_success_feedback_sample_rate"`
	DeliveryPolicy                       string            `json:"delivery_policy"`
	DisplayName                          string            `json:"display_name"`
	HttpSuccessFeedbackRoleArn           string            `json:"http_success_feedback_role_arn"`
	HttpSuccessFeedbackSampleRate        int64             `json:"http_success_feedback_sample_rate"`
	SqsSuccessFeedbackRoleArn            string            `json:"sqs_success_feedback_role_arn"`
	SqsSuccessFeedbackSampleRate         int64             `json:"sqs_success_feedback_sample_rate"`
	ApplicationFailureFeedbackRoleArn    string            `json:"application_failure_feedback_role_arn"`
	KmsMasterKeyId                       string            `json:"kms_master_key_id"`
	NamePrefix                           string            `json:"name_prefix"`
	Policy                               string            `json:"policy"`
}

// A SnsTopicStatus defines the observed state of a SnsTopic
type SnsTopicStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SnsTopicObservation `json:"atProvider"`
}

// A SnsTopicObservation records the observed state of a SnsTopic
type SnsTopicObservation struct {
	Arn string `json:"arn"`
}