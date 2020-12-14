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

// Ec2ClientVpnAuthorizationRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2ClientVpnAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2ClientVpnAuthorizationRuleSpec   `json:"spec"`
	Status Ec2ClientVpnAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2ClientVpnAuthorizationRule contains a list of Ec2ClientVpnAuthorizationRuleList
type Ec2ClientVpnAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2ClientVpnAuthorizationRule `json:"items"`
}

// A Ec2ClientVpnAuthorizationRuleSpec defines the desired state of a Ec2ClientVpnAuthorizationRule
type Ec2ClientVpnAuthorizationRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2ClientVpnAuthorizationRuleParameters `json:"forProvider"`
}

// A Ec2ClientVpnAuthorizationRuleParameters defines the desired state of a Ec2ClientVpnAuthorizationRule
type Ec2ClientVpnAuthorizationRuleParameters struct {
	Id                  string `json:"id"`
	TargetNetworkCidr   string `json:"target_network_cidr"`
	AccessGroupId       string `json:"access_group_id"`
	AuthorizeAllGroups  bool   `json:"authorize_all_groups"`
	ClientVpnEndpointId string `json:"client_vpn_endpoint_id"`
	Description         string `json:"description"`
}

// A Ec2ClientVpnAuthorizationRuleStatus defines the observed state of a Ec2ClientVpnAuthorizationRule
type Ec2ClientVpnAuthorizationRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2ClientVpnAuthorizationRuleObservation `json:"atProvider"`
}

// A Ec2ClientVpnAuthorizationRuleObservation records the observed state of a Ec2ClientVpnAuthorizationRule
type Ec2ClientVpnAuthorizationRuleObservation struct{}