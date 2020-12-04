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

// GuarddutyMember is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GuarddutyMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuarddutyMemberSpec   `json:"spec"`
	Status GuarddutyMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyMember contains a list of GuarddutyMemberList
type GuarddutyMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyMember `json:"items"`
}

// A GuarddutyMemberSpec defines the desired state of a GuarddutyMember
type GuarddutyMemberSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GuarddutyMemberParameters `json:",inline"`
}

// A GuarddutyMemberParameters defines the desired state of a GuarddutyMember
type GuarddutyMemberParameters struct {
	Email                    string   `json:"email"`
	Id                       string   `json:"id"`
	InvitationMessage        string   `json:"invitation_message"`
	Invite                   bool     `json:"invite"`
	AccountId                string   `json:"account_id"`
	DetectorId               string   `json:"detector_id"`
	DisableEmailNotification bool     `json:"disable_email_notification"`
	Timeouts                 Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Update string `json:"update"`
}

// A GuarddutyMemberStatus defines the observed state of a GuarddutyMember
type GuarddutyMemberStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GuarddutyMemberObservation `json:",inline"`
}

// A GuarddutyMemberObservation records the observed state of a GuarddutyMember
type GuarddutyMemberObservation struct {
	RelationshipStatus string `json:"relationship_status"`
}