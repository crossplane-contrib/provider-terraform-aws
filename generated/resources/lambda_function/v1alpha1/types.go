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

// LambdaFunction is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LambdaFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LambdaFunctionSpec   `json:"spec"`
	Status LambdaFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaFunction contains a list of LambdaFunctionList
type LambdaFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaFunction `json:"items"`
}

// A LambdaFunctionSpec defines the desired state of a LambdaFunction
type LambdaFunctionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LambdaFunctionParameters `json:",inline"`
}

// A LambdaFunctionParameters defines the desired state of a LambdaFunction
type LambdaFunctionParameters struct {
	S3Bucket                     string            `json:"s3_bucket"`
	S3Key                        string            `json:"s3_key"`
	Id                           string            `json:"id"`
	Role                         string            `json:"role"`
	Publish                      bool              `json:"publish"`
	ReservedConcurrentExecutions int               `json:"reserved_concurrent_executions"`
	Filename                     string            `json:"filename"`
	KmsKeyArn                    string            `json:"kms_key_arn"`
	MemorySize                   int               `json:"memory_size"`
	Runtime                      string            `json:"runtime"`
	S3ObjectVersion              string            `json:"s3_object_version"`
	SourceCodeHash               string            `json:"source_code_hash"`
	Description                  string            `json:"description"`
	Handler                      string            `json:"handler"`
	Layers                       []string          `json:"layers"`
	Tags                         map[string]string `json:"tags"`
	Timeout                      int               `json:"timeout"`
	FunctionName                 string            `json:"function_name"`
	FileSystemConfig             FileSystemConfig  `json:"file_system_config"`
	Timeouts                     []Timeouts        `json:"timeouts"`
	TracingConfig                TracingConfig     `json:"tracing_config"`
	VpcConfig                    VpcConfig         `json:"vpc_config"`
	DeadLetterConfig             DeadLetterConfig  `json:"dead_letter_config"`
	Environment                  Environment       `json:"environment"`
}

type FileSystemConfig struct {
	Arn            string `json:"arn"`
	LocalMountPath string `json:"local_mount_path"`
}

type Timeouts struct {
	Create string `json:"create"`
}

type TracingConfig struct {
	Mode string `json:"mode"`
}

type VpcConfig struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
}

type DeadLetterConfig struct {
	TargetArn string `json:"target_arn"`
}

type Environment struct {
	Variables map[string]string `json:"variables"`
}

// A LambdaFunctionStatus defines the observed state of a LambdaFunction
type LambdaFunctionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaFunctionObservation `json:",inline"`
}

// A LambdaFunctionObservation records the observed state of a LambdaFunction
type LambdaFunctionObservation struct {
	LastModified   string `json:"last_modified"`
	QualifiedArn   string `json:"qualified_arn"`
	InvokeArn      string `json:"invoke_arn"`
	Version        string `json:"version"`
	SourceCodeSize int    `json:"source_code_size"`
	Arn            string `json:"arn"`
}