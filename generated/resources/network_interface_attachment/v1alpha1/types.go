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

// NetworkInterfaceAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NetworkInterfaceAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkInterfaceAttachmentSpec   `json:"spec"`
	Status NetworkInterfaceAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkInterfaceAttachment contains a list of NetworkInterfaceAttachmentList
type NetworkInterfaceAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterfaceAttachment `json:"items"`
}

// A NetworkInterfaceAttachmentSpec defines the desired state of a NetworkInterfaceAttachment
type NetworkInterfaceAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NetworkInterfaceAttachmentParameters `json:",inline"`
}

// A NetworkInterfaceAttachmentParameters defines the desired state of a NetworkInterfaceAttachment
type NetworkInterfaceAttachmentParameters struct {
	DeviceIndex        int64  `json:"device_index"`
	Id                 string `json:"id"`
	InstanceId         string `json:"instance_id"`
	NetworkInterfaceId string `json:"network_interface_id"`
}

// A NetworkInterfaceAttachmentStatus defines the observed state of a NetworkInterfaceAttachment
type NetworkInterfaceAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NetworkInterfaceAttachmentObservation `json:",inline"`
}

// A NetworkInterfaceAttachmentObservation records the observed state of a NetworkInterfaceAttachment
type NetworkInterfaceAttachmentObservation struct {
	AttachmentId string `json:"attachment_id"`
	Status       string `json:"status"`
}