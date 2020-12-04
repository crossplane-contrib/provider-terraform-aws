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

// Ec2TrafficMirrorFilterRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2TrafficMirrorFilterRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2TrafficMirrorFilterRuleSpec   `json:"spec"`
	Status Ec2TrafficMirrorFilterRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TrafficMirrorFilterRule contains a list of Ec2TrafficMirrorFilterRuleList
type Ec2TrafficMirrorFilterRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TrafficMirrorFilterRule `json:"items"`
}

// A Ec2TrafficMirrorFilterRuleSpec defines the desired state of a Ec2TrafficMirrorFilterRule
type Ec2TrafficMirrorFilterRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2TrafficMirrorFilterRuleParameters `json:",inline"`
}

// A Ec2TrafficMirrorFilterRuleParameters defines the desired state of a Ec2TrafficMirrorFilterRule
type Ec2TrafficMirrorFilterRuleParameters struct {
	Description           string               `json:"description"`
	DestinationCidrBlock  string               `json:"destination_cidr_block"`
	Id                    string               `json:"id"`
	Protocol              int64                `json:"protocol"`
	RuleAction            string               `json:"rule_action"`
	RuleNumber            int64                `json:"rule_number"`
	SourceCidrBlock       string               `json:"source_cidr_block"`
	TrafficDirection      string               `json:"traffic_direction"`
	TrafficMirrorFilterId string               `json:"traffic_mirror_filter_id"`
	DestinationPortRange  DestinationPortRange `json:"destination_port_range"`
	SourcePortRange       SourcePortRange      `json:"source_port_range"`
}

type DestinationPortRange struct {
	FromPort int64 `json:"from_port"`
	ToPort   int64 `json:"to_port"`
}

type SourcePortRange struct {
	FromPort int64 `json:"from_port"`
	ToPort   int64 `json:"to_port"`
}

// A Ec2TrafficMirrorFilterRuleStatus defines the observed state of a Ec2TrafficMirrorFilterRule
type Ec2TrafficMirrorFilterRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2TrafficMirrorFilterRuleObservation `json:",inline"`
}

// A Ec2TrafficMirrorFilterRuleObservation records the observed state of a Ec2TrafficMirrorFilterRule
type Ec2TrafficMirrorFilterRuleObservation struct{}