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

// S3BucketAnalyticsConfiguration is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type S3BucketAnalyticsConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   S3BucketAnalyticsConfigurationSpec   `json:"spec"`
	Status S3BucketAnalyticsConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketAnalyticsConfiguration contains a list of S3BucketAnalyticsConfigurationList
type S3BucketAnalyticsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketAnalyticsConfiguration `json:"items"`
}

// A S3BucketAnalyticsConfigurationSpec defines the desired state of a S3BucketAnalyticsConfiguration
type S3BucketAnalyticsConfigurationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  S3BucketAnalyticsConfigurationParameters `json:",inline"`
}

// A S3BucketAnalyticsConfigurationParameters defines the desired state of a S3BucketAnalyticsConfiguration
type S3BucketAnalyticsConfigurationParameters struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// A S3BucketAnalyticsConfigurationStatus defines the observed state of a S3BucketAnalyticsConfiguration
type S3BucketAnalyticsConfigurationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     S3BucketAnalyticsConfigurationObservation `json:",inline"`
}

// A S3BucketAnalyticsConfigurationObservation records the observed state of a S3BucketAnalyticsConfiguration
type S3BucketAnalyticsConfigurationObservation struct {
	Id string `json:"id"`
}