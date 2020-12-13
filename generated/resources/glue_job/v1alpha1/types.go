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

// GlueJob is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GlueJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlueJobSpec   `json:"spec"`
	Status GlueJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueJob contains a list of GlueJobList
type GlueJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueJob `json:"items"`
}

// A GlueJobSpec defines the desired state of a GlueJob
type GlueJobSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GlueJobParameters `json:"forProvider"`
}

// A GlueJobParameters defines the desired state of a GlueJob
type GlueJobParameters struct {
	SecurityConfiguration   string               `json:"security_configuration"`
	Tags                    map[string]string    `json:"tags"`
	Timeout                 int64                `json:"timeout"`
	Connections             []string             `json:"connections"`
	Description             string               `json:"description"`
	Id                      string               `json:"id"`
	NumberOfWorkers         int64                `json:"number_of_workers"`
	RoleArn                 string               `json:"role_arn"`
	WorkerType              string               `json:"worker_type"`
	DefaultArguments        map[string]string    `json:"default_arguments"`
	GlueVersion             string               `json:"glue_version"`
	MaxCapacity             int64                `json:"max_capacity"`
	MaxRetries              int64                `json:"max_retries"`
	NonOverridableArguments map[string]string    `json:"non_overridable_arguments"`
	Name                    string               `json:"name"`
	Command                 Command              `json:"command"`
	ExecutionProperty       ExecutionProperty    `json:"execution_property"`
	NotificationProperty    NotificationProperty `json:"notification_property"`
}

type Command struct {
	Name           string `json:"name"`
	PythonVersion  string `json:"python_version"`
	ScriptLocation string `json:"script_location"`
}

type ExecutionProperty struct {
	MaxConcurrentRuns int64 `json:"max_concurrent_runs"`
}

type NotificationProperty struct {
	NotifyDelayAfter int64 `json:"notify_delay_after"`
}

// A GlueJobStatus defines the observed state of a GlueJob
type GlueJobStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GlueJobObservation `json:"atProvider"`
}

// A GlueJobObservation records the observed state of a GlueJob
type GlueJobObservation struct {
	Arn string `json:"arn"`
}