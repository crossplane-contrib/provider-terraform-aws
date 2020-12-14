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

// SecretsmanagerSecretVersion is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SecretsmanagerSecretVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretsmanagerSecretVersionSpec   `json:"spec"`
	Status SecretsmanagerSecretVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretsmanagerSecretVersion contains a list of SecretsmanagerSecretVersionList
type SecretsmanagerSecretVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretsmanagerSecretVersion `json:"items"`
}

// A SecretsmanagerSecretVersionSpec defines the desired state of a SecretsmanagerSecretVersion
type SecretsmanagerSecretVersionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SecretsmanagerSecretVersionParameters `json:"forProvider"`
}

// A SecretsmanagerSecretVersionParameters defines the desired state of a SecretsmanagerSecretVersion
type SecretsmanagerSecretVersionParameters struct {
	SecretBinary  string   `json:"secret_binary"`
	SecretId      string   `json:"secret_id"`
	SecretString  string   `json:"secret_string"`
	VersionStages []string `json:"version_stages"`
	Id            string   `json:"id"`
}

// A SecretsmanagerSecretVersionStatus defines the observed state of a SecretsmanagerSecretVersion
type SecretsmanagerSecretVersionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SecretsmanagerSecretVersionObservation `json:"atProvider"`
}

// A SecretsmanagerSecretVersionObservation records the observed state of a SecretsmanagerSecretVersion
type SecretsmanagerSecretVersionObservation struct {
	VersionId string `json:"version_id"`
	Arn       string `json:"arn"`
}