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
	Password                         string   `json:"password"`
	PerformanceInsightsEnabled       bool     `json:"performance_insights_enabled"`
	Iops                             int      `json:"iops"`
	EnabledCloudwatchLogsExports     []string `json:"enabled_cloudwatch_logs_exports"`
	MaxAllocatedStorage              int      `json:"max_allocated_storage"`
	MonitoringInterval               int      `json:"monitoring_interval"`
	PubliclyAccessible               bool     `json:"publicly_accessible"`
	SecurityGroupNames               []string `json:"security_group_names"`
	Domain                           string   `json:"domain"`
	SnapshotIdentifier               string   `json:"snapshot_identifier"`
	CopyTagsToSnapshot               bool     `json:"copy_tags_to_snapshot"`
	DomainIamRoleName                string   `json:"domain_iam_role_name"`
	DeletionProtection               bool     `json:"deletion_protection"`
	AutoMinorVersionUpgrade          bool     `json:"auto_minor_version_upgrade"`
	DeleteAutomatedBackups           bool     `json:"delete_automated_backups"`
	StorageEncrypted                 bool     `json:"storage_encrypted"`
	InstanceClass                    string   `json:"instance_class"`
	ReplicateSourceDb                string   `json:"replicate_source_db"`
	SkipFinalSnapshot                bool     `json:"skip_final_snapshot"`
	AllowMajorVersionUpgrade         bool     `json:"allow_major_version_upgrade"`
	IamDatabaseAuthenticationEnabled bool     `json:"iam_database_authentication_enabled"`
	FinalSnapshotIdentifier          string   `json:"final_snapshot_identifier"`
}

// A DbInstanceStatus defines the observed state of a DbInstance
type DbInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbInstanceObservation `json:",inline"`
}

// A DbInstanceObservation records the observed state of a DbInstance
type DbInstanceObservation struct {
	DbSubnetGroupName                  string   `json:"db_subnet_group_name"`
	MonitoringRoleArn                  string   `json:"monitoring_role_arn"`
	Username                           string   `json:"username"`
	AllocatedStorage                   int      `json:"allocated_storage"`
	CaCertIdentifier                   string   `json:"ca_cert_identifier"`
	StorageType                        string   `json:"storage_type"`
	Address                            string   `json:"address"`
	CharacterSetName                   string   `json:"character_set_name"`
	AvailabilityZone                   string   `json:"availability_zone"`
	MaintenanceWindow                  string   `json:"maintenance_window"`
	OptionGroupName                    string   `json:"option_group_name"`
	ParameterGroupName                 string   `json:"parameter_group_name"`
	VpcSecurityGroupIds                []string `json:"vpc_security_group_ids"`
	IdentifierPrefix                   string   `json:"identifier_prefix"`
	Status                             string   `json:"status"`
	Id                                 string   `json:"id"`
	ResourceId                         string   `json:"resource_id"`
	HostedZoneId                       string   `json:"hosted_zone_id"`
	KmsKeyId                           string   `json:"kms_key_id"`
	Name                               string   `json:"name"`
	PerformanceInsightsRetentionPeriod int      `json:"performance_insights_retention_period"`
	BackupWindow                       string   `json:"backup_window"`
	EngineVersion                      string   `json:"engine_version"`
	LicenseModel                       string   `json:"license_model"`
	Engine                             string   `json:"engine"`
	ApplyImmediately                   bool     `json:"apply_immediately"`
	Identifier                         string   `json:"identifier"`
	MultiAz                            bool     `json:"multi_az"`
	PerformanceInsightsKmsKeyId        string   `json:"performance_insights_kms_key_id"`
	BackupRetentionPeriod              int      `json:"backup_retention_period"`
	Endpoint                           string   `json:"endpoint"`
	Arn                                string   `json:"arn"`
	Port                               int      `json:"port"`
	Replicas                           []string `json:"replicas"`
	Timezone                           string   `json:"timezone"`
}