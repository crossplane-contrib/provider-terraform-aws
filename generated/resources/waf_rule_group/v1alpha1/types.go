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

// WafRuleGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafRuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafRuleGroupSpec   `json:"spec"`
	Status WafRuleGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafRuleGroup contains a list of WafRuleGroupList
type WafRuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafRuleGroup `json:"items"`
}

// A WafRuleGroupSpec defines the desired state of a WafRuleGroup
type WafRuleGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafRuleGroupParameters `json:",inline"`
}

// A WafRuleGroupParameters defines the desired state of a WafRuleGroup
type WafRuleGroupParameters struct {
	Id            string            `json:"id"`
	MetricName    string            `json:"metric_name"`
	Name          string            `json:"name"`
	Tags          map[string]string `json:"tags"`
	ActivatedRule ActivatedRule     `json:"activated_rule"`
}

type ActivatedRule struct {
	RuleId   string `json:"rule_id"`
	Type     string `json:"type"`
	Priority int64  `json:"priority"`
	Action   Action `json:"action"`
}

type Action struct {
	Type string `json:"type"`
}

// A WafRuleGroupStatus defines the observed state of a WafRuleGroup
type WafRuleGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafRuleGroupObservation `json:",inline"`
}

// A WafRuleGroupObservation records the observed state of a WafRuleGroup
type WafRuleGroupObservation struct {
	Arn string `json:"arn"`
}