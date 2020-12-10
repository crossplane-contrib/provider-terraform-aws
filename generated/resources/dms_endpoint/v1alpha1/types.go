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

// DmsEndpoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsEndpointSpec   `json:"spec"`
	Status DmsEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsEndpoint contains a list of DmsEndpointList
type DmsEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsEndpoint `json:"items"`
}

// A DmsEndpointSpec defines the desired state of a DmsEndpoint
type DmsEndpointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsEndpointParameters `json:",inline"`
}

// A DmsEndpointParameters defines the desired state of a DmsEndpoint
type DmsEndpointParameters struct {
	ExtraConnectionAttributes string                `json:"extra_connection_attributes"`
	Password                  string                `json:"password"`
	SslMode                   string                `json:"ssl_mode"`
	Tags                      map[string]string     `json:"tags"`
	CertificateArn            string                `json:"certificate_arn"`
	EngineName                string                `json:"engine_name"`
	ServerName                string                `json:"server_name"`
	ServiceAccessRole         string                `json:"service_access_role"`
	Id                        string                `json:"id"`
	KmsKeyArn                 string                `json:"kms_key_arn"`
	Port                      int64                 `json:"port"`
	Username                  string                `json:"username"`
	DatabaseName              string                `json:"database_name"`
	EndpointId                string                `json:"endpoint_id"`
	EndpointType              string                `json:"endpoint_type"`
	KafkaSettings             KafkaSettings         `json:"kafka_settings"`
	KinesisSettings           KinesisSettings       `json:"kinesis_settings"`
	MongodbSettings           MongodbSettings       `json:"mongodb_settings"`
	S3Settings                S3Settings            `json:"s3_settings"`
	ElasticsearchSettings     ElasticsearchSettings `json:"elasticsearch_settings"`
}

type KafkaSettings struct {
	Broker string `json:"broker"`
	Topic  string `json:"topic"`
}

type KinesisSettings struct {
	MessageFormat        string `json:"message_format"`
	ServiceAccessRoleArn string `json:"service_access_role_arn"`
	StreamArn            string `json:"stream_arn"`
}

type MongodbSettings struct {
	AuthType          string `json:"auth_type"`
	DocsToInvestigate string `json:"docs_to_investigate"`
	ExtractDocId      string `json:"extract_doc_id"`
	NestingLevel      string `json:"nesting_level"`
	AuthMechanism     string `json:"auth_mechanism"`
	AuthSource        string `json:"auth_source"`
}

type S3Settings struct {
	BucketFolder            string `json:"bucket_folder"`
	BucketName              string `json:"bucket_name"`
	CompressionType         string `json:"compression_type"`
	CsvDelimiter            string `json:"csv_delimiter"`
	CsvRowDelimiter         string `json:"csv_row_delimiter"`
	ExternalTableDefinition string `json:"external_table_definition"`
	ServiceAccessRoleArn    string `json:"service_access_role_arn"`
}

type ElasticsearchSettings struct {
	FullLoadErrorPercentage int64  `json:"full_load_error_percentage"`
	ServiceAccessRoleArn    string `json:"service_access_role_arn"`
	EndpointUri             string `json:"endpoint_uri"`
	ErrorRetryDuration      int64  `json:"error_retry_duration"`
}

// A DmsEndpointStatus defines the observed state of a DmsEndpoint
type DmsEndpointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsEndpointObservation `json:",inline"`
}

// A DmsEndpointObservation records the observed state of a DmsEndpoint
type DmsEndpointObservation struct {
	EndpointArn string `json:"endpoint_arn"`
}