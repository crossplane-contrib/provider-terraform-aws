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

// VpcDhcpOptionsAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcDhcpOptionsAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcDhcpOptionsAssociationSpec   `json:"spec"`
	Status VpcDhcpOptionsAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcDhcpOptionsAssociation contains a list of VpcDhcpOptionsAssociationList
type VpcDhcpOptionsAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcDhcpOptionsAssociation `json:"items"`
}

// A VpcDhcpOptionsAssociationSpec defines the desired state of a VpcDhcpOptionsAssociation
type VpcDhcpOptionsAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcDhcpOptionsAssociationParameters `json:"forProvider"`
}

// A VpcDhcpOptionsAssociationParameters defines the desired state of a VpcDhcpOptionsAssociation
type VpcDhcpOptionsAssociationParameters struct {
	DhcpOptionsId string `json:"dhcp_options_id"`
	VpcId         string `json:"vpc_id"`
}

// A VpcDhcpOptionsAssociationStatus defines the observed state of a VpcDhcpOptionsAssociation
type VpcDhcpOptionsAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcDhcpOptionsAssociationObservation `json:"atProvider"`
}

// A VpcDhcpOptionsAssociationObservation records the observed state of a VpcDhcpOptionsAssociation
type VpcDhcpOptionsAssociationObservation struct{}