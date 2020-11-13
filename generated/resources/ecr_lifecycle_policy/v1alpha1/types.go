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

// EcrLifecyclePolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EcrLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EcrLifecyclePolicySpec   `json:"spec"`
	Status EcrLifecyclePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcrLifecyclePolicy contains a list of EcrLifecyclePolicyList
type EcrLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcrLifecyclePolicy `json:"items"`
}

// A EcrLifecyclePolicySpec defines the desired state of a EcrLifecyclePolicy
type EcrLifecyclePolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EcrLifecyclePolicyParameters `json:",inline"`
}

// A EcrLifecyclePolicyParameters defines the desired state of a EcrLifecyclePolicy
type EcrLifecyclePolicyParameters struct {
	Id         string `json:"id"`
	Policy     string `json:"policy"`
	Repository string `json:"repository"`
}

// A EcrLifecyclePolicyStatus defines the observed state of a EcrLifecyclePolicy
type EcrLifecyclePolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EcrLifecyclePolicyObservation `json:",inline"`
}

// A EcrLifecyclePolicyObservation records the observed state of a EcrLifecyclePolicy
type EcrLifecyclePolicyObservation struct {
	RegistryId string `json:"registry_id"`
}