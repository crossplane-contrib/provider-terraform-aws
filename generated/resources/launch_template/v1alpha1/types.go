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
	Name                              string                           `json:"name"`
	UserData                          string                           `json:"user_data"`
	DefaultVersion                    int64                            `json:"default_version"`
	InstanceType                      string                           `json:"instance_type"`
	KernelId                          string                           `json:"kernel_id"`
	RamDiskId                         string                           `json:"ram_disk_id"`
	SecurityGroupNames                []string                         `json:"security_group_names"`
	Tags                              map[string]string                `json:"tags"`
	ImageId                           string                           `json:"image_id"`
	KeyName                           string                           `json:"key_name"`
	EbsOptimized                      string                           `json:"ebs_optimized"`
	DisableApiTermination             bool                             `json:"disable_api_termination"`
	Id                                string                           `json:"id"`
	InstanceInitiatedShutdownBehavior string                           `json:"instance_initiated_shutdown_behavior"`
	NamePrefix                        string                           `json:"name_prefix"`
	UpdateDefaultVersion              bool                             `json:"update_default_version"`
	VpcSecurityGroupIds               []string                         `json:"vpc_security_group_ids"`
	Description                       string                           `json:"description"`
	ElasticGpuSpecifications          ElasticGpuSpecifications         `json:"elastic_gpu_specifications"`
	HibernationOptions                HibernationOptions               `json:"hibernation_options"`
	MetadataOptions                   MetadataOptions                  `json:"metadata_options"`
	Placement                         Placement                        `json:"placement"`
	BlockDeviceMappings               BlockDeviceMappings              `json:"block_device_mappings"`
	ElasticInferenceAccelerator       ElasticInferenceAccelerator      `json:"elastic_inference_accelerator"`
	LicenseSpecification              LicenseSpecification             `json:"license_specification"`
	Monitoring                        Monitoring                       `json:"monitoring"`
	CpuOptions                        CpuOptions                       `json:"cpu_options"`
	CreditSpecification               CreditSpecification              `json:"credit_specification"`
	IamInstanceProfile                IamInstanceProfile               `json:"iam_instance_profile"`
	TagSpecifications                 TagSpecifications                `json:"tag_specifications"`
	CapacityReservationSpecification  CapacityReservationSpecification `json:"capacity_reservation_specification"`
	InstanceMarketOptions             InstanceMarketOptions            `json:"instance_market_options"`
	NetworkInterfaces                 NetworkInterfaces                `json:"network_interfaces"`
}

type ElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type HibernationOptions struct {
	Configured bool `json:"configured"`
}

type MetadataOptions struct {
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int64  `json:"http_put_response_hop_limit"`
	HttpTokens              string `json:"http_tokens"`
}

type Placement struct {
	PartitionNumber  int64  `json:"partition_number"`
	SpreadDomain     string `json:"spread_domain"`
	Tenancy          string `json:"tenancy"`
	Affinity         string `json:"affinity"`
	AvailabilityZone string `json:"availability_zone"`
	GroupName        string `json:"group_name"`
	HostId           string `json:"host_id"`
}

type BlockDeviceMappings struct {
	DeviceName  string `json:"device_name"`
	NoDevice    string `json:"no_device"`
	VirtualName string `json:"virtual_name"`
	Ebs         Ebs    `json:"ebs"`
}

type Ebs struct {
	DeleteOnTermination string `json:"delete_on_termination"`
	Encrypted           string `json:"encrypted"`
	Iops                int64  `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type ElasticInferenceAccelerator struct {
	Type string `json:"type"`
}

type LicenseSpecification struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}

type Monitoring struct {
	Enabled bool `json:"enabled"`
}

type CpuOptions struct {
	CoreCount      int64 `json:"core_count"`
	ThreadsPerCore int64 `json:"threads_per_core"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type IamInstanceProfile struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
}

type TagSpecifications struct {
	Tags         map[string]string `json:"tags"`
	ResourceType string            `json:"resource_type"`
}

type CapacityReservationSpecification struct {
	CapacityReservationPreference string                    `json:"capacity_reservation_preference"`
	CapacityReservationTarget     CapacityReservationTarget `json:"capacity_reservation_target"`
}

type CapacityReservationTarget struct {
	CapacityReservationId string `json:"capacity_reservation_id"`
}

type InstanceMarketOptions struct {
	MarketType  string      `json:"market_type"`
	SpotOptions SpotOptions `json:"spot_options"`
}

type SpotOptions struct {
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
	MaxPrice                     string `json:"max_price"`
	SpotInstanceType             string `json:"spot_instance_type"`
	ValidUntil                   string `json:"valid_until"`
	BlockDurationMinutes         int64  `json:"block_duration_minutes"`
}

type NetworkInterfaces struct {
	Ipv4AddressCount         int64    `json:"ipv4_address_count"`
	NetworkInterfaceId       string   `json:"network_interface_id"`
	SecurityGroups           []string `json:"security_groups"`
	Description              string   `json:"description"`
	DeviceIndex              int64    `json:"device_index"`
	Ipv4Addresses            []string `json:"ipv4_addresses"`
	Ipv6AddressCount         int64    `json:"ipv6_address_count"`
	Ipv6Addresses            []string `json:"ipv6_addresses"`
	PrivateIpAddress         string   `json:"private_ip_address"`
	SubnetId                 string   `json:"subnet_id"`
	AssociatePublicIpAddress string   `json:"associate_public_ip_address"`
	DeleteOnTermination      string   `json:"delete_on_termination"`
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