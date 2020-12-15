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

// SnsPlatformApplication is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SnsPlatformApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnsPlatformApplicationSpec   `json:"spec"`
	Status SnsPlatformApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnsPlatformApplication contains a list of SnsPlatformApplicationList
type SnsPlatformApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnsPlatformApplication `json:"items"`
}

// A SnsPlatformApplicationSpec defines the desired state of a SnsPlatformApplication
type SnsPlatformApplicationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SnsPlatformApplicationParameters `json:"forProvider"`
}

// A SnsPlatformApplicationParameters defines the desired state of a SnsPlatformApplication
type SnsPlatformApplicationParameters struct {
	EventDeliveryFailureTopicArn string `json:"event_delivery_failure_topic_arn"`
	EventEndpointUpdatedTopicArn string `json:"event_endpoint_updated_topic_arn"`
	FailureFeedbackRoleArn       string `json:"failure_feedback_role_arn"`
	Platform                     string `json:"platform"`
	PlatformCredential           string `json:"platform_credential"`
	SuccessFeedbackSampleRate    string `json:"success_feedback_sample_rate"`
	EventEndpointCreatedTopicArn string `json:"event_endpoint_created_topic_arn"`
	EventEndpointDeletedTopicArn string `json:"event_endpoint_deleted_topic_arn"`
	Name                         string `json:"name"`
	PlatformPrincipal            string `json:"platform_principal"`
	SuccessFeedbackRoleArn       string `json:"success_feedback_role_arn"`
}

// A SnsPlatformApplicationStatus defines the observed state of a SnsPlatformApplication
type SnsPlatformApplicationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SnsPlatformApplicationObservation `json:"atProvider"`
}

// A SnsPlatformApplicationObservation records the observed state of a SnsPlatformApplication
type SnsPlatformApplicationObservation struct {
	Arn string `json:"arn"`
}