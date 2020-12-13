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

// DefaultSecurityGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DefaultSecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DefaultSecurityGroupSpec   `json:"spec"`
	Status DefaultSecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DefaultSecurityGroup contains a list of DefaultSecurityGroupList
type DefaultSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DefaultSecurityGroup `json:"items"`
}

// A DefaultSecurityGroupSpec defines the desired state of a DefaultSecurityGroup
type DefaultSecurityGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DefaultSecurityGroupParameters `json:"forProvider"`
}

// A DefaultSecurityGroupParameters defines the desired state of a DefaultSecurityGroup
type DefaultSecurityGroupParameters struct {
	Id                  string            `json:"id"`
	VpcId               string            `json:"vpc_id"`
	Egress              []Egress          `json:"egress"`
	Ingress             []Ingress         `json:"ingress"`
	RevokeRulesOnDelete bool              `json:"revoke_rules_on_delete"`
	Tags                map[string]string `json:"tags"`
}

type Egress struct {
	PrefixListIds  []string `json:"prefix_list_ids"`
	ToPort         int64    `json:"to_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	Self           bool     `json:"self"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Protocol       string   `json:"protocol"`
	SecurityGroups []string `json:"security_groups"`
	FromPort       int64    `json:"from_port"`
	Description    string   `json:"description"`
}

type Ingress struct {
	PrefixListIds  []string `json:"prefix_list_ids"`
	Description    string   `json:"description"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	Protocol       string   `json:"protocol"`
	SecurityGroups []string `json:"security_groups"`
	CidrBlocks     []string `json:"cidr_blocks"`
	FromPort       int64    `json:"from_port"`
	ToPort         int64    `json:"to_port"`
	Self           bool     `json:"self"`
}

// A DefaultSecurityGroupStatus defines the observed state of a DefaultSecurityGroup
type DefaultSecurityGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DefaultSecurityGroupObservation `json:"atProvider"`
}

// A DefaultSecurityGroupObservation records the observed state of a DefaultSecurityGroup
type DefaultSecurityGroupObservation struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	OwnerId     string `json:"owner_id"`
	Arn         string `json:"arn"`
}