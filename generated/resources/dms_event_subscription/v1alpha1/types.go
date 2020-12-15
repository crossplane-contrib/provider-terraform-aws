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

// DmsEventSubscription is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DmsEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DmsEventSubscriptionSpec   `json:"spec"`
	Status DmsEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DmsEventSubscription contains a list of DmsEventSubscriptionList
type DmsEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DmsEventSubscription `json:"items"`
}

// A DmsEventSubscriptionSpec defines the desired state of a DmsEventSubscription
type DmsEventSubscriptionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DmsEventSubscriptionParameters `json:"forProvider"`
}

// A DmsEventSubscriptionParameters defines the desired state of a DmsEventSubscription
type DmsEventSubscriptionParameters struct {
	Enabled         bool              `json:"enabled"`
	EventCategories []string          `json:"event_categories"`
	Id              string            `json:"id"`
	SourceType      string            `json:"source_type"`
	Name            string            `json:"name"`
	SnsTopicArn     string            `json:"sns_topic_arn"`
	SourceIds       []string          `json:"source_ids"`
	Tags            map[string]string `json:"tags"`
	Timeouts        Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DmsEventSubscriptionStatus defines the observed state of a DmsEventSubscription
type DmsEventSubscriptionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DmsEventSubscriptionObservation `json:"atProvider"`
}

// A DmsEventSubscriptionObservation records the observed state of a DmsEventSubscription
type DmsEventSubscriptionObservation struct {
	Arn string `json:"arn"`
}