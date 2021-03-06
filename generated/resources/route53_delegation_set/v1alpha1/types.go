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

// Route53DelegationSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53DelegationSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53DelegationSetSpec   `json:"spec"`
	Status Route53DelegationSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53DelegationSet contains a list of Route53DelegationSetList
type Route53DelegationSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53DelegationSet `json:"items"`
}

// A Route53DelegationSetSpec defines the desired state of a Route53DelegationSet
type Route53DelegationSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53DelegationSetParameters `json:"forProvider"`
}

// A Route53DelegationSetParameters defines the desired state of a Route53DelegationSet
type Route53DelegationSetParameters struct {
	ReferenceName string `json:"reference_name"`
}

// A Route53DelegationSetStatus defines the observed state of a Route53DelegationSet
type Route53DelegationSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53DelegationSetObservation `json:"atProvider"`
}

// A Route53DelegationSetObservation records the observed state of a Route53DelegationSet
type Route53DelegationSetObservation struct {
	NameServers []string `json:"name_servers"`
}