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

// WafregionalRuleGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalRuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalRuleGroupSpec   `json:"spec"`
	Status WafregionalRuleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalRuleGroup contains a list of WafregionalRuleGroupList
type WafregionalRuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalRuleGroup `json:"items"`
}

// A WafregionalRuleGroupSpec defines the desired state of a WafregionalRuleGroup
type WafregionalRuleGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalRuleGroupParameters `json:",inline"`
}

// A WafregionalRuleGroupParameters defines the desired state of a WafregionalRuleGroup
type WafregionalRuleGroupParameters struct {
	Tags          map[string]string `json:"tags"`
	Id            string            `json:"id"`
	MetricName    string            `json:"metric_name"`
	Name          string            `json:"name"`
	ActivatedRule []ActivatedRule   `json:"activated_rule"`
}

type ActivatedRule struct {
	Priority int    `json:"priority"`
	RuleId   string `json:"rule_id"`
	Type     string `json:"type"`
	Action   Action `json:"action"`
}

type Action struct {
	Type string `json:"type"`
}

// A WafregionalRuleGroupStatus defines the observed state of a WafregionalRuleGroup
type WafregionalRuleGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalRuleGroupObservation `json:",inline"`
}

// A WafregionalRuleGroupObservation records the observed state of a WafregionalRuleGroup
type WafregionalRuleGroupObservation struct {
	Arn string `json:"arn"`
}