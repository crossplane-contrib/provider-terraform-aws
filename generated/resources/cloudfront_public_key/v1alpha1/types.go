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

// CloudfrontPublicKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudfrontPublicKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudfrontPublicKeySpec   `json:"spec"`
	Status CloudfrontPublicKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudfrontPublicKey contains a list of CloudfrontPublicKeyList
type CloudfrontPublicKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudfrontPublicKey `json:"items"`
}

// A CloudfrontPublicKeySpec defines the desired state of a CloudfrontPublicKey
type CloudfrontPublicKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudfrontPublicKeyParameters `json:",inline"`
}

// A CloudfrontPublicKeyParameters defines the desired state of a CloudfrontPublicKey
type CloudfrontPublicKeyParameters struct {
	Comment    string `json:"comment"`
	EncodedKey string `json:"encoded_key"`
}

// A CloudfrontPublicKeyStatus defines the observed state of a CloudfrontPublicKey
type CloudfrontPublicKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudfrontPublicKeyObservation `json:",inline"`
}

// A CloudfrontPublicKeyObservation records the observed state of a CloudfrontPublicKey
type CloudfrontPublicKeyObservation struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	NamePrefix      string `json:"name_prefix"`
	CallerReference string `json:"caller_reference"`
	Etag            string `json:"etag"`
}