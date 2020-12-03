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

// DaxCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DaxCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DaxClusterSpec   `json:"spec"`
	Status DaxClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DaxCluster contains a list of DaxClusterList
type DaxClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaxCluster `json:"items"`
}

// A DaxClusterSpec defines the desired state of a DaxCluster
type DaxClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DaxClusterParameters `json:",inline"`
}

// A DaxClusterParameters defines the desired state of a DaxCluster
type DaxClusterParameters struct {
	ClusterName          string               `json:"cluster_name"`
	IamRoleArn           string               `json:"iam_role_arn"`
	MaintenanceWindow    string               `json:"maintenance_window"`
	ParameterGroupName   string               `json:"parameter_group_name"`
	ReplicationFactor    int                  `json:"replication_factor"`
	AvailabilityZones    []string             `json:"availability_zones"`
	Id                   string               `json:"id"`
	SubnetGroupName      string               `json:"subnet_group_name"`
	NodeType             string               `json:"node_type"`
	NotificationTopicArn string               `json:"notification_topic_arn"`
	SecurityGroupIds     []string             `json:"security_group_ids"`
	Tags                 map[string]string    `json:"tags"`
	Description          string               `json:"description"`
	Timeouts             []Timeouts           `json:"timeouts"`
	ServerSideEncryption ServerSideEncryption `json:"server_side_encryption"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type ServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

// A DaxClusterStatus defines the observed state of a DaxCluster
type DaxClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DaxClusterObservation `json:",inline"`
}

// A DaxClusterObservation records the observed state of a DaxCluster
type DaxClusterObservation struct {
	ConfigurationEndpoint string  `json:"configuration_endpoint"`
	Port                  int     `json:"port"`
	Nodes                 []Nodes `json:"nodes"`
	ClusterAddress        string  `json:"cluster_address"`
	Arn                   string  `json:"arn"`
}

type Nodes struct {
	Address          string `json:"address"`
	AvailabilityZone string `json:"availability_zone"`
	Id               string `json:"id"`
	Port             int    `json:"port"`
}