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

// BackupVaultNotifications is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BackupVaultNotifications struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupVaultNotificationsSpec   `json:"spec"`
	Status BackupVaultNotificationsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVaultNotifications contains a list of BackupVaultNotificationsList
type BackupVaultNotificationsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVaultNotifications `json:"items"`
}

// A BackupVaultNotificationsSpec defines the desired state of a BackupVaultNotifications
type BackupVaultNotificationsSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BackupVaultNotificationsParameters `json:",inline"`
}

// A BackupVaultNotificationsParameters defines the desired state of a BackupVaultNotifications
type BackupVaultNotificationsParameters struct {
	Id                string   `json:"id"`
	SnsTopicArn       string   `json:"sns_topic_arn"`
	BackupVaultEvents []string `json:"backup_vault_events"`
	BackupVaultName   string   `json:"backup_vault_name"`
}

// A BackupVaultNotificationsStatus defines the observed state of a BackupVaultNotifications
type BackupVaultNotificationsStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BackupVaultNotificationsObservation `json:",inline"`
}

// A BackupVaultNotificationsObservation records the observed state of a BackupVaultNotifications
type BackupVaultNotificationsObservation struct {
	BackupVaultArn string `json:"backup_vault_arn"`
}