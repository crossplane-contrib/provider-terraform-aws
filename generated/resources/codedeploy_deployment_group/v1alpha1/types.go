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

// CodedeployDeploymentGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CodedeployDeploymentGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CodedeployDeploymentGroupSpec   `json:"spec"`
	Status CodedeployDeploymentGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodedeployDeploymentGroup contains a list of CodedeployDeploymentGroupList
type CodedeployDeploymentGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodedeployDeploymentGroup `json:"items"`
}

// A CodedeployDeploymentGroupSpec defines the desired state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CodedeployDeploymentGroupParameters `json:",inline"`
}

// A CodedeployDeploymentGroupParameters defines the desired state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupParameters struct {
	Id                          string                        `json:"id"`
	ServiceRoleArn              string                        `json:"service_role_arn"`
	AppName                     string                        `json:"app_name"`
	AutoscalingGroups           []string                      `json:"autoscaling_groups"`
	DeploymentConfigName        string                        `json:"deployment_config_name"`
	DeploymentGroupName         string                        `json:"deployment_group_name"`
	Ec2TagFilter                []Ec2TagFilter                `json:"ec2_tag_filter"`
	Ec2TagSet                   []Ec2TagSet                   `json:"ec2_tag_set"`
	LoadBalancerInfo            LoadBalancerInfo              `json:"load_balancer_info"`
	OnPremisesInstanceTagFilter []OnPremisesInstanceTagFilter `json:"on_premises_instance_tag_filter"`
	AutoRollbackConfiguration   AutoRollbackConfiguration     `json:"auto_rollback_configuration"`
	BlueGreenDeploymentConfig   BlueGreenDeploymentConfig     `json:"blue_green_deployment_config"`
	DeploymentStyle             DeploymentStyle               `json:"deployment_style"`
	EcsService                  EcsService                    `json:"ecs_service"`
	TriggerConfiguration        []TriggerConfiguration        `json:"trigger_configuration"`
	AlarmConfiguration          AlarmConfiguration            `json:"alarm_configuration"`
}

type Ec2TagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Ec2TagSet struct {
	Ec2TagFilter []Ec2TagFilter `json:"ec2_tag_filter"`
}

type Ec2TagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type LoadBalancerInfo struct {
	ElbInfo             []ElbInfo           `json:"elb_info"`
	TargetGroupInfo     []TargetGroupInfo   `json:"target_group_info"`
	TargetGroupPairInfo TargetGroupPairInfo `json:"target_group_pair_info"`
}

type ElbInfo struct {
	Name string `json:"name"`
}

type TargetGroupInfo struct {
	Name string `json:"name"`
}

type TargetGroupPairInfo struct {
	ProdTrafficRoute ProdTrafficRoute `json:"prod_traffic_route"`
	TargetGroup      []TargetGroup    `json:"target_group"`
	TestTrafficRoute TestTrafficRoute `json:"test_traffic_route"`
}

type ProdTrafficRoute struct {
	ListenerArns []string `json:"listener_arns"`
}

type TargetGroup struct {
	Name string `json:"name"`
}

type TestTrafficRoute struct {
	ListenerArns []string `json:"listener_arns"`
}

type OnPremisesInstanceTagFilter struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type AutoRollbackConfiguration struct {
	Enabled bool     `json:"enabled"`
	Events  []string `json:"events"`
}

type BlueGreenDeploymentConfig struct {
	DeploymentReadyOption                     DeploymentReadyOption                     `json:"deployment_ready_option"`
	GreenFleetProvisioningOption              GreenFleetProvisioningOption              `json:"green_fleet_provisioning_option"`
	TerminateBlueInstancesOnDeploymentSuccess TerminateBlueInstancesOnDeploymentSuccess `json:"terminate_blue_instances_on_deployment_success"`
}

type DeploymentReadyOption struct {
	ActionOnTimeout   string `json:"action_on_timeout"`
	WaitTimeInMinutes int    `json:"wait_time_in_minutes"`
}

type GreenFleetProvisioningOption struct {
	Action string `json:"action"`
}

type TerminateBlueInstancesOnDeploymentSuccess struct {
	Action                       string `json:"action"`
	TerminationWaitTimeInMinutes int    `json:"termination_wait_time_in_minutes"`
}

type DeploymentStyle struct {
	DeploymentOption string `json:"deployment_option"`
	DeploymentType   string `json:"deployment_type"`
}

type EcsService struct {
	ClusterName string `json:"cluster_name"`
	ServiceName string `json:"service_name"`
}

type TriggerConfiguration struct {
	TriggerEvents    []string `json:"trigger_events"`
	TriggerName      string   `json:"trigger_name"`
	TriggerTargetArn string   `json:"trigger_target_arn"`
}

type AlarmConfiguration struct {
	Alarms                 []string `json:"alarms"`
	Enabled                bool     `json:"enabled"`
	IgnorePollAlarmFailure bool     `json:"ignore_poll_alarm_failure"`
}

// A CodedeployDeploymentGroupStatus defines the observed state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CodedeployDeploymentGroupObservation `json:",inline"`
}

// A CodedeployDeploymentGroupObservation records the observed state of a CodedeployDeploymentGroup
type CodedeployDeploymentGroupObservation struct{}