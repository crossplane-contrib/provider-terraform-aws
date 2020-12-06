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

// EcsCluster is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EcsCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EcsClusterSpec   `json:"spec"`
	Status EcsClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcsCluster contains a list of EcsClusterList
type EcsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcsCluster `json:"items"`
}

// A EcsClusterSpec defines the desired state of a EcsCluster
type EcsClusterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EcsClusterParameters `json:",inline"`
}

// A EcsClusterParameters defines the desired state of a EcsCluster
type EcsClusterParameters struct {
	CapacityProviders               []string                        `json:"capacity_providers"`
	Id                              string                          `json:"id"`
	Name                            string                          `json:"name"`
	Tags                            map[string]string               `json:"tags"`
	DefaultCapacityProviderStrategy DefaultCapacityProviderStrategy `json:"default_capacity_provider_strategy"`
	Setting                         Setting                         `json:"setting"`
}

type DefaultCapacityProviderStrategy struct {
	Base             int64  `json:"base"`
	CapacityProvider string `json:"capacity_provider"`
	Weight           int64  `json:"weight"`
}

type Setting struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// A EcsClusterStatus defines the observed state of a EcsCluster
type EcsClusterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EcsClusterObservation `json:",inline"`
}

// A EcsClusterObservation records the observed state of a EcsCluster
type EcsClusterObservation struct {
	Arn string `json:"arn"`
}