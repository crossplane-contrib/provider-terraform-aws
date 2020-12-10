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

// WafXssMatchSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafXssMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafXssMatchSetSpec   `json:"spec"`
	Status WafXssMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafXssMatchSet contains a list of WafXssMatchSetList
type WafXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafXssMatchSet `json:"items"`
}

// A WafXssMatchSetSpec defines the desired state of a WafXssMatchSet
type WafXssMatchSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafXssMatchSetParameters `json:",inline"`
}

// A WafXssMatchSetParameters defines the desired state of a WafXssMatchSet
type WafXssMatchSetParameters struct {
	Id             string         `json:"id"`
	Name           string         `json:"name"`
	XssMatchTuples XssMatchTuples `json:"xss_match_tuples"`
}

type XssMatchTuples struct {
	TextTransformation string       `json:"text_transformation"`
	FieldToMatch       FieldToMatch `json:"field_to_match"`
}

type FieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// A WafXssMatchSetStatus defines the observed state of a WafXssMatchSet
type WafXssMatchSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafXssMatchSetObservation `json:",inline"`
}

// A WafXssMatchSetObservation records the observed state of a WafXssMatchSet
type WafXssMatchSetObservation struct {
	Arn string `json:"arn"`
}