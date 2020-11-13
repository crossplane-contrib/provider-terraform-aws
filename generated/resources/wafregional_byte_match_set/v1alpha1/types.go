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

// WafregionalByteMatchSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalByteMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalByteMatchSetSpec   `json:"spec"`
	Status WafregionalByteMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalByteMatchSet contains a list of WafregionalByteMatchSetList
type WafregionalByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalByteMatchSet `json:"items"`
}

// A WafregionalByteMatchSetSpec defines the desired state of a WafregionalByteMatchSet
type WafregionalByteMatchSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalByteMatchSetParameters `json:",inline"`
}

// A WafregionalByteMatchSetParameters defines the desired state of a WafregionalByteMatchSet
type WafregionalByteMatchSetParameters struct {
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	ByteMatchTuples []ByteMatchTuples `json:"byte_match_tuples"`
}

type ByteMatchTuples struct {
	PositionalConstraint string       `json:"positional_constraint"`
	TargetString         string       `json:"target_string"`
	TextTransformation   string       `json:"text_transformation"`
	FieldToMatch         FieldToMatch `json:"field_to_match"`
}

type FieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// A WafregionalByteMatchSetStatus defines the observed state of a WafregionalByteMatchSet
type WafregionalByteMatchSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalByteMatchSetObservation `json:",inline"`
}

// A WafregionalByteMatchSetObservation records the observed state of a WafregionalByteMatchSet
type WafregionalByteMatchSetObservation struct{}