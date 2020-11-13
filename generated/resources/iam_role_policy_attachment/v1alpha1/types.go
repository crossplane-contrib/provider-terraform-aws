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

// IamRolePolicyAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamRolePolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamRolePolicyAttachmentSpec   `json:"spec"`
	Status IamRolePolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamRolePolicyAttachment contains a list of IamRolePolicyAttachmentList
type IamRolePolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamRolePolicyAttachment `json:"items"`
}

// A IamRolePolicyAttachmentSpec defines the desired state of a IamRolePolicyAttachment
type IamRolePolicyAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamRolePolicyAttachmentParameters `json:",inline"`
}

// A IamRolePolicyAttachmentParameters defines the desired state of a IamRolePolicyAttachment
type IamRolePolicyAttachmentParameters struct {
	Id        string `json:"id"`
	PolicyArn string `json:"policy_arn"`
	Role      string `json:"role"`
}

// A IamRolePolicyAttachmentStatus defines the observed state of a IamRolePolicyAttachment
type IamRolePolicyAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamRolePolicyAttachmentObservation `json:",inline"`
}

// A IamRolePolicyAttachmentObservation records the observed state of a IamRolePolicyAttachment
type IamRolePolicyAttachmentObservation struct{}