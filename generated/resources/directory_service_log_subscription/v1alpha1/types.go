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

// DirectoryServiceLogSubscription is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DirectoryServiceLogSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DirectoryServiceLogSubscriptionSpec   `json:"spec"`
	Status DirectoryServiceLogSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryServiceLogSubscription contains a list of DirectoryServiceLogSubscriptionList
type DirectoryServiceLogSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectoryServiceLogSubscription `json:"items"`
}

// A DirectoryServiceLogSubscriptionSpec defines the desired state of a DirectoryServiceLogSubscription
type DirectoryServiceLogSubscriptionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DirectoryServiceLogSubscriptionParameters `json:",inline"`
}

// A DirectoryServiceLogSubscriptionParameters defines the desired state of a DirectoryServiceLogSubscription
type DirectoryServiceLogSubscriptionParameters struct {
	DirectoryId  string `json:"directory_id"`
	LogGroupName string `json:"log_group_name"`
}

// A DirectoryServiceLogSubscriptionStatus defines the observed state of a DirectoryServiceLogSubscription
type DirectoryServiceLogSubscriptionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DirectoryServiceLogSubscriptionObservation `json:",inline"`
}

// A DirectoryServiceLogSubscriptionObservation records the observed state of a DirectoryServiceLogSubscription
type DirectoryServiceLogSubscriptionObservation struct {
	Id string `json:"id"`
}