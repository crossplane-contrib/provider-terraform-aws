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

// Wafv2RuleGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Wafv2RuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Wafv2RuleGroupSpec   `json:"spec"`
	Status Wafv2RuleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2RuleGroup contains a list of Wafv2RuleGroupList
type Wafv2RuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2RuleGroup `json:"items"`
}

// A Wafv2RuleGroupSpec defines the desired state of a Wafv2RuleGroup
type Wafv2RuleGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Wafv2RuleGroupParameters `json:",inline"`
}

// A Wafv2RuleGroupParameters defines the desired state of a Wafv2RuleGroup
type Wafv2RuleGroupParameters struct {
	Capacity    int    `json:"capacity"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Scope       string `json:"scope"`
}

// A Wafv2RuleGroupStatus defines the observed state of a Wafv2RuleGroup
type Wafv2RuleGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Wafv2RuleGroupObservation `json:",inline"`
}

// A Wafv2RuleGroupObservation records the observed state of a Wafv2RuleGroup
type Wafv2RuleGroupObservation struct {
	Arn       string `json:"arn"`
	Id        string `json:"id"`
	LockToken string `json:"lock_token"`
}