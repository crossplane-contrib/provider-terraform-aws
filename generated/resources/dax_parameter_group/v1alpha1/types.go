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

// DaxParameterGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DaxParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DaxParameterGroupSpec   `json:"spec"`
	Status DaxParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DaxParameterGroup contains a list of DaxParameterGroupList
type DaxParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaxParameterGroup `json:"items"`
}

// A DaxParameterGroupSpec defines the desired state of a DaxParameterGroup
type DaxParameterGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DaxParameterGroupParameters `json:",inline"`
}

// A DaxParameterGroupParameters defines the desired state of a DaxParameterGroup
type DaxParameterGroupParameters struct {
	Id          string       `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Parameters  []Parameters `json:"parameters"`
}

type Parameters struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A DaxParameterGroupStatus defines the observed state of a DaxParameterGroup
type DaxParameterGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DaxParameterGroupObservation `json:",inline"`
}

// A DaxParameterGroupObservation records the observed state of a DaxParameterGroup
type DaxParameterGroupObservation struct{}