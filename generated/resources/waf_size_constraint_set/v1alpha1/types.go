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

// WafSizeConstraintSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafSizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafSizeConstraintSetSpec   `json:"spec"`
	Status WafSizeConstraintSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafSizeConstraintSet contains a list of WafSizeConstraintSetList
type WafSizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafSizeConstraintSet `json:"items"`
}

// A WafSizeConstraintSetSpec defines the desired state of a WafSizeConstraintSet
type WafSizeConstraintSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafSizeConstraintSetParameters `json:",inline"`
}

// A WafSizeConstraintSetParameters defines the desired state of a WafSizeConstraintSet
type WafSizeConstraintSetParameters struct {
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

// A WafSizeConstraintSetStatus defines the observed state of a WafSizeConstraintSet
type WafSizeConstraintSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafSizeConstraintSetObservation `json:",inline"`
}

// A WafSizeConstraintSetObservation records the observed state of a WafSizeConstraintSet
type WafSizeConstraintSetObservation struct {
	Arn string `json:"arn"`
}