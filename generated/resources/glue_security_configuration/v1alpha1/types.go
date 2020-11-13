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

// GlueSecurityConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueSecurityConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueSecurityConfigurationSpec   `json:"spec"`
	Status GlueSecurityConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueSecurityConfiguration contains a list of GlueSecurityConfigurationList
type GlueSecurityConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueSecurityConfiguration `json:"items"`
}

// A GlueSecurityConfigurationSpec defines the desired state of a GlueSecurityConfiguration
type GlueSecurityConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueSecurityConfigurationParameters `json:",inline"`
}

// A GlueSecurityConfigurationParameters defines the desired state of a GlueSecurityConfiguration
type GlueSecurityConfigurationParameters struct {
	Id                      string                  `json:"id"`
	Name                    string                  `json:"name"`
	EncryptionConfiguration EncryptionConfiguration `json:"encryption_configuration"`
}

type EncryptionConfiguration struct {
	CloudwatchEncryption   CloudwatchEncryption   `json:"cloudwatch_encryption"`
	JobBookmarksEncryption JobBookmarksEncryption `json:"job_bookmarks_encryption"`
	S3Encryption           S3Encryption           `json:"s3_encryption"`
}

type CloudwatchEncryption struct {
	CloudwatchEncryptionMode string `json:"cloudwatch_encryption_mode"`
	KmsKeyArn                string `json:"kms_key_arn"`
}

type JobBookmarksEncryption struct {
	JobBookmarksEncryptionMode string `json:"job_bookmarks_encryption_mode"`
	KmsKeyArn                  string `json:"kms_key_arn"`
}

type S3Encryption struct {
	KmsKeyArn        string `json:"kms_key_arn"`
	S3EncryptionMode string `json:"s3_encryption_mode"`
}

// A GlueSecurityConfigurationStatus defines the observed state of a GlueSecurityConfiguration
type GlueSecurityConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueSecurityConfigurationObservation `json:",inline"`
}

// A GlueSecurityConfigurationObservation records the observed state of a GlueSecurityConfiguration
type GlueSecurityConfigurationObservation struct{}