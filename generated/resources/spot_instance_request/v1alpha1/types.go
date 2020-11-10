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

// SpotInstanceRequest is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SpotInstanceRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotInstanceRequestSpec   `json:"spec"`
	Status SpotInstanceRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotInstanceRequest contains a list of SpotInstanceRequestList
type SpotInstanceRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotInstanceRequest `json:"items"`
}

// A SpotInstanceRequestSpec defines the desired state of a SpotInstanceRequest
type SpotInstanceRequestSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SpotInstanceRequestParameters `json:",inline"`
}

// A SpotInstanceRequestParameters defines the desired state of a SpotInstanceRequest
type SpotInstanceRequestParameters struct {
	GetPasswordData                   bool                   `json:"get_password_data"`
	Monitoring                        bool                   `json:"monitoring"`
	UserData                          string                 `json:"user_data"`
	BlockDurationMinutes              int                    `json:"block_duration_minutes"`
	EbsOptimized                      bool                   `json:"ebs_optimized"`
	IamInstanceProfile                string                 `json:"iam_instance_profile"`
	InstanceInitiatedShutdownBehavior string                 `json:"instance_initiated_shutdown_behavior"`
	SourceDestCheck                   bool                   `json:"source_dest_check"`
	SpotPrice                         string                 `json:"spot_price"`
	WaitForFulfillment                bool                   `json:"wait_for_fulfillment"`
	UserDataBase64                    string                 `json:"user_data_base64"`
	Hibernation                       bool                   `json:"hibernation"`
	InstanceInterruptionBehaviour     string                 `json:"instance_interruption_behaviour"`
	DisableApiTermination             bool                   `json:"disable_api_termination"`
	InstanceType                      string                 `json:"instance_type"`
	LaunchGroup                       string                 `json:"launch_group"`
	Ami                               string                 `json:"ami"`
	SpotType                          string                 `json:"spot_type"`
	Tags                              map[string]string      `json:"tags"`
	VolumeTags                        map[string]string      `json:"volume_tags"`
	RootBlockDevice                   RootBlockDevice        `json:"root_block_device"`
	Timeouts                          []Timeouts             `json:"timeouts"`
	CreditSpecification               CreditSpecification    `json:"credit_specification"`
	EbsBlockDevice                    []EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice              []EphemeralBlockDevice `json:"ephemeral_block_device"`
	MetadataOptions                   MetadataOptions        `json:"metadata_options"`
	NetworkInterface                  []NetworkInterface     `json:"network_interface"`
}

type RootBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	VolumeId            string `json:"volume_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type EbsBlockDevice struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	Iops                int    `json:"iops"`
	KmsKeyId            string `json:"kms_key_id"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeId            string `json:"volume_id"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	NoDevice    bool   `json:"no_device"`
	VirtualName string `json:"virtual_name"`
}

type MetadataOptions struct {
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int    `json:"http_put_response_hop_limit"`
	HttpTokens              string `json:"http_tokens"`
}

type NetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceIndex         int    `json:"device_index"`
	NetworkInterfaceId  string `json:"network_interface_id"`
}

// A SpotInstanceRequestStatus defines the observed state of a SpotInstanceRequest
type SpotInstanceRequestStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SpotInstanceRequestObservation `json:",inline"`
}

// A SpotInstanceRequestObservation records the observed state of a SpotInstanceRequest
type SpotInstanceRequestObservation struct {
	Arn                       string   `json:"arn"`
	PlacementGroup            string   `json:"placement_group"`
	SecurityGroups            []string `json:"security_groups"`
	ValidUntil                string   `json:"valid_until"`
	AssociatePublicIpAddress  bool     `json:"associate_public_ip_address"`
	HostId                    string   `json:"host_id"`
	PasswordData              string   `json:"password_data"`
	CpuCoreCount              int      `json:"cpu_core_count"`
	InstanceState             string   `json:"instance_state"`
	Ipv6AddressCount          int      `json:"ipv6_address_count"`
	PublicIp                  string   `json:"public_ip"`
	Tenancy                   string   `json:"tenancy"`
	AvailabilityZone          string   `json:"availability_zone"`
	KeyName                   string   `json:"key_name"`
	PrivateIp                 string   `json:"private_ip"`
	SpotBidStatus             string   `json:"spot_bid_status"`
	VpcSecurityGroupIds       []string `json:"vpc_security_group_ids"`
	CpuThreadsPerCore         int      `json:"cpu_threads_per_core"`
	Id                        string   `json:"id"`
	SpotRequestState          string   `json:"spot_request_state"`
	SubnetId                  string   `json:"subnet_id"`
	PrimaryNetworkInterfaceId string   `json:"primary_network_interface_id"`
	PrivateDns                string   `json:"private_dns"`
	ValidFrom                 string   `json:"valid_from"`
	Ipv6Addresses             []string `json:"ipv6_addresses"`
	OutpostArn                string   `json:"outpost_arn"`
	PublicDns                 string   `json:"public_dns"`
	SecondaryPrivateIps       []string `json:"secondary_private_ips"`
	SpotInstanceId            string   `json:"spot_instance_id"`
}