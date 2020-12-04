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

// Ec2Fleet is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Ec2Fleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Ec2FleetSpec   `json:"spec"`
	Status Ec2FleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2Fleet contains a list of Ec2FleetList
type Ec2FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2Fleet `json:"items"`
}

// A Ec2FleetSpec defines the desired state of a Ec2Fleet
type Ec2FleetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Ec2FleetParameters `json:",inline"`
}

// A Ec2FleetParameters defines the desired state of a Ec2Fleet
type Ec2FleetParameters struct {
	ReplaceUnhealthyInstances        bool                        `json:"replace_unhealthy_instances"`
	Tags                             map[string]string           `json:"tags"`
	TerminateInstances               bool                        `json:"terminate_instances"`
	TerminateInstancesWithExpiration bool                        `json:"terminate_instances_with_expiration"`
	Type                             string                      `json:"type"`
	ExcessCapacityTerminationPolicy  string                      `json:"excess_capacity_termination_policy"`
	Id                               string                      `json:"id"`
	TargetCapacitySpecification      TargetCapacitySpecification `json:"target_capacity_specification"`
	Timeouts                         Timeouts                    `json:"timeouts"`
	LaunchTemplateConfig             LaunchTemplateConfig        `json:"launch_template_config"`
	OnDemandOptions                  OnDemandOptions             `json:"on_demand_options"`
	SpotOptions                      SpotOptions                 `json:"spot_options"`
}

type TargetCapacitySpecification struct {
	DefaultTargetCapacityType string `json:"default_target_capacity_type"`
	OnDemandTargetCapacity    int64  `json:"on_demand_target_capacity"`
	SpotTargetCapacity        int64  `json:"spot_target_capacity"`
	TotalTargetCapacity       int64  `json:"total_target_capacity"`
}

type Timeouts struct {
	Update string `json:"update"`
	Create string `json:"create"`
	Delete string `json:"delete"`
}

type LaunchTemplateConfig struct {
	LaunchTemplateSpecification LaunchTemplateSpecification `json:"launch_template_specification"`
	Override                    []Override                  `json:"override"`
}

type LaunchTemplateSpecification struct {
	LaunchTemplateId   string `json:"launch_template_id"`
	LaunchTemplateName string `json:"launch_template_name"`
	Version            string `json:"version"`
}

type Override struct {
	InstanceType     string `json:"instance_type"`
	MaxPrice         string `json:"max_price"`
	Priority         int64  `json:"priority"`
	SubnetId         string `json:"subnet_id"`
	WeightedCapacity int64  `json:"weighted_capacity"`
	AvailabilityZone string `json:"availability_zone"`
}

type OnDemandOptions struct {
	AllocationStrategy string `json:"allocation_strategy"`
}

type SpotOptions struct {
	AllocationStrategy           string `json:"allocation_strategy"`
	InstanceInterruptionBehavior string `json:"instance_interruption_behavior"`
	InstancePoolsToUseCount      int64  `json:"instance_pools_to_use_count"`
}

// A Ec2FleetStatus defines the observed state of a Ec2Fleet
type Ec2FleetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Ec2FleetObservation `json:",inline"`
}

// A Ec2FleetObservation records the observed state of a Ec2Fleet
type Ec2FleetObservation struct{}