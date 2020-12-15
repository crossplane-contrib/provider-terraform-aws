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

// S3AccountPublicAccessBlock is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3AccountPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3AccountPublicAccessBlockSpec   `json:"spec"`
	Status S3AccountPublicAccessBlockStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3AccountPublicAccessBlock contains a list of S3AccountPublicAccessBlockList
type S3AccountPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3AccountPublicAccessBlock `json:"items"`
}

// A S3AccountPublicAccessBlockSpec defines the desired state of a S3AccountPublicAccessBlock
type S3AccountPublicAccessBlockSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3AccountPublicAccessBlockParameters `json:"forProvider"`
}

// A S3AccountPublicAccessBlockParameters defines the desired state of a S3AccountPublicAccessBlock
type S3AccountPublicAccessBlockParameters struct {
	BlockPublicPolicy     bool   `json:"block_public_policy"`
	Id                    string `json:"id"`
	IgnorePublicAcls      bool   `json:"ignore_public_acls"`
	RestrictPublicBuckets bool   `json:"restrict_public_buckets"`
	AccountId             string `json:"account_id"`
	BlockPublicAcls       bool   `json:"block_public_acls"`
}

// A S3AccountPublicAccessBlockStatus defines the observed state of a S3AccountPublicAccessBlock
type S3AccountPublicAccessBlockStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3AccountPublicAccessBlockObservation `json:"atProvider"`
}

// A S3AccountPublicAccessBlockObservation records the observed state of a S3AccountPublicAccessBlock
type S3AccountPublicAccessBlockObservation struct{}