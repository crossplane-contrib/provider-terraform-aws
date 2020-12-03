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

// KinesisFirehoseDeliveryStream is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KinesisFirehoseDeliveryStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KinesisFirehoseDeliveryStreamSpec   `json:"spec"`
	Status KinesisFirehoseDeliveryStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisFirehoseDeliveryStream contains a list of KinesisFirehoseDeliveryStreamList
type KinesisFirehoseDeliveryStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisFirehoseDeliveryStream `json:"items"`
}

// A KinesisFirehoseDeliveryStreamSpec defines the desired state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KinesisFirehoseDeliveryStreamParameters `json:",inline"`
}

// A KinesisFirehoseDeliveryStreamParameters defines the desired state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamParameters struct {
	Destination                string                     `json:"destination"`
	DestinationId              string                     `json:"destination_id"`
	Id                         string                     `json:"id"`
	Name                       string                     `json:"name"`
	Tags                       map[string]string          `json:"tags"`
	VersionId                  string                     `json:"version_id"`
	Arn                        string                     `json:"arn"`
	ElasticsearchConfiguration ElasticsearchConfiguration `json:"elasticsearch_configuration"`
	ExtendedS3Configuration    ExtendedS3Configuration    `json:"extended_s3_configuration"`
	KinesisSourceConfiguration KinesisSourceConfiguration `json:"kinesis_source_configuration"`
	RedshiftConfiguration      RedshiftConfiguration      `json:"redshift_configuration"`
	S3Configuration            S3Configuration            `json:"s3_configuration"`
	ServerSideEncryption       ServerSideEncryption       `json:"server_side_encryption"`
	SplunkConfiguration        SplunkConfiguration        `json:"splunk_configuration"`
}

type ElasticsearchConfiguration struct {
	RetryDuration            int                      `json:"retry_duration"`
	RoleArn                  string                   `json:"role_arn"`
	BufferingSize            int                      `json:"buffering_size"`
	ClusterEndpoint          string                   `json:"cluster_endpoint"`
	DomainArn                string                   `json:"domain_arn"`
	IndexName                string                   `json:"index_name"`
	BufferingInterval        int                      `json:"buffering_interval"`
	IndexRotationPeriod      string                   `json:"index_rotation_period"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	TypeName                 string                   `json:"type_name"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  ProcessingConfiguration  `json:"processing_configuration"`
	VpcConfig                VpcConfig                `json:"vpc_config"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type ProcessingConfiguration struct {
	Enabled    bool         `json:"enabled"`
	Processors []Processors `json:"processors"`
}

type Processors struct {
	Type       string       `json:"type"`
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type VpcConfig struct {
	RoleArn          string   `json:"role_arn"`
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
}

type ExtendedS3Configuration struct {
	BucketArn                         string                            `json:"bucket_arn"`
	ErrorOutputPrefix                 string                            `json:"error_output_prefix"`
	KmsKeyArn                         string                            `json:"kms_key_arn"`
	RoleArn                           string                            `json:"role_arn"`
	BufferInterval                    int                               `json:"buffer_interval"`
	BufferSize                        int                               `json:"buffer_size"`
	CompressionFormat                 string                            `json:"compression_format"`
	Prefix                            string                            `json:"prefix"`
	S3BackupMode                      string                            `json:"s3_backup_mode"`
	DataFormatConversionConfiguration DataFormatConversionConfiguration `json:"data_format_conversion_configuration"`
	ProcessingConfiguration           ProcessingConfiguration           `json:"processing_configuration"`
	S3BackupConfiguration             S3BackupConfiguration             `json:"s3_backup_configuration"`
	CloudwatchLoggingOptions          CloudwatchLoggingOptions          `json:"cloudwatch_logging_options"`
}

type DataFormatConversionConfiguration struct {
	Enabled                   bool                      `json:"enabled"`
	SchemaConfiguration       SchemaConfiguration       `json:"schema_configuration"`
	InputFormatConfiguration  InputFormatConfiguration  `json:"input_format_configuration"`
	OutputFormatConfiguration OutputFormatConfiguration `json:"output_format_configuration"`
}

type SchemaConfiguration struct {
	DatabaseName string `json:"database_name"`
	Region       string `json:"region"`
	RoleArn      string `json:"role_arn"`
	TableName    string `json:"table_name"`
	VersionId    string `json:"version_id"`
	CatalogId    string `json:"catalog_id"`
}

type InputFormatConfiguration struct {
	Deserializer Deserializer `json:"deserializer"`
}

type Deserializer struct {
	HiveJsonSerDe  HiveJsonSerDe  `json:"hive_json_ser_de"`
	OpenXJsonSerDe OpenXJsonSerDe `json:"open_x_json_ser_de"`
}

type HiveJsonSerDe struct {
	TimestampFormats []string `json:"timestamp_formats"`
}

type OpenXJsonSerDe struct {
	CaseInsensitive                    bool              `json:"case_insensitive"`
	ColumnToJsonKeyMappings            map[string]string `json:"column_to_json_key_mappings"`
	ConvertDotsInJsonKeysToUnderscores bool              `json:"convert_dots_in_json_keys_to_underscores"`
}

type OutputFormatConfiguration struct {
	Serializer Serializer `json:"serializer"`
}

type Serializer struct {
	OrcSerDe     OrcSerDe     `json:"orc_ser_de"`
	ParquetSerDe ParquetSerDe `json:"parquet_ser_de"`
}

type OrcSerDe struct {
	BloomFilterFalsePositiveProbability int      `json:"bloom_filter_false_positive_probability"`
	RowIndexStride                      int      `json:"row_index_stride"`
	EnablePadding                       bool     `json:"enable_padding"`
	FormatVersion                       string   `json:"format_version"`
	PaddingTolerance                    int      `json:"padding_tolerance"`
	StripeSizeBytes                     int      `json:"stripe_size_bytes"`
	BlockSizeBytes                      int      `json:"block_size_bytes"`
	BloomFilterColumns                  []string `json:"bloom_filter_columns"`
	Compression                         string   `json:"compression"`
	DictionaryKeyThreshold              int      `json:"dictionary_key_threshold"`
}

type ParquetSerDe struct {
	BlockSizeBytes              int    `json:"block_size_bytes"`
	Compression                 string `json:"compression"`
	EnableDictionaryCompression bool   `json:"enable_dictionary_compression"`
	MaxPaddingBytes             int    `json:"max_padding_bytes"`
	PageSizeBytes               int    `json:"page_size_bytes"`
	WriterVersion               string `json:"writer_version"`
}

type ProcessingConfiguration struct {
	Enabled    bool         `json:"enabled"`
	Processors []Processors `json:"processors"`
}

type Processors struct {
	Type       string       `json:"type"`
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type S3BackupConfiguration struct {
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int                      `json:"buffer_interval"`
	BufferSize               int                      `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	Prefix                   string                   `json:"prefix"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type KinesisSourceConfiguration struct {
	RoleArn          string `json:"role_arn"`
	KinesisStreamArn string `json:"kinesis_stream_arn"`
}

type RedshiftConfiguration struct {
	RetryDuration            int                      `json:"retry_duration"`
	RoleArn                  string                   `json:"role_arn"`
	Username                 string                   `json:"username"`
	CopyOptions              string                   `json:"copy_options"`
	DataTableColumns         string                   `json:"data_table_columns"`
	Password                 string                   `json:"password"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	ClusterJdbcurl           string                   `json:"cluster_jdbcurl"`
	DataTableName            string                   `json:"data_table_name"`
	S3BackupConfiguration    S3BackupConfiguration    `json:"s3_backup_configuration"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  ProcessingConfiguration  `json:"processing_configuration"`
}

type S3BackupConfiguration struct {
	BufferInterval           int                      `json:"buffer_interval"`
	BufferSize               int                      `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	Prefix                   string                   `json:"prefix"`
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type ProcessingConfiguration struct {
	Enabled    bool         `json:"enabled"`
	Processors []Processors `json:"processors"`
}

type Processors struct {
	Type       string       `json:"type"`
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type S3Configuration struct {
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int                      `json:"buffer_interval"`
	BufferSize               int                      `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	Prefix                   string                   `json:"prefix"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type ServerSideEncryption struct {
	Enabled bool   `json:"enabled"`
	KeyArn  string `json:"key_arn"`
	KeyType string `json:"key_type"`
}

type SplunkConfiguration struct {
	HecEndpoint              string                   `json:"hec_endpoint"`
	HecEndpointType          string                   `json:"hec_endpoint_type"`
	HecToken                 string                   `json:"hec_token"`
	RetryDuration            int                      `json:"retry_duration"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	HecAcknowledgmentTimeout int                      `json:"hec_acknowledgment_timeout"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  ProcessingConfiguration  `json:"processing_configuration"`
}

type CloudwatchLoggingOptions struct {
	LogStreamName string `json:"log_stream_name"`
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
}

type ProcessingConfiguration struct {
	Enabled    bool         `json:"enabled"`
	Processors []Processors `json:"processors"`
}

type Processors struct {
	Type       string       `json:"type"`
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

// A KinesisFirehoseDeliveryStreamStatus defines the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisFirehoseDeliveryStreamObservation `json:",inline"`
}

// A KinesisFirehoseDeliveryStreamObservation records the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamObservation struct{}