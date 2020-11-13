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
	AvailabilityZones                []string             `json:"availability_zones"`
	BacktrackWindow                  int                  `json:"backtrack_window"`
	DbClusterParameterGroupName      string               `json:"db_cluster_parameter_group_name"`
	IamDatabaseAuthenticationEnabled bool                 `json:"iam_database_authentication_enabled"`
	PreferredBackupWindow            string               `json:"preferred_backup_window"`
	PreferredMaintenanceWindow       string               `json:"preferred_maintenance_window"`
	SourceRegion                     string               `json:"source_region"`
	CopyTagsToSnapshot               bool                 `json:"copy_tags_to_snapshot"`
	EngineMode                       string               `json:"engine_mode"`
	MasterPassword                   string               `json:"master_password"`
	SnapshotIdentifier               string               `json:"snapshot_identifier"`
	ClusterIdentifier                string               `json:"cluster_identifier"`
	DeletionProtection               bool                 `json:"deletion_protection"`
	Tags                             map[string]string    `json:"tags"`
	BackupRetentionPeriod            int                  `json:"backup_retention_period"`
	ClusterMembers                   []string             `json:"cluster_members"`
	EngineVersion                    string               `json:"engine_version"`
	StorageEncrypted                 bool                 `json:"storage_encrypted"`
	ClusterIdentifierPrefix          string               `json:"cluster_identifier_prefix"`
	EnableHttpEndpoint               bool                 `json:"enable_http_endpoint"`
	GlobalClusterIdentifier          string               `json:"global_cluster_identifier"`
	SkipFinalSnapshot                bool                 `json:"skip_final_snapshot"`
	AllowMajorVersionUpgrade         bool                 `json:"allow_major_version_upgrade"`
	Id                               string               `json:"id"`
	Port                             int                  `json:"port"`
	Engine                           string               `json:"engine"`
	DbSubnetGroupName                string               `json:"db_subnet_group_name"`
	EnabledCloudwatchLogsExports     []string             `json:"enabled_cloudwatch_logs_exports"`
	FinalSnapshotIdentifier          string               `json:"final_snapshot_identifier"`
	ReplicationSourceIdentifier      string               `json:"replication_source_identifier"`
	DatabaseName                     string               `json:"database_name"`
	IamRoles                         []string             `json:"iam_roles"`
	KmsKeyId                         string               `json:"kms_key_id"`
	MasterUsername                   string               `json:"master_username"`
	VpcSecurityGroupIds              []string             `json:"vpc_security_group_ids"`
	ApplyImmediately                 bool                 `json:"apply_immediately"`
	S3Import                         S3Import             `json:"s3_import"`
	ScalingConfiguration             ScalingConfiguration `json:"scaling_configuration"`
	Timeouts                         []Timeouts           `json:"timeouts"`
}

type S3Import struct {
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
}

type ScalingConfiguration struct {
	MinCapacity           int    `json:"min_capacity"`
	SecondsUntilAutoPause int    `json:"seconds_until_auto_pause"`
	TimeoutAction         string `json:"timeout_action"`
	AutoPause             bool   `json:"auto_pause"`
	MaxCapacity           int    `json:"max_capacity"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A RdsClusterStatus defines the observed state of a RdsCluster
type RdsClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RdsClusterObservation `json:",inline"`
}

// A RdsClusterObservation records the observed state of a RdsCluster
type RdsClusterObservation struct {
	Arn               string `json:"arn"`
	Endpoint          string `json:"endpoint"`
	ClusterResourceId string `json:"cluster_resource_id"`
	HostedZoneId      string `json:"hosted_zone_id"`
	ReaderEndpoint    string `json:"reader_endpoint"`
}