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

// EbsVolume is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EbsVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EbsVolumeSpec   `json:"spec"`
	Status EbsVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsVolume contains a list of EbsVolumeList
type EbsVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsVolume `json:"items"`
}

// A EbsVolumeSpec defines the desired state of a EbsVolume
type EbsVolumeSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EbsVolumeParameters `json:"forProvider"`
}

// A EbsVolumeParameters defines the desired state of a EbsVolume
type EbsVolumeParameters struct {
	Id                 string            `json:"id"`
	Iops               int64             `json:"iops"`
	KmsKeyId           string            `json:"kms_key_id"`
	MultiAttachEnabled bool              `json:"multi_attach_enabled"`
	OutpostArn         string            `json:"outpost_arn"`
	Encrypted          bool              `json:"encrypted"`
	SnapshotId         string            `json:"snapshot_id"`
	Tags               map[string]string `json:"tags"`
	Type               string            `json:"type"`
	AvailabilityZone   string            `json:"availability_zone"`
	Size               int64             `json:"size"`
}

// A EbsVolumeStatus defines the observed state of a EbsVolume
type EbsVolumeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EbsVolumeObservation `json:"atProvider"`
}

// A EbsVolumeObservation records the observed state of a EbsVolume
type EbsVolumeObservation struct {
	Arn string `json:"arn"`
}