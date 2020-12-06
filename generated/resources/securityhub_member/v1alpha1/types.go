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

// SecurityhubMember is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SecurityhubMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityhubMemberSpec   `json:"spec"`
	Status SecurityhubMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityhubMember contains a list of SecurityhubMemberList
type SecurityhubMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityhubMember `json:"items"`
}

// A SecurityhubMemberSpec defines the desired state of a SecurityhubMember
type SecurityhubMemberSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SecurityhubMemberParameters `json:",inline"`
}

// A SecurityhubMemberParameters defines the desired state of a SecurityhubMember
type SecurityhubMemberParameters struct {
	Id        string `json:"id"`
	Invite    bool   `json:"invite"`
	AccountId string `json:"account_id"`
	Email     string `json:"email"`
}

// A SecurityhubMemberStatus defines the observed state of a SecurityhubMember
type SecurityhubMemberStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SecurityhubMemberObservation `json:",inline"`
}

// A SecurityhubMemberObservation records the observed state of a SecurityhubMember
type SecurityhubMemberObservation struct {
	MasterId     string `json:"master_id"`
	MemberStatus string `json:"member_status"`
}