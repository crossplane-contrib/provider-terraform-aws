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

// GlueCatalogDatabase is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueCatalogDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueCatalogDatabaseSpec   `json:"spec"`
	Status GlueCatalogDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueCatalogDatabase contains a list of GlueCatalogDatabaseList
type GlueCatalogDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueCatalogDatabase `json:"items"`
}

// A GlueCatalogDatabaseSpec defines the desired state of a GlueCatalogDatabase
type GlueCatalogDatabaseSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueCatalogDatabaseParameters `json:",inline"`
}

// A GlueCatalogDatabaseParameters defines the desired state of a GlueCatalogDatabase
type GlueCatalogDatabaseParameters struct {
	CatalogId   string            `json:"catalog_id"`
	Description string            `json:"description"`
	Id          string            `json:"id"`
	LocationUri string            `json:"location_uri"`
	Name        string            `json:"name"`
	Parameters  map[string]string `json:"parameters"`
}

// A GlueCatalogDatabaseStatus defines the observed state of a GlueCatalogDatabase
type GlueCatalogDatabaseStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueCatalogDatabaseObservation `json:",inline"`
}

// A GlueCatalogDatabaseObservation records the observed state of a GlueCatalogDatabase
type GlueCatalogDatabaseObservation struct {
	Arn string `json:"arn"`
}