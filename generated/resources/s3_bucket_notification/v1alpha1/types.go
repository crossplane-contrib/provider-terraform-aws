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

// S3BucketNotification is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3BucketNotification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BucketNotificationSpec   `json:"spec"`
	Status S3BucketNotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketNotification contains a list of S3BucketNotificationList
type S3BucketNotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketNotification `json:"items"`
}

// A S3BucketNotificationSpec defines the desired state of a S3BucketNotification
type S3BucketNotificationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3BucketNotificationParameters `json:",inline"`
}

// A S3BucketNotificationParameters defines the desired state of a S3BucketNotification
type S3BucketNotificationParameters struct {
	Bucket string `json:"bucket"`
}

// A S3BucketNotificationStatus defines the observed state of a S3BucketNotification
type S3BucketNotificationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketNotificationObservation `json:",inline"`
}

// A S3BucketNotificationObservation records the observed state of a S3BucketNotification
type S3BucketNotificationObservation struct {
	Id string `json:"id"`
}