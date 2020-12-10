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
	Domain                             string            `json:"domain"`
	IdentifierPrefix                   string            `json:"identifier_prefix"`
	PerformanceInsightsRetentionPeriod int64             `json:"performance_insights_retention_period"`
	StorageType                        string            `json:"storage_type"`
	SnapshotIdentifier                 string            `json:"snapshot_identifier"`
	EnabledCloudwatchLogsExports       []string          `json:"enabled_cloudwatch_logs_exports"`
	EngineVersion                      string            `json:"engine_version"`
	CharacterSetName                   string            `json:"character_set_name"`
	DbSubnetGroupName                  string            `json:"db_subnet_group_name"`
	FinalSnapshotIdentifier            string            `json:"final_snapshot_identifier"`
	Identifier                         string            `json:"identifier"`
	CopyTagsToSnapshot                 bool              `json:"copy_tags_to_snapshot"`
	DeletionProtection                 bool              `json:"deletion_protection"`
	DomainIamRoleName                  string            `json:"domain_iam_role_name"`
	Engine                             string            `json:"engine"`
	Username                           string            `json:"username"`
	MultiAz                            bool              `json:"multi_az"`
	OptionGroupName                    string            `json:"option_group_name"`
	ParameterGroupName                 string            `json:"parameter_group_name"`
	SkipFinalSnapshot                  bool              `json:"skip_final_snapshot"`
	AvailabilityZone                   string            `json:"availability_zone"`
	DeleteAutomatedBackups             bool              `json:"delete_automated_backups"`
	SecurityGroupNames                 []string          `json:"security_group_names"`
	AutoMinorVersionUpgrade            bool              `json:"auto_minor_version_upgrade"`
	PerformanceInsightsKmsKeyId        string            `json:"performance_insights_kms_key_id"`
	MonitoringInterval                 int64             `json:"monitoring_interval"`
	PerformanceInsightsEnabled         bool              `json:"performance_insights_enabled"`
	ApplyImmediately                   bool              `json:"apply_immediately"`
	Id                                 string            `json:"id"`
	LicenseModel                       string            `json:"license_model"`
	MaxAllocatedStorage                int64             `json:"max_allocated_storage"`
	CaCertIdentifier                   string            `json:"ca_cert_identifier"`
	MaintenanceWindow                  string            `json:"maintenance_window"`
	MonitoringRoleArn                  string            `json:"monitoring_role_arn"`
	IamDatabaseAuthenticationEnabled   bool              `json:"iam_database_authentication_enabled"`
	PubliclyAccessible                 bool              `json:"publicly_accessible"`
	Timezone                           string            `json:"timezone"`
	Name                               string            `json:"name"`
	StorageEncrypted                   bool              `json:"storage_encrypted"`
	Tags                               map[string]string `json:"tags"`
	VpcSecurityGroupIds                []string          `json:"vpc_security_group_ids"`
	AllocatedStorage                   int64             `json:"allocated_storage"`
	AllowMajorVersionUpgrade           bool              `json:"allow_major_version_upgrade"`
	BackupRetentionPeriod              int64             `json:"backup_retention_period"`
	KmsKeyId                           string            `json:"kms_key_id"`
	BackupWindow                       string            `json:"backup_window"`
	InstanceClass                      string            `json:"instance_class"`
	Iops                               int64             `json:"iops"`
	Password                           string            `json:"password"`
	Port                               int64             `json:"port"`
	ReplicateSourceDb                  string            `json:"replicate_source_db"`
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
	Status       string   `json:"status"`
	Arn          string   `json:"arn"`
	Replicas     []string `json:"replicas"`
	ResourceId   string   `json:"resource_id"`
	HostedZoneId string   `json:"hosted_zone_id"`
	Address      string   `json:"address"`
	Endpoint     string   `json:"endpoint"`
}