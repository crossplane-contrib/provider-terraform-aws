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

// ConfigConfigRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigConfigRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigConfigRuleSpec   `json:"spec"`
	Status ConfigConfigRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigConfigRule contains a list of ConfigConfigRuleList
type ConfigConfigRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigConfigRule `json:"items"`
}

// A ConfigConfigRuleSpec defines the desired state of a ConfigConfigRule
type ConfigConfigRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigConfigRuleParameters `json:",inline"`
}

// A ConfigConfigRuleParameters defines the desired state of a ConfigConfigRule
type ConfigConfigRuleParameters struct {
	MaximumExecutionFrequency string `json:"maximum_execution_frequency"`
	Name                      string `json:"name"`
	Description               string `json:"description"`
	InputParameters           string `json:"input_parameters"`
}

// A ConfigConfigRuleStatus defines the observed state of a ConfigConfigRule
type ConfigConfigRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigConfigRuleObservation `json:",inline"`
}

// A ConfigConfigRuleObservation records the observed state of a ConfigConfigRule
type ConfigConfigRuleObservation struct {
	RuleId string `json:"rule_id"`
	Arn    string `json:"arn"`
	Id     string `json:"id"`
}