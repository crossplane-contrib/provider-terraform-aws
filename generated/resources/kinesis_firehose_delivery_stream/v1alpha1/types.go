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
	Name                       string                     `json:"name"`
	Tags                       map[string]string          `json:"tags"`
	SplunkConfiguration        SplunkConfiguration        `json:"splunk_configuration"`
	ElasticsearchConfiguration ElasticsearchConfiguration `json:"elasticsearch_configuration"`
	ExtendedS3Configuration    ExtendedS3Configuration    `json:"extended_s3_configuration"`
	KinesisSourceConfiguration KinesisSourceConfiguration `json:"kinesis_source_configuration"`
	RedshiftConfiguration      RedshiftConfiguration      `json:"redshift_configuration"`
	S3Configuration            S3Configuration            `json:"s3_configuration"`
	ServerSideEncryption       ServerSideEncryption       `json:"server_side_encryption"`
}

type SplunkConfiguration struct {
	S3BackupMode             string                   `json:"s3_backup_mode"`
	HecAcknowledgmentTimeout int                      `json:"hec_acknowledgment_timeout"`
	HecEndpoint              string                   `json:"hec_endpoint"`
	HecEndpointType          string                   `json:"hec_endpoint_type"`
	HecToken                 string                   `json:"hec_token"`
	RetryDuration            int                      `json:"retry_duration"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
	ProcessingConfiguration  ProcessingConfiguration  `json:"processing_configuration"`
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

type ElasticsearchConfiguration struct {
	BufferingInterval        int                      `json:"buffering_interval"`
	BufferingSize            int                      `json:"buffering_size"`
	ClusterEndpoint          string                   `json:"cluster_endpoint"`
	IndexName                string                   `json:"index_name"`
	IndexRotationPeriod      string                   `json:"index_rotation_period"`
	RoleArn                  string                   `json:"role_arn"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	DomainArn                string                   `json:"domain_arn"`
	RetryDuration            int                      `json:"retry_duration"`
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
	VpcId            string   `json:"vpc_id"`
	RoleArn          string   `json:"role_arn"`
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
}

type ExtendedS3Configuration struct {
	BufferInterval                    int                               `json:"buffer_interval"`
	CompressionFormat                 string                            `json:"compression_format"`
	S3BackupMode                      string                            `json:"s3_backup_mode"`
	BucketArn                         string                            `json:"bucket_arn"`
	ErrorOutputPrefix                 string                            `json:"error_output_prefix"`
	KmsKeyArn                         string                            `json:"kms_key_arn"`
	Prefix                            string                            `json:"prefix"`
	RoleArn                           string                            `json:"role_arn"`
	BufferSize                        int                               `json:"buffer_size"`
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
	InputFormatConfiguration  InputFormatConfiguration  `json:"input_format_configuration"`
	OutputFormatConfiguration OutputFormatConfiguration `json:"output_format_configuration"`
	SchemaConfiguration       SchemaConfiguration       `json:"schema_configuration"`
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
	DictionaryKeyThreshold              int      `json:"dictionary_key_threshold"`
	BloomFilterFalsePositiveProbability int      `json:"bloom_filter_false_positive_probability"`
	Compression                         string   `json:"compression"`
	EnablePadding                       bool     `json:"enable_padding"`
	FormatVersion                       string   `json:"format_version"`
	PaddingTolerance                    int      `json:"padding_tolerance"`
	RowIndexStride                      int      `json:"row_index_stride"`
	StripeSizeBytes                     int      `json:"stripe_size_bytes"`
	BlockSizeBytes                      int      `json:"block_size_bytes"`
	BloomFilterColumns                  []string `json:"bloom_filter_columns"`
}

type ParquetSerDe struct {
	PageSizeBytes               int    `json:"page_size_bytes"`
	WriterVersion               string `json:"writer_version"`
	BlockSizeBytes              int    `json:"block_size_bytes"`
	Compression                 string `json:"compression"`
	EnableDictionaryCompression bool   `json:"enable_dictionary_compression"`
	MaxPaddingBytes             int    `json:"max_padding_bytes"`
}

type SchemaConfiguration struct {
	Region       string `json:"region"`
	RoleArn      string `json:"role_arn"`
	TableName    string `json:"table_name"`
	VersionId    string `json:"version_id"`
	CatalogId    string `json:"catalog_id"`
	DatabaseName string `json:"database_name"`
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
	KmsKeyArn                string                   `json:"kms_key_arn"`
	Prefix                   string                   `json:"prefix"`
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int                      `json:"buffer_interval"`
	BufferSize               int                      `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
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
	CopyOptions              string                   `json:"copy_options"`
	DataTableName            string                   `json:"data_table_name"`
	RetryDuration            int                      `json:"retry_duration"`
	RoleArn                  string                   `json:"role_arn"`
	S3BackupMode             string                   `json:"s3_backup_mode"`
	Username                 string                   `json:"username"`
	ClusterJdbcurl           string                   `json:"cluster_jdbcurl"`
	DataTableColumns         string                   `json:"data_table_columns"`
	Password                 string                   `json:"password"`
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
	Prefix                   string                   `json:"prefix"`
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int                      `json:"buffer_interval"`
	BufferSize               int                      `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type S3Configuration struct {
	Prefix                   string                   `json:"prefix"`
	RoleArn                  string                   `json:"role_arn"`
	BucketArn                string                   `json:"bucket_arn"`
	BufferInterval           int                      `json:"buffer_interval"`
	BufferSize               int                      `json:"buffer_size"`
	CompressionFormat        string                   `json:"compression_format"`
	KmsKeyArn                string                   `json:"kms_key_arn"`
	CloudwatchLoggingOptions CloudwatchLoggingOptions `json:"cloudwatch_logging_options"`
}

type CloudwatchLoggingOptions struct {
	Enabled       bool   `json:"enabled"`
	LogGroupName  string `json:"log_group_name"`
	LogStreamName string `json:"log_stream_name"`
}

type ServerSideEncryption struct {
	KeyArn  string `json:"key_arn"`
	KeyType string `json:"key_type"`
	Enabled bool   `json:"enabled"`
}

// A KinesisFirehoseDeliveryStreamStatus defines the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisFirehoseDeliveryStreamObservation `json:",inline"`
}

// A KinesisFirehoseDeliveryStreamObservation records the observed state of a KinesisFirehoseDeliveryStream
type KinesisFirehoseDeliveryStreamObservation struct {
	VersionId     string `json:"version_id"`
	Arn           string `json:"arn"`
	DestinationId string `json:"destination_id"`
	Id            string `json:"id"`
}