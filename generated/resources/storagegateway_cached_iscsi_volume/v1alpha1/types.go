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
	SnapshotId         string            `json:"snapshot_id"`
	SourceVolumeArn    string            `json:"source_volume_arn"`
	Tags               map[string]string `json:"tags"`
	NetworkInterfaceId string            `json:"network_interface_id"`
	TargetName         string            `json:"target_name"`
	VolumeSizeInBytes  int               `json:"volume_size_in_bytes"`
	GatewayArn         string            `json:"gateway_arn"`
	KmsEncrypted       bool              `json:"kms_encrypted"`
	KmsKey             string            `json:"kms_key"`
}

// A StoragegatewayCachedIscsiVolumeStatus defines the observed state of a StoragegatewayCachedIscsiVolume
type StoragegatewayCachedIscsiVolumeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     StoragegatewayCachedIscsiVolumeObservation `json:",inline"`
}

// A StoragegatewayCachedIscsiVolumeObservation records the observed state of a StoragegatewayCachedIscsiVolume
type StoragegatewayCachedIscsiVolumeObservation struct {
	VolumeArn            string `json:"volume_arn"`
	Arn                  string `json:"arn"`
	ChapEnabled          bool   `json:"chap_enabled"`
	Id                   string `json:"id"`
	TargetArn            string `json:"target_arn"`
	VolumeId             string `json:"volume_id"`
	LunNumber            int    `json:"lun_number"`
	NetworkInterfacePort int    `json:"network_interface_port"`
}