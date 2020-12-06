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

// RdsClusterParameterGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RdsClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RdsClusterParameterGroupSpec   `json:"spec"`
	Status RdsClusterParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdsClusterParameterGroup contains a list of RdsClusterParameterGroupList
type RdsClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsClusterParameterGroup `json:"items"`
}

// A RdsClusterParameterGroupSpec defines the desired state of a RdsClusterParameterGroup
type RdsClusterParameterGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RdsClusterParameterGroupParameters `json:",inline"`
}

// A RdsClusterParameterGroupParameters defines the desired state of a RdsClusterParameterGroup
type RdsClusterParameterGroupParameters struct {
	Description string            `json:"description"`
	Family      string            `json:"family"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	NamePrefix  string            `json:"name_prefix"`
	Tags        map[string]string `json:"tags"`
	Parameter   Parameter         `json:"parameter"`
}

type Parameter struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	ApplyMethod string `json:"apply_method"`
}

// A RdsClusterParameterGroupStatus defines the observed state of a RdsClusterParameterGroup
type RdsClusterParameterGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RdsClusterParameterGroupObservation `json:",inline"`
}

// A RdsClusterParameterGroupObservation records the observed state of a RdsClusterParameterGroup
type RdsClusterParameterGroupObservation struct {
	Arn string `json:"arn"`
}