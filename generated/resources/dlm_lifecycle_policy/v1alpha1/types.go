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

// DlmLifecyclePolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DlmLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DlmLifecyclePolicySpec   `json:"spec"`
	Status DlmLifecyclePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DlmLifecyclePolicy contains a list of DlmLifecyclePolicyList
type DlmLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DlmLifecyclePolicy `json:"items"`
}

// A DlmLifecyclePolicySpec defines the desired state of a DlmLifecyclePolicy
type DlmLifecyclePolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DlmLifecyclePolicyParameters `json:",inline"`
}

// A DlmLifecyclePolicyParameters defines the desired state of a DlmLifecyclePolicy
type DlmLifecyclePolicyParameters struct {
	State            string            `json:"state"`
	Tags             map[string]string `json:"tags"`
	Description      string            `json:"description"`
	ExecutionRoleArn string            `json:"execution_role_arn"`
	PolicyDetails    PolicyDetails     `json:"policy_details"`
}

type PolicyDetails struct {
	ResourceTypes []string          `json:"resource_types"`
	TargetTags    map[string]string `json:"target_tags"`
	Schedule      []Schedule        `json:"schedule"`
}

type Schedule struct {
	Name       string            `json:"name"`
	TagsToAdd  map[string]string `json:"tags_to_add"`
	CopyTags   bool              `json:"copy_tags"`
	CreateRule CreateRule        `json:"create_rule"`
	RetainRule RetainRule        `json:"retain_rule"`
}

type CreateRule struct {
	Interval     int      `json:"interval"`
	IntervalUnit string   `json:"interval_unit"`
	Times        []string `json:"times"`
}

type RetainRule struct {
	Count int `json:"count"`
}

// A DlmLifecyclePolicyStatus defines the observed state of a DlmLifecyclePolicy
type DlmLifecyclePolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DlmLifecyclePolicyObservation `json:",inline"`
}

// A DlmLifecyclePolicyObservation records the observed state of a DlmLifecyclePolicy
type DlmLifecyclePolicyObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}