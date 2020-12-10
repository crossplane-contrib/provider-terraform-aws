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

// NetworkAcl is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NetworkAcl struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkAclSpec   `json:"spec"`
	Status NetworkAclStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkAcl contains a list of NetworkAclList
type NetworkAclList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkAcl `json:"items"`
}

// A NetworkAclSpec defines the desired state of a NetworkAcl
type NetworkAclSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NetworkAclParameters `json:",inline"`
}

// A NetworkAclParameters defines the desired state of a NetworkAcl
type NetworkAclParameters struct {
	VpcId     string            `json:"vpc_id"`
	Egress    []Egress          `json:"egress"`
	Id        string            `json:"id"`
	Ingress   []Ingress         `json:"ingress"`
	SubnetIds []string          `json:"subnet_ids"`
	Tags      map[string]string `json:"tags"`
}

type Egress struct {
	Action        string `json:"action"`
	IcmpCode      int64  `json:"icmp_code"`
	IcmpType      int64  `json:"icmp_type"`
	RuleNo        int64  `json:"rule_no"`
	ToPort        int64  `json:"to_port"`
	FromPort      int64  `json:"from_port"`
	CidrBlock     string `json:"cidr_block"`
	Protocol      string `json:"protocol"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
}

type Ingress struct {
	IcmpCode      int64  `json:"icmp_code"`
	IcmpType      int64  `json:"icmp_type"`
	FromPort      int64  `json:"from_port"`
	Protocol      string `json:"protocol"`
	RuleNo        int64  `json:"rule_no"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	Action        string `json:"action"`
	CidrBlock     string `json:"cidr_block"`
	ToPort        int64  `json:"to_port"`
}

// A NetworkAclStatus defines the observed state of a NetworkAcl
type NetworkAclStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NetworkAclObservation `json:",inline"`
}

// A NetworkAclObservation records the observed state of a NetworkAcl
type NetworkAclObservation struct {
	Arn     string `json:"arn"`
	OwnerId string `json:"owner_id"`
}