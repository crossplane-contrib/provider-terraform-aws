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
	Monitoring                        bool   `json:"monitoring"`
	SourceDestCheck                   bool   `json:"source_dest_check"`
	WaitForFulfillment                bool   `json:"wait_for_fulfillment"`
	UserDataBase64                    string `json:"user_data_base64"`
	Ami                               string `json:"ami"`
	IamInstanceProfile                string `json:"iam_instance_profile"`
	SpotType                          string `json:"spot_type"`
	EbsOptimized                      bool   `json:"ebs_optimized"`
	GetPasswordData                   bool   `json:"get_password_data"`
	InstanceType                      string `json:"instance_type"`
	InstanceInitiatedShutdownBehavior string `json:"instance_initiated_shutdown_behavior"`
	UserData                          string `json:"user_data"`
	BlockDurationMinutes              int    `json:"block_duration_minutes"`
	LaunchGroup                       string `json:"launch_group"`
	DisableApiTermination             bool   `json:"disable_api_termination"`
	Hibernation                       bool   `json:"hibernation"`
	InstanceInterruptionBehaviour     string `json:"instance_interruption_behaviour"`
	SpotPrice                         string `json:"spot_price"`
}

// A SpotInstanceRequestStatus defines the observed state of a SpotInstanceRequest
type SpotInstanceRequestStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SpotInstanceRequestObservation `json:",inline"`
}

// A SpotInstanceRequestObservation records the observed state of a SpotInstanceRequest
type SpotInstanceRequestObservation struct {
	CpuCoreCount              int      `json:"cpu_core_count"`
	Ipv6Addresses             []string `json:"ipv6_addresses"`
	KeyName                   string   `json:"key_name"`
	Arn                       string   `json:"arn"`
	AssociatePublicIpAddress  bool     `json:"associate_public_ip_address"`
	AvailabilityZone          string   `json:"availability_zone"`
	Ipv6AddressCount          int      `json:"ipv6_address_count"`
	OutpostArn                string   `json:"outpost_arn"`
	SpotRequestState          string   `json:"spot_request_state"`
	InstanceState             string   `json:"instance_state"`
	PasswordData              string   `json:"password_data"`
	PrivateIp                 string   `json:"private_ip"`
	SpotInstanceId            string   `json:"spot_instance_id"`
	VpcSecurityGroupIds       []string `json:"vpc_security_group_ids"`
	Id                        string   `json:"id"`
	SpotBidStatus             string   `json:"spot_bid_status"`
	SubnetId                  string   `json:"subnet_id"`
	ValidUntil                string   `json:"valid_until"`
	CpuThreadsPerCore         int      `json:"cpu_threads_per_core"`
	PrivateDns                string   `json:"private_dns"`
	PublicDns                 string   `json:"public_dns"`
	SecondaryPrivateIps       []string `json:"secondary_private_ips"`
	HostId                    string   `json:"host_id"`
	PlacementGroup            string   `json:"placement_group"`
	PrimaryNetworkInterfaceId string   `json:"primary_network_interface_id"`
	PublicIp                  string   `json:"public_ip"`
	SecurityGroups            []string `json:"security_groups"`
	Tenancy                   string   `json:"tenancy"`
	ValidFrom                 string   `json:"valid_from"`
}