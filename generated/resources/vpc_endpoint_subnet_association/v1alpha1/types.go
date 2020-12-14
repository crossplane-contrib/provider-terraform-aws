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

// VpcEndpointSubnetAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcEndpointSubnetAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcEndpointSubnetAssociationSpec   `json:"spec"`
	Status VpcEndpointSubnetAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcEndpointSubnetAssociation contains a list of VpcEndpointSubnetAssociationList
type VpcEndpointSubnetAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcEndpointSubnetAssociation `json:"items"`
}

// A VpcEndpointSubnetAssociationSpec defines the desired state of a VpcEndpointSubnetAssociation
type VpcEndpointSubnetAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcEndpointSubnetAssociationParameters `json:"forProvider"`
}

// A VpcEndpointSubnetAssociationParameters defines the desired state of a VpcEndpointSubnetAssociation
type VpcEndpointSubnetAssociationParameters struct {
	VpcEndpointId string   `json:"vpc_endpoint_id"`
	Id            string   `json:"id"`
	SubnetId      string   `json:"subnet_id"`
	Timeouts      Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A VpcEndpointSubnetAssociationStatus defines the observed state of a VpcEndpointSubnetAssociation
type VpcEndpointSubnetAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcEndpointSubnetAssociationObservation `json:"atProvider"`
}

// A VpcEndpointSubnetAssociationObservation records the observed state of a VpcEndpointSubnetAssociation
type VpcEndpointSubnetAssociationObservation struct{}