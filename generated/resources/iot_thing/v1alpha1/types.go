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

// IotThing is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type IotThing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IotThingSpec   `json:"spec"`
	Status IotThingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IotThing contains a list of IotThingList
type IotThingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IotThing `json:"items"`
}

// A IotThingSpec defines the desired state of a IotThing
type IotThingSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  IotThingParameters `json:"forProvider"`
}

// A IotThingParameters defines the desired state of a IotThing
type IotThingParameters struct {
	Attributes    map[string]string `json:"attributes"`
	Name          string            `json:"name"`
	ThingTypeName string            `json:"thing_type_name"`
}

// A IotThingStatus defines the observed state of a IotThing
type IotThingStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     IotThingObservation `json:"atProvider"`
}

// A IotThingObservation records the observed state of a IotThing
type IotThingObservation struct {
	Arn             string `json:"arn"`
	DefaultClientId string `json:"default_client_id"`
	Version         int64  `json:"version"`
}