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

// Subnet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SubnetSpec   `json:"spec"`
	Status SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet contains a list of SubnetList
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// A SubnetSpec defines the desired state of a Subnet
type SubnetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SubnetParameters `json:",inline"`
}

// A SubnetParameters defines the desired state of a Subnet
type SubnetParameters struct {
	AssignIpv6AddressOnCreation bool   `json:"assign_ipv6_address_on_creation"`
	CidrBlock                   string `json:"cidr_block"`
	MapPublicIpOnLaunch         bool   `json:"map_public_ip_on_launch"`
	OutpostArn                  string `json:"outpost_arn"`
	Ipv6CidrBlock               string `json:"ipv6_cidr_block"`
	VpcId                       string `json:"vpc_id"`
}

// A SubnetStatus defines the observed state of a Subnet
type SubnetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SubnetObservation `json:",inline"`
}

// A SubnetObservation records the observed state of a Subnet
type SubnetObservation struct {
	Arn                        string `json:"arn"`
	AvailabilityZone           string `json:"availability_zone"`
	Id                         string `json:"id"`
	Ipv6CidrBlockAssociationId string `json:"ipv6_cidr_block_association_id"`
	AvailabilityZoneId         string `json:"availability_zone_id"`
	OwnerId                    string `json:"owner_id"`
}