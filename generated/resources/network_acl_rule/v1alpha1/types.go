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

// NetworkAclRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type NetworkAclRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkAclRuleSpec   `json:"spec"`
	Status NetworkAclRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkAclRule contains a list of NetworkAclRuleList
type NetworkAclRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkAclRule `json:"items"`
}

// A NetworkAclRuleSpec defines the desired state of a NetworkAclRule
type NetworkAclRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  NetworkAclRuleParameters `json:"forProvider"`
}

// A NetworkAclRuleParameters defines the desired state of a NetworkAclRule
type NetworkAclRuleParameters struct {
	Protocol      string `json:"protocol"`
	RuleNumber    int64  `json:"rule_number"`
	CidrBlock     string `json:"cidr_block"`
	IcmpCode      string `json:"icmp_code"`
	NetworkAclId  string `json:"network_acl_id"`
	Ipv6CidrBlock string `json:"ipv6_cidr_block"`
	RuleAction    string `json:"rule_action"`
	ToPort        int64  `json:"to_port"`
	Egress        bool   `json:"egress"`
	FromPort      int64  `json:"from_port"`
	IcmpType      string `json:"icmp_type"`
}

// A NetworkAclRuleStatus defines the observed state of a NetworkAclRule
type NetworkAclRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     NetworkAclRuleObservation `json:"atProvider"`
}

// A NetworkAclRuleObservation records the observed state of a NetworkAclRule
type NetworkAclRuleObservation struct{}