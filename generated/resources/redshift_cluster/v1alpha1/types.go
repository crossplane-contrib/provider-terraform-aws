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
	Encrypted                        bool   `json:"encrypted"`
	MasterPassword                   string `json:"master_password"`
	OwnerAccount                     string `json:"owner_account"`
	AllowVersionUpgrade              bool   `json:"allow_version_upgrade"`
	FinalSnapshotIdentifier          string `json:"final_snapshot_identifier"`
	NodeType                         string `json:"node_type"`
	ClusterIdentifier                string `json:"cluster_identifier"`
	PubliclyAccessible               bool   `json:"publicly_accessible"`
	SnapshotIdentifier               string `json:"snapshot_identifier"`
	MasterUsername                   string `json:"master_username"`
	ClusterVersion                   string `json:"cluster_version"`
	Port                             int    `json:"port"`
	AutomatedSnapshotRetentionPeriod int    `json:"automated_snapshot_retention_period"`
	NumberOfNodes                    int    `json:"number_of_nodes"`
	SkipFinalSnapshot                bool   `json:"skip_final_snapshot"`
	SnapshotClusterIdentifier        string `json:"snapshot_cluster_identifier"`
	ElasticIp                        string `json:"elastic_ip"`
}

// A RedshiftClusterStatus defines the observed state of a RedshiftCluster
type RedshiftClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RedshiftClusterObservation `json:",inline"`
}

// A RedshiftClusterObservation records the observed state of a RedshiftCluster
type RedshiftClusterObservation struct {
	IamRoles                   []string `json:"iam_roles"`
	KmsKeyId                   string   `json:"kms_key_id"`
	ClusterPublicKey           string   `json:"cluster_public_key"`
	ClusterSecurityGroups      []string `json:"cluster_security_groups"`
	ClusterSubnetGroupName     string   `json:"cluster_subnet_group_name"`
	DnsName                    string   `json:"dns_name"`
	Endpoint                   string   `json:"endpoint"`
	ClusterParameterGroupName  string   `json:"cluster_parameter_group_name"`
	Arn                        string   `json:"arn"`
	DatabaseName               string   `json:"database_name"`
	PreferredMaintenanceWindow string   `json:"preferred_maintenance_window"`
	AvailabilityZone           string   `json:"availability_zone"`
	VpcSecurityGroupIds        []string `json:"vpc_security_group_ids"`
	ClusterType                string   `json:"cluster_type"`
	EnhancedVpcRouting         bool     `json:"enhanced_vpc_routing"`
	Id                         string   `json:"id"`
	ClusterRevisionNumber      string   `json:"cluster_revision_number"`
}