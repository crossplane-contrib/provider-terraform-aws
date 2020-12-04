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

// SsmMaintenanceWindowTask is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SsmMaintenanceWindowTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SsmMaintenanceWindowTaskSpec   `json:"spec"`
	Status SsmMaintenanceWindowTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsmMaintenanceWindowTask contains a list of SsmMaintenanceWindowTaskList
type SsmMaintenanceWindowTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsmMaintenanceWindowTask `json:"items"`
}

// A SsmMaintenanceWindowTaskSpec defines the desired state of a SsmMaintenanceWindowTask
type SsmMaintenanceWindowTaskSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SsmMaintenanceWindowTaskParameters `json:",inline"`
}

// A SsmMaintenanceWindowTaskParameters defines the desired state of a SsmMaintenanceWindowTask
type SsmMaintenanceWindowTaskParameters struct {
	Description              string                   `json:"description"`
	Id                       string                   `json:"id"`
	MaxConcurrency           string                   `json:"max_concurrency"`
	Priority                 int64                    `json:"priority"`
	WindowId                 string                   `json:"window_id"`
	MaxErrors                string                   `json:"max_errors"`
	Name                     string                   `json:"name"`
	ServiceRoleArn           string                   `json:"service_role_arn"`
	TaskArn                  string                   `json:"task_arn"`
	TaskType                 string                   `json:"task_type"`
	Targets                  []Targets                `json:"targets"`
	TaskInvocationParameters TaskInvocationParameters `json:"task_invocation_parameters"`
}

type Targets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type TaskInvocationParameters struct {
	StepFunctionsParameters StepFunctionsParameters `json:"step_functions_parameters"`
	AutomationParameters    AutomationParameters    `json:"automation_parameters"`
	LambdaParameters        LambdaParameters        `json:"lambda_parameters"`
	RunCommandParameters    RunCommandParameters    `json:"run_command_parameters"`
}

type StepFunctionsParameters struct {
	Input string `json:"input"`
	Name  string `json:"name"`
}

type AutomationParameters struct {
	DocumentVersion string    `json:"document_version"`
	Parameter       Parameter `json:"parameter"`
}

type Parameter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type LambdaParameters struct {
	ClientContext string `json:"client_context"`
	Payload       string `json:"payload"`
	Qualifier     string `json:"qualifier"`
}

type RunCommandParameters struct {
	TimeoutSeconds     int64              `json:"timeout_seconds"`
	Comment            string             `json:"comment"`
	DocumentHash       string             `json:"document_hash"`
	DocumentHashType   string             `json:"document_hash_type"`
	OutputS3Bucket     string             `json:"output_s3_bucket"`
	OutputS3KeyPrefix  string             `json:"output_s3_key_prefix"`
	ServiceRoleArn     string             `json:"service_role_arn"`
	NotificationConfig NotificationConfig `json:"notification_config"`
	Parameter          Parameter          `json:"parameter"`
}

type NotificationConfig struct {
	NotificationArn    string   `json:"notification_arn"`
	NotificationEvents []string `json:"notification_events"`
	NotificationType   string   `json:"notification_type"`
}

type Parameter struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

// A SsmMaintenanceWindowTaskStatus defines the observed state of a SsmMaintenanceWindowTask
type SsmMaintenanceWindowTaskStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SsmMaintenanceWindowTaskObservation `json:",inline"`
}

// A SsmMaintenanceWindowTaskObservation records the observed state of a SsmMaintenanceWindowTask
type SsmMaintenanceWindowTaskObservation struct{}