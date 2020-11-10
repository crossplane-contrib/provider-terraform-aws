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

// CodestarnotificationsNotificationRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodestarnotificationsNotificationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodestarnotificationsNotificationRuleSpec   `json:"spec"`
	Status CodestarnotificationsNotificationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodestarnotificationsNotificationRule contains a list of CodestarnotificationsNotificationRuleList
type CodestarnotificationsNotificationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodestarnotificationsNotificationRule `json:"items"`
}

// A CodestarnotificationsNotificationRuleSpec defines the desired state of a CodestarnotificationsNotificationRule
type CodestarnotificationsNotificationRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodestarnotificationsNotificationRuleParameters `json:",inline"`
}

// A CodestarnotificationsNotificationRuleParameters defines the desired state of a CodestarnotificationsNotificationRule
type CodestarnotificationsNotificationRuleParameters struct {
	EventTypeIds []string          `json:"event_type_ids"`
	Name         string            `json:"name"`
	Resource     string            `json:"resource"`
	Status       string            `json:"status"`
	Tags         map[string]string `json:"tags"`
	DetailType   string            `json:"detail_type"`
	Target       []Target          `json:"target"`
}

type Target struct {
	Address string `json:"address"`
	Status  string `json:"status"`
	Type    string `json:"type"`
}

// A CodestarnotificationsNotificationRuleStatus defines the observed state of a CodestarnotificationsNotificationRule
type CodestarnotificationsNotificationRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodestarnotificationsNotificationRuleObservation `json:",inline"`
}

// A CodestarnotificationsNotificationRuleObservation records the observed state of a CodestarnotificationsNotificationRule
type CodestarnotificationsNotificationRuleObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}