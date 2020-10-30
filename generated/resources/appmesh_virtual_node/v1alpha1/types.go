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

// AppmeshVirtualNode is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppmeshVirtualNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppmeshVirtualNodeSpec   `json:"spec"`
	Status AppmeshVirtualNodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshVirtualNode contains a list of AppmeshVirtualNodeList
type AppmeshVirtualNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshVirtualNode `json:"items"`
}

// A AppmeshVirtualNodeSpec defines the desired state of a AppmeshVirtualNode
type AppmeshVirtualNodeSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppmeshVirtualNodeParameters `json:",inline"`
}

// A AppmeshVirtualNodeParameters defines the desired state of a AppmeshVirtualNode
type AppmeshVirtualNodeParameters struct {
	Name     string `json:"name"`
	MeshName string `json:"mesh_name"`
}

// A AppmeshVirtualNodeStatus defines the observed state of a AppmeshVirtualNode
type AppmeshVirtualNodeStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshVirtualNodeObservation `json:",inline"`
}

// A AppmeshVirtualNodeObservation records the observed state of a AppmeshVirtualNode
type AppmeshVirtualNodeObservation struct {
	Id              string `json:"id"`
	ResourceOwner   string `json:"resource_owner"`
	Arn             string `json:"arn"`
	CreatedDate     string `json:"created_date"`
	MeshOwner       string `json:"mesh_owner"`
	LastUpdatedDate string `json:"last_updated_date"`
}