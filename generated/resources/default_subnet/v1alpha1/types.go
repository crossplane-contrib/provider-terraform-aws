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

// DefaultSubnet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DefaultSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultSubnetSpec   `json:"spec"`
	Status DefaultSubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSubnet contains a list of DefaultSubnetList
type DefaultSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSubnet `json:"items"`
}

// A DefaultSubnetSpec defines the desired state of a DefaultSubnet
type DefaultSubnetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DefaultSubnetParameters `json:"forProvider"`
}

// A DefaultSubnetParameters defines the desired state of a DefaultSubnet
type DefaultSubnetParameters struct {
	AvailabilityZone    string            `json:"availability_zone"`
	MapPublicIpOnLaunch bool              `json:"map_public_ip_on_launch"`
	Id                  string            `json:"id"`
	OutpostArn          string            `json:"outpost_arn"`
	Tags                map[string]string `json:"tags"`
	Timeouts            Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A DefaultSubnetStatus defines the observed state of a DefaultSubnet
type DefaultSubnetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DefaultSubnetObservation `json:"atProvider"`
}

// A DefaultSubnetObservation records the observed state of a DefaultSubnet
type DefaultSubnetObservation struct {
	Arn                         string `json:"arn"`
	AssignIpv6AddressOnCreation bool   `json:"assign_ipv6_address_on_creation"`
	Ipv6CidrBlockAssociationId  string `json:"ipv6_cidr_block_association_id"`
	OwnerId                     string `json:"owner_id"`
	AvailabilityZoneId          string `json:"availability_zone_id"`
	CidrBlock                   string `json:"cidr_block"`
	Ipv6CidrBlock               string `json:"ipv6_cidr_block"`
	VpcId                       string `json:"vpc_id"`
}