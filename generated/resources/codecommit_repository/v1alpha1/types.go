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

// CodecommitRepository is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodecommitRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodecommitRepositorySpec   `json:"spec"`
	Status CodecommitRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodecommitRepository contains a list of CodecommitRepositoryList
type CodecommitRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodecommitRepository `json:"items"`
}

// A CodecommitRepositorySpec defines the desired state of a CodecommitRepository
type CodecommitRepositorySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodecommitRepositoryParameters `json:",inline"`
}

// A CodecommitRepositoryParameters defines the desired state of a CodecommitRepository
type CodecommitRepositoryParameters struct {
	DefaultBranch  string `json:"default_branch"`
	Description    string `json:"description"`
	RepositoryName string `json:"repository_name"`
}

// A CodecommitRepositoryStatus defines the observed state of a CodecommitRepository
type CodecommitRepositoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodecommitRepositoryObservation `json:",inline"`
}

// A CodecommitRepositoryObservation records the observed state of a CodecommitRepository
type CodecommitRepositoryObservation struct {
	Id           string `json:"id"`
	Arn          string `json:"arn"`
	CloneUrlHttp string `json:"clone_url_http"`
	CloneUrlSsh  string `json:"clone_url_ssh"`
	RepositoryId string `json:"repository_id"`
}