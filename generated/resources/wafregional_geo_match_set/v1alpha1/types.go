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

// WafregionalGeoMatchSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalGeoMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalGeoMatchSetSpec   `json:"spec"`
	Status WafregionalGeoMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalGeoMatchSet contains a list of WafregionalGeoMatchSetList
type WafregionalGeoMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalGeoMatchSet `json:"items"`
}

// A WafregionalGeoMatchSetSpec defines the desired state of a WafregionalGeoMatchSet
type WafregionalGeoMatchSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalGeoMatchSetParameters `json:",inline"`
}

// A WafregionalGeoMatchSetParameters defines the desired state of a WafregionalGeoMatchSet
type WafregionalGeoMatchSetParameters struct {
	Name               string             `json:"name"`
	Id                 string             `json:"id"`
	GeoMatchConstraint GeoMatchConstraint `json:"geo_match_constraint"`
}

type GeoMatchConstraint struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// A WafregionalGeoMatchSetStatus defines the observed state of a WafregionalGeoMatchSet
type WafregionalGeoMatchSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalGeoMatchSetObservation `json:",inline"`
}

// A WafregionalGeoMatchSetObservation records the observed state of a WafregionalGeoMatchSet
type WafregionalGeoMatchSetObservation struct{}