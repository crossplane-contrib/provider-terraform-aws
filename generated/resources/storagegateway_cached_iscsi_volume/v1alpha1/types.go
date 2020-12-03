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

// StoragegatewayCachedIscsiVolume is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type StoragegatewayCachedIscsiVolume struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StoragegatewayCachedIscsiVolumeSpec   `json:"spec"`
	Status StoragegatewayCachedIscsiVolumeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StoragegatewayCachedIscsiVolume contains a list of StoragegatewayCachedIscsiVolumeList
type StoragegatewayCachedIscsiVolumeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StoragegatewayCachedIscsiVolume `json:"items"`
}

// A StoragegatewayCachedIscsiVolumeSpec defines the desired state of a StoragegatewayCachedIscsiVolume
type StoragegatewayCachedIscsiVolumeSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  StoragegatewayCachedIscsiVolumeParameters `json:",inline"`
}

// A StoragegatewayCachedIscsiVolumeParameters defines the desired state of a StoragegatewayCachedIscsiVolume
type StoragegatewayCachedIscsiVolumeParameters struct {
	KmsEncrypted       bool              `json:"kms_encrypted"`
	NetworkInterfaceId string            `json:"network_interface_id"`
	TargetName         string            `json:"target_name"`
	Tags               map[string]string `json:"tags"`
	KmsKey             string            `json:"kms_key"`
	SnapshotId         string            `json:"snapshot_id"`
	GatewayArn         string            `json:"gateway_arn"`
	Id                 string            `json:"id"`
	SourceVolumeArn    string            `json:"source_volume_arn"`
	VolumeSizeInBytes  int               `json:"volume_size_in_bytes"`
}

// A StoragegatewayCachedIscsiVolumeStatus defines the observed state of a StoragegatewayCachedIscsiVolume
type StoragegatewayCachedIscsiVolumeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     StoragegatewayCachedIscsiVolumeObservation `json:",inline"`
}

// A StoragegatewayCachedIscsiVolumeObservation records the observed state of a StoragegatewayCachedIscsiVolume
type StoragegatewayCachedIscsiVolumeObservation struct {
	LunNumber            int    `json:"lun_number"`
	VolumeId             string `json:"volume_id"`
	ChapEnabled          bool   `json:"chap_enabled"`
	NetworkInterfacePort int    `json:"network_interface_port"`
	TargetArn            string `json:"target_arn"`
	Arn                  string `json:"arn"`
	VolumeArn            string `json:"volume_arn"`
}