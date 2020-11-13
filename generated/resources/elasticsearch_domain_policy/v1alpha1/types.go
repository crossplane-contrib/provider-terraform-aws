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

// ElasticsearchDomainPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElasticsearchDomainPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticsearchDomainPolicySpec   `json:"spec"`
	Status ElasticsearchDomainPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchDomainPolicy contains a list of ElasticsearchDomainPolicyList
type ElasticsearchDomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchDomainPolicy `json:"items"`
}

// A ElasticsearchDomainPolicySpec defines the desired state of a ElasticsearchDomainPolicy
type ElasticsearchDomainPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElasticsearchDomainPolicyParameters `json:",inline"`
}

// A ElasticsearchDomainPolicyParameters defines the desired state of a ElasticsearchDomainPolicy
type ElasticsearchDomainPolicyParameters struct {
	AccessPolicies string `json:"access_policies"`
	DomainName     string `json:"domain_name"`
	Id             string `json:"id"`
}

// A ElasticsearchDomainPolicyStatus defines the observed state of a ElasticsearchDomainPolicy
type ElasticsearchDomainPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElasticsearchDomainPolicyObservation `json:",inline"`
}

// A ElasticsearchDomainPolicyObservation records the observed state of a ElasticsearchDomainPolicy
type ElasticsearchDomainPolicyObservation struct{}