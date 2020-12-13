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

// RamResourceShareAccepter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RamResourceShareAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RamResourceShareAccepterSpec   `json:"spec"`
	Status RamResourceShareAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RamResourceShareAccepter contains a list of RamResourceShareAccepterList
type RamResourceShareAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RamResourceShareAccepter `json:"items"`
}

// A RamResourceShareAccepterSpec defines the desired state of a RamResourceShareAccepter
type RamResourceShareAccepterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RamResourceShareAccepterParameters `json:"forProvider"`
}

// A RamResourceShareAccepterParameters defines the desired state of a RamResourceShareAccepter
type RamResourceShareAccepterParameters struct {
	Id       string   `json:"id"`
	ShareArn string   `json:"share_arn"`
	Timeouts Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A RamResourceShareAccepterStatus defines the observed state of a RamResourceShareAccepter
type RamResourceShareAccepterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RamResourceShareAccepterObservation `json:"atProvider"`
}

// A RamResourceShareAccepterObservation records the observed state of a RamResourceShareAccepter
type RamResourceShareAccepterObservation struct {
	ReceiverAccountId string   `json:"receiver_account_id"`
	Resources         []string `json:"resources"`
	SenderAccountId   string   `json:"sender_account_id"`
	ShareName         string   `json:"share_name"`
	Status            string   `json:"status"`
	InvitationArn     string   `json:"invitation_arn"`
	ShareId           string   `json:"share_id"`
}