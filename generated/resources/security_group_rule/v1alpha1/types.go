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

// SecurityGroupRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SecurityGroupRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecurityGroupRuleSpec   `json:"spec"`
	Status SecurityGroupRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupRule contains a list of SecurityGroupRuleList
type SecurityGroupRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroupRule `json:"items"`
}

// A SecurityGroupRuleSpec defines the desired state of a SecurityGroupRule
type SecurityGroupRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SecurityGroupRuleParameters `json:",inline"`
}

// A SecurityGroupRuleParameters defines the desired state of a SecurityGroupRule
type SecurityGroupRuleParameters struct {
	SecurityGroupId       string   `json:"security_group_id"`
	Self                  bool     `json:"self"`
	Type                  string   `json:"type"`
	Id                    string   `json:"id"`
	Ipv6CidrBlocks        []string `json:"ipv6_cidr_blocks"`
	PrefixListIds         []string `json:"prefix_list_ids"`
	Protocol              string   `json:"protocol"`
	ToPort                int      `json:"to_port"`
	CidrBlocks            []string `json:"cidr_blocks"`
	Description           string   `json:"description"`
	FromPort              int      `json:"from_port"`
	SourceSecurityGroupId string   `json:"source_security_group_id"`
}

// A SecurityGroupRuleStatus defines the observed state of a SecurityGroupRule
type SecurityGroupRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SecurityGroupRuleObservation `json:",inline"`
}

// A SecurityGroupRuleObservation records the observed state of a SecurityGroupRule
type SecurityGroupRuleObservation struct{}