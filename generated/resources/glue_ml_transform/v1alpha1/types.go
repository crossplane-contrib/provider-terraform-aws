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

// GlueMlTransform is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueMlTransform struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueMlTransformSpec   `json:"spec"`
	Status GlueMlTransformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueMlTransform contains a list of GlueMlTransformList
type GlueMlTransformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueMlTransform `json:"items"`
}

// A GlueMlTransformSpec defines the desired state of a GlueMlTransform
type GlueMlTransformSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueMlTransformParameters `json:",inline"`
}

// A GlueMlTransformParameters defines the desired state of a GlueMlTransform
type GlueMlTransformParameters struct {
	Description       string              `json:"description"`
	WorkerType        string              `json:"worker_type"`
	MaxCapacity       int64               `json:"max_capacity"`
	RoleArn           string              `json:"role_arn"`
	Timeout           int64               `json:"timeout"`
	GlueVersion       string              `json:"glue_version"`
	NumberOfWorkers   int64               `json:"number_of_workers"`
	Tags              map[string]string   `json:"tags"`
	Id                string              `json:"id"`
	MaxRetries        int64               `json:"max_retries"`
	Name              string              `json:"name"`
	InputRecordTables []InputRecordTables `json:"input_record_tables"`
	Parameters        Parameters          `json:"parameters"`
}

type InputRecordTables struct {
	CatalogId      string `json:"catalog_id"`
	ConnectionName string `json:"connection_name"`
	DatabaseName   string `json:"database_name"`
	TableName      string `json:"table_name"`
}

type Parameters struct {
	TransformType         string                `json:"transform_type"`
	FindMatchesParameters FindMatchesParameters `json:"find_matches_parameters"`
}

type FindMatchesParameters struct {
	AccuracyCostTradeOff    int64  `json:"accuracy_cost_trade_off"`
	EnforceProvidedLabels   bool   `json:"enforce_provided_labels"`
	PrecisionRecallTradeOff int64  `json:"precision_recall_trade_off"`
	PrimaryKeyColumnName    string `json:"primary_key_column_name"`
}

// A GlueMlTransformStatus defines the observed state of a GlueMlTransform
type GlueMlTransformStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueMlTransformObservation `json:",inline"`
}

// A GlueMlTransformObservation records the observed state of a GlueMlTransform
type GlueMlTransformObservation struct {
	Arn        string   `json:"arn"`
	Schema     []Schema `json:"schema"`
	LabelCount int64    `json:"label_count"`
}

type Schema struct {
	DataType string `json:"data_type"`
	Name     string `json:"name"`
}