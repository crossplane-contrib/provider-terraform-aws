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

// GlueDataCatalogEncryptionSettings is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueDataCatalogEncryptionSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueDataCatalogEncryptionSettingsSpec   `json:"spec"`
	Status GlueDataCatalogEncryptionSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueDataCatalogEncryptionSettings contains a list of GlueDataCatalogEncryptionSettingsList
type GlueDataCatalogEncryptionSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueDataCatalogEncryptionSettings `json:"items"`
}

// A GlueDataCatalogEncryptionSettingsSpec defines the desired state of a GlueDataCatalogEncryptionSettings
type GlueDataCatalogEncryptionSettingsSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueDataCatalogEncryptionSettingsParameters `json:",inline"`
}

// A GlueDataCatalogEncryptionSettingsParameters defines the desired state of a GlueDataCatalogEncryptionSettings
type GlueDataCatalogEncryptionSettingsParameters struct {
	CatalogId                     string                        `json:"catalog_id"`
	Id                            string                        `json:"id"`
	DataCatalogEncryptionSettings DataCatalogEncryptionSettings `json:"data_catalog_encryption_settings"`
}

type DataCatalogEncryptionSettings struct {
	EncryptionAtRest             EncryptionAtRest             `json:"encryption_at_rest"`
	ConnectionPasswordEncryption ConnectionPasswordEncryption `json:"connection_password_encryption"`
}

type EncryptionAtRest struct {
	CatalogEncryptionMode string `json:"catalog_encryption_mode"`
	SseAwsKmsKeyId        string `json:"sse_aws_kms_key_id"`
}

type ConnectionPasswordEncryption struct {
	AwsKmsKeyId                       string `json:"aws_kms_key_id"`
	ReturnConnectionPasswordEncrypted bool   `json:"return_connection_password_encrypted"`
}

// A GlueDataCatalogEncryptionSettingsStatus defines the observed state of a GlueDataCatalogEncryptionSettings
type GlueDataCatalogEncryptionSettingsStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueDataCatalogEncryptionSettingsObservation `json:",inline"`
}

// A GlueDataCatalogEncryptionSettingsObservation records the observed state of a GlueDataCatalogEncryptionSettings
type GlueDataCatalogEncryptionSettingsObservation struct{}