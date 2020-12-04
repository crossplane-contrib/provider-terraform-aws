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

// AppmeshVirtualRouter is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppmeshVirtualRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppmeshVirtualRouterSpec   `json:"spec"`
	Status AppmeshVirtualRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualRouter contains a list of AppmeshVirtualRouterList
type AppmeshVirtualRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshVirtualRouter `json:"items"`
}

// A AppmeshVirtualRouterSpec defines the desired state of a AppmeshVirtualRouter
type AppmeshVirtualRouterSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppmeshVirtualRouterParameters `json:",inline"`
}

// A AppmeshVirtualRouterParameters defines the desired state of a AppmeshVirtualRouter
type AppmeshVirtualRouterParameters struct {
	Name      string            `json:"name"`
	MeshName  string            `json:"mesh_name"`
	MeshOwner string            `json:"mesh_owner"`
	Tags      map[string]string `json:"tags"`
	Id        string            `json:"id"`
	Spec      Spec              `json:"spec"`
}

type Spec struct {
	Listener Listener `json:"listener"`
}

type Listener struct {
	PortMapping PortMapping `json:"port_mapping"`
}

type PortMapping struct {
	Port     int64  `json:"port"`
	Protocol string `json:"protocol"`
}

// A AppmeshVirtualRouterStatus defines the observed state of a AppmeshVirtualRouter
type AppmeshVirtualRouterStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshVirtualRouterObservation `json:",inline"`
}

// A AppmeshVirtualRouterObservation records the observed state of a AppmeshVirtualRouter
type AppmeshVirtualRouterObservation struct {
	CreatedDate     string `json:"created_date"`
	LastUpdatedDate string `json:"last_updated_date"`
	Arn             string `json:"arn"`
	ResourceOwner   string `json:"resource_owner"`
}