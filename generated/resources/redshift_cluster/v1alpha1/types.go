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

// RedshiftCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RedshiftCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedshiftClusterSpec   `json:"spec"`
	Status RedshiftClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedshiftCluster contains a list of RedshiftClusterList
type RedshiftClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedshiftCluster `json:"items"`
}

// A RedshiftClusterSpec defines the desired state of a RedshiftCluster
type RedshiftClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RedshiftClusterParameters `json:",inline"`
}

// A RedshiftClusterParameters defines the desired state of a RedshiftCluster
type RedshiftClusterParameters struct {
	ClusterIdentifier                string            `json:"cluster_identifier"`
	Encrypted                        bool              `json:"encrypted"`
	IamRoles                         []string          `json:"iam_roles"`
	VpcSecurityGroupIds              []string          `json:"vpc_security_group_ids"`
	ClusterPublicKey                 string            `json:"cluster_public_key"`
	ElasticIp                        string            `json:"elastic_ip"`
	KmsKeyId                         string            `json:"kms_key_id"`
	OwnerAccount                     string            `json:"owner_account"`
	SnapshotClusterIdentifier        string            `json:"snapshot_cluster_identifier"`
	AllowVersionUpgrade              bool              `json:"allow_version_upgrade"`
	AutomatedSnapshotRetentionPeriod int               `json:"automated_snapshot_retention_period"`
	MasterPassword                   string            `json:"master_password"`
	MasterUsername                   string            `json:"master_username"`
	NumberOfNodes                    int               `json:"number_of_nodes"`
	PreferredMaintenanceWindow       string            `json:"preferred_maintenance_window"`
	ClusterParameterGroupName        string            `json:"cluster_parameter_group_name"`
	PubliclyAccessible               bool              `json:"publicly_accessible"`
	Tags                             map[string]string `json:"tags"`
	SnapshotIdentifier               string            `json:"snapshot_identifier"`
	AvailabilityZone                 string            `json:"availability_zone"`
	ClusterRevisionNumber            string            `json:"cluster_revision_number"`
	ClusterType                      string            `json:"cluster_type"`
	EnhancedVpcRouting               bool              `json:"enhanced_vpc_routing"`
	FinalSnapshotIdentifier          string            `json:"final_snapshot_identifier"`
	Id                               string            `json:"id"`
	SkipFinalSnapshot                bool              `json:"skip_final_snapshot"`
	DatabaseName                     string            `json:"database_name"`
	Endpoint                         string            `json:"endpoint"`
	Port                             int               `json:"port"`
	ClusterSecurityGroups            []string          `json:"cluster_security_groups"`
	ClusterSubnetGroupName           string            `json:"cluster_subnet_group_name"`
	ClusterVersion                   string            `json:"cluster_version"`
	NodeType                         string            `json:"node_type"`
	Logging                          Logging           `json:"logging"`
	SnapshotCopy                     SnapshotCopy      `json:"snapshot_copy"`
	Timeouts                         []Timeouts        `json:"timeouts"`
}

type Logging struct {
	BucketName  string `json:"bucket_name"`
	Enable      bool   `json:"enable"`
	S3KeyPrefix string `json:"s3_key_prefix"`
}

type SnapshotCopy struct {
	DestinationRegion string `json:"destination_region"`
	GrantName         string `json:"grant_name"`
	RetentionPeriod   int    `json:"retention_period"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A RedshiftClusterStatus defines the observed state of a RedshiftCluster
type RedshiftClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftClusterObservation `json:",inline"`
}

// A RedshiftClusterObservation records the observed state of a RedshiftCluster
type RedshiftClusterObservation struct {
	DnsName string `json:"dns_name"`
	Arn     string `json:"arn"`
}