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

// KinesisAnalyticsApplication is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type KinesisAnalyticsApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KinesisAnalyticsApplicationSpec   `json:"spec"`
	Status KinesisAnalyticsApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KinesisAnalyticsApplication contains a list of KinesisAnalyticsApplicationList
type KinesisAnalyticsApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KinesisAnalyticsApplication `json:"items"`
}

// A KinesisAnalyticsApplicationSpec defines the desired state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  KinesisAnalyticsApplicationParameters `json:",inline"`
}

// A KinesisAnalyticsApplicationParameters defines the desired state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationParameters struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

// A KinesisAnalyticsApplicationStatus defines the observed state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     KinesisAnalyticsApplicationObservation `json:",inline"`
}

// A KinesisAnalyticsApplicationObservation records the observed state of a KinesisAnalyticsApplication
type KinesisAnalyticsApplicationObservation struct {
	Arn                 string `json:"arn"`
	LastUpdateTimestamp string `json:"last_update_timestamp"`
	Status              string `json:"status"`
	Version             int    `json:"version"`
	CreateTimestamp     string `json:"create_timestamp"`
	Id                  string `json:"id"`
}