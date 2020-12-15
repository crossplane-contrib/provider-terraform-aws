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

// SesConfigurationSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SesConfigurationSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SesConfigurationSetSpec   `json:"spec"`
	Status SesConfigurationSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SesConfigurationSet contains a list of SesConfigurationSetList
type SesConfigurationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SesConfigurationSet `json:"items"`
}

// A SesConfigurationSetSpec defines the desired state of a SesConfigurationSet
type SesConfigurationSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SesConfigurationSetParameters `json:"forProvider"`
}

// A SesConfigurationSetParameters defines the desired state of a SesConfigurationSet
type SesConfigurationSetParameters struct {
	Name string `json:"name"`
}

// A SesConfigurationSetStatus defines the observed state of a SesConfigurationSet
type SesConfigurationSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SesConfigurationSetObservation `json:"atProvider"`
}

// A SesConfigurationSetObservation records the observed state of a SesConfigurationSet
type SesConfigurationSetObservation struct{}