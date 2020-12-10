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

// EcrRepository is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EcrRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EcrRepositorySpec   `json:"spec"`
	Status EcrRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcrRepository contains a list of EcrRepositoryList
type EcrRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcrRepository `json:"items"`
}

// A EcrRepositorySpec defines the desired state of a EcrRepository
type EcrRepositorySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EcrRepositoryParameters `json:",inline"`
}

// A EcrRepositoryParameters defines the desired state of a EcrRepository
type EcrRepositoryParameters struct {
	Tags                       map[string]string          `json:"tags"`
	Id                         string                     `json:"id"`
	ImageTagMutability         string                     `json:"image_tag_mutability"`
	Name                       string                     `json:"name"`
	ImageScanningConfiguration ImageScanningConfiguration `json:"image_scanning_configuration"`
	Timeouts                   Timeouts                   `json:"timeouts"`
	EncryptionConfiguration    EncryptionConfiguration    `json:"encryption_configuration"`
}

type ImageScanningConfiguration struct {
	ScanOnPush bool `json:"scan_on_push"`
}

type Timeouts struct {
	Delete string `json:"delete"`
}

type EncryptionConfiguration struct {
	EncryptionType string `json:"encryption_type"`
	KmsKey         string `json:"kms_key"`
}

// A EcrRepositoryStatus defines the observed state of a EcrRepository
type EcrRepositoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EcrRepositoryObservation `json:",inline"`
}

// A EcrRepositoryObservation records the observed state of a EcrRepository
type EcrRepositoryObservation struct {
	RegistryId    string `json:"registry_id"`
	RepositoryUrl string `json:"repository_url"`
	Arn           string `json:"arn"`
}