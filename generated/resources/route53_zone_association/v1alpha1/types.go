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

// Route53ZoneAssociation is a managed resource representing a resource mirrored in the cloud
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster
type Route53ZoneAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Route53ZoneAssociationSpec   `json:"spec"`
	Status Route53ZoneAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Route53ZoneAssociation contains a list of Route53ZoneAssociationList
type Route53ZoneAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route53ZoneAssociation `json:"items"`
}

// A Route53ZoneAssociationSpec defines the desired state of a Route53ZoneAssociation
type Route53ZoneAssociationSpec struct {
	runtimev1alpha1.ResourceSpec `json:",inline"`
	ForProvider                  Route53ZoneAssociationParameters `json:",inline"`
}

// A Route53ZoneAssociationParameters defines the desired state of a Route53ZoneAssociation
type Route53ZoneAssociationParameters struct {
	VpcRegion string `json:"vpc_region"`
	ZoneId    string `json:"zone_id"`
	Id        string `json:"id"`
	VpcId     string `json:"vpc_id"`
}

// A Route53ZoneAssociationStatus defines the observed state of a Route53ZoneAssociation
type Route53ZoneAssociationStatus struct {
	runtimev1alpha1.ResourceStatus `json:",inline"`
	AtProvider                     Route53ZoneAssociationObservation `json:",inline"`
}

// A Route53ZoneAssociationObservation records the observed state of a Route53ZoneAssociation
type Route53ZoneAssociationObservation struct {
	OwningAccount string `json:"owning_account"`
}