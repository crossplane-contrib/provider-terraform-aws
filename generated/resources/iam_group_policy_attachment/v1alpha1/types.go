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

// IamGroupPolicyAttachment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamGroupPolicyAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamGroupPolicyAttachmentSpec   `json:"spec"`
	Status IamGroupPolicyAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamGroupPolicyAttachment contains a list of IamGroupPolicyAttachmentList
type IamGroupPolicyAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamGroupPolicyAttachment `json:"items"`
}

// A IamGroupPolicyAttachmentSpec defines the desired state of a IamGroupPolicyAttachment
type IamGroupPolicyAttachmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamGroupPolicyAttachmentParameters `json:"forProvider"`
}

// A IamGroupPolicyAttachmentParameters defines the desired state of a IamGroupPolicyAttachment
type IamGroupPolicyAttachmentParameters struct {
	PolicyArn string `json:"policy_arn"`
	Group     string `json:"group"`
}

// A IamGroupPolicyAttachmentStatus defines the observed state of a IamGroupPolicyAttachment
type IamGroupPolicyAttachmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamGroupPolicyAttachmentObservation `json:"atProvider"`
}

// A IamGroupPolicyAttachmentObservation records the observed state of a IamGroupPolicyAttachment
type IamGroupPolicyAttachmentObservation struct{}