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

// LightsailKeyPair is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LightsailKeyPair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LightsailKeyPairSpec   `json:"spec"`
	Status LightsailKeyPairStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LightsailKeyPair contains a list of LightsailKeyPairList
type LightsailKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LightsailKeyPair `json:"items"`
}

// A LightsailKeyPairSpec defines the desired state of a LightsailKeyPair
type LightsailKeyPairSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LightsailKeyPairParameters `json:",inline"`
}

// A LightsailKeyPairParameters defines the desired state of a LightsailKeyPair
type LightsailKeyPairParameters struct {
	PgpKey     string `json:"pgp_key"`
	NamePrefix string `json:"name_prefix"`
}

// A LightsailKeyPairStatus defines the observed state of a LightsailKeyPair
type LightsailKeyPairStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LightsailKeyPairObservation `json:",inline"`
}

// A LightsailKeyPairObservation records the observed state of a LightsailKeyPair
type LightsailKeyPairObservation struct {
	PrivateKey           string `json:"private_key"`
	EncryptedPrivateKey  string `json:"encrypted_private_key"`
	Name                 string `json:"name"`
	Fingerprint          string `json:"fingerprint"`
	Id                   string `json:"id"`
	PublicKey            string `json:"public_key"`
	Arn                  string `json:"arn"`
	EncryptedFingerprint string `json:"encrypted_fingerprint"`
}