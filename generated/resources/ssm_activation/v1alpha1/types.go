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

// SsmActivation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmActivation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmActivationSpec   `json:"spec"`
	Status SsmActivationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmActivation contains a list of SsmActivationList
type SsmActivationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmActivation `json:"items"`
}

// A SsmActivationSpec defines the desired state of a SsmActivation
type SsmActivationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmActivationParameters `json:"forProvider"`
}

// A SsmActivationParameters defines the desired state of a SsmActivation
type SsmActivationParameters struct {
	Name              string            `json:"name"`
	IamRole           string            `json:"iam_role"`
	Id                string            `json:"id"`
	RegistrationLimit int64             `json:"registration_limit"`
	Tags              map[string]string `json:"tags"`
	Description       string            `json:"description"`
	ExpirationDate    string            `json:"expiration_date"`
}

// A SsmActivationStatus defines the observed state of a SsmActivation
type SsmActivationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmActivationObservation `json:"atProvider"`
}

// A SsmActivationObservation records the observed state of a SsmActivation
type SsmActivationObservation struct {
	Expired           bool   `json:"expired"`
	RegistrationCount int64  `json:"registration_count"`
	ActivationCode    string `json:"activation_code"`
}