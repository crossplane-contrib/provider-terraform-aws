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

// S3BucketPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3BucketPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BucketPolicySpec   `json:"spec"`
	Status S3BucketPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketPolicy contains a list of S3BucketPolicyList
type S3BucketPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketPolicy `json:"items"`
}

// A S3BucketPolicySpec defines the desired state of a S3BucketPolicy
type S3BucketPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3BucketPolicyParameters `json:"forProvider"`
}

// A S3BucketPolicyParameters defines the desired state of a S3BucketPolicy
type S3BucketPolicyParameters struct {
	Bucket string `json:"bucket"`
	Id     string `json:"id"`
	Policy string `json:"policy"`
}

// A S3BucketPolicyStatus defines the observed state of a S3BucketPolicy
type S3BucketPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketPolicyObservation `json:"atProvider"`
}

// A S3BucketPolicyObservation records the observed state of a S3BucketPolicy
type S3BucketPolicyObservation struct{}