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
	Cpu                     string   `json:"cpu"`
	Family                  string   `json:"family"`
	PidMode                 string   `json:"pid_mode"`
	ContainerDefinitions    string   `json:"container_definitions"`
	Memory                  string   `json:"memory"`
	TaskRoleArn             string   `json:"task_role_arn"`
	ExecutionRoleArn        string   `json:"execution_role_arn"`
	IpcMode                 string   `json:"ipc_mode"`
	RequiresCompatibilities []string `json:"requires_compatibilities"`
}

// A EcsTaskDefinitionStatus defines the observed state of a EcsTaskDefinition
type EcsTaskDefinitionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EcsTaskDefinitionObservation `json:",inline"`
}

// A EcsTaskDefinitionObservation records the observed state of a EcsTaskDefinition
type EcsTaskDefinitionObservation struct {
	Revision    int    `json:"revision"`
	Arn         string `json:"arn"`
	Id          string `json:"id"`
	NetworkMode string `json:"network_mode"`
}