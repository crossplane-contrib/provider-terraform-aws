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
	ForProvider                  NeptuneClusterParameters `json:"forProvider"`
}

// A NeptuneClusterParameters defines the desired state of a NeptuneCluster
type NeptuneClusterParameters struct {
	KmsKeyArn                        string            `json:"kms_key_arn"`
	EngineVersion                    string            `json:"engine_version"`
	PreferredBackupWindow            string            `json:"preferred_backup_window"`
	ReplicationSourceIdentifier      string            `json:"replication_source_identifier"`
	IamRoles                         []string          `json:"iam_roles"`
	NeptuneSubnetGroupName           string            `json:"neptune_subnet_group_name"`
	PreferredMaintenanceWindow       string            `json:"preferred_maintenance_window"`
	SnapshotIdentifier               string            `json:"snapshot_identifier"`
	VpcSecurityGroupIds              []string          `json:"vpc_security_group_ids"`
	DeletionProtection               bool              `json:"deletion_protection"`
	Engine                           string            `json:"engine"`
	IamDatabaseAuthenticationEnabled bool              `json:"iam_database_authentication_enabled"`
	StorageEncrypted                 bool              `json:"storage_encrypted"`
	ApplyImmediately                 bool              `json:"apply_immediately"`
	SkipFinalSnapshot                bool              `json:"skip_final_snapshot"`
	Tags                             map[string]string `json:"tags"`
	AvailabilityZones                []string          `json:"availability_zones"`
	BackupRetentionPeriod            int64             `json:"backup_retention_period"`
	ClusterIdentifierPrefix          string            `json:"cluster_identifier_prefix"`
	EnableCloudwatchLogsExports      []string          `json:"enable_cloudwatch_logs_exports"`
	NeptuneClusterParameterGroupName string            `json:"neptune_cluster_parameter_group_name"`
	ClusterIdentifier                string            `json:"cluster_identifier"`
	FinalSnapshotIdentifier          string            `json:"final_snapshot_identifier"`
	Port                             int64             `json:"port"`
	Timeouts                         Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A NeptuneClusterStatus defines the observed state of a NeptuneCluster
type NeptuneClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NeptuneClusterObservation `json:"atProvider"`
}

// A NeptuneClusterObservation records the observed state of a NeptuneCluster
type NeptuneClusterObservation struct {
	Arn               string   `json:"arn"`
	HostedZoneId      string   `json:"hosted_zone_id"`
	ReaderEndpoint    string   `json:"reader_endpoint"`
	ClusterMembers    []string `json:"cluster_members"`
	ClusterResourceId string   `json:"cluster_resource_id"`
	Endpoint          string   `json:"endpoint"`
}