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

// WafRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafRuleSpec   `json:"spec"`
	Status WafRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafRule contains a list of WafRuleList
type WafRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafRule `json:"items"`
}

// A WafRuleSpec defines the desired state of a WafRule
type WafRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafRuleParameters `json:",inline"`
}

// A WafRuleParameters defines the desired state of a WafRule
type WafRuleParameters struct {
	Id         string            `json:"id"`
	MetricName string            `json:"metric_name"`
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	Predicates Predicates        `json:"predicates"`
}

type Predicates struct {
	DataId  string `json:"data_id"`
	Negated bool   `json:"negated"`
	Type    string `json:"type"`
}

// A WafRuleStatus defines the observed state of a WafRule
type WafRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafRuleObservation `json:",inline"`
}

// A WafRuleObservation records the observed state of a WafRule
type WafRuleObservation struct {
	Arn string `json:"arn"`
}