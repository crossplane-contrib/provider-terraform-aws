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

// KinesisAnalyticsApplication is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KinesisAnalyticsApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KinesisAnalyticsApplicationSpec   `json:"spec"`
	Status KinesisAnalyticsApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisAnalyticsApplication contains a list of KinesisAnalyticsApplicationList
type KinesisAnalyticsApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisAnalyticsApplication `json:"items"`
}

// A KinesisAnalyticsApplicationSpec defines the desired state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KinesisAnalyticsApplicationParameters `json:",inline"`
}

// A KinesisAnalyticsApplicationParameters defines the desired state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationParameters struct {
	Tags                     map[string]string        `json:"tags"`
	Code                     string                   `json:"code"`
	Name                     string                   `json:"name"`
	Description              string                   `json:"description"`
	Inputs                   Inputs                   `json:"inputs"`
	Outputs                  []Outputs                `json:"outputs"`
	ReferenceDataSources     ReferenceDataSources     `json:"reference_data_sources"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type Inputs struct {
	Id                      string                  `json:"id"`
	NamePrefix              string                  `json:"name_prefix"`
	StreamNames             []string                `json:"stream_names"`
	KinesisFirehose         KinesisFirehose         `json:"kinesis_firehose"`
	KinesisStream           KinesisStream           `json:"kinesis_stream"`
	Parallelism             Parallelism             `json:"parallelism"`
	ProcessingConfiguration ProcessingConfiguration `json:"processing_configuration"`
	Schema                  Schema                  `json:"schema"`
}

type KinesisFirehose struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type Parallelism struct {
	Count int `json:"count"`
}

type ProcessingConfiguration struct {
	Lambda Lambda `json:"lambda"`
}

type Lambda struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type Schema struct {
	RecordEncoding string          `json:"record_encoding"`
	RecordFormat   RecordFormat    `json:"record_format"`
	RecordColumns  []RecordColumns `json:"record_columns"`
}

type RecordFormat struct {
	RecordFormatType  string            `json:"record_format_type"`
	MappingParameters MappingParameters `json:"mapping_parameters"`
}

type MappingParameters struct {
	Csv  Csv  `json:"csv"`
	Json Json `json:"json"`
}

type Csv struct {
	RecordColumnDelimiter string `json:"record_column_delimiter"`
	RecordRowDelimiter    string `json:"record_row_delimiter"`
}

type Json struct {
	RecordRowPath string `json:"record_row_path"`
}

type RecordColumns struct {
	Mapping string `json:"mapping"`
	Name    string `json:"name"`
	SqlType string `json:"sql_type"`
}

type Outputs struct {
	Id              string          `json:"id"`
	Name            string          `json:"name"`
	Lambda          Lambda          `json:"lambda"`
	Schema          Schema          `json:"schema"`
	KinesisFirehose KinesisFirehose `json:"kinesis_firehose"`
	KinesisStream   KinesisStream   `json:"kinesis_stream"`
}

type Lambda struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type Schema struct {
	RecordFormatType string `json:"record_format_type"`
}

type KinesisFirehose struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type KinesisStream struct {
	ResourceArn string `json:"resource_arn"`
	RoleArn     string `json:"role_arn"`
}

type ReferenceDataSources struct {
	Id        string `json:"id"`
	TableName string `json:"table_name"`
	S3        S3     `json:"s3"`
	Schema    Schema `json:"schema"`
}

type S3 struct {
	BucketArn string `json:"bucket_arn"`
	FileKey   string `json:"file_key"`
	RoleArn   string `json:"role_arn"`
}

type Schema struct {
	RecordEncoding string          `json:"record_encoding"`
	RecordColumns  []RecordColumns `json:"record_columns"`
	RecordFormat   RecordFormat    `json:"record_format"`
}

type RecordColumns struct {
	Mapping string `json:"mapping"`
	Name    string `json:"name"`
	SqlType string `json:"sql_type"`
}

type RecordFormat struct {
	RecordFormatType  string            `json:"record_format_type"`
	MappingParameters MappingParameters `json:"mapping_parameters"`
}

type MappingParameters struct {
	Csv  Csv  `json:"csv"`
	Json Json `json:"json"`
}

type Csv struct {
	RecordRowDelimiter    string `json:"record_row_delimiter"`
	RecordColumnDelimiter string `json:"record_column_delimiter"`
}

type Json struct {
	RecordRowPath string `json:"record_row_path"`
}

type CloudwatchLoggingOptions struct {
	Id           string `json:"id"`
	LogStreamArn string `json:"log_stream_arn"`
	RoleArn      string `json:"role_arn"`
}

// A KinesisAnalyticsApplicationStatus defines the observed state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisAnalyticsApplicationObservation `json:",inline"`
}

// A KinesisAnalyticsApplicationObservation records the observed state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationObservation struct {
	Id                  string `json:"id"`
	Status              string `json:"status"`
	Arn                 string `json:"arn"`
	LastUpdateTimestamp string `json:"last_update_timestamp"`
	Version             int    `json:"version"`
	CreateTimestamp     string `json:"create_timestamp"`
}