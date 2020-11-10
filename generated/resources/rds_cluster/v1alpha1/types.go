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

// RdsCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RdsCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RdsClusterSpec   `json:"spec"`
	Status RdsClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RdsCluster contains a list of RdsClusterList
type RdsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RdsCluster `json:"items"`
}

// A RdsClusterSpec defines the desired state of a RdsCluster
type RdsClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RdsClusterParameters `json:",inline"`
}

// A RdsClusterParameters defines the desired state of a RdsCluster
type RdsClusterParameters struct {
	EnabledCloudwatchLogsExports     []string             `json:"enabled_cloudwatch_logs_exports"`
	GlobalClusterIdentifier          string               `json:"global_cluster_identifier"`
	AllowMajorVersionUpgrade         bool                 `json:"allow_major_version_upgrade"`
	FinalSnapshotIdentifier          string               `json:"final_snapshot_identifier"`
	ReplicationSourceIdentifier      string               `json:"replication_source_identifier"`
	StorageEncrypted                 bool                 `json:"storage_encrypted"`
	EnableHttpEndpoint               bool                 `json:"enable_http_endpoint"`
	MasterPassword                   string               `json:"master_password"`
	BacktrackWindow                  int                  `json:"backtrack_window"`
	Engine                           string               `json:"engine"`
	SkipFinalSnapshot                bool                 `json:"skip_final_snapshot"`
	SnapshotIdentifier               string               `json:"snapshot_identifier"`
	CopyTagsToSnapshot               bool                 `json:"copy_tags_to_snapshot"`
	IamDatabaseAuthenticationEnabled bool                 `json:"iam_database_authentication_enabled"`
	SourceRegion                     string               `json:"source_region"`
	Tags                             map[string]string    `json:"tags"`
	EngineMode                       string               `json:"engine_mode"`
	BackupRetentionPeriod            int                  `json:"backup_retention_period"`
	DeletionProtection               bool                 `json:"deletion_protection"`
	IamRoles                         []string             `json:"iam_roles"`
	Timeouts                         []Timeouts           `json:"timeouts"`
	S3Import                         S3Import             `json:"s3_import"`
	ScalingConfiguration             ScalingConfiguration `json:"scaling_configuration"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

type S3Import struct {
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
	BucketName          string `json:"bucket_name"`
}

type ScalingConfiguration struct {
	AutoPause             bool   `json:"auto_pause"`
	MaxCapacity           int    `json:"max_capacity"`
	MinCapacity           int    `json:"min_capacity"`
	SecondsUntilAutoPause int    `json:"seconds_until_auto_pause"`
	TimeoutAction         string `json:"timeout_action"`
}

// A RdsClusterStatus defines the observed state of a RdsCluster
type RdsClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RdsClusterObservation `json:",inline"`
}

// A RdsClusterObservation records the observed state of a RdsCluster
type RdsClusterObservation struct {
	ReaderEndpoint              string   `json:"reader_endpoint"`
	ClusterIdentifierPrefix     string   `json:"cluster_identifier_prefix"`
	DatabaseName                string   `json:"database_name"`
	Endpoint                    string   `json:"endpoint"`
	HostedZoneId                string   `json:"hosted_zone_id"`
	MasterUsername              string   `json:"master_username"`
	PreferredMaintenanceWindow  string   `json:"preferred_maintenance_window"`
	Port                        int      `json:"port"`
	PreferredBackupWindow       string   `json:"preferred_backup_window"`
	ClusterIdentifier           string   `json:"cluster_identifier"`
	EngineVersion               string   `json:"engine_version"`
	KmsKeyId                    string   `json:"kms_key_id"`
	Arn                         string   `json:"arn"`
	DbClusterParameterGroupName string   `json:"db_cluster_parameter_group_name"`
	Id                          string   `json:"id"`
	ApplyImmediately            bool     `json:"apply_immediately"`
	AvailabilityZones           []string `json:"availability_zones"`
	ClusterMembers              []string `json:"cluster_members"`
	ClusterResourceId           string   `json:"cluster_resource_id"`
	DbSubnetGroupName           string   `json:"db_subnet_group_name"`
	VpcSecurityGroupIds         []string `json:"vpc_security_group_ids"`
}