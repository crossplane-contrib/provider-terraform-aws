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

// LoadBalancerPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LoadBalancerPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadBalancerPolicySpec   `json:"spec"`
	Status LoadBalancerPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerPolicy contains a list of LoadBalancerPolicyList
type LoadBalancerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerPolicy `json:"items"`
}

// A LoadBalancerPolicySpec defines the desired state of a LoadBalancerPolicy
type LoadBalancerPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LoadBalancerPolicyParameters `json:",inline"`
}

// A LoadBalancerPolicyParameters defines the desired state of a LoadBalancerPolicy
type LoadBalancerPolicyParameters struct {
	LoadBalancerName string            `json:"load_balancer_name"`
	PolicyName       string            `json:"policy_name"`
	PolicyTypeName   string            `json:"policy_type_name"`
	PolicyAttribute  []PolicyAttribute `json:"policy_attribute"`
}

type PolicyAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A LoadBalancerPolicyStatus defines the observed state of a LoadBalancerPolicy
type LoadBalancerPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LoadBalancerPolicyObservation `json:",inline"`
}

// A LoadBalancerPolicyObservation records the observed state of a LoadBalancerPolicy
type LoadBalancerPolicyObservation struct {
	Id string `json:"id"`
}