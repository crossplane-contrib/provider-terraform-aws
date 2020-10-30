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

// GlacierVault is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlacierVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlacierVaultSpec   `json:"spec"`
	Status GlacierVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlacierVault contains a list of GlacierVaultList
type GlacierVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlacierVault `json:"items"`
}

// A GlacierVaultSpec defines the desired state of a GlacierVault
type GlacierVaultSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlacierVaultParameters `json:",inline"`
}

// A GlacierVaultParameters defines the desired state of a GlacierVault
type GlacierVaultParameters struct {
	Name         string `json:"name"`
	AccessPolicy string `json:"access_policy"`
}

// A GlacierVaultStatus defines the observed state of a GlacierVault
type GlacierVaultStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlacierVaultObservation `json:",inline"`
}

// A GlacierVaultObservation records the observed state of a GlacierVault
type GlacierVaultObservation struct {
	Arn      string `json:"arn"`
	Id       string `json:"id"`
	Location string `json:"location"`
}