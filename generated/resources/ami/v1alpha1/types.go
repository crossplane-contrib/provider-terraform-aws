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
	SriovNetSupport    string `json:"sriov_net_support"`
	KernelId           string `json:"kernel_id"`
	Description        string `json:"description"`
	EnaSupport         bool   `json:"ena_support"`
	VirtualizationType string `json:"virtualization_type"`
	RamdiskId          string `json:"ramdisk_id"`
	RootDeviceName     string `json:"root_device_name"`
	Architecture       string `json:"architecture"`
	Name               string `json:"name"`
}

// A AmiStatus defines the observed state of a Ami
type AmiStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AmiObservation `json:",inline"`
}

// A AmiObservation records the observed state of a Ami
type AmiObservation struct {
	ManageEbsSnapshots bool   `json:"manage_ebs_snapshots"`
	RootSnapshotId     string `json:"root_snapshot_id"`
	ImageLocation      string `json:"image_location"`
	Arn                string `json:"arn"`
	Id                 string `json:"id"`
}