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

// RedshiftSnapshotScheduleAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RedshiftSnapshotScheduleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedshiftSnapshotScheduleAssociationSpec   `json:"spec"`
	Status RedshiftSnapshotScheduleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftSnapshotScheduleAssociation contains a list of RedshiftSnapshotScheduleAssociationList
type RedshiftSnapshotScheduleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftSnapshotScheduleAssociation `json:"items"`
}

// A RedshiftSnapshotScheduleAssociationSpec defines the desired state of a RedshiftSnapshotScheduleAssociation
type RedshiftSnapshotScheduleAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RedshiftSnapshotScheduleAssociationParameters `json:"forProvider"`
}

// A RedshiftSnapshotScheduleAssociationParameters defines the desired state of a RedshiftSnapshotScheduleAssociation
type RedshiftSnapshotScheduleAssociationParameters struct {
	ClusterIdentifier  string `json:"cluster_identifier"`
	ScheduleIdentifier string `json:"schedule_identifier"`
}

// A RedshiftSnapshotScheduleAssociationStatus defines the observed state of a RedshiftSnapshotScheduleAssociation
type RedshiftSnapshotScheduleAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftSnapshotScheduleAssociationObservation `json:"atProvider"`
}

// A RedshiftSnapshotScheduleAssociationObservation records the observed state of a RedshiftSnapshotScheduleAssociation
type RedshiftSnapshotScheduleAssociationObservation struct{}