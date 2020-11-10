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

// ElasticacheCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticacheCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticacheClusterSpec   `json:"spec"`
	Status ElasticacheClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticacheCluster contains a list of ElasticacheClusterList
type ElasticacheClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticacheCluster `json:"items"`
}

// A ElasticacheClusterSpec defines the desired state of a ElasticacheCluster
type ElasticacheClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticacheClusterParameters `json:",inline"`
}

// A ElasticacheClusterParameters defines the desired state of a ElasticacheCluster
type ElasticacheClusterParameters struct {
	ClusterId                  string            `json:"cluster_id"`
	NotificationTopicArn       string            `json:"notification_topic_arn"`
	PreferredAvailabilityZones []string          `json:"preferred_availability_zones"`
	SnapshotName               string            `json:"snapshot_name"`
	Tags                       map[string]string `json:"tags"`
	SnapshotRetentionLimit     int               `json:"snapshot_retention_limit"`
	SnapshotArns               []string          `json:"snapshot_arns"`
}

// A ElasticacheClusterStatus defines the observed state of a ElasticacheCluster
type ElasticacheClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticacheClusterObservation `json:",inline"`
}

// A ElasticacheClusterObservation records the observed state of a ElasticacheCluster
type ElasticacheClusterObservation struct {
	ConfigurationEndpoint string   `json:"configuration_endpoint"`
	Port                  int      `json:"port"`
	SecurityGroupNames    []string `json:"security_group_names"`
	ClusterAddress        string   `json:"cluster_address"`
	Engine                string   `json:"engine"`
	NumCacheNodes         int      `json:"num_cache_nodes"`
	ApplyImmediately      bool     `json:"apply_immediately"`
	Id                    string   `json:"id"`
	MaintenanceWindow     string   `json:"maintenance_window"`
	NodeType              string   `json:"node_type"`
	SnapshotWindow        string   `json:"snapshot_window"`
	AvailabilityZone      string   `json:"availability_zone"`
	EngineVersion         string   `json:"engine_version"`
	SecurityGroupIds      []string `json:"security_group_ids"`
	SubnetGroupName       string   `json:"subnet_group_name"`
	Arn                   string   `json:"arn"`
	ParameterGroupName    string   `json:"parameter_group_name"`
	AzMode                string   `json:"az_mode"`
	ReplicationGroupId    string   `json:"replication_group_id"`
}