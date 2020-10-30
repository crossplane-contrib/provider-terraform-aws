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
	Publish                      bool     `json:"publish"`
	S3Bucket                     string   `json:"s3_bucket"`
	Description                  string   `json:"description"`
	Handler                      string   `json:"handler"`
	Runtime                      string   `json:"runtime"`
	S3ObjectVersion              string   `json:"s3_object_version"`
	Filename                     string   `json:"filename"`
	FunctionName                 string   `json:"function_name"`
	Role                         string   `json:"role"`
	S3Key                        string   `json:"s3_key"`
	MemorySize                   int      `json:"memory_size"`
	ReservedConcurrentExecutions int      `json:"reserved_concurrent_executions"`
	Timeout                      int      `json:"timeout"`
	KmsKeyArn                    string   `json:"kms_key_arn"`
	Layers                       []string `json:"layers"`
}

// A LambdaFunctionStatus defines the observed state of a LambdaFunction
type LambdaFunctionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaFunctionObservation `json:",inline"`
}

// A LambdaFunctionObservation records the observed state of a LambdaFunction
type LambdaFunctionObservation struct {
	Id             string `json:"id"`
	InvokeArn      string `json:"invoke_arn"`
	LastModified   string `json:"last_modified"`
	Arn            string `json:"arn"`
	SourceCodeSize int    `json:"source_code_size"`
	Version        string `json:"version"`
	QualifiedArn   string `json:"qualified_arn"`
	SourceCodeHash string `json:"source_code_hash"`
}