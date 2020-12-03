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

// SecurityGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityGroupSpec   `json:"spec"`
	Status SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup contains a list of SecurityGroupList
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// A SecurityGroupSpec defines the desired state of a SecurityGroup
type SecurityGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SecurityGroupParameters `json:",inline"`
}

// A SecurityGroupParameters defines the desired state of a SecurityGroup
type SecurityGroupParameters struct {
	Tags                map[string]string `json:"tags"`
	VpcId               string            `json:"vpc_id"`
	Id                  string            `json:"id"`
	Ingress             []Ingress         `json:"ingress"`
	NamePrefix          string            `json:"name_prefix"`
	RevokeRulesOnDelete bool              `json:"revoke_rules_on_delete"`
	Description         string            `json:"description"`
	Egress              []Egress          `json:"egress"`
	Name                string            `json:"name"`
	Timeouts            []Timeouts        `json:"timeouts"`
}

type Ingress struct {
	ToPort         int      `json:"to_port"`
	CidrBlocks     []string `json:"cidr_blocks"`
	Protocol       string   `json:"protocol"`
	SecurityGroups []string `json:"security_groups"`
	Self           bool     `json:"self"`
	Description    string   `json:"description"`
	FromPort       int      `json:"from_port"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
}

type Egress struct {
	Protocol       string   `json:"protocol"`
	SecurityGroups []string `json:"security_groups"`
	Description    string   `json:"description"`
	Ipv6CidrBlocks []string `json:"ipv6_cidr_blocks"`
	CidrBlocks     []string `json:"cidr_blocks"`
	PrefixListIds  []string `json:"prefix_list_ids"`
	ToPort         int      `json:"to_port"`
	FromPort       int      `json:"from_port"`
	Self           bool     `json:"self"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A SecurityGroupStatus defines the observed state of a SecurityGroup
type SecurityGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SecurityGroupObservation `json:",inline"`
}

// A SecurityGroupObservation records the observed state of a SecurityGroup
type SecurityGroupObservation struct {
	OwnerId string `json:"owner_id"`
	Arn     string `json:"arn"`
}