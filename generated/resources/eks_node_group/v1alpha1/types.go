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

// EksNodeGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EksNodeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EksNodeGroupSpec   `json:"spec"`
	Status EksNodeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EksNodeGroup contains a list of EksNodeGroupList
type EksNodeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EksNodeGroup `json:"items"`
}

// A EksNodeGroupSpec defines the desired state of a EksNodeGroup
type EksNodeGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EksNodeGroupParameters `json:"forProvider"`
}

// A EksNodeGroupParameters defines the desired state of a EksNodeGroup
type EksNodeGroupParameters struct {
	DiskSize           int64             `json:"disk_size"`
	ForceUpdateVersion bool              `json:"force_update_version"`
	NodeRoleArn        string            `json:"node_role_arn"`
	AmiType            string            `json:"ami_type"`
	ClusterName        string            `json:"cluster_name"`
	Labels             map[string]string `json:"labels"`
	ReleaseVersion     string            `json:"release_version"`
	Tags               map[string]string `json:"tags"`
	Version            string            `json:"version"`
	Id                 string            `json:"id"`
	InstanceTypes      []string          `json:"instance_types"`
	NodeGroupName      string            `json:"node_group_name"`
	SubnetIds          []string          `json:"subnet_ids"`
	Timeouts           Timeouts          `json:"timeouts"`
	LaunchTemplate     LaunchTemplate    `json:"launch_template"`
	RemoteAccess       RemoteAccess      `json:"remote_access"`
	ScalingConfig      ScalingConfig     `json:"scaling_config"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type LaunchTemplate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type RemoteAccess struct {
	Ec2SshKey              string   `json:"ec2_ssh_key"`
	SourceSecurityGroupIds []string `json:"source_security_group_ids"`
}

type ScalingConfig struct {
	DesiredSize int64 `json:"desired_size"`
	MaxSize     int64 `json:"max_size"`
	MinSize     int64 `json:"min_size"`
}

// A EksNodeGroupStatus defines the observed state of a EksNodeGroup
type EksNodeGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EksNodeGroupObservation `json:"atProvider"`
}

// A EksNodeGroupObservation records the observed state of a EksNodeGroup
type EksNodeGroupObservation struct {
	Status    string      `json:"status"`
	Arn       string      `json:"arn"`
	Resources []Resources `json:"resources"`
}

type Resources struct {
	AutoscalingGroups           []AutoscalingGroups `json:"autoscaling_groups"`
	RemoteAccessSecurityGroupId string              `json:"remote_access_security_group_id"`
}

type AutoscalingGroups struct {
	Name string `json:"name"`
}