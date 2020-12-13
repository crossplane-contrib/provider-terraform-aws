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

// CodedeployDeploymentConfig is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodedeployDeploymentConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodedeployDeploymentConfigSpec   `json:"spec"`
	Status CodedeployDeploymentConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployDeploymentConfig contains a list of CodedeployDeploymentConfigList
type CodedeployDeploymentConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodedeployDeploymentConfig `json:"items"`
}

// A CodedeployDeploymentConfigSpec defines the desired state of a CodedeployDeploymentConfig
type CodedeployDeploymentConfigSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodedeployDeploymentConfigParameters `json:"forProvider"`
}

// A CodedeployDeploymentConfigParameters defines the desired state of a CodedeployDeploymentConfig
type CodedeployDeploymentConfigParameters struct {
	ComputePlatform      string               `json:"compute_platform"`
	DeploymentConfigName string               `json:"deployment_config_name"`
	Id                   string               `json:"id"`
	MinimumHealthyHosts  MinimumHealthyHosts  `json:"minimum_healthy_hosts"`
	TrafficRoutingConfig TrafficRoutingConfig `json:"traffic_routing_config"`
}

type MinimumHealthyHosts struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

type TrafficRoutingConfig struct {
	Type            string          `json:"type"`
	TimeBasedLinear TimeBasedLinear `json:"time_based_linear"`
	TimeBasedCanary TimeBasedCanary `json:"time_based_canary"`
}

type TimeBasedLinear struct {
	Interval   int64 `json:"interval"`
	Percentage int64 `json:"percentage"`
}

type TimeBasedCanary struct {
	Interval   int64 `json:"interval"`
	Percentage int64 `json:"percentage"`
}

// A CodedeployDeploymentConfigStatus defines the observed state of a CodedeployDeploymentConfig
type CodedeployDeploymentConfigStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodedeployDeploymentConfigObservation `json:"atProvider"`
}

// A CodedeployDeploymentConfigObservation records the observed state of a CodedeployDeploymentConfig
type CodedeployDeploymentConfigObservation struct {
	DeploymentConfigId string `json:"deployment_config_id"`
}