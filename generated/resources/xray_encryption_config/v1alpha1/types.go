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

// XrayEncryptionConfig is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type XrayEncryptionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   XrayEncryptionConfigSpec   `json:"spec"`
	Status XrayEncryptionConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// XrayEncryptionConfig contains a list of XrayEncryptionConfigList
type XrayEncryptionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []XrayEncryptionConfig `json:"items"`
}

// A XrayEncryptionConfigSpec defines the desired state of a XrayEncryptionConfig
type XrayEncryptionConfigSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  XrayEncryptionConfigParameters `json:",inline"`
}

// A XrayEncryptionConfigParameters defines the desired state of a XrayEncryptionConfig
type XrayEncryptionConfigParameters struct {
	Type  string `json:"type"`
	Id    string `json:"id"`
	KeyId string `json:"key_id"`
}

// A XrayEncryptionConfigStatus defines the observed state of a XrayEncryptionConfig
type XrayEncryptionConfigStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     XrayEncryptionConfigObservation `json:",inline"`
}

// A XrayEncryptionConfigObservation records the observed state of a XrayEncryptionConfig
type XrayEncryptionConfigObservation struct{}