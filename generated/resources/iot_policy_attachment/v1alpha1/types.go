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

// IotPolicyAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotPolicyAttachmentSpec   `json:"spec"`
	Status IotPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotPolicyAttachment contains a list of IotPolicyAttachmentList
type IotPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotPolicyAttachment `json:"items"`
}

// A IotPolicyAttachmentSpec defines the desired state of a IotPolicyAttachment
type IotPolicyAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotPolicyAttachmentParameters `json:"forProvider"`
}

// A IotPolicyAttachmentParameters defines the desired state of a IotPolicyAttachment
type IotPolicyAttachmentParameters struct {
	Id     string `json:"id"`
	Policy string `json:"policy"`
	Target string `json:"target"`
}

// A IotPolicyAttachmentStatus defines the observed state of a IotPolicyAttachment
type IotPolicyAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotPolicyAttachmentObservation `json:"atProvider"`
}

// A IotPolicyAttachmentObservation records the observed state of a IotPolicyAttachment
type IotPolicyAttachmentObservation struct{}