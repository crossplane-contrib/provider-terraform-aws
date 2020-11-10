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

// DaxSubnetGroup is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type DaxSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DaxSubnetGroupSpec   `json:"spec"`
	Status DaxSubnetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DaxSubnetGroup contains a list of DaxSubnetGroupList
type DaxSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DaxSubnetGroup `json:"items"`
}

// A DaxSubnetGroupSpec defines the desired state of a DaxSubnetGroup
type DaxSubnetGroupSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  DaxSubnetGroupParameters `json:",inline"`
}

// A DaxSubnetGroupParameters defines the desired state of a DaxSubnetGroup
type DaxSubnetGroupParameters struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	SubnetIds   []string `json:"subnet_ids"`
}

// A DaxSubnetGroupStatus defines the observed state of a DaxSubnetGroup
type DaxSubnetGroupStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     DaxSubnetGroupObservation `json:",inline"`
}

// A DaxSubnetGroupObservation records the observed state of a DaxSubnetGroup
type DaxSubnetGroupObservation struct {
	Id    string `json:"id"`
	VpcId string `json:"vpc_id"`
}