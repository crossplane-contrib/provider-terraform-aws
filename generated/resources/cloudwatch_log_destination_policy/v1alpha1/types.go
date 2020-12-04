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

// CloudwatchLogDestinationPolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchLogDestinationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchLogDestinationPolicySpec   `json:"spec"`
	Status CloudwatchLogDestinationPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogDestinationPolicy contains a list of CloudwatchLogDestinationPolicyList
type CloudwatchLogDestinationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogDestinationPolicy `json:"items"`
}

// A CloudwatchLogDestinationPolicySpec defines the desired state of a CloudwatchLogDestinationPolicy
type CloudwatchLogDestinationPolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchLogDestinationPolicyParameters `json:",inline"`
}

// A CloudwatchLogDestinationPolicyParameters defines the desired state of a CloudwatchLogDestinationPolicy
type CloudwatchLogDestinationPolicyParameters struct {
	DestinationName string `json:"destination_name"`
	Id              string `json:"id"`
	AccessPolicy    string `json:"access_policy"`
}

// A CloudwatchLogDestinationPolicyStatus defines the observed state of a CloudwatchLogDestinationPolicy
type CloudwatchLogDestinationPolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchLogDestinationPolicyObservation `json:",inline"`
}

// A CloudwatchLogDestinationPolicyObservation records the observed state of a CloudwatchLogDestinationPolicy
type CloudwatchLogDestinationPolicyObservation struct{}