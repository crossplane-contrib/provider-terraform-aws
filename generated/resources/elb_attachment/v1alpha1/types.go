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

// ElbAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElbAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElbAttachmentSpec   `json:"spec"`
	Status ElbAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElbAttachment contains a list of ElbAttachmentList
type ElbAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElbAttachment `json:"items"`
}

// A ElbAttachmentSpec defines the desired state of a ElbAttachment
type ElbAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElbAttachmentParameters `json:",inline"`
}

// A ElbAttachmentParameters defines the desired state of a ElbAttachment
type ElbAttachmentParameters struct {
	Elb      string `json:"elb"`
	Instance string `json:"instance"`
}

// A ElbAttachmentStatus defines the observed state of a ElbAttachment
type ElbAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElbAttachmentObservation `json:",inline"`
}

// A ElbAttachmentObservation records the observed state of a ElbAttachment
type ElbAttachmentObservation struct {
	Id string `json:"id"`
}