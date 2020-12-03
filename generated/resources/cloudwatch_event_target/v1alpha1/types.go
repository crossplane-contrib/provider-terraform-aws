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

// CloudwatchEventTarget is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchEventTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchEventTargetSpec   `json:"spec"`
	Status CloudwatchEventTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventTarget contains a list of CloudwatchEventTargetList
type CloudwatchEventTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventTarget `json:"items"`
}

// A CloudwatchEventTargetSpec defines the desired state of a CloudwatchEventTarget
type CloudwatchEventTargetSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchEventTargetParameters `json:",inline"`
}

// A CloudwatchEventTargetParameters defines the desired state of a CloudwatchEventTarget
type CloudwatchEventTargetParameters struct {
	Arn               string              `json:"arn"`
	Id                string              `json:"id"`
	Input             string              `json:"input"`
	InputPath         string              `json:"input_path"`
	RoleArn           string              `json:"role_arn"`
	Rule              string              `json:"rule"`
	TargetId          string              `json:"target_id"`
	BatchTarget       BatchTarget         `json:"batch_target"`
	EcsTarget         EcsTarget           `json:"ecs_target"`
	InputTransformer  InputTransformer    `json:"input_transformer"`
	KinesisTarget     KinesisTarget       `json:"kinesis_target"`
	RunCommandTargets []RunCommandTargets `json:"run_command_targets"`
	SqsTarget         SqsTarget           `json:"sqs_target"`
}

type BatchTarget struct {
	JobName       string `json:"job_name"`
	ArraySize     int    `json:"array_size"`
	JobAttempts   int    `json:"job_attempts"`
	JobDefinition string `json:"job_definition"`
}

type EcsTarget struct {
	Group                string               `json:"group"`
	LaunchType           string               `json:"launch_type"`
	PlatformVersion      string               `json:"platform_version"`
	TaskCount            int                  `json:"task_count"`
	TaskDefinitionArn    string               `json:"task_definition_arn"`
	NetworkConfiguration NetworkConfiguration `json:"network_configuration"`
}

type NetworkConfiguration struct {
	SecurityGroups []string `json:"security_groups"`
	Subnets        []string `json:"subnets"`
	AssignPublicIp bool     `json:"assign_public_ip"`
}

type InputTransformer struct {
	InputPaths    map[string]string `json:"input_paths"`
	InputTemplate string            `json:"input_template"`
}

type KinesisTarget struct {
	PartitionKeyPath string `json:"partition_key_path"`
}

type RunCommandTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SqsTarget struct {
	MessageGroupId string `json:"message_group_id"`
}

// A CloudwatchEventTargetStatus defines the observed state of a CloudwatchEventTarget
type CloudwatchEventTargetStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchEventTargetObservation `json:",inline"`
}

// A CloudwatchEventTargetObservation records the observed state of a CloudwatchEventTarget
type CloudwatchEventTargetObservation struct{}