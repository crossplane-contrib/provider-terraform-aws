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
	Tags                              map[string]string                `json:"tags"`
	UpdateDefaultVersion              bool                             `json:"update_default_version"`
	UserData                          string                           `json:"user_data"`
	VpcSecurityGroupIds               []string                         `json:"vpc_security_group_ids"`
	DisableApiTermination             bool                             `json:"disable_api_termination"`
	InstanceInitiatedShutdownBehavior string                           `json:"instance_initiated_shutdown_behavior"`
	RamDiskId                         string                           `json:"ram_disk_id"`
	SecurityGroupNames                []string                         `json:"security_group_names"`
	EbsOptimized                      string                           `json:"ebs_optimized"`
	Id                                string                           `json:"id"`
	InstanceType                      string                           `json:"instance_type"`
	KeyName                           string                           `json:"key_name"`
	Name                              string                           `json:"name"`
	DefaultVersion                    int64                            `json:"default_version"`
	Description                       string                           `json:"description"`
	KernelId                          string                           `json:"kernel_id"`
	NamePrefix                        string                           `json:"name_prefix"`
	ImageId                           string                           `json:"image_id"`
	CreditSpecification               CreditSpecification              `json:"credit_specification"`
	ElasticGpuSpecifications          ElasticGpuSpecifications         `json:"elastic_gpu_specifications"`
	ElasticInferenceAccelerator       ElasticInferenceAccelerator      `json:"elastic_inference_accelerator"`
	MetadataOptions                   MetadataOptions                  `json:"metadata_options"`
	Monitoring                        Monitoring                       `json:"monitoring"`
	NetworkInterfaces                 NetworkInterfaces                `json:"network_interfaces"`
	TagSpecifications                 TagSpecifications                `json:"tag_specifications"`
	BlockDeviceMappings               BlockDeviceMappings              `json:"block_device_mappings"`
	HibernationOptions                HibernationOptions               `json:"hibernation_options"`
	InstanceMarketOptions             InstanceMarketOptions            `json:"instance_market_options"`
	CpuOptions                        CpuOptions                       `json:"cpu_options"`
	Placement                         Placement                        `json:"placement"`
	CapacityReservationSpecification  CapacityReservationSpecification `json:"capacity_reservation_specification"`
	IamInstanceProfile                IamInstanceProfile               `json:"iam_instance_profile"`
	LicenseSpecification              LicenseSpecification             `json:"license_specification"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type ElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type ElasticInferenceAccelerator struct {
	Type string `json:"type"`
}

type MetadataOptions struct {
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int64  `json:"http_put_response_hop_limit"`
	HttpTokens              string `json:"http_tokens"`
}

type Monitoring struct {
	Enabled bool `json:"enabled"`
}

type NetworkInterfaces struct {
	SecurityGroups           []string `json:"security_groups"`
	DeleteOnTermination      string   `json:"delete_on_termination"`
	Description              string   `json:"description"`
	DeviceIndex              int64    `json:"device_index"`
	Ipv4AddressCount         int64    `json:"ipv4_address_count"`
	PrivateIpAddress         string   `json:"private_ip_address"`
	SubnetId                 string   `json:"subnet_id"`
	AssociatePublicIpAddress string   `json:"associate_public_ip_address"`
	Ipv4Addresses            []string `json:"ipv4_addresses"`
	Ipv6AddressCount         int64    `json:"ipv6_address_count"`
	Ipv6Addresses            []string `json:"ipv6_addresses"`
	NetworkInterfaceId       string   `json:"network_interface_id"`
}

type TagSpecifications struct {
	ResourceType string            `json:"resource_type"`
	Tags         map[string]string `json:"tags"`
}

type BlockDeviceMappings struct {
	DeviceName  string `json:"device_name"`
	NoDevice    string `json:"no_device"`
	VirtualName string `json:"virtual_name"`
	Ebs         Ebs    `json:"ebs"`
}

type Ebs struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination string `json:"delete_on_termination"`
	Encrypted           string `json:"encrypted"`
	Iops                int64  `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
}

type HibernationOptions struct {
	Configured bool `json:"configured"`
}

type InstanceMarketOptions struct {
	MarketType  string      `json:"market_type"`
	SpotOptions SpotOptions `json:"spot_options"`
}

type SpotOptions struct {
	BlockDurationMinutes         int64  `json:"block_duration_minutes"`
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
	MaxPrice                     string `json:"max_price"`
	SpotInstanceType             string `json:"spot_instance_type"`
	ValidUntil                   string `json:"valid_until"`
}

type CpuOptions struct {
	CoreCount      int64 `json:"core_count"`
	ThreadsPerCore int64 `json:"threads_per_core"`
}

type Placement struct {
	AvailabilityZone string `json:"availability_zone"`
	GroupName        string `json:"group_name"`
	HostId           string `json:"host_id"`
	PartitionNumber  int64  `json:"partition_number"`
	SpreadDomain     string `json:"spread_domain"`
	Tenancy          string `json:"tenancy"`
	Affinity         string `json:"affinity"`
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

type LicenseSpecification struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}

// A LaunchTemplateStatus defines the observed state of a LaunchTemplate
type LaunchTemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LaunchTemplateObservation `json:",inline"`
}

// A LaunchTemplateObservation records the observed state of a LaunchTemplate
type LaunchTemplateObservation struct {
	LatestVersion int64  `json:"latest_version"`
	Arn           string `json:"arn"`
}