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
	LaunchConfiguration    string   `json:"launch_configuration"`
	TerminationPolicies    []string `json:"termination_policies"`
	WaitForElbCapacity     int      `json:"wait_for_elb_capacity"`
	WaitForCapacityTimeout string   `json:"wait_for_capacity_timeout"`
	EnabledMetrics         []string `json:"enabled_metrics"`
	MinSize                int      `json:"min_size"`
	TargetGroupArns        []string `json:"target_group_arns"`
	MinElbCapacity         int      `json:"min_elb_capacity"`
	SuspendedProcesses     []string `json:"suspended_processes"`
	MaxInstanceLifetime    int      `json:"max_instance_lifetime"`
	MetricsGranularity     string   `json:"metrics_granularity"`
	HealthCheckGracePeriod int      `json:"health_check_grace_period"`
	MaxSize                int      `json:"max_size"`
	NamePrefix             string   `json:"name_prefix"`
	PlacementGroup         string   `json:"placement_group"`
	ForceDelete            bool     `json:"force_delete"`
	LoadBalancers          []string `json:"load_balancers"`
	ProtectFromScaleIn     bool     `json:"protect_from_scale_in"`
}

// A AutoscalingGroupStatus defines the observed state of a AutoscalingGroup
type AutoscalingGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AutoscalingGroupObservation `json:",inline"`
}

// A AutoscalingGroupObservation records the observed state of a AutoscalingGroup
type AutoscalingGroupObservation struct {
	ServiceLinkedRoleArn string   `json:"service_linked_role_arn"`
	HealthCheckType      string   `json:"health_check_type"`
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Arn                  string   `json:"arn"`
	DesiredCapacity      int      `json:"desired_capacity"`
	AvailabilityZones    []string `json:"availability_zones"`
	VpcZoneIdentifier    []string `json:"vpc_zone_identifier"`
	DefaultCooldown      int      `json:"default_cooldown"`
}