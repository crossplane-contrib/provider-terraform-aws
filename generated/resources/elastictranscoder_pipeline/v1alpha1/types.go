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

// ElastictranscoderPipeline is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ElastictranscoderPipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElastictranscoderPipelineSpec   `json:"spec"`
	Status ElastictranscoderPipelineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElastictranscoderPipeline contains a list of ElastictranscoderPipelineList
type ElastictranscoderPipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElastictranscoderPipeline `json:"items"`
}

// A ElastictranscoderPipelineSpec defines the desired state of a ElastictranscoderPipeline
type ElastictranscoderPipelineSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ElastictranscoderPipelineParameters `json:",inline"`
}

// A ElastictranscoderPipelineParameters defines the desired state of a ElastictranscoderPipeline
type ElastictranscoderPipelineParameters struct {
	AwsKmsKeyArn               string                       `json:"aws_kms_key_arn"`
	Id                         string                       `json:"id"`
	InputBucket                string                       `json:"input_bucket"`
	Name                       string                       `json:"name"`
	OutputBucket               string                       `json:"output_bucket"`
	Role                       string                       `json:"role"`
	ContentConfig              ContentConfig                `json:"content_config"`
	ContentConfigPermissions   []ContentConfigPermissions   `json:"content_config_permissions"`
	Notifications              Notifications                `json:"notifications"`
	ThumbnailConfig            ThumbnailConfig              `json:"thumbnail_config"`
	ThumbnailConfigPermissions []ThumbnailConfigPermissions `json:"thumbnail_config_permissions"`
}

type ContentConfig struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

type ContentConfigPermissions struct {
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
	GranteeType string   `json:"grantee_type"`
}

type Notifications struct {
	Error       string `json:"error"`
	Progressing string `json:"progressing"`
	Warning     string `json:"warning"`
	Completed   string `json:"completed"`
}

type ThumbnailConfig struct {
	Bucket       string `json:"bucket"`
	StorageClass string `json:"storage_class"`
}

type ThumbnailConfigPermissions struct {
	GranteeType string   `json:"grantee_type"`
	Access      []string `json:"access"`
	Grantee     string   `json:"grantee"`
}

// A ElastictranscoderPipelineStatus defines the observed state of a ElastictranscoderPipeline
type ElastictranscoderPipelineStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ElastictranscoderPipelineObservation `json:",inline"`
}

// A ElastictranscoderPipelineObservation records the observed state of a ElastictranscoderPipeline
type ElastictranscoderPipelineObservation struct {
	Arn string `json:"arn"`
}