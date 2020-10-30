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
	Name                   string   `json:"name"`
	OptionalFields         []string `json:"optional_fields"`
	Bucket                 string   `json:"bucket"`
	Enabled                bool     `json:"enabled"`
	IncludedObjectVersions string   `json:"included_object_versions"`
}

// A S3BucketInventoryStatus defines the observed state of a S3BucketInventory
type S3BucketInventoryStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketInventoryObservation `json:",inline"`
}

// A S3BucketInventoryObservation records the observed state of a S3BucketInventory
type S3BucketInventoryObservation struct {
	Id string `json:"id"`
}