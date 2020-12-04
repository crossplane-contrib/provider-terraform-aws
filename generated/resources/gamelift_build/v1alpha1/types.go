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

// GameliftBuild is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type GameliftBuild struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GameliftBuildSpec   `json:"spec"`
	Status GameliftBuildStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GameliftBuild contains a list of GameliftBuildList
type GameliftBuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GameliftBuild `json:"items"`
}

// A GameliftBuildSpec defines the desired state of a GameliftBuild
type GameliftBuildSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  GameliftBuildParameters `json:",inline"`
}

// A GameliftBuildParameters defines the desired state of a GameliftBuild
type GameliftBuildParameters struct {
	Version         string            `json:"version"`
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	OperatingSystem string            `json:"operating_system"`
	Tags            map[string]string `json:"tags"`
	StorageLocation StorageLocation   `json:"storage_location"`
}

type StorageLocation struct {
	RoleArn string `json:"role_arn"`
	Bucket  string `json:"bucket"`
	Key     string `json:"key"`
}

// A GameliftBuildStatus defines the observed state of a GameliftBuild
type GameliftBuildStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     GameliftBuildObservation `json:",inline"`
}

// A GameliftBuildObservation records the observed state of a GameliftBuild
type GameliftBuildObservation struct {
	Arn string `json:"arn"`
}