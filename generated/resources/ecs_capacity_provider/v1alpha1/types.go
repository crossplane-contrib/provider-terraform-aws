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

// EcsCapacityProvider is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EcsCapacityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EcsCapacityProviderSpec   `json:"spec"`
	Status EcsCapacityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcsCapacityProvider contains a list of EcsCapacityProviderList
type EcsCapacityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcsCapacityProvider `json:"items"`
}

// A EcsCapacityProviderSpec defines the desired state of a EcsCapacityProvider
type EcsCapacityProviderSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EcsCapacityProviderParameters `json:",inline"`
}

// A EcsCapacityProviderParameters defines the desired state of a EcsCapacityProvider
type EcsCapacityProviderParameters struct {
	Id                       string                   `json:"id"`
	Name                     string                   `json:"name"`
	Tags                     map[string]string        `json:"tags"`
	AutoScalingGroupProvider AutoScalingGroupProvider `json:"auto_scaling_group_provider"`
}

type AutoScalingGroupProvider struct {
	AutoScalingGroupArn          string         `json:"auto_scaling_group_arn"`
	ManagedTerminationProtection string         `json:"managed_termination_protection"`
	ManagedScaling               ManagedScaling `json:"managed_scaling"`
}

type ManagedScaling struct {
	MaximumScalingStepSize int64  `json:"maximum_scaling_step_size"`
	MinimumScalingStepSize int64  `json:"minimum_scaling_step_size"`
	Status                 string `json:"status"`
	TargetCapacity         int64  `json:"target_capacity"`
}

// A EcsCapacityProviderStatus defines the observed state of a EcsCapacityProvider
type EcsCapacityProviderStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EcsCapacityProviderObservation `json:",inline"`
}

// A EcsCapacityProviderObservation records the observed state of a EcsCapacityProvider
type EcsCapacityProviderObservation struct {
	Arn string `json:"arn"`
}