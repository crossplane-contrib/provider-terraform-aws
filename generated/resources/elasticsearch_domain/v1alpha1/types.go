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

// ElasticsearchDomain is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticsearchDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticsearchDomainSpec   `json:"spec"`
	Status ElasticsearchDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchDomain contains a list of ElasticsearchDomainList
type ElasticsearchDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchDomain `json:"items"`
}

// A ElasticsearchDomainSpec defines the desired state of a ElasticsearchDomain
type ElasticsearchDomainSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticsearchDomainParameters `json:",inline"`
}

// A ElasticsearchDomainParameters defines the desired state of a ElasticsearchDomain
type ElasticsearchDomainParameters struct {
	AccessPolicies          string                  `json:"access_policies"`
	DomainName              string                  `json:"domain_name"`
	Id                      string                  `json:"id"`
	AdvancedOptions         map[string]string       `json:"advanced_options"`
	ElasticsearchVersion    string                  `json:"elasticsearch_version"`
	Tags                    map[string]string       `json:"tags"`
	ClusterConfig           ClusterConfig           `json:"cluster_config"`
	CognitoOptions          CognitoOptions          `json:"cognito_options"`
	DomainEndpointOptions   DomainEndpointOptions   `json:"domain_endpoint_options"`
	LogPublishingOptions    LogPublishingOptions    `json:"log_publishing_options"`
	NodeToNodeEncryption    NodeToNodeEncryption    `json:"node_to_node_encryption"`
	SnapshotOptions         SnapshotOptions         `json:"snapshot_options"`
	Timeouts                Timeouts                `json:"timeouts"`
	VpcOptions              VpcOptions              `json:"vpc_options"`
	AdvancedSecurityOptions AdvancedSecurityOptions `json:"advanced_security_options"`
	EbsOptions              EbsOptions              `json:"ebs_options"`
	EncryptAtRest           EncryptAtRest           `json:"encrypt_at_rest"`
}

type ClusterConfig struct {
	ZoneAwarenessEnabled   bool                `json:"zone_awareness_enabled"`
	DedicatedMasterCount   int64               `json:"dedicated_master_count"`
	DedicatedMasterEnabled bool                `json:"dedicated_master_enabled"`
	InstanceCount          int64               `json:"instance_count"`
	WarmCount              int64               `json:"warm_count"`
	DedicatedMasterType    string              `json:"dedicated_master_type"`
	InstanceType           string              `json:"instance_type"`
	WarmEnabled            bool                `json:"warm_enabled"`
	WarmType               string              `json:"warm_type"`
	ZoneAwarenessConfig    ZoneAwarenessConfig `json:"zone_awareness_config"`
}

type ZoneAwarenessConfig struct {
	AvailabilityZoneCount int64 `json:"availability_zone_count"`
}

type CognitoOptions struct {
	Enabled        bool   `json:"enabled"`
	IdentityPoolId string `json:"identity_pool_id"`
	RoleArn        string `json:"role_arn"`
	UserPoolId     string `json:"user_pool_id"`
}

type DomainEndpointOptions struct {
	EnforceHttps      bool   `json:"enforce_https"`
	TlsSecurityPolicy string `json:"tls_security_policy"`
}

type LogPublishingOptions struct {
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	Enabled               bool   `json:"enabled"`
	LogType               string `json:"log_type"`
}

type NodeToNodeEncryption struct {
	Enabled bool `json:"enabled"`
}

type SnapshotOptions struct {
	AutomatedSnapshotStartHour int64 `json:"automated_snapshot_start_hour"`
}

type Timeouts struct {
	Update string `json:"update"`
}

type VpcOptions struct {
	SecurityGroupIds  []string `json:"security_group_ids"`
	SubnetIds         []string `json:"subnet_ids"`
	VpcId             string   `json:"vpc_id"`
	AvailabilityZones []string `json:"availability_zones"`
}

type AdvancedSecurityOptions struct {
	Enabled                     bool              `json:"enabled"`
	InternalUserDatabaseEnabled bool              `json:"internal_user_database_enabled"`
	MasterUserOptions           MasterUserOptions `json:"master_user_options"`
}

type MasterUserOptions struct {
	MasterUserPassword string `json:"master_user_password"`
	MasterUserArn      string `json:"master_user_arn"`
	MasterUserName     string `json:"master_user_name"`
}

type EbsOptions struct {
	Iops       int64  `json:"iops"`
	VolumeSize int64  `json:"volume_size"`
	VolumeType string `json:"volume_type"`
	EbsEnabled bool   `json:"ebs_enabled"`
}

type EncryptAtRest struct {
	Enabled  bool   `json:"enabled"`
	KmsKeyId string `json:"kms_key_id"`
}

// A ElasticsearchDomainStatus defines the observed state of a ElasticsearchDomain
type ElasticsearchDomainStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticsearchDomainObservation `json:",inline"`
}

// A ElasticsearchDomainObservation records the observed state of a ElasticsearchDomain
type ElasticsearchDomainObservation struct {
	Arn            string `json:"arn"`
	DomainId       string `json:"domain_id"`
	Endpoint       string `json:"endpoint"`
	KibanaEndpoint string `json:"kibana_endpoint"`
}