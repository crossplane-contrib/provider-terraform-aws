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

// LambdaLayerVersion is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LambdaLayerVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LambdaLayerVersionSpec   `json:"spec"`
	Status LambdaLayerVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LambdaLayerVersion contains a list of LambdaLayerVersionList
type LambdaLayerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LambdaLayerVersion `json:"items"`
}

// A LambdaLayerVersionSpec defines the desired state of a LambdaLayerVersion
type LambdaLayerVersionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LambdaLayerVersionParameters `json:",inline"`
}

// A LambdaLayerVersionParameters defines the desired state of a LambdaLayerVersion
type LambdaLayerVersionParameters struct {
	Filename           string   `json:"filename"`
	S3Bucket           string   `json:"s3_bucket"`
	CompatibleRuntimes []string `json:"compatible_runtimes"`
	Description        string   `json:"description"`
	SourceCodeHash     string   `json:"source_code_hash"`
	LayerName          string   `json:"layer_name"`
	LicenseInfo        string   `json:"license_info"`
	S3Key              string   `json:"s3_key"`
	Id                 string   `json:"id"`
	S3ObjectVersion    string   `json:"s3_object_version"`
}

// A LambdaLayerVersionStatus defines the observed state of a LambdaLayerVersion
type LambdaLayerVersionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LambdaLayerVersionObservation `json:",inline"`
}

// A LambdaLayerVersionObservation records the observed state of a LambdaLayerVersion
type LambdaLayerVersionObservation struct {
	SourceCodeSize int64  `json:"source_code_size"`
	Version        string `json:"version"`
	CreatedDate    string `json:"created_date"`
	Arn            string `json:"arn"`
	LayerArn       string `json:"layer_arn"`
}