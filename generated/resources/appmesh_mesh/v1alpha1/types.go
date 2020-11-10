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

// AppmeshMesh is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type AppmeshMesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppmeshMeshSpec   `json:"spec"`
	Status AppmeshMeshStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppmeshMesh contains a list of AppmeshMeshList
type AppmeshMeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppmeshMesh `json:"items"`
}

// A AppmeshMeshSpec defines the desired state of a AppmeshMesh
type AppmeshMeshSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  AppmeshMeshParameters `json:",inline"`
}

// A AppmeshMeshParameters defines the desired state of a AppmeshMesh
type AppmeshMeshParameters struct {
	Name string            `json:"name"`
	Tags map[string]string `json:"tags"`
	Spec Spec              `json:"spec"`
}

type Spec struct {
	EgressFilter EgressFilter `json:"egress_filter"`
}

type EgressFilter struct {
	Type string `json:"type"`
}

// A AppmeshMeshStatus defines the observed state of a AppmeshMesh
type AppmeshMeshStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     AppmeshMeshObservation `json:",inline"`
}

// A AppmeshMeshObservation records the observed state of a AppmeshMesh
type AppmeshMeshObservation struct {
	Id              string `json:"id"`
	LastUpdatedDate string `json:"last_updated_date"`
	MeshOwner       string `json:"mesh_owner"`
	ResourceOwner   string `json:"resource_owner"`
	Arn             string `json:"arn"`
	CreatedDate     string `json:"created_date"`
}