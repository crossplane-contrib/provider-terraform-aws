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

// EcsTaskDefinition is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EcsTaskDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EcsTaskDefinitionSpec   `json:"spec"`
	Status EcsTaskDefinitionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcsTaskDefinition contains a list of EcsTaskDefinitionList
type EcsTaskDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcsTaskDefinition `json:"items"`
}

// A EcsTaskDefinitionSpec defines the desired state of a EcsTaskDefinition
type EcsTaskDefinitionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EcsTaskDefinitionParameters `json:",inline"`
}

// A EcsTaskDefinitionParameters defines the desired state of a EcsTaskDefinition
type EcsTaskDefinitionParameters struct {
	Tags                    map[string]string      `json:"tags"`
	TaskRoleArn             string                 `json:"task_role_arn"`
	ExecutionRoleArn        string                 `json:"execution_role_arn"`
	Memory                  string                 `json:"memory"`
	ContainerDefinitions    string                 `json:"container_definitions"`
	Cpu                     string                 `json:"cpu"`
	Family                  string                 `json:"family"`
	IpcMode                 string                 `json:"ipc_mode"`
	PidMode                 string                 `json:"pid_mode"`
	RequiresCompatibilities []string               `json:"requires_compatibilities"`
	InferenceAccelerator    []InferenceAccelerator `json:"inference_accelerator"`
	PlacementConstraints    []PlacementConstraints `json:"placement_constraints"`
	ProxyConfiguration      ProxyConfiguration     `json:"proxy_configuration"`
	Volume                  []Volume               `json:"volume"`
}

type InferenceAccelerator struct {
	DeviceName string `json:"device_name"`
	DeviceType string `json:"device_type"`
}

type PlacementConstraints struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

type ProxyConfiguration struct {
	ContainerName string            `json:"container_name"`
	Properties    map[string]string `json:"properties"`
	Type          string            `json:"type"`
}

type Volume struct {
	HostPath                  string                    `json:"host_path"`
	Name                      string                    `json:"name"`
	DockerVolumeConfiguration DockerVolumeConfiguration `json:"docker_volume_configuration"`
	EfsVolumeConfiguration    EfsVolumeConfiguration    `json:"efs_volume_configuration"`
}

type DockerVolumeConfiguration struct {
	Autoprovision bool              `json:"autoprovision"`
	Driver        string            `json:"driver"`
	DriverOpts    map[string]string `json:"driver_opts"`
	Labels        map[string]string `json:"labels"`
	Scope         string            `json:"scope"`
}

type EfsVolumeConfiguration struct {
	FileSystemId          string              `json:"file_system_id"`
	RootDirectory         string              `json:"root_directory"`
	TransitEncryption     string              `json:"transit_encryption"`
	TransitEncryptionPort int                 `json:"transit_encryption_port"`
	AuthorizationConfig   AuthorizationConfig `json:"authorization_config"`
}

type AuthorizationConfig struct {
	AccessPointId string `json:"access_point_id"`
	Iam           string `json:"iam"`
}

// A EcsTaskDefinitionStatus defines the observed state of a EcsTaskDefinition
type EcsTaskDefinitionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EcsTaskDefinitionObservation `json:",inline"`
}

// A EcsTaskDefinitionObservation records the observed state of a EcsTaskDefinition
type EcsTaskDefinitionObservation struct {
	Arn         string `json:"arn"`
	NetworkMode string `json:"network_mode"`
	Revision    int    `json:"revision"`
	Id          string `json:"id"`
}