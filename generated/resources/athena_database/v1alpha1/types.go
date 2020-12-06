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

// AthenaDatabase is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AthenaDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AthenaDatabaseSpec   `json:"spec"`
	Status AthenaDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AthenaDatabase contains a list of AthenaDatabaseList
type AthenaDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AthenaDatabase `json:"items"`
}

// A AthenaDatabaseSpec defines the desired state of a AthenaDatabase
type AthenaDatabaseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AthenaDatabaseParameters `json:",inline"`
}

// A AthenaDatabaseParameters defines the desired state of a AthenaDatabase
type AthenaDatabaseParameters struct {
	Id                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	Bucket                  string                  `json:"bucket"`
	ForceDestroy            bool                    `json:"force_destroy"`
	EncryptionConfiguration EncryptionConfiguration `json:"encryption_configuration"`
}

type EncryptionConfiguration struct {
	EncryptionOption string `json:"encryption_option"`
	KmsKey           string `json:"kms_key"`
}

// A AthenaDatabaseStatus defines the observed state of a AthenaDatabase
type AthenaDatabaseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AthenaDatabaseObservation `json:",inline"`
}

// A AthenaDatabaseObservation records the observed state of a AthenaDatabase
type AthenaDatabaseObservation struct{}