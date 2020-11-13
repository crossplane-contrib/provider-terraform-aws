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

// S3BucketInventory is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3BucketInventory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BucketInventorySpec   `json:"spec"`
	Status S3BucketInventoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketInventory contains a list of S3BucketInventoryList
type S3BucketInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketInventory `json:"items"`
}

// A S3BucketInventorySpec defines the desired state of a S3BucketInventory
type S3BucketInventorySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3BucketInventoryParameters `json:",inline"`
}

// A S3BucketInventoryParameters defines the desired state of a S3BucketInventory
type S3BucketInventoryParameters struct {
	OptionalFields         []string    `json:"optional_fields"`
	Bucket                 string      `json:"bucket"`
	Enabled                bool        `json:"enabled"`
	Id                     string      `json:"id"`
	IncludedObjectVersions string      `json:"included_object_versions"`
	Name                   string      `json:"name"`
	Destination            Destination `json:"destination"`
	Filter                 Filter      `json:"filter"`
	Schedule               Schedule    `json:"schedule"`
}

type Destination struct {
	Bucket Bucket `json:"bucket"`
}

type Bucket struct {
	Format     string     `json:"format"`
	Prefix     string     `json:"prefix"`
	AccountId  string     `json:"account_id"`
	BucketArn  string     `json:"bucket_arn"`
	Encryption Encryption `json:"encryption"`
}

type Encryption struct {
	SseKms SseKms `json:"sse_kms"`
	SseS3  SseS3  `json:"sse_s3"`
}

type SseKms struct {
	KeyId string `json:"key_id"`
}

type SseS3 struct{}

type Filter struct {
	Prefix string `json:"prefix"`
}

type Schedule struct {
	Frequency string `json:"frequency"`
}

// A S3BucketInventoryStatus defines the observed state of a S3BucketInventory
type S3BucketInventoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketInventoryObservation `json:",inline"`
}

// A S3BucketInventoryObservation records the observed state of a S3BucketInventory
type S3BucketInventoryObservation struct{}