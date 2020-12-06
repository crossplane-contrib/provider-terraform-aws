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

// MskCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MskCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MskClusterSpec   `json:"spec"`
	Status MskClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MskCluster contains a list of MskClusterList
type MskClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MskCluster `json:"items"`
}

// A MskClusterSpec defines the desired state of a MskCluster
type MskClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MskClusterParameters `json:",inline"`
}

// A MskClusterParameters defines the desired state of a MskCluster
type MskClusterParameters struct {
	EnhancedMonitoring   string               `json:"enhanced_monitoring"`
	Tags                 map[string]string    `json:"tags"`
	Id                   string               `json:"id"`
	KafkaVersion         string               `json:"kafka_version"`
	NumberOfBrokerNodes  int64                `json:"number_of_broker_nodes"`
	ClusterName          string               `json:"cluster_name"`
	BrokerNodeGroupInfo  BrokerNodeGroupInfo  `json:"broker_node_group_info"`
	ClientAuthentication ClientAuthentication `json:"client_authentication"`
	ConfigurationInfo    ConfigurationInfo    `json:"configuration_info"`
	EncryptionInfo       EncryptionInfo       `json:"encryption_info"`
	LoggingInfo          LoggingInfo          `json:"logging_info"`
	OpenMonitoring       OpenMonitoring       `json:"open_monitoring"`
}

type BrokerNodeGroupInfo struct {
	EbsVolumeSize  int64    `json:"ebs_volume_size"`
	InstanceType   string   `json:"instance_type"`
	SecurityGroups []string `json:"security_groups"`
	AzDistribution string   `json:"az_distribution"`
	ClientSubnets  []string `json:"client_subnets"`
}

type ClientAuthentication struct {
	Tls Tls `json:"tls"`
}

type Tls struct {
	CertificateAuthorityArns []string `json:"certificate_authority_arns"`
}

type ConfigurationInfo struct {
	Arn      string `json:"arn"`
	Revision int64  `json:"revision"`
}

type EncryptionInfo struct {
	EncryptionAtRestKmsKeyArn string              `json:"encryption_at_rest_kms_key_arn"`
	EncryptionInTransit       EncryptionInTransit `json:"encryption_in_transit"`
}

type EncryptionInTransit struct {
	ClientBroker string `json:"client_broker"`
	InCluster    bool   `json:"in_cluster"`
}

type LoggingInfo struct {
	BrokerLogs BrokerLogs `json:"broker_logs"`
}

type BrokerLogs struct {
	CloudwatchLogs CloudwatchLogs `json:"cloudwatch_logs"`
	Firehose       Firehose       `json:"firehose"`
	S3             S3             `json:"s3"`
}

type CloudwatchLogs struct {
	Enabled  bool   `json:"enabled"`
	LogGroup string `json:"log_group"`
}

type Firehose struct {
	DeliveryStream string `json:"delivery_stream"`
	Enabled        bool   `json:"enabled"`
}

type S3 struct {
	Bucket  string `json:"bucket"`
	Enabled bool   `json:"enabled"`
	Prefix  string `json:"prefix"`
}

type OpenMonitoring struct {
	Prometheus Prometheus `json:"prometheus"`
}

type Prometheus struct {
	NodeExporter NodeExporter `json:"node_exporter"`
	JmxExporter  JmxExporter  `json:"jmx_exporter"`
}

type NodeExporter struct {
	EnabledInBroker bool `json:"enabled_in_broker"`
}

type JmxExporter struct {
	EnabledInBroker bool `json:"enabled_in_broker"`
}

// A MskClusterStatus defines the observed state of a MskCluster
type MskClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MskClusterObservation `json:",inline"`
}

// A MskClusterObservation records the observed state of a MskCluster
type MskClusterObservation struct {
	CurrentVersion         string `json:"current_version"`
	ZookeeperConnectString string `json:"zookeeper_connect_string"`
	Arn                    string `json:"arn"`
	BootstrapBrokers       string `json:"bootstrap_brokers"`
	BootstrapBrokersTls    string `json:"bootstrap_brokers_tls"`
}