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

// SnsSmsPreferences is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SnsSmsPreferences struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnsSmsPreferencesSpec   `json:"spec"`
	Status SnsSmsPreferencesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnsSmsPreferences contains a list of SnsSmsPreferencesList
type SnsSmsPreferencesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnsSmsPreferences `json:"items"`
}

// A SnsSmsPreferencesSpec defines the desired state of a SnsSmsPreferences
type SnsSmsPreferencesSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SnsSmsPreferencesParameters `json:"forProvider"`
}

// A SnsSmsPreferencesParameters defines the desired state of a SnsSmsPreferences
type SnsSmsPreferencesParameters struct {
	DefaultSenderId                   string `json:"default_sender_id"`
	DefaultSmsType                    string `json:"default_sms_type"`
	DeliveryStatusIamRoleArn          string `json:"delivery_status_iam_role_arn"`
	DeliveryStatusSuccessSamplingRate string `json:"delivery_status_success_sampling_rate"`
	Id                                string `json:"id"`
	MonthlySpendLimit                 string `json:"monthly_spend_limit"`
	UsageReportS3Bucket               string `json:"usage_report_s3_bucket"`
}

// A SnsSmsPreferencesStatus defines the observed state of a SnsSmsPreferences
type SnsSmsPreferencesStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SnsSmsPreferencesObservation `json:"atProvider"`
}

// A SnsSmsPreferencesObservation records the observed state of a SnsSmsPreferences
type SnsSmsPreferencesObservation struct{}