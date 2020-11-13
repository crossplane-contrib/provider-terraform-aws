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

// EbsEncryptionByDefault is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EbsEncryptionByDefault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EbsEncryptionByDefaultSpec   `json:"spec"`
	Status EbsEncryptionByDefaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EbsEncryptionByDefault contains a list of EbsEncryptionByDefaultList
type EbsEncryptionByDefaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EbsEncryptionByDefault `json:"items"`
}

// A EbsEncryptionByDefaultSpec defines the desired state of a EbsEncryptionByDefault
type EbsEncryptionByDefaultSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EbsEncryptionByDefaultParameters `json:",inline"`
}

// A EbsEncryptionByDefaultParameters defines the desired state of a EbsEncryptionByDefault
type EbsEncryptionByDefaultParameters struct {
	Enabled bool   `json:"enabled"`
	Id      string `json:"id"`
}

// A EbsEncryptionByDefaultStatus defines the observed state of a EbsEncryptionByDefault
type EbsEncryptionByDefaultStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EbsEncryptionByDefaultObservation `json:",inline"`
}

// A EbsEncryptionByDefaultObservation records the observed state of a EbsEncryptionByDefault
type EbsEncryptionByDefaultObservation struct{}