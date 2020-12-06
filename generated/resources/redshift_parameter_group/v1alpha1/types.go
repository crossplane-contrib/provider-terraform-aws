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

// RedshiftParameterGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RedshiftParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedshiftParameterGroupSpec   `json:"spec"`
	Status RedshiftParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftParameterGroup contains a list of RedshiftParameterGroupList
type RedshiftParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftParameterGroup `json:"items"`
}

// A RedshiftParameterGroupSpec defines the desired state of a RedshiftParameterGroup
type RedshiftParameterGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RedshiftParameterGroupParameters `json:",inline"`
}

// A RedshiftParameterGroupParameters defines the desired state of a RedshiftParameterGroup
type RedshiftParameterGroupParameters struct {
	Tags        map[string]string `json:"tags"`
	Description string            `json:"description"`
	Family      string            `json:"family"`
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Parameter   Parameter         `json:"parameter"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A RedshiftParameterGroupStatus defines the observed state of a RedshiftParameterGroup
type RedshiftParameterGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftParameterGroupObservation `json:",inline"`
}

// A RedshiftParameterGroupObservation records the observed state of a RedshiftParameterGroup
type RedshiftParameterGroupObservation struct {
	Arn string `json:"arn"`
}