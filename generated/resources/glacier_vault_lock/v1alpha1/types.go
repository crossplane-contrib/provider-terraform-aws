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

// GlacierVaultLock is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlacierVaultLock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlacierVaultLockSpec   `json:"spec"`
	Status GlacierVaultLockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlacierVaultLock contains a list of GlacierVaultLockList
type GlacierVaultLockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlacierVaultLock `json:"items"`
}

// A GlacierVaultLockSpec defines the desired state of a GlacierVaultLock
type GlacierVaultLockSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlacierVaultLockParameters `json:",inline"`
}

// A GlacierVaultLockParameters defines the desired state of a GlacierVaultLock
type GlacierVaultLockParameters struct {
	CompleteLock        bool   `json:"complete_lock"`
	IgnoreDeletionError bool   `json:"ignore_deletion_error"`
	Policy              string `json:"policy"`
	VaultName           string `json:"vault_name"`
}

// A GlacierVaultLockStatus defines the observed state of a GlacierVaultLock
type GlacierVaultLockStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlacierVaultLockObservation `json:",inline"`
}

// A GlacierVaultLockObservation records the observed state of a GlacierVaultLock
type GlacierVaultLockObservation struct {
	Id string `json:"id"`
}