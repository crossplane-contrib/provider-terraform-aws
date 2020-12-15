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

// EbsSnapshot is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EbsSnapshot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EbsSnapshotSpec   `json:"spec"`
	Status EbsSnapshotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsSnapshot contains a list of EbsSnapshotList
type EbsSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsSnapshot `json:"items"`
}

// A EbsSnapshotSpec defines the desired state of a EbsSnapshot
type EbsSnapshotSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EbsSnapshotParameters `json:"forProvider"`
}

// A EbsSnapshotParameters defines the desired state of a EbsSnapshot
type EbsSnapshotParameters struct {
	Tags        map[string]string `json:"tags"`
	VolumeId    string            `json:"volume_id"`
	Description string            `json:"description"`
	Timeouts    Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A EbsSnapshotStatus defines the observed state of a EbsSnapshot
type EbsSnapshotStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EbsSnapshotObservation `json:"atProvider"`
}

// A EbsSnapshotObservation records the observed state of a EbsSnapshot
type EbsSnapshotObservation struct {
	VolumeSize          int64  `json:"volume_size"`
	DataEncryptionKeyId string `json:"data_encryption_key_id"`
	KmsKeyId            string `json:"kms_key_id"`
	OwnerAlias          string `json:"owner_alias"`
	Arn                 string `json:"arn"`
	Encrypted           bool   `json:"encrypted"`
	OwnerId             string `json:"owner_id"`
}