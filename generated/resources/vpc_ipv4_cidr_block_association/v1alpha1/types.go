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

// VpcIpv4CidrBlockAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type VpcIpv4CidrBlockAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpcIpv4CidrBlockAssociationSpec   `json:"spec"`
	Status VpcIpv4CidrBlockAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VpcIpv4CidrBlockAssociation contains a list of VpcIpv4CidrBlockAssociationList
type VpcIpv4CidrBlockAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpcIpv4CidrBlockAssociation `json:"items"`
}

// A VpcIpv4CidrBlockAssociationSpec defines the desired state of a VpcIpv4CidrBlockAssociation
type VpcIpv4CidrBlockAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  VpcIpv4CidrBlockAssociationParameters `json:",inline"`
}

// A VpcIpv4CidrBlockAssociationParameters defines the desired state of a VpcIpv4CidrBlockAssociation
type VpcIpv4CidrBlockAssociationParameters struct {
	CidrBlock string     `json:"cidr_block"`
	Id        string     `json:"id"`
	VpcId     string     `json:"vpc_id"`
	Timeouts  []Timeouts `json:"timeouts"`
}

type Timeouts struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}

// A VpcIpv4CidrBlockAssociationStatus defines the observed state of a VpcIpv4CidrBlockAssociation
type VpcIpv4CidrBlockAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     VpcIpv4CidrBlockAssociationObservation `json:",inline"`
}

// A VpcIpv4CidrBlockAssociationObservation records the observed state of a VpcIpv4CidrBlockAssociation
type VpcIpv4CidrBlockAssociationObservation struct{}