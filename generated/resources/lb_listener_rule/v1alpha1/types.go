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

// LbListenerRule is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type LbListenerRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LbListenerRuleSpec   `json:"spec"`
	Status LbListenerRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LbListenerRule contains a list of LbListenerRuleList
type LbListenerRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LbListenerRule `json:"items"`
}

// A LbListenerRuleSpec defines the desired state of a LbListenerRule
type LbListenerRuleSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  LbListenerRuleParameters `json:",inline"`
}

// A LbListenerRuleParameters defines the desired state of a LbListenerRule
type LbListenerRuleParameters struct {
	ListenerArn string `json:"listener_arn"`
}

// A LbListenerRuleStatus defines the observed state of a LbListenerRule
type LbListenerRuleStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     LbListenerRuleObservation `json:",inline"`
}

// A LbListenerRuleObservation records the observed state of a LbListenerRule
type LbListenerRuleObservation struct {
	Id       string `json:"id"`
	Priority int    `json:"priority"`
	Arn      string `json:"arn"`
}