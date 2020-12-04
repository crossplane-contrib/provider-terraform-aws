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
	EngineVersion                      string            `json:"engine_version"`
	IamDatabaseAuthenticationEnabled   bool              `json:"iam_database_authentication_enabled"`
	Password                           string            `json:"password"`
	SkipFinalSnapshot                  bool              `json:"skip_final_snapshot"`
	PerformanceInsightsKmsKeyId        string            `json:"performance_insights_kms_key_id"`
	PerformanceInsightsRetentionPeriod int64             `json:"performance_insights_retention_period"`
	Iops                               int64             `json:"iops"`
	AllocatedStorage                   int64             `json:"allocated_storage"`
	CopyTagsToSnapshot                 bool              `json:"copy_tags_to_snapshot"`
	Identifier                         string            `json:"identifier"`
	FinalSnapshotIdentifier            string            `json:"final_snapshot_identifier"`
	MonitoringInterval                 int64             `json:"monitoring_interval"`
	MonitoringRoleArn                  string            `json:"monitoring_role_arn"`
	Tags                               map[string]string `json:"tags"`
	Engine                             string            `json:"engine"`
	Name                               string            `json:"name"`
	AvailabilityZone                   string            `json:"availability_zone"`
	BackupRetentionPeriod              int64             `json:"backup_retention_period"`
	InstanceClass                      string            `json:"instance_class"`
	ParameterGroupName                 string            `json:"parameter_group_name"`
	Username                           string            `json:"username"`
	BackupWindow                       string            `json:"backup_window"`
	CharacterSetName                   string            `json:"character_set_name"`
	MaxAllocatedStorage                int64             `json:"max_allocated_storage"`
	MultiAz                            bool              `json:"multi_az"`
	SecurityGroupNames                 []string          `json:"security_group_names"`
	VpcSecurityGroupIds                []string          `json:"vpc_security_group_ids"`
	DeletionProtection                 bool              `json:"deletion_protection"`
	EnabledCloudwatchLogsExports       []string          `json:"enabled_cloudwatch_logs_exports"`
	Id                                 string            `json:"id"`
	LicenseModel                       string            `json:"license_model"`
	DomainIamRoleName                  string            `json:"domain_iam_role_name"`
	KmsKeyId                           string            `json:"kms_key_id"`
	ApplyImmediately                   bool              `json:"apply_immediately"`
	OptionGroupName                    string            `json:"option_group_name"`
	ReplicateSourceDb                  string            `json:"replicate_source_db"`
	StorageType                        string            `json:"storage_type"`
	AllowMajorVersionUpgrade           bool              `json:"allow_major_version_upgrade"`
	AutoMinorVersionUpgrade            bool              `json:"auto_minor_version_upgrade"`
	StorageEncrypted                   bool              `json:"storage_encrypted"`
	Domain                             string            `json:"domain"`
	Port                               int64             `json:"port"`
	Timezone                           string            `json:"timezone"`
	DbSubnetGroupName                  string            `json:"db_subnet_group_name"`
	PerformanceInsightsEnabled         bool              `json:"performance_insights_enabled"`
	PubliclyAccessible                 bool              `json:"publicly_accessible"`
	DeleteAutomatedBackups             bool              `json:"delete_automated_backups"`
	IdentifierPrefix                   string            `json:"identifier_prefix"`
	CaCertIdentifier                   string            `json:"ca_cert_identifier"`
	MaintenanceWindow                  string            `json:"maintenance_window"`
	SnapshotIdentifier                 string            `json:"snapshot_identifier"`
	S3Import                           S3Import          `json:"s3_import"`
	Timeouts                           Timeouts          `json:"timeouts"`
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
	Arn          string   `json:"arn"`
	HostedZoneId string   `json:"hosted_zone_id"`
	Address      string   `json:"address"`
	ResourceId   string   `json:"resource_id"`
	Endpoint     string   `json:"endpoint"`
	Replicas     []string `json:"replicas"`
	Status       string   `json:"status"`
}