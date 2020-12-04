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

// ConfigOrganizationCustomRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigOrganizationCustomRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigOrganizationCustomRuleSpec   `json:"spec"`
	Status ConfigOrganizationCustomRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigOrganizationCustomRule contains a list of ConfigOrganizationCustomRuleList
type ConfigOrganizationCustomRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigOrganizationCustomRule `json:"items"`
}

// A ConfigOrganizationCustomRuleSpec defines the desired state of a ConfigOrganizationCustomRule
type ConfigOrganizationCustomRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigOrganizationCustomRuleParameters `json:",inline"`
}

// A ConfigOrganizationCustomRuleParameters defines the desired state of a ConfigOrganizationCustomRule
type ConfigOrganizationCustomRuleParameters struct {
	Description               string   `json:"description"`
	MaximumExecutionFrequency string   `json:"maximum_execution_frequency"`
	TagKeyScope               string   `json:"tag_key_scope"`
	TriggerTypes              []string `json:"trigger_types"`
	ResourceIdScope           string   `json:"resource_id_scope"`
	ResourceTypesScope        []string `json:"resource_types_scope"`
	TagValueScope             string   `json:"tag_value_scope"`
	ExcludedAccounts          []string `json:"excluded_accounts"`
	Id                        string   `json:"id"`
	InputParameters           string   `json:"input_parameters"`
	LambdaFunctionArn         string   `json:"lambda_function_arn"`
	Name                      string   `json:"name"`
	Timeouts                  Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A ConfigOrganizationCustomRuleStatus defines the observed state of a ConfigOrganizationCustomRule
type ConfigOrganizationCustomRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigOrganizationCustomRuleObservation `json:",inline"`
}

// A ConfigOrganizationCustomRuleObservation records the observed state of a ConfigOrganizationCustomRule
type ConfigOrganizationCustomRuleObservation struct {
	Arn string `json:"arn"`
}