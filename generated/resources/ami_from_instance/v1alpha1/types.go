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
	SourceInstanceId      string `json:"source_instance_id"`
	Description           string `json:"description"`
	Name                  string `json:"name"`
	SnapshotWithoutReboot bool   `json:"snapshot_without_reboot"`
}

// A AmiFromInstanceStatus defines the observed state of a AmiFromInstance
type AmiFromInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AmiFromInstanceObservation `json:",inline"`
}

// A AmiFromInstanceObservation records the observed state of a AmiFromInstance
type AmiFromInstanceObservation struct {
	RamdiskId          string `json:"ramdisk_id"`
	RootDeviceName     string `json:"root_device_name"`
	SriovNetSupport    string `json:"sriov_net_support"`
	Architecture       string `json:"architecture"`
	Arn                string `json:"arn"`
	Id                 string `json:"id"`
	ManageEbsSnapshots bool   `json:"manage_ebs_snapshots"`
	RootSnapshotId     string `json:"root_snapshot_id"`
	VirtualizationType string `json:"virtualization_type"`
	EnaSupport         bool   `json:"ena_support"`
	ImageLocation      string `json:"image_location"`
	KernelId           string `json:"kernel_id"`
}