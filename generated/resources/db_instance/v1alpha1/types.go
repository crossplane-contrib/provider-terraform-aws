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

// DbInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbInstanceSpec   `json:"spec"`
	Status DbInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbInstance contains a list of DbInstanceList
type DbInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbInstance `json:"items"`
}

// A DbInstanceSpec defines the desired state of a DbInstance
type DbInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbInstanceParameters `json:",inline"`
}

// A DbInstanceParameters defines the desired state of a DbInstance
type DbInstanceParameters struct {
	AutoMinorVersionUpgrade          bool              `json:"auto_minor_version_upgrade"`
	Domain                           string            `json:"domain"`
	PubliclyAccessible               bool              `json:"publicly_accessible"`
	Password                         string            `json:"password"`
	SecurityGroupNames               []string          `json:"security_group_names"`
	AllowMajorVersionUpgrade         bool              `json:"allow_major_version_upgrade"`
	MaxAllocatedStorage              int               `json:"max_allocated_storage"`
	InstanceClass                    string            `json:"instance_class"`
	Iops                             int               `json:"iops"`
	IamDatabaseAuthenticationEnabled bool              `json:"iam_database_authentication_enabled"`
	StorageEncrypted                 bool              `json:"storage_encrypted"`
	Tags                             map[string]string `json:"tags"`
	DeletionProtection               bool              `json:"deletion_protection"`
	DomainIamRoleName                string            `json:"domain_iam_role_name"`
	MonitoringInterval               int               `json:"monitoring_interval"`
	ReplicateSourceDb                string            `json:"replicate_source_db"`
	SkipFinalSnapshot                bool              `json:"skip_final_snapshot"`
	EnabledCloudwatchLogsExports     []string          `json:"enabled_cloudwatch_logs_exports"`
	FinalSnapshotIdentifier          string            `json:"final_snapshot_identifier"`
	SnapshotIdentifier               string            `json:"snapshot_identifier"`
	CopyTagsToSnapshot               bool              `json:"copy_tags_to_snapshot"`
	DeleteAutomatedBackups           bool              `json:"delete_automated_backups"`
	PerformanceInsightsEnabled       bool              `json:"performance_insights_enabled"`
	S3Import                         S3Import          `json:"s3_import"`
	Timeouts                         []Timeouts        `json:"timeouts"`
}

type S3Import struct {
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DbInstanceStatus defines the observed state of a DbInstance
type DbInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbInstanceObservation `json:",inline"`
}

// A DbInstanceObservation records the observed state of a DbInstance
type DbInstanceObservation struct {
	PerformanceInsightsKmsKeyId        string   `json:"performance_insights_kms_key_id"`
	Arn                                string   `json:"arn"`
	ResourceId                         string   `json:"resource_id"`
	Port                               int      `json:"port"`
	Timezone                           string   `json:"timezone"`
	BackupRetentionPeriod              int      `json:"backup_retention_period"`
	MultiAz                            bool     `json:"multi_az"`
	VpcSecurityGroupIds                []string `json:"vpc_security_group_ids"`
	Address                            string   `json:"address"`
	BackupWindow                       string   `json:"backup_window"`
	HostedZoneId                       string   `json:"hosted_zone_id"`
	Replicas                           []string `json:"replicas"`
	Endpoint                           string   `json:"endpoint"`
	CaCertIdentifier                   string   `json:"ca_cert_identifier"`
	PerformanceInsightsRetentionPeriod int      `json:"performance_insights_retention_period"`
	ApplyImmediately                   bool     `json:"apply_immediately"`
	Username                           string   `json:"username"`
	DbSubnetGroupName                  string   `json:"db_subnet_group_name"`
	AllocatedStorage                   int      `json:"allocated_storage"`
	Id                                 string   `json:"id"`
	CharacterSetName                   string   `json:"character_set_name"`
	Name                               string   `json:"name"`
	KmsKeyId                           string   `json:"kms_key_id"`
	MonitoringRoleArn                  string   `json:"monitoring_role_arn"`
	EngineVersion                      string   `json:"engine_version"`
	IdentifierPrefix                   string   `json:"identifier_prefix"`
	MaintenanceWindow                  string   `json:"maintenance_window"`
	Status                             string   `json:"status"`
	Engine                             string   `json:"engine"`
	OptionGroupName                    string   `json:"option_group_name"`
	AvailabilityZone                   string   `json:"availability_zone"`
	Identifier                         string   `json:"identifier"`
	ParameterGroupName                 string   `json:"parameter_group_name"`
	StorageType                        string   `json:"storage_type"`
	LicenseModel                       string   `json:"license_model"`
}