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

// Wafv2RegexPatternSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Wafv2RegexPatternSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Wafv2RegexPatternSetSpec   `json:"spec"`
	Status Wafv2RegexPatternSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Wafv2RegexPatternSet contains a list of Wafv2RegexPatternSetList
type Wafv2RegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Wafv2RegexPatternSet `json:"items"`
}

// A Wafv2RegexPatternSetSpec defines the desired state of a Wafv2RegexPatternSet
type Wafv2RegexPatternSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Wafv2RegexPatternSetParameters `json:"forProvider"`
}

// A Wafv2RegexPatternSetParameters defines the desired state of a Wafv2RegexPatternSet
type Wafv2RegexPatternSetParameters struct {
	Description       string              `json:"description"`
	Id                string              `json:"id"`
	Name              string              `json:"name"`
	Scope             string              `json:"scope"`
	Tags              map[string]string   `json:"tags"`
	RegularExpression []RegularExpression `json:"regular_expression"`
}

type RegularExpression struct {
	RegexString string `json:"regex_string"`
}

// A Wafv2RegexPatternSetStatus defines the observed state of a Wafv2RegexPatternSet
type Wafv2RegexPatternSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Wafv2RegexPatternSetObservation `json:"atProvider"`
}

// A Wafv2RegexPatternSetObservation records the observed state of a Wafv2RegexPatternSet
type Wafv2RegexPatternSetObservation struct {
	LockToken string `json:"lock_token"`
	Arn       string `json:"arn"`
}