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
	Ami                               string `json:"ami"`
	Monitoring                        bool   `json:"monitoring"`
	DisableApiTermination             bool   `json:"disable_api_termination"`
	GetPasswordData                   bool   `json:"get_password_data"`
	IamInstanceProfile                string `json:"iam_instance_profile"`
	InstanceType                      string `json:"instance_type"`
	Hibernation                       bool   `json:"hibernation"`
	InstanceInitiatedShutdownBehavior string `json:"instance_initiated_shutdown_behavior"`
	EbsOptimized                      bool   `json:"ebs_optimized"`
	UserDataBase64                    string `json:"user_data_base64"`
	SourceDestCheck                   bool   `json:"source_dest_check"`
	UserData                          string `json:"user_data"`
}

// A InstanceStatus defines the observed state of a Instance
type InstanceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     InstanceObservation `json:",inline"`
}

// A InstanceObservation records the observed state of a Instance
type InstanceObservation struct {
	CpuCoreCount              int      `json:"cpu_core_count"`
	Id                        string   `json:"id"`
	VpcSecurityGroupIds       []string `json:"vpc_security_group_ids"`
	Ipv6AddressCount          int      `json:"ipv6_address_count"`
	Ipv6Addresses             []string `json:"ipv6_addresses"`
	PublicIp                  string   `json:"public_ip"`
	AssociatePublicIpAddress  bool     `json:"associate_public_ip_address"`
	AvailabilityZone          string   `json:"availability_zone"`
	PrimaryNetworkInterfaceId string   `json:"primary_network_interface_id"`
	PrivateIp                 string   `json:"private_ip"`
	SubnetId                  string   `json:"subnet_id"`
	InstanceState             string   `json:"instance_state"`
	SecondaryPrivateIps       []string `json:"secondary_private_ips"`
	Arn                       string   `json:"arn"`
	CpuThreadsPerCore         int      `json:"cpu_threads_per_core"`
	HostId                    string   `json:"host_id"`
	PlacementGroup            string   `json:"placement_group"`
	SecurityGroups            []string `json:"security_groups"`
	PrivateDns                string   `json:"private_dns"`
	PublicDns                 string   `json:"public_dns"`
	PasswordData              string   `json:"password_data"`
	KeyName                   string   `json:"key_name"`
	OutpostArn                string   `json:"outpost_arn"`
	Tenancy                   string   `json:"tenancy"`
}