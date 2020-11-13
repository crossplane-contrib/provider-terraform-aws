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

// GuarddutyFilter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GuarddutyFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuarddutyFilterSpec   `json:"spec"`
	Status GuarddutyFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuarddutyFilter contains a list of GuarddutyFilterList
type GuarddutyFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuarddutyFilter `json:"items"`
}

// A GuarddutyFilterSpec defines the desired state of a GuarddutyFilter
type GuarddutyFilterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GuarddutyFilterParameters `json:",inline"`
}

// A GuarddutyFilterParameters defines the desired state of a GuarddutyFilter
type GuarddutyFilterParameters struct {
	Tags            map[string]string `json:"tags"`
	Action          string            `json:"action"`
	Description     string            `json:"description"`
	DetectorId      string            `json:"detector_id"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	Rank            int               `json:"rank"`
	FindingCriteria FindingCriteria   `json:"finding_criteria"`
}

type FindingCriteria struct {
	Criterion []Criterion `json:"criterion"`
}

type Criterion struct {
	Equals             []string `json:"equals"`
	Field              string   `json:"field"`
	GreaterThan        string   `json:"greater_than"`
	GreaterThanOrEqual string   `json:"greater_than_or_equal"`
	LessThan           string   `json:"less_than"`
	LessThanOrEqual    string   `json:"less_than_or_equal"`
	NotEquals          []string `json:"not_equals"`
}

// A GuarddutyFilterStatus defines the observed state of a GuarddutyFilter
type GuarddutyFilterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GuarddutyFilterObservation `json:",inline"`
}

// A GuarddutyFilterObservation records the observed state of a GuarddutyFilter
type GuarddutyFilterObservation struct {
	Arn string `json:"arn"`
}