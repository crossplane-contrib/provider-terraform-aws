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

// IamUserLoginProfile is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamUserLoginProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamUserLoginProfileSpec   `json:"spec"`
	Status IamUserLoginProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamUserLoginProfile contains a list of IamUserLoginProfileList
type IamUserLoginProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamUserLoginProfile `json:"items"`
}

// A IamUserLoginProfileSpec defines the desired state of a IamUserLoginProfile
type IamUserLoginProfileSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamUserLoginProfileParameters `json:",inline"`
}

// A IamUserLoginProfileParameters defines the desired state of a IamUserLoginProfile
type IamUserLoginProfileParameters struct {
	PasswordLength        int64  `json:"password_length"`
	PasswordResetRequired bool   `json:"password_reset_required"`
	PgpKey                string `json:"pgp_key"`
	User                  string `json:"user"`
	Id                    string `json:"id"`
}

// A IamUserLoginProfileStatus defines the observed state of a IamUserLoginProfile
type IamUserLoginProfileStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamUserLoginProfileObservation `json:",inline"`
}

// A IamUserLoginProfileObservation records the observed state of a IamUserLoginProfile
type IamUserLoginProfileObservation struct {
	EncryptedPassword string `json:"encrypted_password"`
	KeyFingerprint    string `json:"key_fingerprint"`
}