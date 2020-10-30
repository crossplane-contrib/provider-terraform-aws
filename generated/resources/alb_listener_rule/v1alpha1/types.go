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

// AlbListenerRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AlbListenerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlbListenerRuleSpec   `json:"spec"`
	Status AlbListenerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AlbListenerRule contains a list of AlbListenerRuleList
type AlbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlbListenerRule `json:"items"`
}

// A AlbListenerRuleSpec defines the desired state of a AlbListenerRule
type AlbListenerRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AlbListenerRuleParameters `json:",inline"`
}

// A AlbListenerRuleParameters defines the desired state of a AlbListenerRule
type AlbListenerRuleParameters struct {
	ListenerArn string `json:"listener_arn"`
}

// A AlbListenerRuleStatus defines the observed state of a AlbListenerRule
type AlbListenerRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AlbListenerRuleObservation `json:",inline"`
}

// A AlbListenerRuleObservation records the observed state of a AlbListenerRule
type AlbListenerRuleObservation struct {
	Priority int    `json:"priority"`
	Arn      string `json:"arn"`
	Id       string `json:"id"`
}