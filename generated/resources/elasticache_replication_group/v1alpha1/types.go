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

// ElasticacheReplicationGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticacheReplicationGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticacheReplicationGroupSpec   `json:"spec"`
	Status ElasticacheReplicationGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheReplicationGroup contains a list of ElasticacheReplicationGroupList
type ElasticacheReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheReplicationGroup `json:"items"`
}

// A ElasticacheReplicationGroupSpec defines the desired state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticacheReplicationGroupParameters `json:",inline"`
}

// A ElasticacheReplicationGroupParameters defines the desired state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupParameters struct {
	AtRestEncryptionEnabled     bool              `json:"at_rest_encryption_enabled"`
	Port                        int64             `json:"port"`
	SecurityGroupNames          []string          `json:"security_group_names"`
	SubnetGroupName             string            `json:"subnet_group_name"`
	AuthToken                   string            `json:"auth_token"`
	SnapshotArns                []string          `json:"snapshot_arns"`
	ApplyImmediately            bool              `json:"apply_immediately"`
	AutomaticFailoverEnabled    bool              `json:"automatic_failover_enabled"`
	EngineVersion               string            `json:"engine_version"`
	MaintenanceWindow           string            `json:"maintenance_window"`
	NotificationTopicArn        string            `json:"notification_topic_arn"`
	AutoMinorVersionUpgrade     bool              `json:"auto_minor_version_upgrade"`
	SecurityGroupIds            []string          `json:"security_group_ids"`
	SnapshotRetentionLimit      int64             `json:"snapshot_retention_limit"`
	NumberCacheClusters         int64             `json:"number_cache_clusters"`
	ReplicationGroupDescription string            `json:"replication_group_description"`
	TransitEncryptionEnabled    bool              `json:"transit_encryption_enabled"`
	AvailabilityZones           []string          `json:"availability_zones"`
	Id                          string            `json:"id"`
	ParameterGroupName          string            `json:"parameter_group_name"`
	ReplicationGroupId          string            `json:"replication_group_id"`
	SnapshotName                string            `json:"snapshot_name"`
	Engine                      string            `json:"engine"`
	KmsKeyId                    string            `json:"kms_key_id"`
	NodeType                    string            `json:"node_type"`
	SnapshotWindow              string            `json:"snapshot_window"`
	Tags                        map[string]string `json:"tags"`
	ClusterMode                 ClusterMode       `json:"cluster_mode"`
	Timeouts                    Timeouts          `json:"timeouts"`
}

type ClusterMode struct {
	NumNodeGroups        int64 `json:"num_node_groups"`
	ReplicasPerNodeGroup int64 `json:"replicas_per_node_group"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A ElasticacheReplicationGroupStatus defines the observed state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticacheReplicationGroupObservation `json:",inline"`
}

// A ElasticacheReplicationGroupObservation records the observed state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupObservation struct {
	PrimaryEndpointAddress       string   `json:"primary_endpoint_address"`
	ConfigurationEndpointAddress string   `json:"configuration_endpoint_address"`
	MemberClusters               []string `json:"member_clusters"`
}