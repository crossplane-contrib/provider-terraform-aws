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

// Route53ResolverRuleAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53ResolverRuleAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53ResolverRuleAssociationSpec   `json:"spec"`
	Status Route53ResolverRuleAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ResolverRuleAssociation contains a list of Route53ResolverRuleAssociationList
type Route53ResolverRuleAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ResolverRuleAssociation `json:"items"`
}

// A Route53ResolverRuleAssociationSpec defines the desired state of a Route53ResolverRuleAssociation
type Route53ResolverRuleAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53ResolverRuleAssociationParameters `json:",inline"`
}

// A Route53ResolverRuleAssociationParameters defines the desired state of a Route53ResolverRuleAssociation
type Route53ResolverRuleAssociationParameters struct {
	ResolverRuleId string     `json:"resolver_rule_id"`
	VpcId          string     `json:"vpc_id"`
	Name           string     `json:"name"`
	Timeouts       []Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Delete string `json:"delete"`
	Create string `json:"create"`
}

// A Route53ResolverRuleAssociationStatus defines the observed state of a Route53ResolverRuleAssociation
type Route53ResolverRuleAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53ResolverRuleAssociationObservation `json:",inline"`
}

// A Route53ResolverRuleAssociationObservation records the observed state of a Route53ResolverRuleAssociation
type Route53ResolverRuleAssociationObservation struct {
	Id string `json:"id"`
}