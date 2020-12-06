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

// KmsKey is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KmsKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KmsKeySpec   `json:"spec"`
	Status KmsKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KmsKey contains a list of KmsKeyList
type KmsKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KmsKey `json:"items"`
}

// A KmsKeySpec defines the desired state of a KmsKey
type KmsKeySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KmsKeyParameters `json:",inline"`
}

// A KmsKeyParameters defines the desired state of a KmsKey
type KmsKeyParameters struct {
	Tags                  map[string]string `json:"tags"`
	CustomerMasterKeySpec string            `json:"customer_master_key_spec"`
	DeletionWindowInDays  int64             `json:"deletion_window_in_days"`
	EnableKeyRotation     bool              `json:"enable_key_rotation"`
	Id                    string            `json:"id"`
	IsEnabled             bool              `json:"is_enabled"`
	KeyUsage              string            `json:"key_usage"`
	Description           string            `json:"description"`
	Policy                string            `json:"policy"`
}

// A KmsKeyStatus defines the observed state of a KmsKey
type KmsKeyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KmsKeyObservation `json:",inline"`
}

// A KmsKeyObservation records the observed state of a KmsKey
type KmsKeyObservation struct {
	Arn   string `json:"arn"`
	KeyId string `json:"key_id"`
}