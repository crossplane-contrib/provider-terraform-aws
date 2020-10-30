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

// MskConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MskConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MskConfigurationSpec   `json:"spec"`
	Status MskConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MskConfiguration contains a list of MskConfigurationList
type MskConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MskConfiguration `json:"items"`
}

// A MskConfigurationSpec defines the desired state of a MskConfiguration
type MskConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MskConfigurationParameters `json:",inline"`
}

// A MskConfigurationParameters defines the desired state of a MskConfiguration
type MskConfigurationParameters struct {
	Description      string   `json:"description"`
	KafkaVersions    []string `json:"kafka_versions"`
	Name             string   `json:"name"`
	ServerProperties string   `json:"server_properties"`
}

// A MskConfigurationStatus defines the observed state of a MskConfiguration
type MskConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MskConfigurationObservation `json:",inline"`
}

// A MskConfigurationObservation records the observed state of a MskConfiguration
type MskConfigurationObservation struct {
	Arn            string `json:"arn"`
	Id             string `json:"id"`
	LatestRevision int    `json:"latest_revision"`
}