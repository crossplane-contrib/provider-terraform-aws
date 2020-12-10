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

// WafregionalRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalRuleSpec   `json:"spec"`
	Status WafregionalRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRule contains a list of WafregionalRuleList
type WafregionalRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalRule `json:"items"`
}

// A WafregionalRuleSpec defines the desired state of a WafregionalRule
type WafregionalRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalRuleParameters `json:",inline"`
}

// A WafregionalRuleParameters defines the desired state of a WafregionalRule
type WafregionalRuleParameters struct {
	MetricName string            `json:"metric_name"`
	Name       string            `json:"name"`
	Tags       map[string]string `json:"tags"`
	Id         string            `json:"id"`
	Predicate  Predicate         `json:"predicate"`
}

type Predicate struct {
	DataId  string `json:"data_id"`
	Negated bool   `json:"negated"`
	Type    string `json:"type"`
}

// A WafregionalRuleStatus defines the observed state of a WafregionalRule
type WafregionalRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalRuleObservation `json:",inline"`
}

// A WafregionalRuleObservation records the observed state of a WafregionalRule
type WafregionalRuleObservation struct {
	Arn string `json:"arn"`
}