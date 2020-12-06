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

// WafGeoMatchSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafGeoMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafGeoMatchSetSpec   `json:"spec"`
	Status WafGeoMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafGeoMatchSet contains a list of WafGeoMatchSetList
type WafGeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafGeoMatchSet `json:"items"`
}

// A WafGeoMatchSetSpec defines the desired state of a WafGeoMatchSet
type WafGeoMatchSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafGeoMatchSetParameters `json:",inline"`
}

// A WafGeoMatchSetParameters defines the desired state of a WafGeoMatchSet
type WafGeoMatchSetParameters struct {
	Name               string             `json:"name"`
	Id                 string             `json:"id"`
	GeoMatchConstraint GeoMatchConstraint `json:"geo_match_constraint"`
}

type GeoMatchConstraint struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// A WafGeoMatchSetStatus defines the observed state of a WafGeoMatchSet
type WafGeoMatchSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafGeoMatchSetObservation `json:",inline"`
}

// A WafGeoMatchSetObservation records the observed state of a WafGeoMatchSet
type WafGeoMatchSetObservation struct {
	Arn string `json:"arn"`
}