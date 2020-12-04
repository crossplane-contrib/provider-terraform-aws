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

// WafregionalSizeConstraintSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafregionalSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafregionalSizeConstraintSetSpec   `json:"spec"`
	Status WafregionalSizeConstraintSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalSizeConstraintSet contains a list of WafregionalSizeConstraintSetList
type WafregionalSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalSizeConstraintSet `json:"items"`
}

// A WafregionalSizeConstraintSetSpec defines the desired state of a WafregionalSizeConstraintSet
type WafregionalSizeConstraintSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafregionalSizeConstraintSetParameters `json:",inline"`
}

// A WafregionalSizeConstraintSetParameters defines the desired state of a WafregionalSizeConstraintSet
type WafregionalSizeConstraintSetParameters struct {
	Id              string          `json:"id"`
	Name            string          `json:"name"`
	SizeConstraints SizeConstraints `json:"size_constraints"`
}

type SizeConstraints struct {
	ComparisonOperator string       `json:"comparison_operator"`
	Size               int64        `json:"size"`
	TextTransformation string       `json:"text_transformation"`
	FieldToMatch       FieldToMatch `json:"field_to_match"`
}

type FieldToMatch struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

// A WafregionalSizeConstraintSetStatus defines the observed state of a WafregionalSizeConstraintSet
type WafregionalSizeConstraintSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafregionalSizeConstraintSetObservation `json:",inline"`
}

// A WafregionalSizeConstraintSetObservation records the observed state of a WafregionalSizeConstraintSet
type WafregionalSizeConstraintSetObservation struct {
	Arn string `json:"arn"`
}