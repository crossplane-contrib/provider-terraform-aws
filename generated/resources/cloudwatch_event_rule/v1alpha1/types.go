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

// CloudwatchEventRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchEventRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchEventRuleSpec   `json:"spec"`
	Status CloudwatchEventRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventRule contains a list of CloudwatchEventRuleList
type CloudwatchEventRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventRule `json:"items"`
}

// A CloudwatchEventRuleSpec defines the desired state of a CloudwatchEventRule
type CloudwatchEventRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchEventRuleParameters `json:"forProvider"`
}

// A CloudwatchEventRuleParameters defines the desired state of a CloudwatchEventRule
type CloudwatchEventRuleParameters struct {
	EventPattern       string            `json:"event_pattern"`
	IsEnabled          bool              `json:"is_enabled"`
	RoleArn            string            `json:"role_arn"`
	ScheduleExpression string            `json:"schedule_expression"`
	Description        string            `json:"description"`
	Name               string            `json:"name"`
	NamePrefix         string            `json:"name_prefix"`
	Tags               map[string]string `json:"tags"`
}

// A CloudwatchEventRuleStatus defines the observed state of a CloudwatchEventRule
type CloudwatchEventRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchEventRuleObservation `json:"atProvider"`
}

// A CloudwatchEventRuleObservation records the observed state of a CloudwatchEventRule
type CloudwatchEventRuleObservation struct {
	Arn string `json:"arn"`
}