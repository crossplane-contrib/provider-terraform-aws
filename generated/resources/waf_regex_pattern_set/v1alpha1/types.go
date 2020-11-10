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

// WafRegexPatternSet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type WafRegexPatternSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WafRegexPatternSetSpec   `json:"spec"`
	Status WafRegexPatternSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafRegexPatternSet contains a list of WafRegexPatternSetList
type WafRegexPatternSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafRegexPatternSet `json:"items"`
}

// A WafRegexPatternSetSpec defines the desired state of a WafRegexPatternSet
type WafRegexPatternSetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  WafRegexPatternSetParameters `json:",inline"`
}

// A WafRegexPatternSetParameters defines the desired state of a WafRegexPatternSet
type WafRegexPatternSetParameters struct {
	Name                string   `json:"name"`
	RegexPatternStrings []string `json:"regex_pattern_strings"`
}

// A WafRegexPatternSetStatus defines the observed state of a WafRegexPatternSet
type WafRegexPatternSetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     WafRegexPatternSetObservation `json:",inline"`
}

// A WafRegexPatternSetObservation records the observed state of a WafRegexPatternSet
type WafRegexPatternSetObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}