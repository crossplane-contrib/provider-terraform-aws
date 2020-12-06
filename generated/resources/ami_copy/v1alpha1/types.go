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

// AmiCopy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AmiCopy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AmiCopySpec   `json:"spec"`
	Status AmiCopyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AmiCopy contains a list of AmiCopyList
type AmiCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AmiCopy `json:"items"`
}

// A AmiCopySpec defines the desired state of a AmiCopy
type AmiCopySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AmiCopyParameters `json:",inline"`
}

// A AmiCopyParameters defines the desired state of a AmiCopy
type AmiCopyParameters struct {
	Description          string               `json:"description"`
	SourceAmiId          string               `json:"source_ami_id"`
	SourceAmiRegion      string               `json:"source_ami_region"`
	KmsKeyId             string               `json:"kms_key_id"`
	Tags                 map[string]string    `json:"tags"`
	Id                   string               `json:"id"`
	Name                 string               `json:"name"`
	Encrypted            bool                 `json:"encrypted"`
	EbsBlockDevice       EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice EphemeralBlockDevice `json:"ephemeral_block_device"`
	Timeouts             Timeouts             `json:"timeouts"`
}

type EbsBlockDevice struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A AmiCopyStatus defines the observed state of a AmiCopy
type AmiCopyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AmiCopyObservation `json:",inline"`
}

// A AmiCopyObservation records the observed state of a AmiCopy
type AmiCopyObservation struct {
	VirtualizationType string `json:"virtualization_type"`
	EnaSupport         bool   `json:"ena_support"`
	KernelId           string `json:"kernel_id"`
	RootDeviceName     string `json:"root_device_name"`
	ManageEbsSnapshots bool   `json:"manage_ebs_snapshots"`
	Architecture       string `json:"architecture"`
	ImageLocation      string `json:"image_location"`
	RootSnapshotId     string `json:"root_snapshot_id"`
	Arn                string `json:"arn"`
	RamdiskId          string `json:"ramdisk_id"`
	SriovNetSupport    string `json:"sriov_net_support"`
}