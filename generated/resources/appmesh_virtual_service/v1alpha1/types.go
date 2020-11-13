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

// AppmeshVirtualService is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppmeshVirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppmeshVirtualServiceSpec   `json:"spec"`
	Status AppmeshVirtualServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualService contains a list of AppmeshVirtualServiceList
type AppmeshVirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshVirtualService `json:"items"`
}

// A AppmeshVirtualServiceSpec defines the desired state of a AppmeshVirtualService
type AppmeshVirtualServiceSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppmeshVirtualServiceParameters `json:",inline"`
}

// A AppmeshVirtualServiceParameters defines the desired state of a AppmeshVirtualService
type AppmeshVirtualServiceParameters struct {
	Tags      map[string]string `json:"tags"`
	MeshName  string            `json:"mesh_name"`
	MeshOwner string            `json:"mesh_owner"`
	Name      string            `json:"name"`
	Id        string            `json:"id"`
	Spec      Spec              `json:"spec"`
}

type Spec struct {
	Provider Provider `json:"provider"`
}

type Provider struct {
	VirtualNode   VirtualNode   `json:"virtual_node"`
	VirtualRouter VirtualRouter `json:"virtual_router"`
}

type VirtualNode struct {
	VirtualNodeName string `json:"virtual_node_name"`
}

type VirtualRouter struct {
	VirtualRouterName string `json:"virtual_router_name"`
}

// A AppmeshVirtualServiceStatus defines the observed state of a AppmeshVirtualService
type AppmeshVirtualServiceStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshVirtualServiceObservation `json:",inline"`
}

// A AppmeshVirtualServiceObservation records the observed state of a AppmeshVirtualService
type AppmeshVirtualServiceObservation struct {
	Arn             string `json:"arn"`
	CreatedDate     string `json:"created_date"`
	LastUpdatedDate string `json:"last_updated_date"`
	ResourceOwner   string `json:"resource_owner"`
}