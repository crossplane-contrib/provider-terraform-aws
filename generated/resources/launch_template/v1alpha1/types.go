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
	ForProvider                  LaunchTemplateParameters `json:"forProvider"`
}

// A LaunchTemplateParameters defines the desired state of a LaunchTemplate
type LaunchTemplateParameters struct {
	InstanceType                      string                           `json:"instance_type"`
	NamePrefix                        string                           `json:"name_prefix"`
	RamDiskId                         string                           `json:"ram_disk_id"`
	Tags                              map[string]string                `json:"tags"`
	VpcSecurityGroupIds               []string                         `json:"vpc_security_group_ids"`
	Description                       string                           `json:"description"`
	ImageId                           string                           `json:"image_id"`
	Id                                string                           `json:"id"`
	InstanceInitiatedShutdownBehavior string                           `json:"instance_initiated_shutdown_behavior"`
	KernelId                          string                           `json:"kernel_id"`
	Name                              string                           `json:"name"`
	SecurityGroupNames                []string                         `json:"security_group_names"`
	EbsOptimized                      string                           `json:"ebs_optimized"`
	DefaultVersion                    int64                            `json:"default_version"`
	UpdateDefaultVersion              bool                             `json:"update_default_version"`
	UserData                          string                           `json:"user_data"`
	DisableApiTermination             bool                             `json:"disable_api_termination"`
	KeyName                           string                           `json:"key_name"`
	CpuOptions                        CpuOptions                       `json:"cpu_options"`
	ElasticInferenceAccelerator       ElasticInferenceAccelerator      `json:"elastic_inference_accelerator"`
	MetadataOptions                   MetadataOptions                  `json:"metadata_options"`
	BlockDeviceMappings               BlockDeviceMappings              `json:"block_device_mappings"`
	CapacityReservationSpecification  CapacityReservationSpecification `json:"capacity_reservation_specification"`
	ElasticGpuSpecifications          ElasticGpuSpecifications         `json:"elastic_gpu_specifications"`
	IamInstanceProfile                IamInstanceProfile               `json:"iam_instance_profile"`
	InstanceMarketOptions             InstanceMarketOptions            `json:"instance_market_options"`
	LicenseSpecification              LicenseSpecification             `json:"license_specification"`
	HibernationOptions                HibernationOptions               `json:"hibernation_options"`
	Monitoring                        Monitoring                       `json:"monitoring"`
	Placement                         Placement                        `json:"placement"`
	TagSpecifications                 TagSpecifications                `json:"tag_specifications"`
	CreditSpecification               CreditSpecification              `json:"credit_specification"`
	NetworkInterfaces                 NetworkInterfaces                `json:"network_interfaces"`
}

type CpuOptions struct {
	CoreCount      int64 `json:"core_count"`
	ThreadsPerCore int64 `json:"threads_per_core"`
}

type ElasticInferenceAccelerator struct {
	Type string `json:"type"`
}

type MetadataOptions struct {
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int64  `json:"http_put_response_hop_limit"`
	HttpTokens              string `json:"http_tokens"`
}

type BlockDeviceMappings struct {
	DeviceName  string `json:"device_name"`
	NoDevice    string `json:"no_device"`
	VirtualName string `json:"virtual_name"`
	Ebs         Ebs    `json:"ebs"`
}

type Ebs struct {
	Encrypted           string `json:"encrypted"`
	Iops                int64  `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination string `json:"delete_on_termination"`
}

type CapacityReservationSpecification struct {
	CapacityReservationPreference string                    `json:"capacity_reservation_preference"`
	CapacityReservationTarget     CapacityReservationTarget `json:"capacity_reservation_target"`
}

type CapacityReservationTarget struct {
	CapacityReservationId string `json:"capacity_reservation_id"`
}

type ElasticGpuSpecifications struct {
	Type string `json:"type"`
}

type IamInstanceProfile struct {
	Arn  string `json:"arn"`
	Name string `json:"name"`
}

type InstanceMarketOptions struct {
	MarketType  string      `json:"market_type"`
	SpotOptions SpotOptions `json:"spot_options"`
}

type SpotOptions struct {
	MaxPrice                     string `json:"max_price"`
	SpotInstanceType             string `json:"spot_instance_type"`
	ValidUntil                   string `json:"valid_until"`
	BlockDurationMinutes         int64  `json:"block_duration_minutes"`
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
}

type LicenseSpecification struct {
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}

type HibernationOptions struct {
	Configured bool `json:"configured"`
}

type Monitoring struct {
	Enabled bool `json:"enabled"`
}

type Placement struct {
	Affinity         string `json:"affinity"`
	AvailabilityZone string `json:"availability_zone"`
	GroupName        string `json:"group_name"`
	HostId           string `json:"host_id"`
	PartitionNumber  int64  `json:"partition_number"`
	SpreadDomain     string `json:"spread_domain"`
	Tenancy          string `json:"tenancy"`
}

type TagSpecifications struct {
	ResourceType string            `json:"resource_type"`
	Tags         map[string]string `json:"tags"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type NetworkInterfaces struct {
	DeleteOnTermination      string   `json:"delete_on_termination"`
	Description              string   `json:"description"`
	DeviceIndex              int64    `json:"device_index"`
	Ipv4AddressCount         int64    `json:"ipv4_address_count"`
	Ipv4Addresses            []string `json:"ipv4_addresses"`
	SecurityGroups           []string `json:"security_groups"`
	AssociatePublicIpAddress string   `json:"associate_public_ip_address"`
	Ipv6AddressCount         int64    `json:"ipv6_address_count"`
	Ipv6Addresses            []string `json:"ipv6_addresses"`
	NetworkInterfaceId       string   `json:"network_interface_id"`
	PrivateIpAddress         string   `json:"private_ip_address"`
	SubnetId                 string   `json:"subnet_id"`
}

// A LaunchTemplateStatus defines the observed state of a LaunchTemplate
type LaunchTemplateStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LaunchTemplateObservation `json:"atProvider"`
}

// A LaunchTemplateObservation records the observed state of a LaunchTemplate
type LaunchTemplateObservation struct {
	Arn           string `json:"arn"`
	LatestVersion int64  `json:"latest_version"`
}