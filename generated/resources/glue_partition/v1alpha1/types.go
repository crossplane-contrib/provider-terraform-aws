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

// GluePartition is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GluePartition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GluePartitionSpec   `json:"spec"`
	Status GluePartitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GluePartition contains a list of GluePartitionList
type GluePartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GluePartition `json:"items"`
}

// A GluePartitionSpec defines the desired state of a GluePartition
type GluePartitionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GluePartitionParameters `json:",inline"`
}

// A GluePartitionParameters defines the desired state of a GluePartition
type GluePartitionParameters struct {
	TableName         string            `json:"table_name"`
	Parameters        map[string]string `json:"parameters"`
	PartitionValues   []string          `json:"partition_values"`
	CatalogId         string            `json:"catalog_id"`
	DatabaseName      string            `json:"database_name"`
	Id                string            `json:"id"`
	StorageDescriptor StorageDescriptor `json:"storage_descriptor"`
}

type StorageDescriptor struct {
	InputFormat            string            `json:"input_format"`
	Location               string            `json:"location"`
	NumberOfBuckets        int64             `json:"number_of_buckets"`
	OutputFormat           string            `json:"output_format"`
	Parameters             map[string]string `json:"parameters"`
	StoredAsSubDirectories bool              `json:"stored_as_sub_directories"`
	BucketColumns          []string          `json:"bucket_columns"`
	Compressed             bool              `json:"compressed"`
	SerDeInfo              SerDeInfo         `json:"ser_de_info"`
	SkewedInfo             SkewedInfo        `json:"skewed_info"`
	SortColumns            SortColumns       `json:"sort_columns"`
	Columns                Columns           `json:"columns"`
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
	SortOrder int64  `json:"sort_order"`
}

type Columns struct {
	Comment string `json:"comment"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

// A GluePartitionStatus defines the observed state of a GluePartition
type GluePartitionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GluePartitionObservation `json:",inline"`
}

// A GluePartitionObservation records the observed state of a GluePartition
type GluePartitionObservation struct {
	CreationTime     string `json:"creation_time"`
	LastAccessedTime string `json:"last_accessed_time"`
	LastAnalyzedTime string `json:"last_analyzed_time"`
}