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

// ElasticacheParameterGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticacheParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticacheParameterGroupSpec   `json:"spec"`
	Status ElasticacheParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheParameterGroup contains a list of ElasticacheParameterGroupList
type ElasticacheParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheParameterGroup `json:"items"`
}

// A ElasticacheParameterGroupSpec defines the desired state of a ElasticacheParameterGroup
type ElasticacheParameterGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticacheParameterGroupParameters `json:",inline"`
}

// A ElasticacheParameterGroupParameters defines the desired state of a ElasticacheParameterGroup
type ElasticacheParameterGroupParameters struct {
	Description string      `json:"description"`
	Family      string      `json:"family"`
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Parameter   []Parameter `json:"parameter"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A ElasticacheParameterGroupStatus defines the observed state of a ElasticacheParameterGroup
type ElasticacheParameterGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticacheParameterGroupObservation `json:",inline"`
}

// A ElasticacheParameterGroupObservation records the observed state of a ElasticacheParameterGroup
type ElasticacheParameterGroupObservation struct{}