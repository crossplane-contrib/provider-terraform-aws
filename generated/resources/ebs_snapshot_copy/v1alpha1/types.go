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

// EbsSnapshotCopy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EbsSnapshotCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EbsSnapshotCopySpec   `json:"spec"`
	Status EbsSnapshotCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsSnapshotCopy contains a list of EbsSnapshotCopyList
type EbsSnapshotCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsSnapshotCopy `json:"items"`
}

// A EbsSnapshotCopySpec defines the desired state of a EbsSnapshotCopy
type EbsSnapshotCopySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EbsSnapshotCopyParameters `json:",inline"`
}

// A EbsSnapshotCopyParameters defines the desired state of a EbsSnapshotCopy
type EbsSnapshotCopyParameters struct {
	SourceRegion     string            `json:"source_region"`
	SourceSnapshotId string            `json:"source_snapshot_id"`
	Encrypted        bool              `json:"encrypted"`
	KmsKeyId         string            `json:"kms_key_id"`
	Tags             map[string]string `json:"tags"`
	Description      string            `json:"description"`
	Id               string            `json:"id"`
}

// A EbsSnapshotCopyStatus defines the observed state of a EbsSnapshotCopy
type EbsSnapshotCopyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EbsSnapshotCopyObservation `json:",inline"`
}

// A EbsSnapshotCopyObservation records the observed state of a EbsSnapshotCopy
type EbsSnapshotCopyObservation struct {
	OwnerId             string `json:"owner_id"`
	VolumeId            string `json:"volume_id"`
	DataEncryptionKeyId string `json:"data_encryption_key_id"`
	OwnerAlias          string `json:"owner_alias"`
	VolumeSize          int    `json:"volume_size"`
	Arn                 string `json:"arn"`
}