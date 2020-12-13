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

// VpcEndpoint is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcEndpointSpec   `json:"spec"`
	Status VpcEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpoint contains a list of VpcEndpointList
type VpcEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcEndpoint `json:"items"`
}

// A VpcEndpointSpec defines the desired state of a VpcEndpoint
type VpcEndpointSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcEndpointParameters `json:"forProvider"`
}

// A VpcEndpointParameters defines the desired state of a VpcEndpoint
type VpcEndpointParameters struct {
	PrivateDnsEnabled bool              `json:"private_dns_enabled"`
	SecurityGroupIds  []string          `json:"security_group_ids"`
	Tags              map[string]string `json:"tags"`
	VpcEndpointType   string            `json:"vpc_endpoint_type"`
	VpcId             string            `json:"vpc_id"`
	AutoAccept        bool              `json:"auto_accept"`
	Id                string            `json:"id"`
	ServiceName       string            `json:"service_name"`
	SubnetIds         []string          `json:"subnet_ids"`
	Policy            string            `json:"policy"`
	RouteTableIds     []string          `json:"route_table_ids"`
	Timeouts          Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A VpcEndpointStatus defines the observed state of a VpcEndpoint
type VpcEndpointStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcEndpointObservation `json:"atProvider"`
}

// A VpcEndpointObservation records the observed state of a VpcEndpoint
type VpcEndpointObservation struct {
	CidrBlocks          []string   `json:"cidr_blocks"`
	DnsEntry            []DnsEntry `json:"dns_entry"`
	NetworkInterfaceIds []string   `json:"network_interface_ids"`
	Arn                 string     `json:"arn"`
	OwnerId             string     `json:"owner_id"`
	PrefixListId        string     `json:"prefix_list_id"`
	RequesterManaged    bool       `json:"requester_managed"`
	State               string     `json:"state"`
}

type DnsEntry struct {
	DnsName      string `json:"dns_name"`
	HostedZoneId string `json:"hosted_zone_id"`
}