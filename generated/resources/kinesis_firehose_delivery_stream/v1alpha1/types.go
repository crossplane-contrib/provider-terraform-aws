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
	VersionId                  string                     `json:"version_id"`
	Arn                        string                     `json:"arn"`
	Destination                string                     `json:"destination"`
	DestinationId              string                     `json:"destination_id"`
	Id                         string                     `json:"id"`
	Name                       string                     `json:"name"`
	Tags                       map[string]string          `json:"tags"`
	ExtendedS3Configuration    ExtendedS3Configuration    `json:"extended_s3_configuration"`
	KinesisSourceConfiguration KinesisSourceConfiguration `json:"kinesis_source_configuration"`
	RedshiftConfiguration      RedshiftConfiguration      `json:"redshift_configuration"`
	S3Configuration            S3Configuration            `json:"s3_configuration"`
	ServerSideEncryption       ServerSideEncryption       `json:"server_side_encryption"`
	SplunkConfiguration        SplunkConfiguration        `json:"splunk_configuration"`
	ElasticsearchConfiguration ElasticsearchConfiguration `json:"elasticsearch_configuration"`
}

type ExtendedS3Configuration struct {
	BucketArn                         string                            `json:"bucket_arn"`
	BufferInterval                    int64                             `json:"buffer_interval"`
	CompressionFormat                 string                            `json:"compression_format"`
	KmsKeyArn                         string                            `json:"kms_key_arn"`
	S3BackupMode                      string                            `json:"s3_backup_mode"`
	BufferSize                        int64                             `json:"buffer_size"`
	ErrorOutputPrefix                 string                            `json:"error_output_prefix"`
	Prefix                            string                            `json:"prefix"`
	RoleArn                           string                            `json:"role_arn"`
	CloudwatchLoggingOptions          CloudwatchLoggingOptions          `json:"cloudwatch_logging_options"`
	DataFormatConversionConfiguration DataFormatConversionConfiguration `json:"data_format_conversion_configuration"`
	ProcessingConfiguration           ProcessingConfiguration           `json:"processing_configuration"`
	S3BackupConfiguration             S3BackupConfiguration             `json:"s3_backup_configuration"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type DataFormatConversionConfiguration struct {
	Enabled                   bool                      `json:"enabled"`
	SchemaConfiguration       SchemaConfiguration       `json:"schema_configuration"`
	InputFormatConfiguration  InputFormatConfiguration  `json:"input_format_configuration"`
	OutputFormatConfiguration OutputFormatConfiguration `json:"output_format_configuration"`
}

type SchemaConfiguration struct {
	CatalogId    string `json:"catalog_id"`
	DatabaseName string `json:"database_name"`
	Region       string `json:"region"`
	RoleArn      string `json:"role_arn"`
	TableName    string `json:"table_name"`
	VersionId    string `json:"version_id"`
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
	StripeSizeBytes                     int64    `json:"stripe_size_bytes"`
	BlockSizeBytes                      int64    `json:"block_size_bytes"`
	Compression                         string   `json:"compression"`
	PaddingTolerance                    int64    `json:"padding_tolerance"`
	EnablePadding                       bool     `json:"enable_padding"`
	FormatVersion                       string   `json:"format_version"`
	RowIndexStride                      int64    `json:"row_index_stride"`
	BloomFilterColumns                  []string `json:"bloom_filter_columns"`
	BloomFilterFalsePositiveProbability int64    `json:"bloom_filter_false_positive_probability"`
	DictionaryKeyThreshold              int64    `json:"dictionary_key_threshold"`
}

type ParquetSerDe struct {
	WriterVersion               string `json:"writer_version"`
	BlockSizeBytes              int64  `json:"block_size_bytes"`
	Compression                 string `json:"compression"`
	EnableDictionaryCompression bool   `json:"enable_dictionary_compression"`
	MaxPaddingBytes             int64  `json:"max_padding_bytes"`
	PageSizeBytes               int64  `json:"page_size_bytes"`
}

type ProcessingConfiguration struct {
	Enabled    bool       `json:"enabled"`
	Processors Processors `json:"processors"`
}

type Processors struct {
	Type       string     `json:"type"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type S3BackupConfiguration struct {
	BufferSize               int64                    `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	Prefix                   string                   `json:"prefix"`
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int64                    `json:"buffer_interval"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type KinesisSourceConfiguration struct {
	KinesisStreamArn string `json:"kinesis_stream_arn"`
	RoleArn          string `json:"role_arn"`
}

type RedshiftConfiguration struct {
	ClusterJdbcurl           string                   `json:"cluster_jdbcurl"`
	DataTableColumns         string                   `json:"data_table_columns"`
	DataTableName            string                   `json:"data_table_name"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	Username                 string                   `json:"username"`
	CopyOptions              string                   `json:"copy_options"`
	Password                 string                   `json:"password"`
	RetryDuration            int64                    `json:"retry_duration"`
	RoleArn                  string                   `json:"role_arn"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  ProcessingConfiguration  `json:"processing_configuration"`
	S3BackupConfiguration    S3BackupConfiguration    `json:"s3_backup_configuration"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type ProcessingConfiguration struct {
	Enabled    bool       `json:"enabled"`
	Processors Processors `json:"processors"`
}

type Processors struct {
	Type       string     `json:"type"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterValue string `json:"parameter_value"`
	ParameterName  string `json:"parameter_name"`
}

type S3BackupConfiguration struct {
	BufferSize               int64                    `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	Prefix                   string                   `json:"prefix"`
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int64                    `json:"buffer_interval"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type S3Configuration struct {
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int64                    `json:"buffer_interval"`
	BufferSize               int64                    `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	Prefix                   string                   `json:"prefix"`
	RoleArn                  string                   `json:"role_arn"`
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
	HecAcknowledgmentTimeout int64                    `json:"hec_acknowledgment_timeout"`
	HecEndpoint              string                   `json:"hec_endpoint"`
	HecEndpointType          string                   `json:"hec_endpoint_type"`
	HecToken                 string                   `json:"hec_token"`
	RetryDuration            int64                    `json:"retry_duration"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  ProcessingConfiguration  `json:"processing_configuration"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type ProcessingConfiguration struct {
	Enabled    bool       `json:"enabled"`
	Processors Processors `json:"processors"`
}

type Processors struct {
	Type       string     `json:"type"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type ElasticsearchConfiguration struct {
	ClusterEndpoint          string                   `json:"cluster_endpoint"`
	IndexName                string                   `json:"index_name"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	BufferingInterval        int64                    `json:"buffering_interval"`
	BufferingSize            int64                    `json:"buffering_size"`
	DomainArn                string                   `json:"domain_arn"`
	IndexRotationPeriod      string                   `json:"index_rotation_period"`
	RetryDuration            int64                    `json:"retry_duration"`
	RoleArn                  string                   `json:"role_arn"`
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
	Enabled    bool       `json:"enabled"`
	Processors Processors `json:"processors"`
}

type Processors struct {
	Type       string     `json:"type"`
	Parameters Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName  string `json:"parameter_name"`
	ParameterValue string `json:"parameter_value"`
}

type VpcConfig struct {
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
	RoleArn          string   `json:"role_arn"`
	SecurityGroupIds []string `json:"security_group_ids"`
}

// A KinesisFirehoseDeliveryStreamStatus defines the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisFirehoseDeliveryStreamObservation `json:",inline"`
}

// A KinesisFirehoseDeliveryStreamObservation records the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamObservation struct{}