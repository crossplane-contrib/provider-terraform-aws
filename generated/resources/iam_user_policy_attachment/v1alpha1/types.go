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

// IamUserPolicyAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamUserPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamUserPolicyAttachmentSpec   `json:"spec"`
	Status IamUserPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamUserPolicyAttachment contains a list of IamUserPolicyAttachmentList
type IamUserPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamUserPolicyAttachment `json:"items"`
}

// A IamUserPolicyAttachmentSpec defines the desired state of a IamUserPolicyAttachment
type IamUserPolicyAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamUserPolicyAttachmentParameters `json:",inline"`
}

// A IamUserPolicyAttachmentParameters defines the desired state of a IamUserPolicyAttachment
type IamUserPolicyAttachmentParameters struct {
	Id        string `json:"id"`
	PolicyArn string `json:"policy_arn"`
	User      string `json:"user"`
}

// A IamUserPolicyAttachmentStatus defines the observed state of a IamUserPolicyAttachment
type IamUserPolicyAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamUserPolicyAttachmentObservation `json:",inline"`
}

// A IamUserPolicyAttachmentObservation records the observed state of a IamUserPolicyAttachment
type IamUserPolicyAttachmentObservation struct{}