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

// LbSslNegotiationPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LbSslNegotiationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LbSslNegotiationPolicySpec   `json:"spec"`
	Status LbSslNegotiationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbSslNegotiationPolicy contains a list of LbSslNegotiationPolicyList
type LbSslNegotiationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbSslNegotiationPolicy `json:"items"`
}

// A LbSslNegotiationPolicySpec defines the desired state of a LbSslNegotiationPolicy
type LbSslNegotiationPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LbSslNegotiationPolicyParameters `json:",inline"`
}

// A LbSslNegotiationPolicyParameters defines the desired state of a LbSslNegotiationPolicy
type LbSslNegotiationPolicyParameters struct {
	Name         string      `json:"name"`
	LbPort       int         `json:"lb_port"`
	LoadBalancer string      `json:"load_balancer"`
	Attribute    []Attribute `json:"attribute"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A LbSslNegotiationPolicyStatus defines the observed state of a LbSslNegotiationPolicy
type LbSslNegotiationPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LbSslNegotiationPolicyObservation `json:",inline"`
}

// A LbSslNegotiationPolicyObservation records the observed state of a LbSslNegotiationPolicy
type LbSslNegotiationPolicyObservation struct {
	Id string `json:"id"`
}