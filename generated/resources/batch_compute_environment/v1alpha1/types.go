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

// BatchComputeEnvironment is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type BatchComputeEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BatchComputeEnvironmentSpec   `json:"spec"`
	Status BatchComputeEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BatchComputeEnvironment contains a list of BatchComputeEnvironmentList
type BatchComputeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BatchComputeEnvironment `json:"items"`
}

// A BatchComputeEnvironmentSpec defines the desired state of a BatchComputeEnvironment
type BatchComputeEnvironmentSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  BatchComputeEnvironmentParameters `json:",inline"`
}

// A BatchComputeEnvironmentParameters defines the desired state of a BatchComputeEnvironment
type BatchComputeEnvironmentParameters struct {
	ServiceRole                  string           `json:"service_role"`
	ComputeEnvironmentNamePrefix string           `json:"compute_environment_name_prefix"`
	State                        string           `json:"state"`
	Type                         string           `json:"type"`
	ComputeResources             ComputeResources `json:"compute_resources"`
}

type ComputeResources struct {
	BidPercentage      int               `json:"bid_percentage"`
	Ec2KeyPair         string            `json:"ec2_key_pair"`
	ImageId            string            `json:"image_id"`
	AllocationStrategy string            `json:"allocation_strategy"`
	InstanceType       []string          `json:"instance_type"`
	MinVcpus           int               `json:"min_vcpus"`
	Subnets            []string          `json:"subnets"`
	Type               string            `json:"type"`
	InstanceRole       string            `json:"instance_role"`
	SpotIamFleetRole   string            `json:"spot_iam_fleet_role"`
	DesiredVcpus       int               `json:"desired_vcpus"`
	MaxVcpus           int               `json:"max_vcpus"`
	SecurityGroupIds   []string          `json:"security_group_ids"`
	Tags               map[string]string `json:"tags"`
	LaunchTemplate     LaunchTemplate    `json:"launch_template"`
}

type LaunchTemplate struct {
	LaunchTemplateName string `json:"launch_template_name"`
	Version            string `json:"version"`
	LaunchTemplateId   string `json:"launch_template_id"`
}

// A BatchComputeEnvironmentStatus defines the observed state of a BatchComputeEnvironment
type BatchComputeEnvironmentStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     BatchComputeEnvironmentObservation `json:",inline"`
}

// A BatchComputeEnvironmentObservation records the observed state of a BatchComputeEnvironment
type BatchComputeEnvironmentObservation struct {
	Arn                    string `json:"arn"`
	ComputeEnvironmentName string `json:"compute_environment_name"`
	EcsClusterArn          string `json:"ecs_cluster_arn"`
	Status                 string `json:"status"`
	Id                     string `json:"id"`
	StatusReason           string `json:"status_reason"`
}