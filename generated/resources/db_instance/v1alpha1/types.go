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
	PerformanceInsightsRetentionPeriod int               `json:"performance_insights_retention_period"`
	BackupRetentionPeriod              int               `json:"backup_retention_period"`
	Domain                             string            `json:"domain"`
	CharacterSetName                   string            `json:"character_set_name"`
	FinalSnapshotIdentifier            string            `json:"final_snapshot_identifier"`
	StorageEncrypted                   bool              `json:"storage_encrypted"`
	AvailabilityZone                   string            `json:"availability_zone"`
	Password                           string            `json:"password"`
	PerformanceInsightsKmsKeyId        string            `json:"performance_insights_kms_key_id"`
	StorageType                        string            `json:"storage_type"`
	EngineVersion                      string            `json:"engine_version"`
	Identifier                         string            `json:"identifier"`
	ParameterGroupName                 string            `json:"parameter_group_name"`
	CopyTagsToSnapshot                 bool              `json:"copy_tags_to_snapshot"`
	InstanceClass                      string            `json:"instance_class"`
	Iops                               int               `json:"iops"`
	PubliclyAccessible                 bool              `json:"publicly_accessible"`
	Tags                               map[string]string `json:"tags"`
	AutoMinorVersionUpgrade            bool              `json:"auto_minor_version_upgrade"`
	Engine                             string            `json:"engine"`
	Timezone                           string            `json:"timezone"`
	VpcSecurityGroupIds                []string          `json:"vpc_security_group_ids"`
	ApplyImmediately                   bool              `json:"apply_immediately"`
	IamDatabaseAuthenticationEnabled   bool              `json:"iam_database_authentication_enabled"`
	MaxAllocatedStorage                int               `json:"max_allocated_storage"`
	SkipFinalSnapshot                  bool              `json:"skip_final_snapshot"`
	BackupWindow                       string            `json:"backup_window"`
	Id                                 string            `json:"id"`
	KmsKeyId                           string            `json:"kms_key_id"`
	Username                           string            `json:"username"`
	DeletionProtection                 bool              `json:"deletion_protection"`
	IdentifierPrefix                   string            `json:"identifier_prefix"`
	MaintenanceWindow                  string            `json:"maintenance_window"`
	MonitoringInterval                 int               `json:"monitoring_interval"`
	MultiAz                            bool              `json:"multi_az"`
	Name                               string            `json:"name"`
	ReplicateSourceDb                  string            `json:"replicate_source_db"`
	CaCertIdentifier                   string            `json:"ca_cert_identifier"`
	DbSubnetGroupName                  string            `json:"db_subnet_group_name"`
	OptionGroupName                    string            `json:"option_group_name"`
	SecurityGroupNames                 []string          `json:"security_group_names"`
	PerformanceInsightsEnabled         bool              `json:"performance_insights_enabled"`
	AllocatedStorage                   int               `json:"allocated_storage"`
	AllowMajorVersionUpgrade           bool              `json:"allow_major_version_upgrade"`
	EnabledCloudwatchLogsExports       []string          `json:"enabled_cloudwatch_logs_exports"`
	LicenseModel                       string            `json:"license_model"`
	MonitoringRoleArn                  string            `json:"monitoring_role_arn"`
	SnapshotIdentifier                 string            `json:"snapshot_identifier"`
	DeleteAutomatedBackups             bool              `json:"delete_automated_backups"`
	DomainIamRoleName                  string            `json:"domain_iam_role_name"`
	Port                               int               `json:"port"`
	S3Import                           S3Import          `json:"s3_import"`
	Timeouts                           []Timeouts        `json:"timeouts"`
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
	Replicas     []string `json:"replicas"`
	Endpoint     string   `json:"endpoint"`
	Address      string   `json:"address"`
	Status       string   `json:"status"`
	Arn          string   `json:"arn"`
	HostedZoneId string   `json:"hosted_zone_id"`
	ResourceId   string   `json:"resource_id"`
}