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

// Ami is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ami struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AmiSpec   `json:"spec"`
	Status AmiStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ami contains a list of AmiList
type AmiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ami `json:"items"`
}

// A AmiSpec defines the desired state of a Ami
type AmiSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AmiParameters `json:",inline"`
}

// A AmiParameters defines the desired state of a Ami
type AmiParameters struct {
	Id                   string               `json:"id"`
	Tags                 map[string]string    `json:"tags"`
	Architecture         string               `json:"architecture"`
	VirtualizationType   string               `json:"virtualization_type"`
	Description          string               `json:"description"`
	EnaSupport           bool                 `json:"ena_support"`
	RamdiskId            string               `json:"ramdisk_id"`
	RootDeviceName       string               `json:"root_device_name"`
	SriovNetSupport      string               `json:"sriov_net_support"`
	ImageLocation        string               `json:"image_location"`
	KernelId             string               `json:"kernel_id"`
	Name                 string               `json:"name"`
	EbsBlockDevice       EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice EphemeralBlockDevice `json:"ephemeral_block_device"`
	Timeouts             Timeouts             `json:"timeouts"`
}

type EbsBlockDevice struct {
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A AmiStatus defines the observed state of a Ami
type AmiStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AmiObservation `json:",inline"`
}

// A AmiObservation records the observed state of a Ami
type AmiObservation struct {
	Arn                string `json:"arn"`
	RootSnapshotId     string `json:"root_snapshot_id"`
	ManageEbsSnapshots bool   `json:"manage_ebs_snapshots"`
}