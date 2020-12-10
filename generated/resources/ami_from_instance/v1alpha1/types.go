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

// AmiFromInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AmiFromInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AmiFromInstanceSpec   `json:"spec"`
	Status AmiFromInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AmiFromInstance contains a list of AmiFromInstanceList
type AmiFromInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AmiFromInstance `json:"items"`
}

// A AmiFromInstanceSpec defines the desired state of a AmiFromInstance
type AmiFromInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AmiFromInstanceParameters `json:",inline"`
}

// A AmiFromInstanceParameters defines the desired state of a AmiFromInstance
type AmiFromInstanceParameters struct {
	SnapshotWithoutReboot bool                 `json:"snapshot_without_reboot"`
	Name                  string               `json:"name"`
	Description           string               `json:"description"`
	Id                    string               `json:"id"`
	Tags                  map[string]string    `json:"tags"`
	SourceInstanceId      string               `json:"source_instance_id"`
	EbsBlockDevice        EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice  EphemeralBlockDevice `json:"ephemeral_block_device"`
	Timeouts              Timeouts             `json:"timeouts"`
}

type EbsBlockDevice struct {
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
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

// A AmiFromInstanceStatus defines the observed state of a AmiFromInstance
type AmiFromInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AmiFromInstanceObservation `json:",inline"`
}

// A AmiFromInstanceObservation records the observed state of a AmiFromInstance
type AmiFromInstanceObservation struct {
	Architecture       string `json:"architecture"`
	EnaSupport         bool   `json:"ena_support"`
	VirtualizationType string `json:"virtualization_type"`
	Arn                string `json:"arn"`
	ImageLocation      string `json:"image_location"`
	ManageEbsSnapshots bool   `json:"manage_ebs_snapshots"`
	RamdiskId          string `json:"ramdisk_id"`
	RootSnapshotId     string `json:"root_snapshot_id"`
	SriovNetSupport    string `json:"sriov_net_support"`
	KernelId           string `json:"kernel_id"`
	RootDeviceName     string `json:"root_device_name"`
}