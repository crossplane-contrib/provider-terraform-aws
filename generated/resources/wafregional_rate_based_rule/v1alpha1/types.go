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

// WafregionalRateBasedRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalRateBasedRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalRateBasedRuleSpec   `json:"spec"`
	Status WafregionalRateBasedRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRateBasedRule contains a list of WafregionalRateBasedRuleList
type WafregionalRateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalRateBasedRule `json:"items"`
}

// A WafregionalRateBasedRuleSpec defines the desired state of a WafregionalRateBasedRule
type WafregionalRateBasedRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalRateBasedRuleParameters `json:",inline"`
}

// A WafregionalRateBasedRuleParameters defines the desired state of a WafregionalRateBasedRule
type WafregionalRateBasedRuleParameters struct {
	MetricName string `json:"metric_name"`
	Name       string `json:"name"`
	RateKey    string `json:"rate_key"`
	RateLimit  int    `json:"rate_limit"`
}

// A WafregionalRateBasedRuleStatus defines the observed state of a WafregionalRateBasedRule
type WafregionalRateBasedRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalRateBasedRuleObservation `json:",inline"`
}

// A WafregionalRateBasedRuleObservation records the observed state of a WafregionalRateBasedRule
type WafregionalRateBasedRuleObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}