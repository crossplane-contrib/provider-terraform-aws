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

// DbEventSubscription is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DbEventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DbEventSubscriptionSpec   `json:"spec"`
	Status DbEventSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DbEventSubscription contains a list of DbEventSubscriptionList
type DbEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DbEventSubscription `json:"items"`
}

// A DbEventSubscriptionSpec defines the desired state of a DbEventSubscription
type DbEventSubscriptionSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DbEventSubscriptionParameters `json:"forProvider"`
}

// A DbEventSubscriptionParameters defines the desired state of a DbEventSubscription
type DbEventSubscriptionParameters struct {
	EventCategories []string          `json:"event_categories"`
	Name            string            `json:"name"`
	NamePrefix      string            `json:"name_prefix"`
	SnsTopic        string            `json:"sns_topic"`
	Tags            map[string]string `json:"tags"`
	Enabled         bool              `json:"enabled"`
	Id              string            `json:"id"`
	SourceIds       []string          `json:"source_ids"`
	SourceType      string            `json:"source_type"`
	Timeouts        Timeouts          `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

// A DbEventSubscriptionStatus defines the observed state of a DbEventSubscription
type DbEventSubscriptionStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DbEventSubscriptionObservation `json:"atProvider"`
}

// A DbEventSubscriptionObservation records the observed state of a DbEventSubscription
type DbEventSubscriptionObservation struct {
	Arn           string `json:"arn"`
	CustomerAwsId string `json:"customer_aws_id"`
}