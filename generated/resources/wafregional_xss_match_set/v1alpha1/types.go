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

// WafregionalXssMatchSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalXssMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalXssMatchSetSpec   `json:"spec"`
	Status WafregionalXssMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalXssMatchSet contains a list of WafregionalXssMatchSetList
type WafregionalXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalXssMatchSet `json:"items"`
}

// A WafregionalXssMatchSetSpec defines the desired state of a WafregionalXssMatchSet
type WafregionalXssMatchSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalXssMatchSetParameters `json:",inline"`
}

// A WafregionalXssMatchSetParameters defines the desired state of a WafregionalXssMatchSet
type WafregionalXssMatchSetParameters struct {
	Name          string          `json:"name"`
	XssMatchTuple []XssMatchTuple `json:"xss_match_tuple"`
}

type XssMatchTuple struct {
	TextTransformation string       `json:"text_transformation"`
	FieldToMatch       FieldToMatch `json:"field_to_match"`
}

type FieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// A WafregionalXssMatchSetStatus defines the observed state of a WafregionalXssMatchSet
type WafregionalXssMatchSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalXssMatchSetObservation `json:",inline"`
}

// A WafregionalXssMatchSetObservation records the observed state of a WafregionalXssMatchSet
type WafregionalXssMatchSetObservation struct {
	Id string `json:"id"`
}