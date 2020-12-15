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

// DevicefarmProject is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DevicefarmProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DevicefarmProjectSpec   `json:"spec"`
	Status DevicefarmProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DevicefarmProject contains a list of DevicefarmProjectList
type DevicefarmProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DevicefarmProject `json:"items"`
}

// A DevicefarmProjectSpec defines the desired state of a DevicefarmProject
type DevicefarmProjectSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DevicefarmProjectParameters `json:"forProvider"`
}

// A DevicefarmProjectParameters defines the desired state of a DevicefarmProject
type DevicefarmProjectParameters struct {
	Name string `json:"name"`
}

// A DevicefarmProjectStatus defines the observed state of a DevicefarmProject
type DevicefarmProjectStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DevicefarmProjectObservation `json:"atProvider"`
}

// A DevicefarmProjectObservation records the observed state of a DevicefarmProject
type DevicefarmProjectObservation struct {
	Arn string `json:"arn"`
}