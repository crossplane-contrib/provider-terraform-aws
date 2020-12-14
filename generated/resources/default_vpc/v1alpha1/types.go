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

// DefaultVpc is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DefaultVpc struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultVpcSpec   `json:"spec"`
	Status DefaultVpcStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultVpc contains a list of DefaultVpcList
type DefaultVpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultVpc `json:"items"`
}

// A DefaultVpcSpec defines the desired state of a DefaultVpc
type DefaultVpcSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DefaultVpcParameters `json:"forProvider"`
}

// A DefaultVpcParameters defines the desired state of a DefaultVpc
type DefaultVpcParameters struct {
	EnableDnsSupport            bool              `json:"enable_dns_support"`
	EnableClassiclink           bool              `json:"enable_classiclink"`
	Id                          string            `json:"id"`
	EnableClassiclinkDnsSupport bool              `json:"enable_classiclink_dns_support"`
	EnableDnsHostnames          bool              `json:"enable_dns_hostnames"`
	Tags                        map[string]string `json:"tags"`
}

// A DefaultVpcStatus defines the observed state of a DefaultVpc
type DefaultVpcStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DefaultVpcObservation `json:"atProvider"`
}

// A DefaultVpcObservation records the observed state of a DefaultVpc
type DefaultVpcObservation struct {
	CidrBlock                    string `json:"cidr_block"`
	DefaultSecurityGroupId       string `json:"default_security_group_id"`
	DefaultRouteTableId          string `json:"default_route_table_id"`
	DhcpOptionsId                string `json:"dhcp_options_id"`
	InstanceTenancy              string `json:"instance_tenancy"`
	Ipv6CidrBlock                string `json:"ipv6_cidr_block"`
	OwnerId                      string `json:"owner_id"`
	Arn                          string `json:"arn"`
	AssignGeneratedIpv6CidrBlock bool   `json:"assign_generated_ipv6_cidr_block"`
	DefaultNetworkAclId          string `json:"default_network_acl_id"`
	Ipv6AssociationId            string `json:"ipv6_association_id"`
	MainRouteTableId             string `json:"main_route_table_id"`
}