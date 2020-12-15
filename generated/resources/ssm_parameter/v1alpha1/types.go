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

// SsmParameter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmParameter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmParameterSpec   `json:"spec"`
	Status SsmParameterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmParameter contains a list of SsmParameterList
type SsmParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmParameter `json:"items"`
}

// A SsmParameterSpec defines the desired state of a SsmParameter
type SsmParameterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmParameterParameters `json:"forProvider"`
}

// A SsmParameterParameters defines the desired state of a SsmParameter
type SsmParameterParameters struct {
	Tier           string            `json:"tier"`
	Type           string            `json:"type"`
	Value          string            `json:"value"`
	Description    string            `json:"description"`
	Name           string            `json:"name"`
	DataType       string            `json:"data_type"`
	Id             string            `json:"id"`
	KeyId          string            `json:"key_id"`
	Overwrite      bool              `json:"overwrite"`
	Tags           map[string]string `json:"tags"`
	AllowedPattern string            `json:"allowed_pattern"`
	Arn            string            `json:"arn"`
}

// A SsmParameterStatus defines the observed state of a SsmParameter
type SsmParameterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmParameterObservation `json:"atProvider"`
}

// A SsmParameterObservation records the observed state of a SsmParameter
type SsmParameterObservation struct {
	Version int64 `json:"version"`
}