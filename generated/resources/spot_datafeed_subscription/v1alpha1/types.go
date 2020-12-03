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

// SpotDatafeedSubscription is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type SpotDatafeedSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SpotDatafeedSubscriptionSpec   `json:"spec"`
	Status SpotDatafeedSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpotDatafeedSubscription contains a list of SpotDatafeedSubscriptionList
type SpotDatafeedSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpotDatafeedSubscription `json:"items"`
}

// A SpotDatafeedSubscriptionSpec defines the desired state of a SpotDatafeedSubscription
type SpotDatafeedSubscriptionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  SpotDatafeedSubscriptionParameters `json:",inline"`
}

// A SpotDatafeedSubscriptionParameters defines the desired state of a SpotDatafeedSubscription
type SpotDatafeedSubscriptionParameters struct {
	Id     string `json:"id"`
	Prefix string `json:"prefix"`
	Bucket string `json:"bucket"`
}

// A SpotDatafeedSubscriptionStatus defines the observed state of a SpotDatafeedSubscription
type SpotDatafeedSubscriptionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     SpotDatafeedSubscriptionObservation `json:",inline"`
}

// A SpotDatafeedSubscriptionObservation records the observed state of a SpotDatafeedSubscription
type SpotDatafeedSubscriptionObservation struct{}