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

// AppmeshRoute is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppmeshRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppmeshRouteSpec   `json:"spec"`
	Status AppmeshRouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshRoute contains a list of AppmeshRouteList
type AppmeshRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshRoute `json:"items"`
}

// A AppmeshRouteSpec defines the desired state of a AppmeshRoute
type AppmeshRouteSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppmeshRouteParameters `json:",inline"`
}

// A AppmeshRouteParameters defines the desired state of a AppmeshRoute
type AppmeshRouteParameters struct {
	VirtualRouterName string `json:"virtual_router_name"`
	MeshName          string `json:"mesh_name"`
	Name              string `json:"name"`
}

// A AppmeshRouteStatus defines the observed state of a AppmeshRoute
type AppmeshRouteStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshRouteObservation `json:",inline"`
}

// A AppmeshRouteObservation records the observed state of a AppmeshRoute
type AppmeshRouteObservation struct {
	Arn             string `json:"arn"`
	LastUpdatedDate string `json:"last_updated_date"`
	CreatedDate     string `json:"created_date"`
	Id              string `json:"id"`
	MeshOwner       string `json:"mesh_owner"`
	ResourceOwner   string `json:"resource_owner"`
}