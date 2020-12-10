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

// KeyPair is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KeyPair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyPairSpec   `json:"spec"`
	Status KeyPairStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyPair contains a list of KeyPairList
type KeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyPair `json:"items"`
}

// A KeyPairSpec defines the desired state of a KeyPair
type KeyPairSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KeyPairParameters `json:",inline"`
}

// A KeyPairParameters defines the desired state of a KeyPair
type KeyPairParameters struct {
	Id            string            `json:"id"`
	KeyName       string            `json:"key_name"`
	KeyNamePrefix string            `json:"key_name_prefix"`
	PublicKey     string            `json:"public_key"`
	Tags          map[string]string `json:"tags"`
}

// A KeyPairStatus defines the observed state of a KeyPair
type KeyPairStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KeyPairObservation `json:",inline"`
}

// A KeyPairObservation records the observed state of a KeyPair
type KeyPairObservation struct {
	Arn         string `json:"arn"`
	Fingerprint string `json:"fingerprint"`
	KeyPairId   string `json:"key_pair_id"`
}