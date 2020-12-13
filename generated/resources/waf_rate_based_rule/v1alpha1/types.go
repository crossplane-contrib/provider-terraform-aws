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

// WafRateBasedRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafRateBasedRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafRateBasedRuleSpec   `json:"spec"`
	Status WafRateBasedRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafRateBasedRule contains a list of WafRateBasedRuleList
type WafRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafRateBasedRule `json:"items"`
}

// A WafRateBasedRuleSpec defines the desired state of a WafRateBasedRule
type WafRateBasedRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafRateBasedRuleParameters `json:"forProvider"`
}

// A WafRateBasedRuleParameters defines the desired state of a WafRateBasedRule
type WafRateBasedRuleParameters struct {
	Id         string            `json:"id"`
	MetricName string            `json:"metric_name"`
	Name       string            `json:"name"`
	RateKey    string            `json:"rate_key"`
	RateLimit  int64             `json:"rate_limit"`
	Tags       map[string]string `json:"tags"`
	Predicates Predicates        `json:"predicates"`
}

type Predicates struct {
	DataId  string `json:"data_id"`
	Negated bool   `json:"negated"`
	Type    string `json:"type"`
}

// A WafRateBasedRuleStatus defines the observed state of a WafRateBasedRule
type WafRateBasedRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafRateBasedRuleObservation `json:"atProvider"`
}

// A WafRateBasedRuleObservation records the observed state of a WafRateBasedRule
type WafRateBasedRuleObservation struct {
	Arn string `json:"arn"`
}