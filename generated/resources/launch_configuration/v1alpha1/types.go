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
	ForProvider                  LaunchConfigurationParameters `json:"forProvider"`
}

// A LaunchConfigurationParameters defines the desired state of a LaunchConfiguration
type LaunchConfigurationParameters struct {
	AssociatePublicIpAddress     bool                 `json:"associate_public_ip_address"`
	KeyName                      string               `json:"key_name"`
	SecurityGroups               []string             `json:"security_groups"`
	SpotPrice                    string               `json:"spot_price"`
	UserData                     string               `json:"user_data"`
	EbsOptimized                 bool                 `json:"ebs_optimized"`
	EnableMonitoring             bool                 `json:"enable_monitoring"`
	InstanceType                 string               `json:"instance_type"`
	UserDataBase64               string               `json:"user_data_base64"`
	VpcClassicLinkId             string               `json:"vpc_classic_link_id"`
	VpcClassicLinkSecurityGroups []string             `json:"vpc_classic_link_security_groups"`
	PlacementTenancy             string               `json:"placement_tenancy"`
	IamInstanceProfile           string               `json:"iam_instance_profile"`
	Id                           string               `json:"id"`
	Name                         string               `json:"name"`
	NamePrefix                   string               `json:"name_prefix"`
	ImageId                      string               `json:"image_id"`
	EphemeralBlockDevice         EphemeralBlockDevice `json:"ephemeral_block_device"`
	RootBlockDevice              RootBlockDevice      `json:"root_block_device"`
	EbsBlockDevice               EbsBlockDevice       `json:"ebs_block_device"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type RootBlockDevice struct {
	Iops                int64  `json:"iops"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Encrypted           bool   `json:"encrypted"`
}

type EbsBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
	NoDevice            bool   `json:"no_device"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
}

// A LaunchConfigurationStatus defines the observed state of a LaunchConfiguration
type LaunchConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LaunchConfigurationObservation `json:"atProvider"`
}

// A LaunchConfigurationObservation records the observed state of a LaunchConfiguration
type LaunchConfigurationObservation struct {
	Arn string `json:"arn"`
}