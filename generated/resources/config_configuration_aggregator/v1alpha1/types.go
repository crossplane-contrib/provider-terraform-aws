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

// ConfigConfigurationAggregator is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigConfigurationAggregator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigConfigurationAggregatorSpec   `json:"spec"`
	Status ConfigConfigurationAggregatorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigConfigurationAggregator contains a list of ConfigConfigurationAggregatorList
type ConfigConfigurationAggregatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigConfigurationAggregator `json:"items"`
}

// A ConfigConfigurationAggregatorSpec defines the desired state of a ConfigConfigurationAggregator
type ConfigConfigurationAggregatorSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigConfigurationAggregatorParameters `json:",inline"`
}

// A ConfigConfigurationAggregatorParameters defines the desired state of a ConfigConfigurationAggregator
type ConfigConfigurationAggregatorParameters struct {
	Name                          string                        `json:"name"`
	Tags                          map[string]string             `json:"tags"`
	AccountAggregationSource      AccountAggregationSource      `json:"account_aggregation_source"`
	OrganizationAggregationSource OrganizationAggregationSource `json:"organization_aggregation_source"`
}

type AccountAggregationSource struct {
	AccountIds []string `json:"account_ids"`
	AllRegions bool     `json:"all_regions"`
	Regions    []string `json:"regions"`
}

type OrganizationAggregationSource struct {
	AllRegions bool     `json:"all_regions"`
	Regions    []string `json:"regions"`
	RoleArn    string   `json:"role_arn"`
}

// A ConfigConfigurationAggregatorStatus defines the observed state of a ConfigConfigurationAggregator
type ConfigConfigurationAggregatorStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigConfigurationAggregatorObservation `json:",inline"`
}

// A ConfigConfigurationAggregatorObservation records the observed state of a ConfigConfigurationAggregator
type ConfigConfigurationAggregatorObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}