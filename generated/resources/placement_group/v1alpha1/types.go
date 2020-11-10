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

// PlacementGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type PlacementGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlacementGroupSpec   `json:"spec"`
	Status PlacementGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlacementGroup contains a list of PlacementGroupList
type PlacementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlacementGroup `json:"items"`
}

// A PlacementGroupSpec defines the desired state of a PlacementGroup
type PlacementGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  PlacementGroupParameters `json:",inline"`
}

// A PlacementGroupParameters defines the desired state of a PlacementGroup
type PlacementGroupParameters struct {
	Strategy string            `json:"strategy"`
	Tags     map[string]string `json:"tags"`
	Name     string            `json:"name"`
}

// A PlacementGroupStatus defines the observed state of a PlacementGroup
type PlacementGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     PlacementGroupObservation `json:",inline"`
}

// A PlacementGroupObservation records the observed state of a PlacementGroup
type PlacementGroupObservation struct {
	PlacementGroupId string `json:"placement_group_id"`
	Arn              string `json:"arn"`
	Id               string `json:"id"`
}