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
	Name            string `json:"name"`
	SourceAmiId     string `json:"source_ami_id"`
	Encrypted       bool   `json:"encrypted"`
	SourceAmiRegion string `json:"source_ami_region"`
	Description     string `json:"description"`
}

// A AmiCopyStatus defines the observed state of a AmiCopy
type AmiCopyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AmiCopyObservation `json:",inline"`
}

// A AmiCopyObservation records the observed state of a AmiCopy
type AmiCopyObservation struct {
	EnaSupport         bool   `json:"ena_support"`
	ImageLocation      string `json:"image_location"`
	SriovNetSupport    string `json:"sriov_net_support"`
	KernelId           string `json:"kernel_id"`
	VirtualizationType string `json:"virtualization_type"`
	Id                 string `json:"id"`
	RamdiskId          string `json:"ramdisk_id"`
	RootSnapshotId     string `json:"root_snapshot_id"`
	RootDeviceName     string `json:"root_device_name"`
	Architecture       string `json:"architecture"`
	Arn                string `json:"arn"`
	KmsKeyId           string `json:"kms_key_id"`
	ManageEbsSnapshots bool   `json:"manage_ebs_snapshots"`
}