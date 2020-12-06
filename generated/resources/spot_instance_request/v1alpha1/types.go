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
	SecondaryPrivateIps               []string             `json:"secondary_private_ips"`
	SpotType                          string               `json:"spot_type"`
	UserDataBase64                    string               `json:"user_data_base64"`
	CpuCoreCount                      int64                `json:"cpu_core_count"`
	GetPasswordData                   bool                 `json:"get_password_data"`
	Ipv6AddressCount                  int64                `json:"ipv6_address_count"`
	VolumeTags                        map[string]string    `json:"volume_tags"`
	Hibernation                       bool                 `json:"hibernation"`
	Tenancy                           string               `json:"tenancy"`
	UserData                          string               `json:"user_data"`
	CpuThreadsPerCore                 int64                `json:"cpu_threads_per_core"`
	EbsOptimized                      bool                 `json:"ebs_optimized"`
	Id                                string               `json:"id"`
	PlacementGroup                    string               `json:"placement_group"`
	ValidUntil                        string               `json:"valid_until"`
	VpcSecurityGroupIds               []string             `json:"vpc_security_group_ids"`
	WaitForFulfillment                bool                 `json:"wait_for_fulfillment"`
	SecurityGroups                    []string             `json:"security_groups"`
	SourceDestCheck                   bool                 `json:"source_dest_check"`
	Ami                               string               `json:"ami"`
	AssociatePublicIpAddress          bool                 `json:"associate_public_ip_address"`
	AvailabilityZone                  string               `json:"availability_zone"`
	Ipv6Addresses                     []string             `json:"ipv6_addresses"`
	DisableApiTermination             bool                 `json:"disable_api_termination"`
	InstanceType                      string               `json:"instance_type"`
	PrivateIp                         string               `json:"private_ip"`
	SpotPrice                         string               `json:"spot_price"`
	InstanceInterruptionBehaviour     string               `json:"instance_interruption_behaviour"`
	Monitoring                        bool                 `json:"monitoring"`
	Tags                              map[string]string    `json:"tags"`
	SubnetId                          string               `json:"subnet_id"`
	ValidFrom                         string               `json:"valid_from"`
	BlockDurationMinutes              int64                `json:"block_duration_minutes"`
	HostId                            string               `json:"host_id"`
	IamInstanceProfile                string               `json:"iam_instance_profile"`
	KeyName                           string               `json:"key_name"`
	LaunchGroup                       string               `json:"launch_group"`
	InstanceInitiatedShutdownBehavior string               `json:"instance_initiated_shutdown_behavior"`
	NetworkInterface                  NetworkInterface     `json:"network_interface"`
	RootBlockDevice                   RootBlockDevice      `json:"root_block_device"`
	Timeouts                          Timeouts             `json:"timeouts"`
	CreditSpecification               CreditSpecification  `json:"credit_specification"`
	EbsBlockDevice                    EbsBlockDevice       `json:"ebs_block_device"`
	EphemeralBlockDevice              EphemeralBlockDevice `json:"ephemeral_block_device"`
	MetadataOptions                   MetadataOptions      `json:"metadata_options"`
}

type NetworkInterface struct {
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceIndex         int64  `json:"device_index"`
	NetworkInterfaceId  string `json:"network_interface_id"`
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
	Create string `json:"create"`
	Delete string `json:"delete"`
}

type CreditSpecification struct {
	CpuCredits string `json:"cpu_credits"`
}

type EbsBlockDevice struct {
	Encrypted           bool   `json:"encrypted"`
	Iops                int64  `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int64  `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	KmsKeyId            string `json:"kms_key_id"`
	VolumeId            string `json:"volume_id"`
}

type EphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	NoDevice    bool   `json:"no_device"`
	VirtualName string `json:"virtual_name"`
}

type MetadataOptions struct {
	HttpEndpoint            string `json:"http_endpoint"`
	HttpPutResponseHopLimit int64  `json:"http_put_response_hop_limit"`
	HttpTokens              string `json:"http_tokens"`
}

// A SpotInstanceRequestStatus defines the observed state of a SpotInstanceRequest
type SpotInstanceRequestStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SpotInstanceRequestObservation `json:",inline"`
}

// A SpotInstanceRequestObservation records the observed state of a SpotInstanceRequest
type SpotInstanceRequestObservation struct {
	PrimaryNetworkInterfaceId string `json:"primary_network_interface_id"`
	PrivateDns                string `json:"private_dns"`
	SpotInstanceId            string `json:"spot_instance_id"`
	PublicIp                  string `json:"public_ip"`
	Arn                       string `json:"arn"`
	PasswordData              string `json:"password_data"`
	OutpostArn                string `json:"outpost_arn"`
	SpotBidStatus             string `json:"spot_bid_status"`
	InstanceState             string `json:"instance_state"`
	SpotRequestState          string `json:"spot_request_state"`
	PublicDns                 string `json:"public_dns"`
}