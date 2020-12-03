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

// NeptuneCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NeptuneCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NeptuneClusterSpec   `json:"spec"`
	Status NeptuneClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NeptuneCluster contains a list of NeptuneClusterList
type NeptuneClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NeptuneCluster `json:"items"`
}

// A NeptuneClusterSpec defines the desired state of a NeptuneCluster
type NeptuneClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NeptuneClusterParameters `json:",inline"`
}

// A NeptuneClusterParameters defines the desired state of a NeptuneCluster
type NeptuneClusterParameters struct {
	IamDatabaseAuthenticationEnabled bool              `json:"iam_database_authentication_enabled"`
	ClusterIdentifier                string            `json:"cluster_identifier"`
	AvailabilityZones                []string          `json:"availability_zones"`
	Engine                           string            `json:"engine"`
	FinalSnapshotIdentifier          string            `json:"final_snapshot_identifier"`
	StorageEncrypted                 bool              `json:"storage_encrypted"`
	ApplyImmediately                 bool              `json:"apply_immediately"`
	EnableCloudwatchLogsExports      []string          `json:"enable_cloudwatch_logs_exports"`
	KmsKeyArn                        string            `json:"kms_key_arn"`
	Id                               string            `json:"id"`
	PreferredBackupWindow            string            `json:"preferred_backup_window"`
	IamRoles                         []string          `json:"iam_roles"`
	PreferredMaintenanceWindow       string            `json:"preferred_maintenance_window"`
	ReplicationSourceIdentifier      string            `json:"replication_source_identifier"`
	NeptuneClusterParameterGroupName string            `json:"neptune_cluster_parameter_group_name"`
	BackupRetentionPeriod            int               `json:"backup_retention_period"`
	NeptuneSubnetGroupName           string            `json:"neptune_subnet_group_name"`
	Tags                             map[string]string `json:"tags"`
	VpcSecurityGroupIds              []string          `json:"vpc_security_group_ids"`
	Port                             int               `json:"port"`
	SkipFinalSnapshot                bool              `json:"skip_final_snapshot"`
	EngineVersion                    string            `json:"engine_version"`
	DeletionProtection               bool              `json:"deletion_protection"`
	SnapshotIdentifier               string            `json:"snapshot_identifier"`
	ClusterIdentifierPrefix          string            `json:"cluster_identifier_prefix"`
	Timeouts                         []Timeouts        `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

// A NeptuneClusterStatus defines the observed state of a NeptuneCluster
type NeptuneClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NeptuneClusterObservation `json:",inline"`
}

// A NeptuneClusterObservation records the observed state of a NeptuneCluster
type NeptuneClusterObservation struct {
	ClusterResourceId string   `json:"cluster_resource_id"`
	ClusterMembers    []string `json:"cluster_members"`
	Endpoint          string   `json:"endpoint"`
	Arn               string   `json:"arn"`
	ReaderEndpoint    string   `json:"reader_endpoint"`
	HostedZoneId      string   `json:"hosted_zone_id"`
}