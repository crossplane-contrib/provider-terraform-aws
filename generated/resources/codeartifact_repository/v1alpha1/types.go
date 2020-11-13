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

// CodeartifactRepository is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodeartifactRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodeartifactRepositorySpec   `json:"spec"`
	Status CodeartifactRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodeartifactRepository contains a list of CodeartifactRepositoryList
type CodeartifactRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodeartifactRepository `json:"items"`
}

// A CodeartifactRepositorySpec defines the desired state of a CodeartifactRepository
type CodeartifactRepositorySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodeartifactRepositoryParameters `json:",inline"`
}

// A CodeartifactRepositoryParameters defines the desired state of a CodeartifactRepository
type CodeartifactRepositoryParameters struct {
	Id          string     `json:"id"`
	Repository  string     `json:"repository"`
	Description string     `json:"description"`
	Domain      string     `json:"domain"`
	DomainOwner string     `json:"domain_owner"`
	Upstream    []Upstream `json:"upstream"`
}

type Upstream struct {
	RepositoryName string `json:"repository_name"`
}

// A CodeartifactRepositoryStatus defines the observed state of a CodeartifactRepository
type CodeartifactRepositoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodeartifactRepositoryObservation `json:",inline"`
}

// A CodeartifactRepositoryObservation records the observed state of a CodeartifactRepository
type CodeartifactRepositoryObservation struct {
	ExternalConnections  []ExternalConnections `json:"external_connections"`
	AdministratorAccount string                `json:"administrator_account"`
	Arn                  string                `json:"arn"`
}

type ExternalConnections struct {
	PackageFormat          string `json:"package_format"`
	Status                 string `json:"status"`
	ExternalConnectionName string `json:"external_connection_name"`
}