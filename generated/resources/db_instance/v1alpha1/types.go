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
	Tags                               map[string]string `json:"tags"`
	AllocatedStorage                   int               `json:"allocated_storage"`
	Id                                 string            `json:"id"`
	DeleteAutomatedBackups             bool              `json:"delete_automated_backups"`
	EnabledCloudwatchLogsExports       []string          `json:"enabled_cloudwatch_logs_exports"`
	PerformanceInsightsKmsKeyId        string            `json:"performance_insights_kms_key_id"`
	Port                               int               `json:"port"`
	StorageType                        string            `json:"storage_type"`
	AllowMajorVersionUpgrade           bool              `json:"allow_major_version_upgrade"`
	CopyTagsToSnapshot                 bool              `json:"copy_tags_to_snapshot"`
	PubliclyAccessible                 bool              `json:"publicly_accessible"`
	AutoMinorVersionUpgrade            bool              `json:"auto_minor_version_upgrade"`
	AvailabilityZone                   string            `json:"availability_zone"`
	IdentifierPrefix                   string            `json:"identifier_prefix"`
	PerformanceInsightsEnabled         bool              `json:"performance_insights_enabled"`
	StorageEncrypted                   bool              `json:"storage_encrypted"`
	BackupRetentionPeriod              int               `json:"backup_retention_period"`
	Identifier                         string            `json:"identifier"`
	SkipFinalSnapshot                  bool              `json:"skip_final_snapshot"`
	ApplyImmediately                   bool              `json:"apply_immediately"`
	FinalSnapshotIdentifier            string            `json:"final_snapshot_identifier"`
	Iops                               int               `json:"iops"`
	KmsKeyId                           string            `json:"kms_key_id"`
	MultiAz                            bool              `json:"multi_az"`
	Name                               string            `json:"name"`
	Password                           string            `json:"password"`
	DomainIamRoleName                  string            `json:"domain_iam_role_name"`
	IamDatabaseAuthenticationEnabled   bool              `json:"iam_database_authentication_enabled"`
	SnapshotIdentifier                 string            `json:"snapshot_identifier"`
	VpcSecurityGroupIds                []string          `json:"vpc_security_group_ids"`
	MonitoringInterval                 int               `json:"monitoring_interval"`
	SecurityGroupNames                 []string          `json:"security_group_names"`
	EngineVersion                      string            `json:"engine_version"`
	Engine                             string            `json:"engine"`
	InstanceClass                      string            `json:"instance_class"`
	ReplicateSourceDb                  string            `json:"replicate_source_db"`
	Timezone                           string            `json:"timezone"`
	CharacterSetName                   string            `json:"character_set_name"`
	DbSubnetGroupName                  string            `json:"db_subnet_group_name"`
	MonitoringRoleArn                  string            `json:"monitoring_role_arn"`
	Domain                             string            `json:"domain"`
	MaxAllocatedStorage                int               `json:"max_allocated_storage"`
	LicenseModel                       string            `json:"license_model"`
	OptionGroupName                    string            `json:"option_group_name"`
	ParameterGroupName                 string            `json:"parameter_group_name"`
	CaCertIdentifier                   string            `json:"ca_cert_identifier"`
	DeletionProtection                 bool              `json:"deletion_protection"`
	Username                           string            `json:"username"`
	MaintenanceWindow                  string            `json:"maintenance_window"`
	PerformanceInsightsRetentionPeriod int               `json:"performance_insights_retention_period"`
	BackupWindow                       string            `json:"backup_window"`
	Timeouts                           []Timeouts        `json:"timeouts"`
	S3Import                           S3Import          `json:"s3_import"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

type S3Import struct {
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
}

// A DbInstanceStatus defines the observed state of a DbInstance
type DbInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbInstanceObservation `json:",inline"`
}

// A DbInstanceObservation records the observed state of a DbInstance
type DbInstanceObservation struct {
	Endpoint     string   `json:"endpoint"`
	Status       string   `json:"status"`
	ResourceId   string   `json:"resource_id"`
	Arn          string   `json:"arn"`
	Replicas     []string `json:"replicas"`
	HostedZoneId string   `json:"hosted_zone_id"`
	Address      string   `json:"address"`
}