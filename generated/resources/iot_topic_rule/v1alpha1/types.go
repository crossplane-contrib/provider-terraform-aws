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

// IotTopicRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotTopicRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotTopicRuleSpec   `json:"spec"`
	Status IotTopicRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotTopicRule contains a list of IotTopicRuleList
type IotTopicRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotTopicRule `json:"items"`
}

// A IotTopicRuleSpec defines the desired state of a IotTopicRule
type IotTopicRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotTopicRuleParameters `json:",inline"`
}

// A IotTopicRuleParameters defines the desired state of a IotTopicRule
type IotTopicRuleParameters struct {
	Name        string `json:"name"`
	Sql         string `json:"sql"`
	SqlVersion  string `json:"sql_version"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
}

// A IotTopicRuleStatus defines the observed state of a IotTopicRule
type IotTopicRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotTopicRuleObservation `json:",inline"`
}

// A IotTopicRuleObservation records the observed state of a IotTopicRule
type IotTopicRuleObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}