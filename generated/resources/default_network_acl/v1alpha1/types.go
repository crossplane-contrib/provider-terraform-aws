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

// DefaultNetworkAcl is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DefaultNetworkAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultNetworkAclSpec   `json:"spec"`
	Status DefaultNetworkAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultNetworkAcl contains a list of DefaultNetworkAclList
type DefaultNetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultNetworkAcl `json:"items"`
}

// A DefaultNetworkAclSpec defines the desired state of a DefaultNetworkAcl
type DefaultNetworkAclSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DefaultNetworkAclParameters `json:"forProvider"`
}

// A DefaultNetworkAclParameters defines the desired state of a DefaultNetworkAcl
type DefaultNetworkAclParameters struct {
	Tags                map[string]string `json:"tags"`
	DefaultNetworkAclId string            `json:"default_network_acl_id"`
	Id                  string            `json:"id"`
	SubnetIds           []string          `json:"subnet_ids"`
	Egress              Egress            `json:"egress"`
	Ingress             Ingress           `json:"ingress"`
}

type Egress struct {
	Action        string `json:"action"`
	FromPort      int64  `json:"from_port"`
	ToPort        int64  `json:"to_port"`
	RuleNo        int64  `json:"rule_no"`
	CidrBlock     string `json:"cidr_block"`
	IcmpCode      int64  `json:"icmp_code"`
	IcmpType      int64  `json:"icmp_type"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	Protocol      string `json:"protocol"`
}

type Ingress struct {
	CidrBlock     string `json:"cidr_block"`
	IcmpCode      int64  `json:"icmp_code"`
	IcmpType      int64  `json:"icmp_type"`
	RuleNo        int64  `json:"rule_no"`
	Action        string `json:"action"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	Protocol      string `json:"protocol"`
	ToPort        int64  `json:"to_port"`
	FromPort      int64  `json:"from_port"`
}

// A DefaultNetworkAclStatus defines the observed state of a DefaultNetworkAcl
type DefaultNetworkAclStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DefaultNetworkAclObservation `json:"atProvider"`
}

// A DefaultNetworkAclObservation records the observed state of a DefaultNetworkAcl
type DefaultNetworkAclObservation struct {
	VpcId   string `json:"vpc_id"`
	Arn     string `json:"arn"`
	OwnerId string `json:"owner_id"`
}