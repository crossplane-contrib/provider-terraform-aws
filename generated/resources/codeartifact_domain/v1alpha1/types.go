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

// CodeartifactDomain is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodeartifactDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodeartifactDomainSpec   `json:"spec"`
	Status CodeartifactDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeartifactDomain contains a list of CodeartifactDomainList
type CodeartifactDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeartifactDomain `json:"items"`
}

// A CodeartifactDomainSpec defines the desired state of a CodeartifactDomain
type CodeartifactDomainSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodeartifactDomainParameters `json:"forProvider"`
}

// A CodeartifactDomainParameters defines the desired state of a CodeartifactDomain
type CodeartifactDomainParameters struct {
	Domain        string `json:"domain"`
	EncryptionKey string `json:"encryption_key"`
}

// A CodeartifactDomainStatus defines the observed state of a CodeartifactDomain
type CodeartifactDomainStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodeartifactDomainObservation `json:"atProvider"`
}

// A CodeartifactDomainObservation records the observed state of a CodeartifactDomain
type CodeartifactDomainObservation struct {
	Arn             string `json:"arn"`
	AssetSizeBytes  int64  `json:"asset_size_bytes"`
	CreatedTime     string `json:"created_time"`
	Owner           string `json:"owner"`
	RepositoryCount int64  `json:"repository_count"`
}