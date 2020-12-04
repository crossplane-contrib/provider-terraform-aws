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
	Ami                               string               `json:"ami"`
	EbsOptimized                      bool                 `json:"ebs_optimized"`
	Tenancy                           string               `json:"tenancy"`
	DisableApiTermination             bool                 `json:"disable_api_termination"`
	IamInstanceProfile                string               `json:"iam_instance_profile"`
	KeyName                           string               `json:"key_name"`
	SecondaryPrivateIps               []string             `json:"secondary_private_ips"`
	UserData                          string               `json:"user_data"`
	VpcSecurityGroupIds               []string             `json:"vpc_security_group_ids"`
	HostId                            string               `json:"host_id"`
	Ipv6Addresses                     []string             `json:"ipv6_addresses"`
	SubnetId                          string               `json:"subnet_id"`
	GetPasswordData                   bool                 `json:"get_password_data"`
	InstanceType                      string               `json:"instance_type"`
	PlacementGroup                    string               `json:"placement_group"`
	SourceDestCheck                   bool                 `json:"source_dest_check"`
	Hibernation                       bool                 `json:"hibernation"`
	SecurityGroups                    []string             `json:"security_groups"`
	Tags                              map[string]string    `json:"tags"`
	UserDataBase64                    string               `json:"user_data_base64"`
	Ipv6AddressCount                  int64                `json:"ipv6_address_count"`
	VolumeTags                        map[string]string    `json:"volume_tags"`
	AssociatePublicIpAddress          bool                 `json:"associate_public_ip_address"`
	AvailabilityZone                  string               `json:"availability_zone"`
	CpuCoreCount                      int64                `json:"cpu_core_count"`
	CpuThreadsPerCore                 int64                `json:"cpu_threads_per_core"`
	Id                                string               `json:"id"`
	InstanceInitiatedShutdownBehavior string               `json:"instance_initiated_shutdown_behavior"`
	Monitoring                        bool                 `json:"monitoring"`
	PrivateIp                         string               `json:"private_ip"`
	RootBlockDevice                   RootBlockDevice      `json:"root_block_device"`
	Timeouts                          Timeouts             `json:"timeouts"`
	CreditSpecification               CreditSpecification  `json:"credit_specification"`
	EbsBlockDevice                    EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice              EphemeralBlockDevice `json:"ephemeral_block_device"`
	MetadataOptions                   MetadataOptions      `json:"metadata_options"`
	NetworkInterface                  NetworkInterface     `json:"network_interface"`
}

type RootBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	VolumeId            string `json:"volume_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Update string `json:"update"`
	Create string `json:"create"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type EbsBlockDevice struct {
	Iops                int64  `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeId            string `json:"volume_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	Encrypted           bool   `json:"encrypted"`
	DeviceName          string `json:"device_name"`
	KmsKeyId            string `json:"kms_key_id"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
}

type EphemeralBlockDevice struct {
	VirtualName string `json:"virtual_name"`
	DeviceName  string `json:"device_name"`
	NoDevice    bool   `json:"no_device"`
}

type MetadataOptions struct {
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int64  `json:"http_put_response_hop_limit"`
	HttpTokens              string `json:"http_tokens"`
}

type NetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceIndex         int64  `json:"device_index"`
	NetworkInterfaceId  string `json:"network_interface_id"`
}

// A InstanceStatus defines the observed state of a Instance
type InstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     InstanceObservation `json:",inline"`
}

// A InstanceObservation records the observed state of a Instance
type InstanceObservation struct {
	PrimaryNetworkInterfaceId string `json:"primary_network_interface_id"`
	InstanceState             string `json:"instance_state"`
	PublicIp                  string `json:"public_ip"`
	OutpostArn                string `json:"outpost_arn"`
	PrivateDns                string `json:"private_dns"`
	PasswordData              string `json:"password_data"`
	Arn                       string `json:"arn"`
	PublicDns                 string `json:"public_dns"`
}