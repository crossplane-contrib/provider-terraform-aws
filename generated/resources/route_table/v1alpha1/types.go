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

// RouteTable is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteTableSpec   `json:"spec"`
	Status RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTable contains a list of RouteTableList
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// A RouteTableSpec defines the desired state of a RouteTable
type RouteTableSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  RouteTableParameters `json:",inline"`
}

// A RouteTableParameters defines the desired state of a RouteTable
type RouteTableParameters struct {
	Tags  map[string]string `json:"tags"`
	VpcId string            `json:"vpc_id"`
}

// A RouteTableStatus defines the observed state of a RouteTable
type RouteTableStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     RouteTableObservation `json:",inline"`
}

// A RouteTableObservation records the observed state of a RouteTable
type RouteTableObservation struct {
	Id              string   `json:"id"`
	OwnerId         string   `json:"owner_id"`
	PropagatingVgws []string `json:"propagating_vgws"`
}