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

// WafByteMatchSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafByteMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafByteMatchSetSpec   `json:"spec"`
	Status WafByteMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafByteMatchSet contains a list of WafByteMatchSetList
type WafByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafByteMatchSet `json:"items"`
}

// A WafByteMatchSetSpec defines the desired state of a WafByteMatchSet
type WafByteMatchSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafByteMatchSetParameters `json:",inline"`
}

// A WafByteMatchSetParameters defines the desired state of a WafByteMatchSet
type WafByteMatchSetParameters struct {
	Name string `json:"name"`
}

// A WafByteMatchSetStatus defines the observed state of a WafByteMatchSet
type WafByteMatchSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafByteMatchSetObservation `json:",inline"`
}

// A WafByteMatchSetObservation records the observed state of a WafByteMatchSet
type WafByteMatchSetObservation struct {
	Id string `json:"id"`
}