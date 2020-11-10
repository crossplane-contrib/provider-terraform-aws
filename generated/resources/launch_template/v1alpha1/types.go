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
	UpdateDefaultVersion              bool                             `json:"update_default_version"`
	VpcSecurityGroupIds               []string                         `json:"vpc_security_group_ids"`
	DisableApiTermination             bool                             `json:"disable_api_termination"`
	KernelId                          string                           `json:"kernel_id"`
	NamePrefix                        string                           `json:"name_prefix"`
	EbsOptimized                      string                           `json:"ebs_optimized"`
	RamDiskId                         string                           `json:"ram_disk_id"`
	Tags                              map[string]string                `json:"tags"`
	InstanceInitiatedShutdownBehavior string                           `json:"instance_initiated_shutdown_behavior"`
	UserData                          string                           `json:"user_data"`
	KeyName                           string                           `json:"key_name"`
	SecurityGroupNames                []string                         `json:"security_group_names"`
	Description                       string                           `json:"description"`
	ImageId                           string                           `json:"image_id"`
	InstanceType                      string                           `json:"instance_type"`
	CapacityReservationSpecification  CapacityReservationSpecification `json:"capacity_reservation_specification"`
	IamInstanceProfile                IamInstanceProfile               `json:"iam_instance_profile"`
	CreditSpecification               CreditSpecification              `json:"credit_specification"`
	InstanceMarketOptions             InstanceMarketOptions            `json:"instance_market_options"`
	LicenseSpecification              []LicenseSpecification           `json:"license_specification"`
	MetadataOptions                   MetadataOptions                  `json:"metadata_options"`
	Monitoring                        Monitoring                       `json:"monitoring"`
	Placement                         Placement                        `json:"placement"`
	BlockDeviceMappings               []BlockDeviceMappings            `json:"block_device_mappings"`
	CpuOptions                        CpuOptions                       `json:"cpu_options"`
	ElasticInferenceAccelerator       ElasticInferenceAccelerator      `json:"elastic_inference_accelerator"`
	HibernationOptions                HibernationOptions               `json:"hibernation_options"`
	TagSpecifications                 []TagSpecifications              `json:"tag_specifications"`
	ElasticGpuSpecifications          []ElasticGpuSpecifications       `json:"elastic_gpu_specifications"`
	NetworkInterfaces                 []NetworkInterfaces              `json:"network_interfaces"`
}

type CapacityReservationSpecification struct {
	CapacityReservationPreference string                    `json:"capacity_reservation_preference"`
	CapacityReservationTarget     CapacityReservationTarget `json:"capacity_reservation_target"`
}

type CapacityReservationTarget struct {
	CapacityReservationId string `json:"capacity_reservation_id"`
}

type IamInstanceProfile struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type InstanceMarketOptions struct {
	MarketType  string      `json:"market_type"`
	SpotOptions SpotOptions `json:"spot_options"`
}

type SpotOptions struct {
	BlockDurationMinutes         int    `json:"block_duration_minutes"`
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
	MaxPrice                     string `json:"max_price"`
	SpotInstanceType             string `json:"spot_instance_type"`
	ValidUntil                   string `json:"valid_until"`
}

type LicenseSpecification struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}

type MetadataOptions struct {
	HttpTokens              string `json:"http_tokens"`
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int    `json:"http_put_response_hop_limit"`
}

type Monitoring struct {
	Enabled bool `json:"enabled"`
}

type Placement struct {
	HostId           string `json:"host_id"`
	PartitionNumber  int    `json:"partition_number"`
	SpreadDomain     string `json:"spread_domain"`
	Tenancy          string `json:"tenancy"`
	Affinity         string `json:"affinity"`
	AvailabilityZone string `json:"availability_zone"`
	GroupName        string `json:"group_name"`
}

type BlockDeviceMappings struct {
	DeviceName  string `json:"device_name"`
	NoDevice    string `json:"no_device"`
	VirtualName string `json:"virtual_name"`
	Ebs         Ebs    `json:"ebs"`
}

type Ebs struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination string `json:"delete_on_termination"`
	Encrypted           string `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
}

type CpuOptions struct {
	CoreCount      int `json:"core_count"`
	ThreadsPerCore int `json:"threads_per_core"`
}

type ElasticInferenceAccelerator struct {
	Type string `json:"type"`
}

type HibernationOptions struct {
	Configured bool `json:"configured"`
}

type TagSpecifications struct {
	ResourceType string            `json:"resource_type"`
	Tags         map[string]string `json:"tags"`
}

type ElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type NetworkInterfaces struct {
	DeviceIndex              int      `json:"device_index"`
	Ipv6AddressCount         int      `json:"ipv6_address_count"`
	NetworkInterfaceId       string   `json:"network_interface_id"`
	PrivateIpAddress         string   `json:"private_ip_address"`
	AssociatePublicIpAddress string   `json:"associate_public_ip_address"`
	DeleteOnTermination      string   `json:"delete_on_termination"`
	Description              string   `json:"description"`
	Ipv4AddressCount         int      `json:"ipv4_address_count"`
	Ipv4Addresses            []string `json:"ipv4_addresses"`
	Ipv6Addresses            []string `json:"ipv6_addresses"`
	SecurityGroups           []string `json:"security_groups"`
	SubnetId                 string   `json:"subnet_id"`
}

// A LaunchTemplateStatus defines the observed state of a LaunchTemplate
type LaunchTemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LaunchTemplateObservation `json:",inline"`
}

// A LaunchTemplateObservation records the observed state of a LaunchTemplate
type LaunchTemplateObservation struct {
	Name           string `json:"name"`
	DefaultVersion int    `json:"default_version"`
	Id             string `json:"id"`
	LatestVersion  int    `json:"latest_version"`
	Arn            string `json:"arn"`
}