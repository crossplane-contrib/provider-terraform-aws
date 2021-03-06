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

// MediaStoreContainer is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type MediaStoreContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MediaStoreContainerSpec   `json:"spec"`
	Status MediaStoreContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MediaStoreContainer contains a list of MediaStoreContainerList
type MediaStoreContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MediaStoreContainer `json:"items"`
}

// A MediaStoreContainerSpec defines the desired state of a MediaStoreContainer
type MediaStoreContainerSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  MediaStoreContainerParameters `json:"forProvider"`
}

// A MediaStoreContainerParameters defines the desired state of a MediaStoreContainer
type MediaStoreContainerParameters struct {
	Name string            `json:"name"`
	Tags map[string]string `json:"tags"`
}

// A MediaStoreContainerStatus defines the observed state of a MediaStoreContainer
type MediaStoreContainerStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     MediaStoreContainerObservation `json:"atProvider"`
}

// A MediaStoreContainerObservation records the observed state of a MediaStoreContainer
type MediaStoreContainerObservation struct {
	Arn      string `json:"arn"`
	Endpoint string `json:"endpoint"`
}