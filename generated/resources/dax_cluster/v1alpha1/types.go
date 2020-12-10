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
	Id                   string               `json:"id"`
	NotificationTopicArn string               `json:"notification_topic_arn"`
	SubnetGroupName      string               `json:"subnet_group_name"`
	ClusterName          string               `json:"cluster_name"`
	Description          string               `json:"description"`
	ReplicationFactor    int64                `json:"replication_factor"`
	IamRoleArn           string               `json:"iam_role_arn"`
	NodeType             string               `json:"node_type"`
	Tags                 map[string]string    `json:"tags"`
	AvailabilityZones    []string             `json:"availability_zones"`
	MaintenanceWindow    string               `json:"maintenance_window"`
	ParameterGroupName   string               `json:"parameter_group_name"`
	SecurityGroupIds     []string             `json:"security_group_ids"`
	ServerSideEncryption ServerSideEncryption `json:"server_side_encryption"`
	Timeouts             Timeouts             `json:"timeouts"`
}

type ServerSideEncryption struct {
	Enabled bool `json:"enabled"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DaxClusterStatus defines the observed state of a DaxCluster
type DaxClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DaxClusterObservation `json:",inline"`
}

// A DaxClusterObservation records the observed state of a DaxCluster
type DaxClusterObservation struct {
	ConfigurationEndpoint string  `json:"configuration_endpoint"`
	Port                  int64   `json:"port"`
	Arn                   string  `json:"arn"`
	Nodes                 []Nodes `json:"nodes"`
	ClusterAddress        string  `json:"cluster_address"`
}

type Nodes struct {
	AvailabilityZone string `json:"availability_zone"`
	Id               string `json:"id"`
	Port             int64  `json:"port"`
	Address          string `json:"address"`
}