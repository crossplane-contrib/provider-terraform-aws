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

// VolumeAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VolumeAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VolumeAttachmentSpec   `json:"spec"`
	Status VolumeAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VolumeAttachment contains a list of VolumeAttachmentList
type VolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VolumeAttachment `json:"items"`
}

// A VolumeAttachmentSpec defines the desired state of a VolumeAttachment
type VolumeAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VolumeAttachmentParameters `json:"forProvider"`
}

// A VolumeAttachmentParameters defines the desired state of a VolumeAttachment
type VolumeAttachmentParameters struct {
	DeviceName  string `json:"device_name"`
	ForceDetach bool   `json:"force_detach"`
	InstanceId  string `json:"instance_id"`
	SkipDestroy bool   `json:"skip_destroy"`
	VolumeId    string `json:"volume_id"`
}

// A VolumeAttachmentStatus defines the observed state of a VolumeAttachment
type VolumeAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VolumeAttachmentObservation `json:"atProvider"`
}

// A VolumeAttachmentObservation records the observed state of a VolumeAttachment
type VolumeAttachmentObservation struct{}