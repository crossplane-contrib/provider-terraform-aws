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

// OpsworksUserProfile is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type OpsworksUserProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpsworksUserProfileSpec   `json:"spec"`
	Status OpsworksUserProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpsworksUserProfile contains a list of OpsworksUserProfileList
type OpsworksUserProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpsworksUserProfile `json:"items"`
}

// A OpsworksUserProfileSpec defines the desired state of a OpsworksUserProfile
type OpsworksUserProfileSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  OpsworksUserProfileParameters `json:",inline"`
}

// A OpsworksUserProfileParameters defines the desired state of a OpsworksUserProfile
type OpsworksUserProfileParameters struct {
	AllowSelfManagement bool   `json:"allow_self_management"`
	SshPublicKey        string `json:"ssh_public_key"`
	SshUsername         string `json:"ssh_username"`
	UserArn             string `json:"user_arn"`
}

// A OpsworksUserProfileStatus defines the observed state of a OpsworksUserProfile
type OpsworksUserProfileStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     OpsworksUserProfileObservation `json:",inline"`
}

// A OpsworksUserProfileObservation records the observed state of a OpsworksUserProfile
type OpsworksUserProfileObservation struct {
	Id string `json:"id"`
}