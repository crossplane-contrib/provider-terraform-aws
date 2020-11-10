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

// ConfigRemediationConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigRemediationConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigRemediationConfigurationSpec   `json:"spec"`
	Status ConfigRemediationConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigRemediationConfiguration contains a list of ConfigRemediationConfigurationList
type ConfigRemediationConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigRemediationConfiguration `json:"items"`
}

// A ConfigRemediationConfigurationSpec defines the desired state of a ConfigRemediationConfiguration
type ConfigRemediationConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigRemediationConfigurationParameters `json:",inline"`
}

// A ConfigRemediationConfigurationParameters defines the desired state of a ConfigRemediationConfiguration
type ConfigRemediationConfigurationParameters struct {
	ConfigRuleName string      `json:"config_rule_name"`
	ResourceType   string      `json:"resource_type"`
	TargetId       string      `json:"target_id"`
	TargetType     string      `json:"target_type"`
	TargetVersion  string      `json:"target_version"`
	Parameter      []Parameter `json:"parameter"`
}

type Parameter struct {
	StaticValue   string `json:"static_value"`
	Name          string `json:"name"`
	ResourceValue string `json:"resource_value"`
}

// A ConfigRemediationConfigurationStatus defines the observed state of a ConfigRemediationConfiguration
type ConfigRemediationConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigRemediationConfigurationObservation `json:",inline"`
}

// A ConfigRemediationConfigurationObservation records the observed state of a ConfigRemediationConfiguration
type ConfigRemediationConfigurationObservation struct {
	Arn string `json:"arn"`
	Id  string `json:"id"`
}