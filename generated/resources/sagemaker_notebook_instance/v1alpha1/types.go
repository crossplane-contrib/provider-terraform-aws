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

// SagemakerNotebookInstance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SagemakerNotebookInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SagemakerNotebookInstanceSpec   `json:"spec"`
	Status SagemakerNotebookInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerNotebookInstance contains a list of SagemakerNotebookInstanceList
type SagemakerNotebookInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerNotebookInstance `json:"items"`
}

// A SagemakerNotebookInstanceSpec defines the desired state of a SagemakerNotebookInstance
type SagemakerNotebookInstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SagemakerNotebookInstanceParameters `json:",inline"`
}

// A SagemakerNotebookInstanceParameters defines the desired state of a SagemakerNotebookInstance
type SagemakerNotebookInstanceParameters struct {
	Id                   string            `json:"id"`
	InstanceType         string            `json:"instance_type"`
	KmsKeyId             string            `json:"kms_key_id"`
	Name                 string            `json:"name"`
	RoleArn              string            `json:"role_arn"`
	RootAccess           string            `json:"root_access"`
	SecurityGroups       []string          `json:"security_groups"`
	Tags                 map[string]string `json:"tags"`
	SubnetId             string            `json:"subnet_id"`
	LifecycleConfigName  string            `json:"lifecycle_config_name"`
	DirectInternetAccess string            `json:"direct_internet_access"`
}

// A SagemakerNotebookInstanceStatus defines the observed state of a SagemakerNotebookInstance
type SagemakerNotebookInstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SagemakerNotebookInstanceObservation `json:",inline"`
}

// A SagemakerNotebookInstanceObservation records the observed state of a SagemakerNotebookInstance
type SagemakerNotebookInstanceObservation struct {
	Arn string `json:"arn"`
}