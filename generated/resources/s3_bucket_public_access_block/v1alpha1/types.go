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

// S3BucketPublicAccessBlock is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3BucketPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BucketPublicAccessBlockSpec   `json:"spec"`
	Status S3BucketPublicAccessBlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketPublicAccessBlock contains a list of S3BucketPublicAccessBlockList
type S3BucketPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketPublicAccessBlock `json:"items"`
}

// A S3BucketPublicAccessBlockSpec defines the desired state of a S3BucketPublicAccessBlock
type S3BucketPublicAccessBlockSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3BucketPublicAccessBlockParameters `json:",inline"`
}

// A S3BucketPublicAccessBlockParameters defines the desired state of a S3BucketPublicAccessBlock
type S3BucketPublicAccessBlockParameters struct {
	BlockPublicAcls       bool   `json:"block_public_acls"`
	BlockPublicPolicy     bool   `json:"block_public_policy"`
	Bucket                string `json:"bucket"`
	Id                    string `json:"id"`
	IgnorePublicAcls      bool   `json:"ignore_public_acls"`
	RestrictPublicBuckets bool   `json:"restrict_public_buckets"`
}

// A S3BucketPublicAccessBlockStatus defines the observed state of a S3BucketPublicAccessBlock
type S3BucketPublicAccessBlockStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketPublicAccessBlockObservation `json:",inline"`
}

// A S3BucketPublicAccessBlockObservation records the observed state of a S3BucketPublicAccessBlock
type S3BucketPublicAccessBlockObservation struct{}