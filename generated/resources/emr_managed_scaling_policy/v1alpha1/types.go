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

// EmrManagedScalingPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EmrManagedScalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmrManagedScalingPolicySpec   `json:"spec"`
	Status EmrManagedScalingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrManagedScalingPolicy contains a list of EmrManagedScalingPolicyList
type EmrManagedScalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrManagedScalingPolicy `json:"items"`
}

// A EmrManagedScalingPolicySpec defines the desired state of a EmrManagedScalingPolicy
type EmrManagedScalingPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EmrManagedScalingPolicyParameters `json:",inline"`
}

// A EmrManagedScalingPolicyParameters defines the desired state of a EmrManagedScalingPolicy
type EmrManagedScalingPolicyParameters struct {
	ClusterId     string          `json:"cluster_id"`
	ComputeLimits []ComputeLimits `json:"compute_limits"`
}

type ComputeLimits struct {
	MaximumCapacityUnits         int    `json:"maximum_capacity_units"`
	MaximumCoreCapacityUnits     int    `json:"maximum_core_capacity_units"`
	MaximumOndemandCapacityUnits int    `json:"maximum_ondemand_capacity_units"`
	MinimumCapacityUnits         int    `json:"minimum_capacity_units"`
	UnitType                     string `json:"unit_type"`
}

// A EmrManagedScalingPolicyStatus defines the observed state of a EmrManagedScalingPolicy
type EmrManagedScalingPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EmrManagedScalingPolicyObservation `json:",inline"`
}

// A EmrManagedScalingPolicyObservation records the observed state of a EmrManagedScalingPolicy
type EmrManagedScalingPolicyObservation struct {
	Id string `json:"id"`
}