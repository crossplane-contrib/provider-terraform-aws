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

// IamAccountPasswordPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IamAccountPasswordPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IamAccountPasswordPolicySpec   `json:"spec"`
	Status IamAccountPasswordPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamAccountPasswordPolicy contains a list of IamAccountPasswordPolicyList
type IamAccountPasswordPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamAccountPasswordPolicy `json:"items"`
}

// A IamAccountPasswordPolicySpec defines the desired state of a IamAccountPasswordPolicy
type IamAccountPasswordPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IamAccountPasswordPolicyParameters `json:",inline"`
}

// A IamAccountPasswordPolicyParameters defines the desired state of a IamAccountPasswordPolicy
type IamAccountPasswordPolicyParameters struct {
	AllowUsersToChangePassword bool   `json:"allow_users_to_change_password"`
	Id                         string `json:"id"`
	PasswordReusePrevention    int64  `json:"password_reuse_prevention"`
	RequireUppercaseCharacters bool   `json:"require_uppercase_characters"`
	RequireLowercaseCharacters bool   `json:"require_lowercase_characters"`
	RequireNumbers             bool   `json:"require_numbers"`
	RequireSymbols             bool   `json:"require_symbols"`
	HardExpiry                 bool   `json:"hard_expiry"`
	MaxPasswordAge             int64  `json:"max_password_age"`
	MinimumPasswordLength      int64  `json:"minimum_password_length"`
}

// A IamAccountPasswordPolicyStatus defines the observed state of a IamAccountPasswordPolicy
type IamAccountPasswordPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IamAccountPasswordPolicyObservation `json:",inline"`
}

// A IamAccountPasswordPolicyObservation records the observed state of a IamAccountPasswordPolicy
type IamAccountPasswordPolicyObservation struct {
	ExpirePasswords bool `json:"expire_passwords"`
}