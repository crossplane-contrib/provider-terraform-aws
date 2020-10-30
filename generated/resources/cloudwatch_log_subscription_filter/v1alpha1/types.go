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

// CloudwatchLogSubscriptionFilter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchLogSubscriptionFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchLogSubscriptionFilterSpec   `json:"spec"`
	Status CloudwatchLogSubscriptionFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogSubscriptionFilter contains a list of CloudwatchLogSubscriptionFilterList
type CloudwatchLogSubscriptionFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogSubscriptionFilter `json:"items"`
}

// A CloudwatchLogSubscriptionFilterSpec defines the desired state of a CloudwatchLogSubscriptionFilter
type CloudwatchLogSubscriptionFilterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchLogSubscriptionFilterParameters `json:",inline"`
}

// A CloudwatchLogSubscriptionFilterParameters defines the desired state of a CloudwatchLogSubscriptionFilter
type CloudwatchLogSubscriptionFilterParameters struct {
	LogGroupName   string `json:"log_group_name"`
	Name           string `json:"name"`
	DestinationArn string `json:"destination_arn"`
	Distribution   string `json:"distribution"`
	FilterPattern  string `json:"filter_pattern"`
}

// A CloudwatchLogSubscriptionFilterStatus defines the observed state of a CloudwatchLogSubscriptionFilter
type CloudwatchLogSubscriptionFilterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchLogSubscriptionFilterObservation `json:",inline"`
}

// A CloudwatchLogSubscriptionFilterObservation records the observed state of a CloudwatchLogSubscriptionFilter
type CloudwatchLogSubscriptionFilterObservation struct {
	Id      string `json:"id"`
	RoleArn string `json:"role_arn"`
}