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

// EcsService is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type EcsService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EcsServiceSpec   `json:"spec"`
	Status EcsServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcsService contains a list of EcsServiceList
type EcsServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcsService `json:"items"`
}

// A EcsServiceSpec defines the desired state of a EcsService
type EcsServiceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  EcsServiceParameters `json:",inline"`
}

// A EcsServiceParameters defines the desired state of a EcsService
type EcsServiceParameters struct {
	PropagateTags                   string                     `json:"propagate_tags"`
	EnableEcsManagedTags            bool                       `json:"enable_ecs_managed_tags"`
	LaunchType                      string                     `json:"launch_type"`
	ForceNewDeployment              bool                       `json:"force_new_deployment"`
	SchedulingStrategy              string                     `json:"scheduling_strategy"`
	TaskDefinition                  string                     `json:"task_definition"`
	Cluster                         string                     `json:"cluster"`
	DeploymentMinimumHealthyPercent int                        `json:"deployment_minimum_healthy_percent"`
	PlatformVersion                 string                     `json:"platform_version"`
	Tags                            map[string]string          `json:"tags"`
	DeploymentMaximumPercent        int                        `json:"deployment_maximum_percent"`
	Id                              string                     `json:"id"`
	IamRole                         string                     `json:"iam_role"`
	Name                            string                     `json:"name"`
	DesiredCount                    int                        `json:"desired_count"`
	HealthCheckGracePeriodSeconds   int                        `json:"health_check_grace_period_seconds"`
	CapacityProviderStrategy        []CapacityProviderStrategy `json:"capacity_provider_strategy"`
	DeploymentController            DeploymentController       `json:"deployment_controller"`
	LoadBalancer                    []LoadBalancer             `json:"load_balancer"`
	NetworkConfiguration            NetworkConfiguration       `json:"network_configuration"`
	OrderedPlacementStrategy        []OrderedPlacementStrategy `json:"ordered_placement_strategy"`
	PlacementConstraints            []PlacementConstraints     `json:"placement_constraints"`
	ServiceRegistries               ServiceRegistries          `json:"service_registries"`
	Timeouts                        []Timeouts                 `json:"timeouts"`
}

type CapacityProviderStrategy struct {
	Base             int    `json:"base"`
	CapacityProvider string `json:"capacity_provider"`
	Weight           int    `json:"weight"`
}

type DeploymentController struct {
	Type string `json:"type"`
}

type LoadBalancer struct {
	ContainerName  string `json:"container_name"`
	ContainerPort  int    `json:"container_port"`
	ElbName        string `json:"elb_name"`
	TargetGroupArn string `json:"target_group_arn"`
}

type NetworkConfiguration struct {
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
	AssignPublicIp bool     `json:"assign_public_ip"`
}

type OrderedPlacementStrategy struct {
	Field string `json:"field"`
	Type  string `json:"type"`
}

type PlacementConstraints struct {
	Expression string `json:"expression"`
	Type       string `json:"type"`
}

type ServiceRegistries struct {
	RegistryArn   string `json:"registry_arn"`
	ContainerName string `json:"container_name"`
	ContainerPort int    `json:"container_port"`
	Port          int    `json:"port"`
}

type Timeouts struct {
	Delete string `json:"delete"`
}

// A EcsServiceStatus defines the observed state of a EcsService
type EcsServiceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     EcsServiceObservation `json:",inline"`
}

// A EcsServiceObservation records the observed state of a EcsService
type EcsServiceObservation struct{}