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

// GlueCatalogTable is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueCatalogTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueCatalogTableSpec   `json:"spec"`
	Status GlueCatalogTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueCatalogTable contains a list of GlueCatalogTableList
type GlueCatalogTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueCatalogTable `json:"items"`
}

// A GlueCatalogTableSpec defines the desired state of a GlueCatalogTable
type GlueCatalogTableSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueCatalogTableParameters `json:",inline"`
}

// A GlueCatalogTableParameters defines the desired state of a GlueCatalogTable
type GlueCatalogTableParameters struct {
	TableType         string            `json:"table_type"`
	DatabaseName      string            `json:"database_name"`
	Description       string            `json:"description"`
	Name              string            `json:"name"`
	Owner             string            `json:"owner"`
	Parameters        map[string]string `json:"parameters"`
	Retention         int               `json:"retention"`
	CatalogId         string            `json:"catalog_id"`
	Id                string            `json:"id"`
	ViewExpandedText  string            `json:"view_expanded_text"`
	ViewOriginalText  string            `json:"view_original_text"`
	PartitionKeys     []PartitionKeys   `json:"partition_keys"`
	StorageDescriptor StorageDescriptor `json:"storage_descriptor"`
}

type PartitionKeys struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type StorageDescriptor struct {
	StoredAsSubDirectories bool              `json:"stored_as_sub_directories"`
	BucketColumns          []string          `json:"bucket_columns"`
	Compressed             bool              `json:"compressed"`
	InputFormat            string            `json:"input_format"`
	Location               string            `json:"location"`
	NumberOfBuckets        int               `json:"number_of_buckets"`
	OutputFormat           string            `json:"output_format"`
	Parameters             map[string]string `json:"parameters"`
	SerDeInfo              SerDeInfo         `json:"ser_de_info"`
	SkewedInfo             SkewedInfo        `json:"skewed_info"`
	SortColumns            []SortColumns     `json:"sort_columns"`
	Columns                []Columns         `json:"columns"`
}

type SerDeInfo struct {
	Name                 string            `json:"name"`
	Parameters           map[string]string `json:"parameters"`
	SerializationLibrary string            `json:"serialization_library"`
}

type SkewedInfo struct {
	SkewedColumnNames             []string          `json:"skewed_column_names"`
	SkewedColumnValueLocationMaps map[string]string `json:"skewed_column_value_location_maps"`
	SkewedColumnValues            []string          `json:"skewed_column_values"`
}

type SortColumns struct {
	Column    string `json:"column"`
	SortOrder int    `json:"sort_order"`
}

type Columns struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

// A GlueCatalogTableStatus defines the observed state of a GlueCatalogTable
type GlueCatalogTableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueCatalogTableObservation `json:",inline"`
}

// A GlueCatalogTableObservation records the observed state of a GlueCatalogTable
type GlueCatalogTableObservation struct {
	Arn string `json:"arn"`
}