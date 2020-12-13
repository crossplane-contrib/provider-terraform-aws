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

// DatasyncLocationS3 is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DatasyncLocationS3 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasyncLocationS3Spec   `json:"spec"`
	Status DatasyncLocationS3Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncLocationS3 contains a list of DatasyncLocationS3List
type DatasyncLocationS3List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncLocationS3 `json:"items"`
}

// A DatasyncLocationS3Spec defines the desired state of a DatasyncLocationS3
type DatasyncLocationS3Spec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DatasyncLocationS3Parameters `json:"forProvider"`
}

// A DatasyncLocationS3Parameters defines the desired state of a DatasyncLocationS3
type DatasyncLocationS3Parameters struct {
	Id           string            `json:"id"`
	S3BucketArn  string            `json:"s3_bucket_arn"`
	Subdirectory string            `json:"subdirectory"`
	Tags         map[string]string `json:"tags"`
	S3Config     S3Config          `json:"s3_config"`
}

type S3Config struct {
	BucketAccessRoleArn string `json:"bucket_access_role_arn"`
}

// A DatasyncLocationS3Status defines the observed state of a DatasyncLocationS3
type DatasyncLocationS3Status struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DatasyncLocationS3Observation `json:"atProvider"`
}

// A DatasyncLocationS3Observation records the observed state of a DatasyncLocationS3
type DatasyncLocationS3Observation struct {
	Uri string `json:"uri"`
	Arn string `json:"arn"`
}