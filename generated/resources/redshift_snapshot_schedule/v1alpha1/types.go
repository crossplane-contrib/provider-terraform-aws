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

// RedshiftSnapshotSchedule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RedshiftSnapshotSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedshiftSnapshotScheduleSpec   `json:"spec"`
	Status RedshiftSnapshotScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftSnapshotSchedule contains a list of RedshiftSnapshotScheduleList
type RedshiftSnapshotScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftSnapshotSchedule `json:"items"`
}

// A RedshiftSnapshotScheduleSpec defines the desired state of a RedshiftSnapshotSchedule
type RedshiftSnapshotScheduleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RedshiftSnapshotScheduleParameters `json:",inline"`
}

// A RedshiftSnapshotScheduleParameters defines the desired state of a RedshiftSnapshotSchedule
type RedshiftSnapshotScheduleParameters struct {
	Tags             map[string]string `json:"tags"`
	Definitions      []string          `json:"definitions"`
	Description      string            `json:"description"`
	ForceDestroy     bool              `json:"force_destroy"`
	Id               string            `json:"id"`
	Identifier       string            `json:"identifier"`
	IdentifierPrefix string            `json:"identifier_prefix"`
}

// A RedshiftSnapshotScheduleStatus defines the observed state of a RedshiftSnapshotSchedule
type RedshiftSnapshotScheduleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftSnapshotScheduleObservation `json:",inline"`
}

// A RedshiftSnapshotScheduleObservation records the observed state of a RedshiftSnapshotSchedule
type RedshiftSnapshotScheduleObservation struct {
	Arn string `json:"arn"`
}