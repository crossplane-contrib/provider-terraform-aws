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

// IotRoleAlias is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotRoleAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotRoleAliasSpec   `json:"spec"`
	Status IotRoleAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotRoleAlias contains a list of IotRoleAliasList
type IotRoleAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotRoleAlias `json:"items"`
}

// A IotRoleAliasSpec defines the desired state of a IotRoleAlias
type IotRoleAliasSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotRoleAliasParameters `json:"forProvider"`
}

// A IotRoleAliasParameters defines the desired state of a IotRoleAlias
type IotRoleAliasParameters struct {
	RoleArn            string `json:"role_arn"`
	Alias              string `json:"alias"`
	CredentialDuration int64  `json:"credential_duration"`
}

// A IotRoleAliasStatus defines the observed state of a IotRoleAlias
type IotRoleAliasStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotRoleAliasObservation `json:"atProvider"`
}

// A IotRoleAliasObservation records the observed state of a IotRoleAlias
type IotRoleAliasObservation struct {
	Arn string `json:"arn"`
}