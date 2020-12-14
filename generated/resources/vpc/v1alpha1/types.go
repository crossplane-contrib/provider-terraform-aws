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

// Vpc is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Vpc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcSpec   `json:"spec"`
	Status VpcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Vpc contains a list of VpcList
type VpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vpc `json:"items"`
}

// A VpcSpec defines the desired state of a Vpc
type VpcSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcParameters `json:"forProvider"`
}

// A VpcParameters defines the desired state of a Vpc
type VpcParameters struct {
	EnableClassiclinkDnsSupport  bool              `json:"enable_classiclink_dns_support"`
	EnableDnsSupport             bool              `json:"enable_dns_support"`
	AssignGeneratedIpv6CidrBlock bool              `json:"assign_generated_ipv6_cidr_block"`
	EnableClassiclink            bool              `json:"enable_classiclink"`
	EnableDnsHostnames           bool              `json:"enable_dns_hostnames"`
	InstanceTenancy              string            `json:"instance_tenancy"`
	Tags                         map[string]string `json:"tags"`
	CidrBlock                    string            `json:"cidr_block"`
	Id                           string            `json:"id"`
}

// A VpcStatus defines the observed state of a Vpc
type VpcStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcObservation `json:"atProvider"`
}

// A VpcObservation records the observed state of a Vpc
type VpcObservation struct {
	DefaultRouteTableId    string `json:"default_route_table_id"`
	Ipv6AssociationId      string `json:"ipv6_association_id"`
	MainRouteTableId       string `json:"main_route_table_id"`
	Ipv6CidrBlock          string `json:"ipv6_cidr_block"`
	OwnerId                string `json:"owner_id"`
	Arn                    string `json:"arn"`
	DefaultNetworkAclId    string `json:"default_network_acl_id"`
	DefaultSecurityGroupId string `json:"default_security_group_id"`
	DhcpOptionsId          string `json:"dhcp_options_id"`
}