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

// Route53ResolverRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53ResolverRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53ResolverRuleSpec   `json:"spec"`
	Status Route53ResolverRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverRule contains a list of Route53ResolverRuleList
type Route53ResolverRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverRule `json:"items"`
}

// A Route53ResolverRuleSpec defines the desired state of a Route53ResolverRule
type Route53ResolverRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53ResolverRuleParameters `json:",inline"`
}

// A Route53ResolverRuleParameters defines the desired state of a Route53ResolverRule
type Route53ResolverRuleParameters struct {
	Id                 string            `json:"id"`
	ResolverEndpointId string            `json:"resolver_endpoint_id"`
	RuleType           string            `json:"rule_type"`
	DomainName         string            `json:"domain_name"`
	Name               string            `json:"name"`
	Tags               map[string]string `json:"tags"`
	TargetIp           TargetIp          `json:"target_ip"`
	Timeouts           Timeouts          `json:"timeouts"`
}

type TargetIp struct {
	Ip   string `json:"ip"`
	Port int64  `json:"port"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A Route53ResolverRuleStatus defines the observed state of a Route53ResolverRule
type Route53ResolverRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53ResolverRuleObservation `json:",inline"`
}

// A Route53ResolverRuleObservation records the observed state of a Route53ResolverRule
type Route53ResolverRuleObservation struct {
	Arn         string `json:"arn"`
	ShareStatus string `json:"share_status"`
	OwnerId     string `json:"owner_id"`
}