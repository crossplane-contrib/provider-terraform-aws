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

// LaunchTemplate is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LaunchTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LaunchTemplateSpec   `json:"spec"`
	Status LaunchTemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchTemplate contains a list of LaunchTemplateList
type LaunchTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchTemplate `json:"items"`
}

// A LaunchTemplateSpec defines the desired state of a LaunchTemplate
type LaunchTemplateSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LaunchTemplateParameters `json:",inline"`
}

// A LaunchTemplateParameters defines the desired state of a LaunchTemplate
type LaunchTemplateParameters struct {
	InstanceType                      string   `json:"instance_type"`
	NamePrefix                        string   `json:"name_prefix"`
	DisableApiTermination             bool     `json:"disable_api_termination"`
	ImageId                           string   `json:"image_id"`
	KernelId                          string   `json:"kernel_id"`
	RamDiskId                         string   `json:"ram_disk_id"`
	SecurityGroupNames                []string `json:"security_group_names"`
	Description                       string   `json:"description"`
	UserData                          string   `json:"user_data"`
	VpcSecurityGroupIds               []string `json:"vpc_security_group_ids"`
	InstanceInitiatedShutdownBehavior string   `json:"instance_initiated_shutdown_behavior"`
	UpdateDefaultVersion              bool     `json:"update_default_version"`
	EbsOptimized                      string   `json:"ebs_optimized"`
	KeyName                           string   `json:"key_name"`
}

// A LaunchTemplateStatus defines the observed state of a LaunchTemplate
type LaunchTemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LaunchTemplateObservation `json:",inline"`
}

// A LaunchTemplateObservation records the observed state of a LaunchTemplate
type LaunchTemplateObservation struct {
	Arn            string `json:"arn"`
	DefaultVersion int    `json:"default_version"`
	Id             string `json:"id"`
	Name           string `json:"name"`
	LatestVersion  int    `json:"latest_version"`
}