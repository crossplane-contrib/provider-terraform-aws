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

// GameliftAlias is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GameliftAlias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GameliftAliasSpec   `json:"spec"`
	Status GameliftAliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftAlias contains a list of GameliftAliasList
type GameliftAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameliftAlias `json:"items"`
}

// A GameliftAliasSpec defines the desired state of a GameliftAlias
type GameliftAliasSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GameliftAliasParameters `json:",inline"`
}

// A GameliftAliasParameters defines the desired state of a GameliftAlias
type GameliftAliasParameters struct {
	Name            string            `json:"name"`
	Tags            map[string]string `json:"tags"`
	Description     string            `json:"description"`
	RoutingStrategy RoutingStrategy   `json:"routing_strategy"`
}

type RoutingStrategy struct {
	FleetId string `json:"fleet_id"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

// A GameliftAliasStatus defines the observed state of a GameliftAlias
type GameliftAliasStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GameliftAliasObservation `json:",inline"`
}

// A GameliftAliasObservation records the observed state of a GameliftAlias
type GameliftAliasObservation struct {
	Id  string `json:"id"`
	Arn string `json:"arn"`
}