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

// SagemakerModel is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SagemakerModel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SagemakerModelSpec   `json:"spec"`
	Status SagemakerModelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerModel contains a list of SagemakerModelList
type SagemakerModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerModel `json:"items"`
}

// A SagemakerModelSpec defines the desired state of a SagemakerModel
type SagemakerModelSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SagemakerModelParameters `json:",inline"`
}

// A SagemakerModelParameters defines the desired state of a SagemakerModel
type SagemakerModelParameters struct {
	EnableNetworkIsolation bool              `json:"enable_network_isolation"`
	ExecutionRoleArn       string            `json:"execution_role_arn"`
	Id                     string            `json:"id"`
	Name                   string            `json:"name"`
	Tags                   map[string]string `json:"tags"`
	Container              Container         `json:"container"`
	PrimaryContainer       PrimaryContainer  `json:"primary_container"`
	VpcConfig              VpcConfig         `json:"vpc_config"`
}

type Container struct {
	ModelDataUrl      string            `json:"model_data_url"`
	ContainerHostname string            `json:"container_hostname"`
	Environment       map[string]string `json:"environment"`
	Image             string            `json:"image"`
}

type PrimaryContainer struct {
	Image             string            `json:"image"`
	ModelDataUrl      string            `json:"model_data_url"`
	ContainerHostname string            `json:"container_hostname"`
	Environment       map[string]string `json:"environment"`
}

type VpcConfig struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	Subnets          []string `json:"subnets"`
}

// A SagemakerModelStatus defines the observed state of a SagemakerModel
type SagemakerModelStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SagemakerModelObservation `json:",inline"`
}

// A SagemakerModelObservation records the observed state of a SagemakerModel
type SagemakerModelObservation struct {
	Arn string `json:"arn"`
}