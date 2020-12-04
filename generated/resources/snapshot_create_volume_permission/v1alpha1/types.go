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

// SnapshotCreateVolumePermission is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SnapshotCreateVolumePermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SnapshotCreateVolumePermissionSpec   `json:"spec"`
	Status SnapshotCreateVolumePermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SnapshotCreateVolumePermission contains a list of SnapshotCreateVolumePermissionList
type SnapshotCreateVolumePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SnapshotCreateVolumePermission `json:"items"`
}

// A SnapshotCreateVolumePermissionSpec defines the desired state of a SnapshotCreateVolumePermission
type SnapshotCreateVolumePermissionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SnapshotCreateVolumePermissionParameters `json:",inline"`
}

// A SnapshotCreateVolumePermissionParameters defines the desired state of a SnapshotCreateVolumePermission
type SnapshotCreateVolumePermissionParameters struct {
	AccountId  string `json:"account_id"`
	Id         string `json:"id"`
	SnapshotId string `json:"snapshot_id"`
}

// A SnapshotCreateVolumePermissionStatus defines the observed state of a SnapshotCreateVolumePermission
type SnapshotCreateVolumePermissionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SnapshotCreateVolumePermissionObservation `json:",inline"`
}

// A SnapshotCreateVolumePermissionObservation records the observed state of a SnapshotCreateVolumePermission
type SnapshotCreateVolumePermissionObservation struct{}