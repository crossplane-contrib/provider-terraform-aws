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
	ForProvider                  ElasticsearchDomainParameters `json:"forProvider"`
}

// A ElasticsearchDomainParameters defines the desired state of a ElasticsearchDomain
type ElasticsearchDomainParameters struct {
	Id                      string                  `json:"id"`
	Tags                    map[string]string       `json:"tags"`
	AccessPolicies          string                  `json:"access_policies"`
	ElasticsearchVersion    string                  `json:"elasticsearch_version"`
	AdvancedOptions         map[string]string       `json:"advanced_options"`
	DomainName              string                  `json:"domain_name"`
	LogPublishingOptions    LogPublishingOptions    `json:"log_publishing_options"`
	Timeouts                Timeouts                `json:"timeouts"`
	AdvancedSecurityOptions AdvancedSecurityOptions `json:"advanced_security_options"`
	ClusterConfig           ClusterConfig           `json:"cluster_config"`
	CognitoOptions          CognitoOptions          `json:"cognito_options"`
	NodeToNodeEncryption    NodeToNodeEncryption    `json:"node_to_node_encryption"`
	SnapshotOptions         SnapshotOptions         `json:"snapshot_options"`
	VpcOptions              VpcOptions              `json:"vpc_options"`
	DomainEndpointOptions   DomainEndpointOptions   `json:"domain_endpoint_options"`
	EbsOptions              EbsOptions              `json:"ebs_options"`
	EncryptAtRest           EncryptAtRest           `json:"encrypt_at_rest"`
}

type LogPublishingOptions struct {
	CloudwatchLogGroupArn string `json:"cloudwatch_log_group_arn"`
	Enabled               bool   `json:"enabled"`
	LogType               string `json:"log_type"`
}

type Timeouts struct {
	Update string `json:"update"`
}

type AdvancedSecurityOptions struct {
	InternalUserDatabaseEnabled bool              `json:"internal_user_database_enabled"`
	Enabled                     bool              `json:"enabled"`
	MasterUserOptions           MasterUserOptions `json:"master_user_options"`
}

type MasterUserOptions struct {
	MasterUserArn      string `json:"master_user_arn"`
	MasterUserName     string `json:"master_user_name"`
	MasterUserPassword string `json:"master_user_password"`
}

type ClusterConfig struct {
	DedicatedMasterCount   int64               `json:"dedicated_master_count"`
	DedicatedMasterEnabled bool                `json:"dedicated_master_enabled"`
	DedicatedMasterType    string              `json:"dedicated_master_type"`
	InstanceCount          int64               `json:"instance_count"`
	WarmCount              int64               `json:"warm_count"`
	WarmType               string              `json:"warm_type"`
	InstanceType           string              `json:"instance_type"`
	WarmEnabled            bool                `json:"warm_enabled"`
	ZoneAwarenessEnabled   bool                `json:"zone_awareness_enabled"`
	ZoneAwarenessConfig    ZoneAwarenessConfig `json:"zone_awareness_config"`
}

type ZoneAwarenessConfig struct {
	AvailabilityZoneCount int64 `json:"availability_zone_count"`
}

type CognitoOptions struct {
	RoleArn        string `json:"role_arn"`
	UserPoolId     string `json:"user_pool_id"`
	Enabled        bool   `json:"enabled"`
	IdentityPoolId string `json:"identity_pool_id"`
}

type NodeToNodeEncryption struct {
	Enabled bool `json:"enabled"`
}

type SnapshotOptions struct {
	AutomatedSnapshotStartHour int64 `json:"automated_snapshot_start_hour"`
}

type VpcOptions struct {
	AvailabilityZones []string `json:"availability_zones"`
	SecurityGroupIds  []string `json:"security_group_ids"`
	SubnetIds         []string `json:"subnet_ids"`
	VpcId             string   `json:"vpc_id"`
}

type DomainEndpointOptions struct {
	EnforceHttps      bool   `json:"enforce_https"`
	TlsSecurityPolicy string `json:"tls_security_policy"`
}

type EbsOptions struct {
	VolumeSize int64  `json:"volume_size"`
	VolumeType string `json:"volume_type"`
	EbsEnabled bool   `json:"ebs_enabled"`
	Iops       int64  `json:"iops"`
}

type EncryptAtRest struct {
	KmsKeyId string `json:"kms_key_id"`
	Enabled  bool   `json:"enabled"`
}

// A ElasticsearchDomainStatus defines the observed state of a ElasticsearchDomain
type ElasticsearchDomainStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticsearchDomainObservation `json:"atProvider"`
}

// A ElasticsearchDomainObservation records the observed state of a ElasticsearchDomain
type ElasticsearchDomainObservation struct {
	DomainId       string `json:"domain_id"`
	KibanaEndpoint string `json:"kibana_endpoint"`
	Arn            string `json:"arn"`
	Endpoint       string `json:"endpoint"`
}