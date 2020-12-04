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

// BackupPlan is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BackupPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupPlanSpec   `json:"spec"`
	Status BackupPlanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPlan contains a list of BackupPlanList
type BackupPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPlan `json:"items"`
}

// A BackupPlanSpec defines the desired state of a BackupPlan
type BackupPlanSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BackupPlanParameters `json:",inline"`
}

// A BackupPlanParameters defines the desired state of a BackupPlan
type BackupPlanParameters struct {
	Id   string            `json:"id"`
	Name string            `json:"name"`
	Tags map[string]string `json:"tags"`
	Rule []Rule            `json:"rule"`
}

type Rule struct {
	TargetVaultName   string            `json:"target_vault_name"`
	CompletionWindow  int64             `json:"completion_window"`
	RecoveryPointTags map[string]string `json:"recovery_point_tags"`
	RuleName          string            `json:"rule_name"`
	Schedule          string            `json:"schedule"`
	StartWindow       int64             `json:"start_window"`
	CopyAction        CopyAction        `json:"copy_action"`
	Lifecycle         Lifecycle         `json:"lifecycle"`
}

type CopyAction struct {
	DestinationVaultArn string    `json:"destination_vault_arn"`
	Lifecycle           Lifecycle `json:"lifecycle"`
}

type Lifecycle struct {
	ColdStorageAfter int64 `json:"cold_storage_after"`
	DeleteAfter      int64 `json:"delete_after"`
}

type Lifecycle struct {
	ColdStorageAfter int64 `json:"cold_storage_after"`
	DeleteAfter      int64 `json:"delete_after"`
}

// A BackupPlanStatus defines the observed state of a BackupPlan
type BackupPlanStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BackupPlanObservation `json:",inline"`
}

// A BackupPlanObservation records the observed state of a BackupPlan
type BackupPlanObservation struct {
	Arn     string `json:"arn"`
	Version string `json:"version"`
}