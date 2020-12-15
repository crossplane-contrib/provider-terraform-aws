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

// IamGroupMembership is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamGroupMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamGroupMembershipSpec   `json:"spec"`
	Status IamGroupMembershipStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamGroupMembership contains a list of IamGroupMembershipList
type IamGroupMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamGroupMembership `json:"items"`
}

// A IamGroupMembershipSpec defines the desired state of a IamGroupMembership
type IamGroupMembershipSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamGroupMembershipParameters `json:"forProvider"`
}

// A IamGroupMembershipParameters defines the desired state of a IamGroupMembership
type IamGroupMembershipParameters struct {
	Name  string   `json:"name"`
	Users []string `json:"users"`
	Group string   `json:"group"`
}

// A IamGroupMembershipStatus defines the observed state of a IamGroupMembership
type IamGroupMembershipStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamGroupMembershipObservation `json:"atProvider"`
}

// A IamGroupMembershipObservation records the observed state of a IamGroupMembership
type IamGroupMembershipObservation struct{}