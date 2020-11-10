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

// AutoscalingGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AutoscalingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutoscalingGroupSpec   `json:"spec"`
	Status AutoscalingGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AutoscalingGroup contains a list of AutoscalingGroupList
type AutoscalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoscalingGroup `json:"items"`
}

// A AutoscalingGroupSpec defines the desired state of a AutoscalingGroup
type AutoscalingGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AutoscalingGroupParameters `json:",inline"`
}

// A AutoscalingGroupParameters defines the desired state of a AutoscalingGroup
type AutoscalingGroupParameters struct {
	EnabledMetrics         []string               `json:"enabled_metrics"`
	MetricsGranularity     string                 `json:"metrics_granularity"`
	ForceDelete            bool                   `json:"force_delete"`
	MaxInstanceLifetime    int                    `json:"max_instance_lifetime"`
	ProtectFromScaleIn     bool                   `json:"protect_from_scale_in"`
	TargetGroupArns        []string               `json:"target_group_arns"`
	TerminationPolicies    []string               `json:"termination_policies"`
	WaitForElbCapacity     int                    `json:"wait_for_elb_capacity"`
	LaunchConfiguration    string                 `json:"launch_configuration"`
	MinSize                int                    `json:"min_size"`
	PlacementGroup         string                 `json:"placement_group"`
	MaxSize                int                    `json:"max_size"`
	MinElbCapacity         int                    `json:"min_elb_capacity"`
	NamePrefix             string                 `json:"name_prefix"`
	SuspendedProcesses     []string               `json:"suspended_processes"`
	WaitForCapacityTimeout string                 `json:"wait_for_capacity_timeout"`
	HealthCheckGracePeriod int                    `json:"health_check_grace_period"`
	LoadBalancers          []string               `json:"load_balancers"`
	Timeouts               []Timeouts             `json:"timeouts"`
	InitialLifecycleHook   []InitialLifecycleHook `json:"initial_lifecycle_hook"`
	LaunchTemplate         LaunchTemplate         `json:"launch_template"`
	MixedInstancesPolicy   MixedInstancesPolicy   `json:"mixed_instances_policy"`
	Tag                    []Tag                  `json:"tag"`
}

type Timeouts struct {
	Delete string `json:"delete"`
}

type InitialLifecycleHook struct {
	DefaultResult         string `json:"default_result"`
	HeartbeatTimeout      int    `json:"heartbeat_timeout"`
	LifecycleTransition   string `json:"lifecycle_transition"`
	Name                  string `json:"name"`
	NotificationMetadata  string `json:"notification_metadata"`
	NotificationTargetArn string `json:"notification_target_arn"`
	RoleArn               string `json:"role_arn"`
}

type LaunchTemplate struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type MixedInstancesPolicy struct {
	LaunchTemplate        LaunchTemplate        `json:"launch_template"`
	InstancesDistribution InstancesDistribution `json:"instances_distribution"`
}

type LaunchTemplate struct {
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
	WeightedCapacity string `json:"weighted_capacity"`
}

type InstancesDistribution struct {
	OnDemandAllocationStrategy          string `json:"on_demand_allocation_strategy"`
	OnDemandBaseCapacity                int    `json:"on_demand_base_capacity"`
	OnDemandPercentageAboveBaseCapacity int    `json:"on_demand_percentage_above_base_capacity"`
	SpotAllocationStrategy              string `json:"spot_allocation_strategy"`
	SpotInstancePools                   int    `json:"spot_instance_pools"`
	SpotMaxPrice                        string `json:"spot_max_price"`
}

type Tag struct {
	Key               string `json:"key"`
	PropagateAtLaunch bool   `json:"propagate_at_launch"`
	Value             string `json:"value"`
}

// A AutoscalingGroupStatus defines the observed state of a AutoscalingGroup
type AutoscalingGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AutoscalingGroupObservation `json:",inline"`
}

// A AutoscalingGroupObservation records the observed state of a AutoscalingGroup
type AutoscalingGroupObservation struct {
	DesiredCapacity      int      `json:"desired_capacity"`
	Name                 string   `json:"name"`
	Id                   string   `json:"id"`
	ServiceLinkedRoleArn string   `json:"service_linked_role_arn"`
	Arn                  string   `json:"arn"`
	AvailabilityZones    []string `json:"availability_zones"`
	DefaultCooldown      int      `json:"default_cooldown"`
	HealthCheckType      string   `json:"health_check_type"`
	VpcZoneIdentifier    []string `json:"vpc_zone_identifier"`
}