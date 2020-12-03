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

// KmsGrant is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KmsGrant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KmsGrantSpec   `json:"spec"`
	Status KmsGrantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KmsGrant contains a list of KmsGrantList
type KmsGrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsGrant `json:"items"`
}

// A KmsGrantSpec defines the desired state of a KmsGrant
type KmsGrantSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KmsGrantParameters `json:",inline"`
}

// A KmsGrantParameters defines the desired state of a KmsGrant
type KmsGrantParameters struct {
	Operations          []string      `json:"operations"`
	RetiringPrincipal   string        `json:"retiring_principal"`
	GranteePrincipal    string        `json:"grantee_principal"`
	Id                  string        `json:"id"`
	KeyId               string        `json:"key_id"`
	Name                string        `json:"name"`
	RetireOnDelete      bool          `json:"retire_on_delete"`
	GrantCreationTokens []string      `json:"grant_creation_tokens"`
	Constraints         []Constraints `json:"constraints"`
}

type Constraints struct {
	EncryptionContextEquals map[string]string `json:"encryption_context_equals"`
	EncryptionContextSubset map[string]string `json:"encryption_context_subset"`
}

// A KmsGrantStatus defines the observed state of a KmsGrant
type KmsGrantStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KmsGrantObservation `json:",inline"`
}

// A KmsGrantObservation records the observed state of a KmsGrant
type KmsGrantObservation struct {
	GrantId    string `json:"grant_id"`
	GrantToken string `json:"grant_token"`
}