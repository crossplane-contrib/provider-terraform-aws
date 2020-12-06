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

// ConfigAggregateAuthorization is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type ConfigAggregateAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigAggregateAuthorizationSpec   `json:"spec"`
	Status ConfigAggregateAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigAggregateAuthorization contains a list of ConfigAggregateAuthorizationList
type ConfigAggregateAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigAggregateAuthorization `json:"items"`
}

// A ConfigAggregateAuthorizationSpec defines the desired state of a ConfigAggregateAuthorization
type ConfigAggregateAuthorizationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  ConfigAggregateAuthorizationParameters `json:",inline"`
}

// A ConfigAggregateAuthorizationParameters defines the desired state of a ConfigAggregateAuthorization
type ConfigAggregateAuthorizationParameters struct {
	AccountId string            `json:"account_id"`
	Id        string            `json:"id"`
	Region    string            `json:"region"`
	Tags      map[string]string `json:"tags"`
}

// A ConfigAggregateAuthorizationStatus defines the observed state of a ConfigAggregateAuthorization
type ConfigAggregateAuthorizationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     ConfigAggregateAuthorizationObservation `json:",inline"`
}

// A ConfigAggregateAuthorizationObservation records the observed state of a ConfigAggregateAuthorization
type ConfigAggregateAuthorizationObservation struct {
	Arn string `json:"arn"`
}