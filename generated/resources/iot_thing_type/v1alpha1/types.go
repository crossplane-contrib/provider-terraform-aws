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

// IotThingType is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotThingType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotThingTypeSpec   `json:"spec"`
	Status IotThingTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotThingType contains a list of IotThingTypeList
type IotThingTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotThingType `json:"items"`
}

// A IotThingTypeSpec defines the desired state of a IotThingType
type IotThingTypeSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotThingTypeParameters `json:"forProvider"`
}

// A IotThingTypeParameters defines the desired state of a IotThingType
type IotThingTypeParameters struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Deprecated bool       `json:"deprecated"`
	Properties Properties `json:"properties"`
}

type Properties struct {
	SearchableAttributes []string `json:"searchable_attributes"`
	Description          string   `json:"description"`
}

// A IotThingTypeStatus defines the observed state of a IotThingType
type IotThingTypeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotThingTypeObservation `json:"atProvider"`
}

// A IotThingTypeObservation records the observed state of a IotThingType
type IotThingTypeObservation struct {
	Arn string `json:"arn"`
}