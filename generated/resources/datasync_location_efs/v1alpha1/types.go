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

// DatasyncLocationEfs is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatasyncLocationEfs struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasyncLocationEfsSpec   `json:"spec"`
	Status DatasyncLocationEfsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncLocationEfs contains a list of DatasyncLocationEfsList
type DatasyncLocationEfsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncLocationEfs `json:"items"`
}

// A DatasyncLocationEfsSpec defines the desired state of a DatasyncLocationEfs
type DatasyncLocationEfsSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatasyncLocationEfsParameters `json:",inline"`
}

// A DatasyncLocationEfsParameters defines the desired state of a DatasyncLocationEfs
type DatasyncLocationEfsParameters struct {
	Id               string            `json:"id"`
	Subdirectory     string            `json:"subdirectory"`
	Tags             map[string]string `json:"tags"`
	EfsFileSystemArn string            `json:"efs_file_system_arn"`
	Ec2Config        Ec2Config         `json:"ec2_config"`
}

type Ec2Config struct {
	SecurityGroupArns []string `json:"security_group_arns"`
	SubnetArn         string   `json:"subnet_arn"`
}

// A DatasyncLocationEfsStatus defines the observed state of a DatasyncLocationEfs
type DatasyncLocationEfsStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatasyncLocationEfsObservation `json:",inline"`
}

// A DatasyncLocationEfsObservation records the observed state of a DatasyncLocationEfs
type DatasyncLocationEfsObservation struct {
	Uri string `json:"uri"`
	Arn string `json:"arn"`
}