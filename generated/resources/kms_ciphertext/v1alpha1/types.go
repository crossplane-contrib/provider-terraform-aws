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

// KmsCiphertext is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KmsCiphertext struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KmsCiphertextSpec   `json:"spec"`
	Status KmsCiphertextStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KmsCiphertext contains a list of KmsCiphertextList
type KmsCiphertextList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsCiphertext `json:"items"`
}

// A KmsCiphertextSpec defines the desired state of a KmsCiphertext
type KmsCiphertextSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KmsCiphertextParameters `json:",inline"`
}

// A KmsCiphertextParameters defines the desired state of a KmsCiphertext
type KmsCiphertextParameters struct {
	KeyId     string            `json:"key_id"`
	Plaintext string            `json:"plaintext"`
	Context   map[string]string `json:"context"`
	Id        string            `json:"id"`
}

// A KmsCiphertextStatus defines the observed state of a KmsCiphertext
type KmsCiphertextStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KmsCiphertextObservation `json:",inline"`
}

// A KmsCiphertextObservation records the observed state of a KmsCiphertext
type KmsCiphertextObservation struct {
	CiphertextBlob string `json:"ciphertext_blob"`
}