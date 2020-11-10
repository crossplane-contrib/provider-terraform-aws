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

// ConfigOrganizationManagedRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigOrganizationManagedRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigOrganizationManagedRuleSpec   `json:"spec"`
	Status ConfigOrganizationManagedRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigOrganizationManagedRule contains a list of ConfigOrganizationManagedRuleList
type ConfigOrganizationManagedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigOrganizationManagedRule `json:"items"`
}

// A ConfigOrganizationManagedRuleSpec defines the desired state of a ConfigOrganizationManagedRule
type ConfigOrganizationManagedRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigOrganizationManagedRuleParameters `json:",inline"`
}

// A ConfigOrganizationManagedRuleParameters defines the desired state of a ConfigOrganizationManagedRule
type ConfigOrganizationManagedRuleParameters struct {
	ResourceIdScope           string     `json:"resource_id_scope"`
	ResourceTypesScope        []string   `json:"resource_types_scope"`
	RuleIdentifier            string     `json:"rule_identifier"`
	ExcludedAccounts          []string   `json:"excluded_accounts"`
	Name                      string     `json:"name"`
	MaximumExecutionFrequency string     `json:"maximum_execution_frequency"`
	TagKeyScope               string     `json:"tag_key_scope"`
	TagValueScope             string     `json:"tag_value_scope"`
	Description               string     `json:"description"`
	InputParameters           string     `json:"input_parameters"`
	Timeouts                  []Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A ConfigOrganizationManagedRuleStatus defines the observed state of a ConfigOrganizationManagedRule
type ConfigOrganizationManagedRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigOrganizationManagedRuleObservation `json:",inline"`
}

// A ConfigOrganizationManagedRuleObservation records the observed state of a ConfigOrganizationManagedRule
type ConfigOrganizationManagedRuleObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}