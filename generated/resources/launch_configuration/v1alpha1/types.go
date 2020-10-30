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

// LaunchConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LaunchConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LaunchConfigurationSpec   `json:"spec"`
	Status LaunchConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LaunchConfiguration contains a list of LaunchConfigurationList
type LaunchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LaunchConfiguration `json:"items"`
}

// A LaunchConfigurationSpec defines the desired state of a LaunchConfiguration
type LaunchConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LaunchConfigurationParameters `json:",inline"`
}

// A LaunchConfigurationParameters defines the desired state of a LaunchConfiguration
type LaunchConfigurationParameters struct {
	VpcClassicLinkId             string   `json:"vpc_classic_link_id"`
	IamInstanceProfile           string   `json:"iam_instance_profile"`
	SecurityGroups               []string `json:"security_groups"`
	SpotPrice                    string   `json:"spot_price"`
	NamePrefix                   string   `json:"name_prefix"`
	UserDataBase64               string   `json:"user_data_base64"`
	VpcClassicLinkSecurityGroups []string `json:"vpc_classic_link_security_groups"`
	AssociatePublicIpAddress     bool     `json:"associate_public_ip_address"`
	EnableMonitoring             bool     `json:"enable_monitoring"`
	InstanceType                 string   `json:"instance_type"`
	ImageId                      string   `json:"image_id"`
	PlacementTenancy             string   `json:"placement_tenancy"`
	UserData                     string   `json:"user_data"`
}

// A LaunchConfigurationStatus defines the observed state of a LaunchConfiguration
type LaunchConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LaunchConfigurationObservation `json:",inline"`
}

// A LaunchConfigurationObservation records the observed state of a LaunchConfiguration
type LaunchConfigurationObservation struct {
	Arn          string `json:"arn"`
	Id           string `json:"id"`
	EbsOptimized bool   `json:"ebs_optimized"`
	KeyName      string `json:"key_name"`
	Name         string `json:"name"`
}