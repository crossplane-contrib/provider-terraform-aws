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

// CloudwatchLogResourcePolicy is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type CloudwatchLogResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudwatchLogResourcePolicySpec   `json:"spec"`
	Status CloudwatchLogResourcePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchLogResourcePolicy contains a list of CloudwatchLogResourcePolicyList
type CloudwatchLogResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchLogResourcePolicy `json:"items"`
}

// A CloudwatchLogResourcePolicySpec defines the desired state of a CloudwatchLogResourcePolicy
type CloudwatchLogResourcePolicySpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  CloudwatchLogResourcePolicyParameters `json:"forProvider"`
}

// A CloudwatchLogResourcePolicyParameters defines the desired state of a CloudwatchLogResourcePolicy
type CloudwatchLogResourcePolicyParameters struct {
	Id             string `json:"id"`
	PolicyDocument string `json:"policy_document"`
	PolicyName     string `json:"policy_name"`
}

// A CloudwatchLogResourcePolicyStatus defines the observed state of a CloudwatchLogResourcePolicy
type CloudwatchLogResourcePolicyStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     CloudwatchLogResourcePolicyObservation `json:"atProvider"`
}

// A CloudwatchLogResourcePolicyObservation records the observed state of a CloudwatchLogResourcePolicy
type CloudwatchLogResourcePolicyObservation struct{}