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

// Wafv2WebAclAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Wafv2WebAclAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Wafv2WebAclAssociationSpec   `json:"spec"`
	Status Wafv2WebAclAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2WebAclAssociation contains a list of Wafv2WebAclAssociationList
type Wafv2WebAclAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2WebAclAssociation `json:"items"`
}

// A Wafv2WebAclAssociationSpec defines the desired state of a Wafv2WebAclAssociation
type Wafv2WebAclAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Wafv2WebAclAssociationParameters `json:",inline"`
}

// A Wafv2WebAclAssociationParameters defines the desired state of a Wafv2WebAclAssociation
type Wafv2WebAclAssociationParameters struct {
	WebAclArn   string `json:"web_acl_arn"`
	ResourceArn string `json:"resource_arn"`
}

// A Wafv2WebAclAssociationStatus defines the observed state of a Wafv2WebAclAssociation
type Wafv2WebAclAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Wafv2WebAclAssociationObservation `json:",inline"`
}

// A Wafv2WebAclAssociationObservation records the observed state of a Wafv2WebAclAssociation
type Wafv2WebAclAssociationObservation struct {
	Id string `json:"id"`
}