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

// NeptuneClusterParameterGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NeptuneClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NeptuneClusterParameterGroupSpec   `json:"spec"`
	Status NeptuneClusterParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneClusterParameterGroup contains a list of NeptuneClusterParameterGroupList
type NeptuneClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneClusterParameterGroup `json:"items"`
}

// A NeptuneClusterParameterGroupSpec defines the desired state of a NeptuneClusterParameterGroup
type NeptuneClusterParameterGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NeptuneClusterParameterGroupParameters `json:",inline"`
}

// A NeptuneClusterParameterGroupParameters defines the desired state of a NeptuneClusterParameterGroup
type NeptuneClusterParameterGroupParameters struct {
	Family      string            `json:"family"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Parameter   []Parameter       `json:"parameter"`
}

type Parameter struct {
	ApplyMethod string `json:"apply_method"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

// A NeptuneClusterParameterGroupStatus defines the observed state of a NeptuneClusterParameterGroup
type NeptuneClusterParameterGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NeptuneClusterParameterGroupObservation `json:",inline"`
}

// A NeptuneClusterParameterGroupObservation records the observed state of a NeptuneClusterParameterGroup
type NeptuneClusterParameterGroupObservation struct {
	Arn string `json:"arn"`
}