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

// BackupVault is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BackupVault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupVaultSpec   `json:"spec"`
	Status BackupVaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVault contains a list of BackupVaultList
type BackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVault `json:"items"`
}

// A BackupVaultSpec defines the desired state of a BackupVault
type BackupVaultSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BackupVaultParameters `json:",inline"`
}

// A BackupVaultParameters defines the desired state of a BackupVault
type BackupVaultParameters struct {
	Name string            `json:"name"`
	Tags map[string]string `json:"tags"`
}

// A BackupVaultStatus defines the observed state of a BackupVault
type BackupVaultStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BackupVaultObservation `json:",inline"`
}

// A BackupVaultObservation records the observed state of a BackupVault
type BackupVaultObservation struct {
	KmsKeyArn      string `json:"kms_key_arn"`
	RecoveryPoints int    `json:"recovery_points"`
	Arn            string `json:"arn"`
	Id             string `json:"id"`
}