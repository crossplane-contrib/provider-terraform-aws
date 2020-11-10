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

// Instance is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceSpec   `json:"spec"`
	Status InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Instance contains a list of InstanceList
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// A InstanceSpec defines the desired state of a Instance
type InstanceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  InstanceParameters `json:",inline"`
}

// A InstanceParameters defines the desired state of a Instance
type InstanceParameters struct {
	DisableApiTermination             bool                   `json:"disable_api_termination"`
	GetPasswordData                   bool                   `json:"get_password_data"`
	Monitoring                        bool                   `json:"monitoring"`
	InstanceInitiatedShutdownBehavior string                 `json:"instance_initiated_shutdown_behavior"`
	Ami                               string                 `json:"ami"`
	Hibernation                       bool                   `json:"hibernation"`
	UserData                          string                 `json:"user_data"`
	IamInstanceProfile                string                 `json:"iam_instance_profile"`
	UserDataBase64                    string                 `json:"user_data_base64"`
	InstanceType                      string                 `json:"instance_type"`
	EbsOptimized                      bool                   `json:"ebs_optimized"`
	SourceDestCheck                   bool                   `json:"source_dest_check"`
	Tags                              map[string]string      `json:"tags"`
	NetworkInterface                  []NetworkInterface     `json:"network_interface"`
	RootBlockDevice                   RootBlockDevice        `json:"root_block_device"`
	Timeouts                          []Timeouts             `json:"timeouts"`
	CreditSpecification               CreditSpecification    `json:"credit_specification"`
	EbsBlockDevice                    []EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice              []EphemeralBlockDevice `json:"ephemeral_block_device"`
	MetadataOptions                   MetadataOptions        `json:"metadata_options"`
}

type NetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceIndex         int    `json:"device_index"`
	NetworkInterfaceId  string `json:"network_interface_id"`
}

type RootBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	VolumeId            string `json:"volume_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type EbsBlockDevice struct {
	DeviceName          string `json:"device_name"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	KmsKeyId            string `json:"kms_key_id"`
	VolumeId            string `json:"volume_id"`
	VolumeSize          int    `json:"volume_size"`
	Encrypted           bool   `json:"encrypted"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	NoDevice    bool   `json:"no_device"`
	VirtualName string `json:"virtual_name"`
}

type MetadataOptions struct {
	HttpTokens              string `json:"http_tokens"`
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int    `json:"http_put_response_hop_limit"`
}

// A InstanceStatus defines the observed state of a Instance
type InstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     InstanceObservation `json:",inline"`
}

// A InstanceObservation records the observed state of a Instance
type InstanceObservation struct {
	OutpostArn                string            `json:"outpost_arn"`
	VpcSecurityGroupIds       []string          `json:"vpc_security_group_ids"`
	AvailabilityZone          string            `json:"availability_zone"`
	HostId                    string            `json:"host_id"`
	Id                        string            `json:"id"`
	PlacementGroup            string            `json:"placement_group"`
	PrivateDns                string            `json:"private_dns"`
	PublicIp                  string            `json:"public_ip"`
	SecurityGroups            []string          `json:"security_groups"`
	Arn                       string            `json:"arn"`
	CpuThreadsPerCore         int               `json:"cpu_threads_per_core"`
	PublicDns                 string            `json:"public_dns"`
	Tenancy                   string            `json:"tenancy"`
	PrivateIp                 string            `json:"private_ip"`
	PasswordData              string            `json:"password_data"`
	PrimaryNetworkInterfaceId string            `json:"primary_network_interface_id"`
	KeyName                   string            `json:"key_name"`
	InstanceState             string            `json:"instance_state"`
	Ipv6AddressCount          int               `json:"ipv6_address_count"`
	SecondaryPrivateIps       []string          `json:"secondary_private_ips"`
	SubnetId                  string            `json:"subnet_id"`
	AssociatePublicIpAddress  bool              `json:"associate_public_ip_address"`
	CpuCoreCount              int               `json:"cpu_core_count"`
	VolumeTags                map[string]string `json:"volume_tags"`
	Ipv6Addresses             []string          `json:"ipv6_addresses"`
}