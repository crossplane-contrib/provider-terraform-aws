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

// OrganizationsPolicyAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OrganizationsPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OrganizationsPolicyAttachmentSpec   `json:"spec"`
	Status OrganizationsPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrganizationsPolicyAttachment contains a list of OrganizationsPolicyAttachmentList
type OrganizationsPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrganizationsPolicyAttachment `json:"items"`
}

// A OrganizationsPolicyAttachmentSpec defines the desired state of a OrganizationsPolicyAttachment
type OrganizationsPolicyAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OrganizationsPolicyAttachmentParameters `json:",inline"`
}

// A OrganizationsPolicyAttachmentParameters defines the desired state of a OrganizationsPolicyAttachment
type OrganizationsPolicyAttachmentParameters struct {
	PolicyId string `json:"policy_id"`
	TargetId string `json:"target_id"`
}

// A OrganizationsPolicyAttachmentStatus defines the observed state of a OrganizationsPolicyAttachment
type OrganizationsPolicyAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OrganizationsPolicyAttachmentObservation `json:",inline"`
}

// A OrganizationsPolicyAttachmentObservation records the observed state of a OrganizationsPolicyAttachment
type OrganizationsPolicyAttachmentObservation struct {
	Id string `json:"id"`
}