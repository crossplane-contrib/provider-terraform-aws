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
	ForProvider                  DbInstanceParameters `json:"forProvider"`
}

// A DbInstanceParameters defines the desired state of a DbInstance
type DbInstanceParameters struct {
	Iops                               int64             `json:"iops"`
	ParameterGroupName                 string            `json:"parameter_group_name"`
	AllowMajorVersionUpgrade           bool              `json:"allow_major_version_upgrade"`
	BackupWindow                       string            `json:"backup_window"`
	PerformanceInsightsEnabled         bool              `json:"performance_insights_enabled"`
	StorageType                        string            `json:"storage_type"`
	ApplyImmediately                   bool              `json:"apply_immediately"`
	CharacterSetName                   string            `json:"character_set_name"`
	InstanceClass                      string            `json:"instance_class"`
	MonitoringRoleArn                  string            `json:"monitoring_role_arn"`
	SnapshotIdentifier                 string            `json:"snapshot_identifier"`
	AvailabilityZone                   string            `json:"availability_zone"`
	IamDatabaseAuthenticationEnabled   bool              `json:"iam_database_authentication_enabled"`
	Timezone                           string            `json:"timezone"`
	KmsKeyId                           string            `json:"kms_key_id"`
	OptionGroupName                    string            `json:"option_group_name"`
	VpcSecurityGroupIds                []string          `json:"vpc_security_group_ids"`
	BackupRetentionPeriod              int64             `json:"backup_retention_period"`
	DeleteAutomatedBackups             bool              `json:"delete_automated_backups"`
	Identifier                         string            `json:"identifier"`
	LicenseModel                       string            `json:"license_model"`
	AllocatedStorage                   int64             `json:"allocated_storage"`
	EnabledCloudwatchLogsExports       []string          `json:"enabled_cloudwatch_logs_exports"`
	FinalSnapshotIdentifier            string            `json:"final_snapshot_identifier"`
	CopyTagsToSnapshot                 bool              `json:"copy_tags_to_snapshot"`
	DbSubnetGroupName                  string            `json:"db_subnet_group_name"`
	EngineVersion                      string            `json:"engine_version"`
	Password                           string            `json:"password"`
	CaCertIdentifier                   string            `json:"ca_cert_identifier"`
	MonitoringInterval                 int64             `json:"monitoring_interval"`
	MultiAz                            bool              `json:"multi_az"`
	Name                               string            `json:"name"`
	ReplicateSourceDb                  string            `json:"replicate_source_db"`
	SecurityGroupNames                 []string          `json:"security_group_names"`
	DomainIamRoleName                  string            `json:"domain_iam_role_name"`
	IdentifierPrefix                   string            `json:"identifier_prefix"`
	PubliclyAccessible                 bool              `json:"publicly_accessible"`
	StorageEncrypted                   bool              `json:"storage_encrypted"`
	DeletionProtection                 bool              `json:"deletion_protection"`
	SkipFinalSnapshot                  bool              `json:"skip_final_snapshot"`
	MaxAllocatedStorage                int64             `json:"max_allocated_storage"`
	PerformanceInsightsKmsKeyId        string            `json:"performance_insights_kms_key_id"`
	AutoMinorVersionUpgrade            bool              `json:"auto_minor_version_upgrade"`
	Engine                             string            `json:"engine"`
	MaintenanceWindow                  string            `json:"maintenance_window"`
	PerformanceInsightsRetentionPeriod int64             `json:"performance_insights_retention_period"`
	Domain                             string            `json:"domain"`
	Id                                 string            `json:"id"`
	Port                               int64             `json:"port"`
	Tags                               map[string]string `json:"tags"`
	Username                           string            `json:"username"`
	S3Import                           S3Import          `json:"s3_import"`
	Timeouts                           Timeouts          `json:"timeouts"`
}

type S3Import struct {
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
	BucketName          string `json:"bucket_name"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DbInstanceStatus defines the observed state of a DbInstance
type DbInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbInstanceObservation `json:"atProvider"`
}

// A DbInstanceObservation records the observed state of a DbInstance
type DbInstanceObservation struct {
	Address      string   `json:"address"`
	Replicas     []string `json:"replicas"`
	ResourceId   string   `json:"resource_id"`
	Endpoint     string   `json:"endpoint"`
	Arn          string   `json:"arn"`
	HostedZoneId string   `json:"hosted_zone_id"`
	Status       string   `json:"status"`
}