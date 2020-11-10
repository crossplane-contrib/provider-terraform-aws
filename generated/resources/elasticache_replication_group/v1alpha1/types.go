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
	NotificationTopicArn        string            `json:"notification_topic_arn"`
	KmsKeyId                    string            `json:"kms_key_id"`
	Tags                        map[string]string `json:"tags"`
	SnapshotRetentionLimit      int               `json:"snapshot_retention_limit"`
	TransitEncryptionEnabled    bool              `json:"transit_encryption_enabled"`
	SnapshotArns                []string          `json:"snapshot_arns"`
	AutoMinorVersionUpgrade     bool              `json:"auto_minor_version_upgrade"`
	AvailabilityZones           []string          `json:"availability_zones"`
	Port                        int               `json:"port"`
	SnapshotName                string            `json:"snapshot_name"`
	AtRestEncryptionEnabled     bool              `json:"at_rest_encryption_enabled"`
	AutomaticFailoverEnabled    bool              `json:"automatic_failover_enabled"`
	Engine                      string            `json:"engine"`
	ReplicationGroupDescription string            `json:"replication_group_description"`
	ReplicationGroupId          string            `json:"replication_group_id"`
	AuthToken                   string            `json:"auth_token"`
	ClusterMode                 ClusterMode       `json:"cluster_mode"`
	Timeouts                    []Timeouts        `json:"timeouts"`
}

type ClusterMode struct {
	ReplicasPerNodeGroup int `json:"replicas_per_node_group"`
	NumNodeGroups        int `json:"num_node_groups"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A ElasticacheReplicationGroupStatus defines the observed state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticacheReplicationGroupObservation `json:",inline"`
}

// A ElasticacheReplicationGroupObservation records the observed state of a ElasticacheReplicationGroup
type ElasticacheReplicationGroupObservation struct {
	SubnetGroupName              string   `json:"subnet_group_name"`
	ConfigurationEndpointAddress string   `json:"configuration_endpoint_address"`
	SecurityGroupNames           []string `json:"security_group_names"`
	MemberClusters               []string `json:"member_clusters"`
	SecurityGroupIds             []string `json:"security_group_ids"`
	EngineVersion                string   `json:"engine_version"`
	NumberCacheClusters          int      `json:"number_cache_clusters"`
	ParameterGroupName           string   `json:"parameter_group_name"`
	SnapshotWindow               string   `json:"snapshot_window"`
	Id                           string   `json:"id"`
	ApplyImmediately             bool     `json:"apply_immediately"`
	PrimaryEndpointAddress       string   `json:"primary_endpoint_address"`
	MaintenanceWindow            string   `json:"maintenance_window"`
	NodeType                     string   `json:"node_type"`
}